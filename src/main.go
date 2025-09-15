package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)

    fmt.Println("🏚️ Bienvenue dans COCHLEA - RPG terminal")

    c1 := characterCreation(scanner)

    for {
        fmt.Println("\n--- MENU PRINCIPAL ---")
        fmt.Println("1. Afficher les informations du personnage")
        fmt.Println("2. Accéder à l’inventaire")
        fmt.Println("3. Quitter")

        fmt.Print("Ton choix : ")
        scanner.Scan()
        choix := strings.TrimSpace(scanner.Text())

        switch choix {
        case "1":
            displayInfo(c1)
        case "2":
            accessInventory(&c1, scanner)
        case "3":
            fmt.Println("👋 Fin de la session. À bientôt.")
            return
        default:
            fmt.Println("❌ Choix invalide.")
        }
    }
}
