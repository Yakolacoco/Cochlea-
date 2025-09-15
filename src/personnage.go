package main

import (
	"fmt"
	"strings"
)

type Equipment struct {
	Head  string
	Torso string
	Feet  string
}

type Character struct {
	Name           string
	Class          string
	Level          int
	MaxHP          int
	CurrentHP      int
	Inventory      []string
	Skills         []string
	Gold           int
	MaxInventory   int
	InventoryBoost int
	Mana           int
	MaxMana        int
	Experience     int
	MaxExperience  int
	Equipment      Equipment
	Initiative     int
}

// Crée un personnage initialisé manuellement

func initCharacter(name, class string) Character {
	maxHP := 100
	initiative := 0
	gold := 100
	mana := 50
	maxMana := 50
	skills := []string{"Coup de poing"}

	switch strings.ToLower(class) {
	case "elfe":
		maxHP = 80
	case "humain":
		maxHP = 100
	case "nain":
		maxHP = 120
	}

	currentHP := maxHP / 2

	return Character{
		Name:           formatName(name),
		Class:          class,
		Level:          1,
		MaxHP:          maxHP,
		CurrentHP:      currentHP,
		Inventory:      []string{"Potion", "Potion", "Potion"},
		Skills:         skills,
		Gold:           gold,
		MaxInventory:   10,
		InventoryBoost: 0,
		Mana:           mana,
		MaxMana:        maxMana,
		Experience:     0,
		MaxExperience:  100,
		Initiative:     initiative,
		Equipment:      Equipment{},
	}
}

func formatName(name string) string {
	if len(name) == 0 {
		return "Inconnu"
	}
	name = strings.ToLower(name)
	return strings.ToUpper(string(name[0])) + name[1:]
}

// Affiche les informations du personnage

func displayInfo(c Character) {
	fmt.Println("---------- FICHE PERSONNAGE ----------")
	fmt.Printf("Nom : %s\n", c.Name)
	fmt.Printf("Classe : %s\n", c.Class)
	fmt.Printf("Niveau : %d\n", c.Level)
	fmt.Printf("PV : %d / %d\n", c.CurrentHP, c.MaxHP)
	fmt.Printf("Mana : %d / %d\n", c.Mana, c.MaxMana)
	fmt.Printf("XP : %d / %d\n", c.Experience, c.MaxExperience)
	fmt.Printf("Or : %d\n", c.Gold)
	fmt.Println("Compétences :", c.Skills)
	fmt.Println("Équipement :", c.Equipment)
	fmt.Println("--------------------------------------")
}

// Affiche l’inventaire

func accessInventory(c Character) {
	fmt.Println("------ INVENTAIRE ------")
	if len(c.Inventory) == 0 {
		fmt.Println("Inventaire vide.")
	} else {
		for i, item := range c.Inventory {
			fmt.Printf("%d. %s\n", i+1, item)
		}
	}
	fmt.Println("------------------------")
}
