package main

import (
	"bufio"
	"fmt"
	"strconv"
)

// Définition des items consommables disponibles en jeu
var ItemsList = []Item{
	{Name: "Potion de soin", Effect: "soin", Value: 30, Price: 25},                   // Rend 30 PV
	{Name: "Pain sec", Effect: "faim", Value: 5, Price: 10},                          // Restaure 5 de faim
	{Name: "Potion de poison", Effect: "poison", Value: 20, Price: 20},               // Nouveau : inflige du poison (utilisable en combat)
	{Name: "Livre de Sort : Boule de Feu", Effect: "spellbook", Value: 0, Price: 50}, // Nouveau : apprend un sort
}

// Affiche le contenu de la boutique (équipements + consommables)
func afficherBoutique() {
	fmt.Println("\n🛒 Boutique — Bienvenue !")

	// Affichage des équipements disponibles
	fmt.Println("Équipements disponibles :")
	for i, eq := range Equipments {
		fmt.Printf("%d. [%s] %s (+%d PV / +%d Dégâts) - %d capsules\n",
			i+1, eq.Type, eq.Name, eq.BonusHP, eq.BonusDmg, eq.Price)
	}

	// Affichage des objets consommables disponibles
	fmt.Println("\nConsommables :")
	for i, item := range ItemsList {
		fmt.Printf("%d. %s (%s +%d) - %d capsules\n",
			i+len(Equipments)+1, item.Name, item.Effect, item.Value, item.Price)
	}
}

// Gestion de l’achat dans la boutique

func acheterDansBoutique(c *Character, scanner *bufio.Scanner) {
	for {
		// Affiche la boutique à chaque tour
		afficherBoutique()
		fmt.Printf("\n💰 Capsules : %d\n", c.Argent)
		fmt.Print("Entre le numéro de l’objet à acheter (ou '0' pour quitter) : ")

		// Récupère l’entrée utilisateur
		scanner.Scan()
		input := scanner.Text()
		choix, err := strconv.Atoi(input)
		if err != nil || choix < 0 {
			fmt.Println("❌ Entrée invalide.")
			continue
		}

		// Si le joueur tape "0", il quitte la boutique
		if choix == 0 {
			fmt.Println("🚪 Tu quittes la boutique.")
			return
		}

		// Achat d’un équipement
		if choix >= 1 && choix <= len(Equipments) {
			item := Equipments[choix-1]
			if c.Argent >= item.Price {
				c.Inventaire = append(c.Inventaire, item.Name)
				c.Argent -= item.Price
				fmt.Printf("✅ Tu as acheté : %s\n", item.Name)
			} else {
				fmt.Println("❌ Pas assez de capsules.")
			}

			// Achat d’un consommable
		} else if choix <= len(Equipments)+len(ItemsList) {
			item := ItemsList[choix-len(Equipments)-1]
			if c.Argent >= item.Price {
				c.Inventaire = append(c.Inventaire, item.Name)
				c.Argent -= item.Price
				fmt.Printf("✅ Tu as acheté : %s\n", item.Name)
			} else {
				fmt.Println("❌ Pas assez de capsules.")
			}

			// Numéro invalide
		} else {
			fmt.Println("❌ Numéro invalide.")
		}
	}
}
