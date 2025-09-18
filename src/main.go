package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// main : point d'entrée du jeu
// initialise le jeu, crée le personnage, génère la tour et gère le menu principal
func main() {
	rand.Seed(time.Now().UnixNano()) // initialise le générateur de hasard
	scanner := bufio.NewScanner(os.Stdin)

	// affichage du titre du jeu
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

	// création du personnage via la fonction characterCreation
	joueur := characterCreation(scanner)

	// génération de la tour (ici 20 étages)
	tower := genererTour(20)
	currentFloor := 1

	// boucle principale du menu
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
		choix := strings.TrimSpace(scanner.Text()) // nettoie l'entrée utilisateur

		switch choix {
		case "1":
			displayInfo(joueur) // affiche toutes les infos du joueur
		case "2":
			accessInventory(&joueur, scanner) // gestion de l'inventaire
		case "3":
			acheterDansBoutique(&joueur, scanner) // fonction boutique
		case "4":
			manger(&joueur) // consomme un item nourriture si disponible
		case "5":
			if currentFloor > tower.MaxFloor {
				fmt.Println("🏁 Vous avez atteint le sommet de la tour !")
				continue
			}

			fmt.Printf("\n🔼 Tu montes à l'étage %d\n", currentFloor)

			// affiche l'étage et gère les coffres
			tower.afficherEtage(currentFloor, &joueur)

			// récupère le monstre ou boss à combattre
			monstre := tower.getMonsterForCombat(currentFloor)
			if monstre != nil {
				combat(&joueur, *monstre) // lance le combat, déférencer le pointeur

				// si le joueur meurt
				if joueur.PVActuels <= 0 {
					fmt.Println("💀 Tu es tombé...")

					if rand.Intn(100) < 25 { // 25% de chance de résurrection
						joueur.PVActuels = joueur.PVMax / 4
						fmt.Printf("✨ Une force mystérieuse te réanime avec %d PV !\n", joueur.PVActuels)
					} else {
						fmt.Println("💀 Game Over !")
						return
					}
				} else {
					// récompense pour avoir vaincu
					fmt.Println("💰 Tu gagnes 5 capsules pour avoir vaincu l'ennemi !")
					joueur.Argent += 5
				}
			}

			currentFloor++ // passe à l'étage suivant

		case "6":
			fmt.Println("👋 Fin de la session. À bientôt.")
			return
		default:
			fmt.Println("❌ Choix invalide.") // entrée invalide dans le menu
		}
	}
}
