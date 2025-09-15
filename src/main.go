package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Bienvenue dans Prison RPG (Prototype)")

	fmt.Print("Entrez votre nom : ")
	scanner.Scan()
	name := scanner.Text()

	class := ""
	for {
		fmt.Print("Choisissez votre classe (Humain, Elfe, Nain) : ")
		scanner.Scan()
		class = scanner.Text()
		classLower := strings.ToLower(class)
		if classLower == "humain" || classLower == "elfe" || classLower == "nain" {
			break
		}
		fmt.Println("Classe invalide.")
	}

	player := initCharacter(name, class)

	for {
		fmt.Println("\n--- MENU PRINCIPAL ---")
		fmt.Println("1. Afficher les informations du personnage")
		fmt.Println("2. Accéder à l’inventaire")
		fmt.Println("3. Quitter")

		fmt.Print("Votre choix : ")
		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			displayInfo(player)
		case "2":
			accessInventory(player)
		case "3":
			fmt.Println("À bientôt !")
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}
}