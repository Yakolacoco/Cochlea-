package main

import (
	"bufio"
	"fmt"
	"math/rand" // <-- Ajoute ceci
	"os"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("🏚️ Bienvenue dans COCHLEA - RPG terminal")
	fmt.Println(` _______  _______  _______           _______  _        _______ 
(  ____ \(  ___  )(  ____ \|\     /|(  ____ \( \      (  ___  )
| (    \/| (   ) || (    \/| )   ( || (    \/| (      | (   ) |
| |      | |   | || |      | (___) || (__    | |      | (___) |
| |      | |   | || |      |  ___  ||  __)   | |      |  ___  |
| |      | |   | || |      | (   ) || (      | |      | (   ) |
| (____/\| (___) || (____/\| )   ( || (____/\| (____/\| )   ( |
(_______/(_______)(_______/|/     \|(_______/(_______/|/     \|
                                                               `)

	// Création du personnage
	joueur := characterCreation(scanner)

	// Génération de la tour avec 10 étages (modifiable)
	tower := genererTour(10)
	currentFloor := 1

	for {
		fmt.Println("\n--- MENU PRINCIPAL ---")
		fmt.Println("1. Afficher les informations du personnage")
		fmt.Println("2. Accéder à l’inventaire")
		fmt.Println("3. Aller à la boutique")
		fmt.Println("4. Manger (si nourriture)")
		fmt.Println("5. Monter à l'étage suivant")
		fmt.Println("6. Quitter")

		fmt.Print("Ton choix : ")
		scanner.Scan()
		choix := strings.TrimSpace(scanner.Text())

		switch choix {
		case "1":
			displayInfo(joueur)
		case "2":
			accessInventory(&joueur, scanner)
		case "3":
			acheterDansBoutique(&joueur, scanner)
		case "4":
			manger(&joueur)
		case "5":
			if currentFloor > tower.MaxFloor {
				fmt.Println("🏁 Vous avez atteint le sommet de la tour !")
				continue
			}

			fmt.Printf("\n🔼 Tu montes à l'étage %d\n", currentFloor)

			// Afficher l'étage et gérer coffres
			tower.afficherEtage(currentFloor, &joueur)

			// Récupérer le monstre ou boss pour le combat
			monstre := tower.getMonsterForCombat(currentFloor)
			if monstre != nil {
				combat(&joueur, *monstre) // déférencer le pointeur

				if joueur.PVActuels <= 0 {
					fmt.Println("💀 Tu es tombé...")

					if rand.Intn(100) < 25 { // 25% de chance
						joueur.PVActuels = joueur.PVMax / 4
						fmt.Printf("✨ Une force mystérieuse te réanime avec %d PV !\n", joueur.PVActuels)
					} else {
						fmt.Println("💀 Game Over !")
						return
					}
				} else {
					fmt.Println("💰 Tu gagnes 5 capsules pour avoir vaincu l'ennemi !")
					joueur.Argent += 5
				}
			}

			currentFloor++

		case "6":
			fmt.Println("👋 Fin de la session. À bientôt.")
			return
		default:
			fmt.Println("❌ Choix invalide.")
		}

	}
}
