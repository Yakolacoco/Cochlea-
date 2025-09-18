package main

import (
	"bufio"  // pour lire l'entrée utilisateur
	"fmt"    // pour afficher des messages
	"strconv" // pour convertir string -> int
)

// ItemsList : liste des items consommables disponibles dans la boutique
// Value = valeur de soin/faim/poison
// Price = prix en capsules
var ItemsList = []Item{
	{Name: "Potion de soin", Effect: "soin", Value: 30, Price: 25},                  // soigne 30 PV
	{Name: "Pain sec", Effect: "faim", Value: 5, Price: 10},                         // +5 faim
	{Name: "Potion de poison", Effect: "poison", Value: 20, Price: 20},              // inflige poison en combat
	{Name: "Livre de Sort : Boule de Feu", Effect: "spellbook", Value: 0, Price: 50}, // apprend un sort
	{Name: "Potion de soin majeure", Effect: "soin", Value: 60, Price: 60},           // soigne beaucoup
	{Name: "Sandwich frais", Effect: "faim", Value: 10, Price: 15},                   // nourriture + faim
	{Name: "Élixir de régénération", Effect: "soin", Value: 100, Price: 150},        // soigne totalement
	{Name: "Potion de force", Effect: "buff", Value: 5, Price: 50},                   // augmente dégâts temporairement
	{Name: "Barre énergétique", Effect: "faim", Value: 8, Price: 12},                 // +8 faim
	{Name: "Potion de rapidité", Effect: "initiative", Value: 3, Price: 60},          // +3 initiative temporaire
	{Name: "Herbe médicinale", Effect: "soin", Value: 20, Price: 15},                 // soin léger
	{Name: "Livre de Sort : Éclair", Effect: "spellbook", Value: 0, Price: 60},       // nouveau sort
}


// afficherBoutique : affiche tout le contenu de la boutique (équipements + items)
// n'utilise pas de retour, juste affiche dans le terminal
func afficherBoutique() {
	fmt.Println("\n🛒 Boutique — Bienvenue !")

	// affiche les équipements disponibles
	fmt.Println("Équipements disponibles :")
	for i, eq := range Equipments {
		fmt.Printf("%d. [%s] %s (+%d PV / +%d Dégâts) - %d capsules\n",
			i+1, eq.Type, eq.Name, eq.BonusHP, eq.BonusDmg, eq.Price)
	}

	// affiche les objets consommables
	fmt.Println("\nConsommables :")
	for i, item := range ItemsList {
		fmt.Printf("%d. %s (%s +%d) - %d capsules\n",
			i+len(Equipments)+1, item.Name, item.Effect, item.Value, item.Price)
	}
}

// acheterDansBoutique : gère l'achat d'un objet par le joueur
// utilise scanner pour récupérer la saisie
// modifie le Character passé en pointeur
func acheterDansBoutique(c *Character, scanner *bufio.Scanner) {
	for {
		// affiche la boutique à chaque tour pour voir les prix et inventaire
		afficherBoutique()
		fmt.Printf("\n💰 Capsules : %d\n", c.Argent)
		fmt.Print("Entre le numéro de l’objet à acheter (ou '0' pour quitter) : ")

		// récupère la saisie
		scanner.Scan()
		input := scanner.Text()
		choix, err := strconv.Atoi(input) // convert string -> int
		if err != nil || choix < 0 {      // check si valide
			fmt.Println("❌ Entrée invalide.")
			continue
		}

		// quitte la boutique
		if choix == 0 {
			fmt.Println("🚪 Tu quittes la boutique.")
			return
		}

		// achat d'un équipement
		if choix >= 1 && choix <= len(Equipments) {
			item := Equipments[choix-1]
			if c.Argent >= item.Price {
				c.Inventaire = append(c.Inventaire, item.Name)
				c.Argent -= item.Price
				fmt.Printf("✅ Tu as acheté : %s\n", item.Name)
			} else {
				fmt.Println("❌ Pas assez de capsules.")
			}

			// achat d'un consommable
		} else if choix <= len(Equipments)+len(ItemsList) {
			item := ItemsList[choix-len(Equipments)-1]
			if c.Argent >= item.Price {
				c.Inventaire = append(c.Inventaire, item.Name)
				c.Argent -= item.Price
				fmt.Printf("✅ Tu as acheté : %s\n", item.Name)
			} else {
				fmt.Println("❌ Pas assez de capsules.")
			}

			// numéro invalide
		} else {
			fmt.Println("❌ Numéro invalide.")
		}
	}
}
