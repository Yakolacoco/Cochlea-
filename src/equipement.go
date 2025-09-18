package main

import "fmt"

// Equipment : structure pour représenter une arme ou une armure
// Name = nom de l'équipement
// Type = "arme" ou "armure"
// BonusHP = combien de PV ça ajoute si armure
// BonusDmg = combien de dégâts ça ajoute si arme
// Price = prix en capsules
type Equipment struct {
	Name     string
	Type     string
	BonusHP  int
	BonusDmg int
	Price    int
}

// Item : objet consommable (potion, poison, etc.)
// Effect = type d'effet (soin, poison, sort)
// Value = valeur de l'effet (ex: 30 PV)
// Price = prix en capsules
type Item struct {
	Name   string
	Effect string
	Value  int
	Price  int
}

// Equipments : liste de tous les équipements disponibles dans le jeu
var Equipments = []Equipment{
	{"Couteau artisanal", "arme", 0, 15, 20},
	{"Bâton en bois", "arme", 0, 8, 10},
	{"Gilet pare-balles léger", "armure", 20, 0, 40},
	{"Barre de fer", "arme", 0, 30, 30},
	{"Gilet pare-balles lourd", "armure", 50, 0, 80},
	{"Katana affûté", "arme", 0, 45, 70},             
	{"Bouclier en acier", "armure", 40, 0, 60},       
	{"Armure de combat", "armure", 60, 0, 100},        
	{"Fusil à pompe", "arme", 0, 60, 120},            
	{"Casque renforcé", "armure", 25, 0, 45},         
	{"Épée longue", "arme", 0, 35, 50},                
	{"Gilet tactique", "armure", 35, 0, 55},           
}


// EquiperItem : équipe un item sur le personnage
// c = personnage
// equip = équipement à équiper
// index = position dans l'inventaire pour le retirer après usage
func EquiperItem(c *Character, equip *Equipment, index int) {
	switch equip.Type {
	case "arme":
		if c.ÉquipementArme != nil {
			// remettre l'ancienne arme dans l'inventaire et retirer son bonus
			c.Inventaire = append(c.Inventaire, c.ÉquipementArme.Name)
			c.DegatsBase -= c.ÉquipementArme.BonusDmg
		}
		c.ÉquipementArme = equip
		c.DegatsBase += equip.BonusDmg
		fmt.Println("✅ Arme équipée :", equip.Name)

	case "armure":
		if c.ÉquipementArmure != nil {
			// remettre l'ancienne armure dans l'inventaire et retirer son bonus
			c.Inventaire = append(c.Inventaire, c.ÉquipementArmure.Name)
			c.PVMax -= c.ÉquipementArmure.BonusHP
		}
		c.ÉquipementArmure = equip
		c.PVMax += equip.BonusHP
		if c.PVActuels > c.PVMax { // ajuster les PV si trop élevés
			c.PVActuels = c.PVMax
		}
		fmt.Println("✅ Armure équipée :", equip.Name)
	}

	// supprimer l'objet utilisé de l'inventaire
	c.Inventaire = append(c.Inventaire[:index], c.Inventaire[index+1:]...)
}

// GetEquipmentByName : récupère un pointeur vers un équipement selon son nom
// renvoie nil si non trouvé
func GetEquipmentByName(name string) *Equipment {
	for i := range Equipments {
		if Equipments[i].Name == name { // compare le nom
			return &Equipments[i]
		}
	}
	return nil
}
