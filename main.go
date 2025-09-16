package main

import (
	"fmt"
	"strings"
	"unicode"
)

type Character struct {
	Nom        string
	Classe     string
	Niveau     int
	PVMax      int
	PVActuels  int
	Inventaire []string
	Argent     int
	DegatsBase int
	Initiative int
	Passif     string
}

func characterCreation(scanner Scanner) Character {
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
		Inventaire: []string{"Potion de soin", "Potion de soin", "Potion de soin"},
		DegatsBase: 10,
		Initiative: 0,
		Argent:     100,
	}

	switch classe {
	case "Meurtrier":
		c.PVMax = 120
		c.Passif = "+20 PV Max, mais +20% fatigue par étage."
		c.DegatsBase += 5
	case "Voleur":
		c.PVMax = 80
		c.Passif = "+5 initiative et +100 or, mais -20 PV Max."
		c.Initiative += 5
		c.Argent += 100
	case "Hacker":
		c.PVMax = 100
		c.Passif = "Sort passif : Pirater (monstre confus 1/2 chance de rater). Moins d’équipement."
	case "Psychopathe":
		c.PVMax = 100
		c.Passif = "+10 dégâts de base, mais la faim/fatigue augmente 2× plus vite. 50% de chance de faire x2 dégâts."
		c.DegatsBase += 10
	}

	c.PVActuels = c.PVMax / 2
	return c
}

func displayInfo(c Character) {
	fmt.Println("--- Infos Personnage ---")
	fmt.Println("Nom :", c.Nom)
	fmt.Println("Classe :", c.Classe)
	fmt.Println("Niveau :", c.Niveau)
	fmt.Printf("PV : %d / %d\n", c.PVActuels, c.PVMax)
	fmt.Println("Argent :", c.Argent)
	fmt.Println("Dégâts de base :", c.DegatsBase)
	fmt.Println("Inventaire :", c.Inventaire)
	fmt.Println("Passif :", c.Passif)
}

func accessInventory(c *Character, scanner Scanner) {
	fmt.Println("--- Inventaire ---")
	if len(c.Inventaire) == 0 {
		fmt.Println("Inventaire vide.")
		return
	}
	for i, item := range c.Inventaire {
		fmt.Printf("%d. %s\n", i+1, item)
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
