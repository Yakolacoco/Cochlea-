package main

import "fmt"

// Structure de l'équipement
type Equipment struct {
	Name     string // nom de l'équipement
	Type     string // "arme" ou "armure"
	BonusHP  int    // bonus de points de vie
	BonusDmg int    // bonus de dégâts
	Price    int    // prix en capsules
}
	
type Item struct {
	Name   string
	Effect string // "soin", "poison", etc.
	Value  int
	Price  int
}

// Liste des équipements disponibles
var Equipments = []Equipment{
	{"Couteau artisanal", "arme", 0, 15, 20},
	{"Bâton en bois", "arme", 0, 8, 10},
	{"Gilet pare-balles léger", "armure", 20, 0, 40},
	{"Barre de fer", "arme", 0, 30, 30},
	{"Gilet pare-balles lourd", "armure", 50, 0, 80},
}

// Équipe un équipement pour le personnage
func EquiperItem(c *Character, equip *Equipment, index int) {
	switch equip.Type {
	case "arme":
		if c.ÉquipementArme != nil {
			c.Inventaire = append(c.Inventaire, c.ÉquipementArme.Name)
			c.DegatsBase -= c.ÉquipementArme.BonusDmg
		}
		c.ÉquipementArme = equip
		c.DegatsBase += equip.BonusDmg
		fmt.Println("✅ Arme équipée :", equip.Name)
	case "armure":
		if c.ÉquipementArmure != nil {
			c.Inventaire = append(c.Inventaire, c.ÉquipementArmure.Name)
			c.PVMax -= c.ÉquipementArmure.BonusHP
		}
		c.ÉquipementArmure = equip
		c.PVMax += equip.BonusHP
		if c.PVActuels > c.PVMax {
			c.PVActuels = c.PVMax
		}
		fmt.Println("✅ Armure équipée :", equip.Name)
	}

	// Supprimer l'item de l'inventaire après équipement
	c.Inventaire = append(c.Inventaire[:index], c.Inventaire[index+1:]...)
}

// Récupère un équipement par son nom
func GetEquipmentByName(name string) *Equipment {
	for i := range Equipments {
		if Equipments[i].Name == name {
			return &Equipments[i]
		}
	}
	return nil
}
