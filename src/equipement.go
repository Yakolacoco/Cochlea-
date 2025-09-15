package main

// Les équipements
type Equipment struct {
	Name     string // nom de l'équipement
	Type     string // "arme", "armure", "accessoire"
	BonusHP  int    // bonus de points de vie
	BonusDmg int    // bonus de dégâts
	Price    int    // prix en monnaie du jeu
}

// Les items consommables
type Item struct {
	Name   string // nom de l'item
	Effect string // "soin", "attaque", "argent"
	Value  int    // valeur de l'effet (PV, dégâts ou argent)
	Price  int    // prix en boutique
}

// Liste des équipements
var Equipments = []Equipment{
	{"Couteau artisanal", "arme", 0, 15, 20}, // 0 = defance bonnus 15 = bonnus de degat 20 = prix de litem
	{"Bâton en bois", "arme", 0, 8, 10},
	{"Gilet pare-balles léger", "armure", 20, 0, 40},
	{"Barre de fer", "arme", 0, 30, 30},
	{"Gilet pare-balles lourd", "armure", 50, 0, 80},
}

// Liste des items consommables
var Items = []Item{
	{"Bandage", "soin", 20, 5},        // soigne 20 HP
	{"Mord au rat", "attaque", 10, 3}, // inflige 10 PV à l'ennemi
	{"Poisson", "soin", 15, 4},        // soigne 15 HP
	{"capsules", "argent", 1, 1},      // 1 capsule = 1 unité d'argent
}
