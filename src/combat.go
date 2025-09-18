package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Monster struct
type Monster struct {
	Name       string
	PVMax      int
	PVActuels  int
	DegatsBase int
	Initiative int
	IsUltimate bool
}

// Combat principal
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

		// Récompense normale
		rewardArgent := 5
		rewardXP := 5 + rand.Intn(6) // 5 à 10 XP

		// Si boss ou ultimate, bonus argent et XP
		if monstre.DegatsBase >= 25 || monstre.IsUltimate { 
			rewardArgent = 20 + rand.Intn(11) // 20 à 30 capsules
			rewardXP = 20 + rand.Intn(11)     // 20 à 30 XP
			fmt.Println("🎉 Boss vaincu ! Récompenses bonus !")
		}

		joueur.Argent += rewardArgent
		fmt.Printf("💰 Tu gagnes %d capsules.\n", rewardArgent)

		gagnerXP(joueur, rewardXP)
		fmt.Printf("⭐ Tu gagnes %d XP.\n", rewardXP)
	}

	// Appliquer la perte de faim/fatigue après combat
	apresCombat(joueur)
}


// Tour du joueur
func playerTurn(joueur *Character, monstre *Monster) {
	// Appliquer brûlure si actif
	if joueur.BurnTurns > 0 {
		joueur.PVActuels -= joueur.BurnDmg
		if joueur.PVActuels < 0 {
			joueur.PVActuels = 0
		}
		fmt.Printf("🔥 Brûlure : tu subis %d dégâts (%d PV restants)\n", joueur.BurnDmg, joueur.PVActuels)
		joueur.BurnTurns--
	}

	fmt.Printf("\n🎯 Ton tour ! (%d PV)\n", joueur.PVActuels)
	fmt.Println("1. Attaquer (Coup de poing)")
	for i, skill := range joueur.Skills {
		fmt.Printf("%d. %s\n", i+2, skill)
	}
	fmt.Println("0. Utiliser un objet")
	fmt.Print("Choix : ")

	var choix string
	fmt.Scanln(&choix)

	switch choix {
	case "0":
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

	case "1":
		// Attaque de base
		degats := rand.Intn(joueur.DegatsBase/2) + joueur.DegatsBase/2
		if joueur.Classe == "Psychopathe" && rand.Intn(2) == 0 {
			degats *= 2
			fmt.Println("😈 Passif 'Psychopathe' activé ! Coup critique !")
		}
		monstre.PVActuels -= degats
		if monstre.PVActuels < 0 {
			monstre.PVActuels = 0
		}
		fmt.Printf("🗡️ Tu infliges %d dégâts à %s. (%d PV restants)\n", degats, monstre.Name, monstre.PVActuels)
		if joueur.Classe == "Hacker" && rand.Intn(2) == 0 {
			fmt.Println("💻 Passif 'Hacker' activé ! Le monstre est confus et rate son prochain tour !")
			monstre.Initiative = -1
		}

	default:
		// Compétences
		idx := parseInt(choix) - 2
		if idx < 0 || idx >= len(joueur.Skills) {
			fmt.Println("❌ Choix invalide, tu rates ton tour !")
			return
		}
		skill := joueur.Skills[idx]
		degats := 0

		switch skill {
		case "Boule de Feu":
			degats = 20 + rand.Intn(11)
		case "Éclair":
			degats = 15 + rand.Intn(11)
			if rand.Intn(3) == 0 {
				fmt.Println("⚡ Éclair paralyse le monstre ! Il rate son prochain tour.")
				monstre.Initiative = -1
			}
		case "Coup de poing":
			degats = rand.Intn(joueur.DegatsBase/2) + joueur.DegatsBase/2
		default:
			degats = rand.Intn(joueur.DegatsBase/2) + joueur.DegatsBase/2
		}

		if joueur.Classe == "Psychopathe" && rand.Intn(2) == 0 {
			degats *= 2
			fmt.Println("😈 Passif 'Psychopathe' activé ! Coup critique !")
		}

		monstre.PVActuels -= degats
		if monstre.PVActuels < 0 {
			monstre.PVActuels = 0
		}
		fmt.Printf("🔥 Tu utilises '%s' et infliges %d dégâts à %s. (%d PV restants)\n",
			skill, degats, monstre.Name, monstre.PVActuels)
	}

	// Poison
	if joueur.PoisonNextAttack {
		poisonDmg := 15
		monstre.PVActuels -= poisonDmg
		if monstre.PVActuels < 0 {
			monstre.PVActuels = 0
		}
		fmt.Printf("☠️ Le poison inflige %d PV de dégâts à %s ! (%d PV restants)\n",
			poisonDmg, monstre.Name, monstre.PVActuels)
		joueur.PoisonNextAttack = false
	}
}

// Tour du monstre
func monsterTurn(joueur *Character, monstre *Monster) {
	if monstre.Initiative < 0 {
		fmt.Printf("😵 %s est confus et rate son tour !\n", monstre.Name)
		monstre.Initiative = 10
		return
	}

	degats := rand.Intn(monstre.DegatsBase/2) + monstre.DegatsBase/2

	// Boss ultime Nael
	if monstre.IsUltimate {
		extra := rand.Intn(10) + 5
		degats += extra
		fmt.Println("🌶️ Nael lance des épices ! Dégâts supplémentaires :", extra)
		joueur.BurnTurns = 3
		joueur.BurnDmg = 5
		fmt.Println("🔥 Le joueur subit brûlure !")
	}

	joueur.PVActuels -= degats
	if joueur.PVActuels < 0 {
		joueur.PVActuels = 0
	}
	fmt.Printf("👹 %s attaque et inflige %d dégâts. (%d PV restants)\n", monstre.Name, degats, joueur.PVActuels)
}
