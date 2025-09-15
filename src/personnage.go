package main

import (
    "bufio"
    "fmt"
    "strings"
    "unicode"
)

type Character struct {
    Nom           string
    Classe        string
    Niveau        int
    PVMax         int
    PVActuels     int
    Inventaire    []string
    Argent        int
    DegatsBase    int
    Initiative    int
    Passif        string
    Faim          int
    Fatigue       int
}

func characterCreation(scanner *bufio.Scanner) Character {
    var nom, classe string
    fmt.Print("Quel est ton nom ? ")
    scanner.Scan()
    nom = formatNom(scanner.Text())

    for {
        fmt.Print("Choisis ta peine (Meurtrier, Voleur, Hacker, Psychopathe) : ")
        scanner.Scan()
        classe = strings.Title(strings.ToLower(scanner.Text()))
        if classe == "Meurtrier" || classe == "Voleur" || classe == "Hacker" || classe == "Psychopathe" {
            break
        }
        fmt.Println("❌ Classe invalide.")
    }

    return initCharacter(nom, classe)
}

func initCharacter(nom string, classe string) Character {
    c := Character{
        Nom:        nom,
        Classe:     classe,
        Niveau:     1,
        Inventaire: []string{"Potion de soin", "Pain sec", "Potion de soin"},
        DegatsBase: 10,
        Initiative: 10,
        Argent:     100,
        Faim:       20,
        Fatigue:    20,
    }

    switch classe {
    case "Meurtrier":
        c.PVMax = 120
        c.Passif = "+20 PV Max, mais +10% fatigue par étage."
        c.DegatsBase += 5
    case "Voleur":
        c.PVMax = 100
        c.Passif = "+5 initiative et +100 or."
        c.Initiative += 5
        c.Argent += 100
    case "Hacker":
        c.PVMax = 80
        c.Passif = "Sort passif : Pirater (monstre confus 1/2 chance de rater)."
    case "Psychopathe":
        c.PVMax = 100
        c.Passif = "+10 dégâts de base, mais faim/fatigue augmentent 2× plus vite. 50% de chance de faire x2 dégâts."
        c.DegatsBase += 10
    }

    c.PVActuels = c.PVMax / 2
    return c
}

func apresCombat(c *Character) {
    perte := 3
    if c.Classe == "Psychopathe" {
        perte *= 2
    }

    c.Faim -= perte
    c.Fatigue -= perte

    if c.Faim < 10 || c.Fatigue < 10 {
        if c.Initiative > 0 {
            c.Initiative--
            fmt.Println("⚠️ Tu te sens faible... ton initiative diminue.")
        }
    }

    if c.Faim < 0 {
        c.Faim = 0
    }
    if c.Fatigue < 0 {
        c.Fatigue = 0
    }
}

func manger(c *Character) {
    for i, item := range c.Inventaire {
        if item == "Pain sec" {
            fmt.Println("🍞 Tu manges un pain sec. +5 faim.")
            c.Faim += 5
            if c.Faim > 20 {
                c.Faim = 20
            }
            c.Inventaire = append(c.Inventaire[:i], c.Inventaire[i+1:]...)
            return
        }
    }
    fmt.Println("❌ Tu n’as rien à manger.")
}

func displayInfo(c Character) {
    fmt.Println("--- Infos Personnage ---")
    fmt.Println("Nom :", c.Nom)
    fmt.Println("Classe :", c.Classe)
    fmt.Println("Niveau :", c.Niveau)
    fmt.Printf("PV : %d / %d\n", c.PVActuels, c.PVMax)
    fmt.Println("Argent :", c.Argent)
    fmt.Println("Dégâts de base :", c.DegatsBase)
    fmt.Println("Initiative :", c.Initiative)
    fmt.Println("Faim :", c.Faim, "/ 20")
    fmt.Println("Fatigue :", c.Fatigue, "/ 20")
    fmt.Println("Inventaire :", c.Inventaire)
    fmt.Println("Passif :", c.Passif)
}

func accessInventory(c *Character, scanner *bufio.Scanner) {
    fmt.Println("--- Inventaire ---")
    if len(c.Inventaire) == 0 {
        fmt.Println("Inventaire vide.")
        return
    }

    for i, item := range c.Inventaire {
        fmt.Printf("%d. %s\n", i+1, item)
    }

    fmt.Print("Veux-tu manger un aliment ? (oui/non) : ")
    scanner.Scan()
    choix := strings.ToLower(scanner.Text())
    if choix == "oui" {
        manger(c)
    }
}

func formatNom(input string) string {
    input = strings.ToLower(input)
    runes := []rune(input)
    for i, r := range runes {
        if unicode.IsLetter(r) {
            runes[i] = unicode.ToUpper(r)
            break
        }
    }
    return string(runes)
}
