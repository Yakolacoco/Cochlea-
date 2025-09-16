package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Monster struct {
	Name       string
	PVMax      int
	PVActuels  int
	DegatsBase int
	Initiative int
}

// Fonction pour lancer un combat
func combat(joueur *Character, monstre Monster) {
	rand.Seed(time.Now().UnixNano())
	fmt.Printf("\n⚔️ Un %s apparaît !\n", monstre.Name)

	for joueur.PVActuels > 0 && monstre.PVActuels > 0 {
		if joueur.Initiative >= monstre.Initiative {
			playerTurn(joueur, &monstre)
			if monstre.PVActuels <= 0 {
				break
			}
			monsterTurn(joueur, &monstre)
		} else {
			monsterTurn(joueur, &monstre)
			if joueur.PVActuels <= 0 {
				break
			}
			playerTurn(joueur, &monstre)
		}
	}

	if monstre.PVActuels <= 0 {
		fmt.Printf("🏆 Tu as vaincu %s !\n", monstre.Name)
		joueur.Argent += 5
		fmt.Println("💰 Tu gagnes 5 capsules.")
		gagnerXP(joueur, 5+rand.Intn(6)) // 5 à 10 XP
	}



	// Appliquer la perte de faim/fatigue après combat
	apresCombat(joueur)
}

// Tour du joueur
func playerTurn(joueur *Character, monstre *Monster) {
	fmt.Printf("\n🎯 Ton tour ! (%d PV)\n", joueur.PVActuels)
	fmt.Printf("1. Attaquer\n2. Utiliser un objet\nChoix : ")
	var choix string
	fmt.Scanln(&choix)

	switch choix {
	case "1":
		degats := rand.Intn(joueur.DegatsBase/2) + joueur.DegatsBase/2
		// Bonus Psychopathe : 50% chance x2 dégâts
		if joueur.Classe == "Psychopathe" && rand.Intn(2) == 0 {
			degats *= 2
			fmt.Println("💥 Coup critique Psychopathe !")
		}
		monstre.PVActuels -= degats
		if monstre.PVActuels < 0 {
			monstre.PVActuels = 0
		}
		fmt.Printf("🗡️ Tu infliges %d dégâts à %s. (%d PV restants)\n", degats, monstre.Name, monstre.PVActuels)
	case "2":
		if len(joueur.Inventaire) == 0 {
			fmt.Println("❌ Inventaire vide !")
			return
		}
		fmt.Println("--- Inventaire ---")
		for i, item := range joueur.Inventaire {
			fmt.Printf("%d. %s\n", i+1, item)
		}
		fmt.Print("Choisis un objet : ")
		var idx int
		fmt.Scanln(&idx)
		if idx < 1 || idx > len(joueur.Inventaire) {
			fmt.Println("❌ Choix invalide.")
			return
		}
		useConsumable(joueur, joueur.Inventaire[idx-1], idx-1)
	default:
		fmt.Println("❌ Choix invalide, tu rates ton tour !")
	}
}

// Tour du monstre
func monsterTurn(joueur *Character, monstre *Monster) {
	degats := rand.Intn(monstre.DegatsBase/2) + monstre.DegatsBase/2
	joueur.PVActuels -= degats
	if joueur.PVActuels < 0 {
		joueur.PVActuels = 0
	}
	fmt.Printf("👹 %s attaque et inflige %d dégâts. (%d PV restants)\n", monstre.Name, degats, joueur.PVActuels)
}

// Générer un monstre aléatoire
func genererMonstreTour() Monster {
    mobs := []Monster{
        {Name: "Rat géant", PVMax: 40, PVActuels: 40, DegatsBase: 7, Initiative: 8},
        {Name: "Zombie affamé", PVMax: 60, PVActuels: 60, DegatsBase: 10, Initiative: 9},
        {Name: "Garde armé", PVMax: 100, PVActuels: 100, DegatsBase: 13, Initiative: 10},
        {Name: "Prisonnier fou", PVMax: 60, PVActuels: 35, DegatsBase: 20, Initiative: 12},
        {Name: "Chien dressé", PVMax: 40, PVActuels: 40, DegatsBase: 10, Initiative: 12},
    }
    rand.Seed(time.Now().UnixNano())
    return mobs[rand.Intn(len(mobs))]
}


