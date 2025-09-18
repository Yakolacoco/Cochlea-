package main

import "fmt"

// Inventory : structure qui regroupe tout ce que le joueur possède
// Equipements = armes/armures, Items = consommables ou objets divers
type Inventory struct {
	Equipements []Equipment // liste des équipements possédés
	Items       []Item      // liste des objets consommables
}

// AddEquipment : ajoute un équipement dans l'inventaire
// eq : équipement à ajouter
func (inv *Inventory) AddEquipment(eq Equipment) {
	inv.Equipements = append(inv.Equipements, eq) // utilise append pour ajouter à la slice
}

// AddItem : ajoute un objet consommable dans l'inventaire
// it : item à ajouter
func (inv *Inventory) AddItem(it Item) {
	inv.Items = append(inv.Items, it)
}

// Show : affiche tout l'inventaire du joueur
// parcours les slices Equipements et Items et print chaque objet avec ses stats
func (inv *Inventory) Show() {
	fmt.Println("\n=== Équipements ===")
	for i, eq := range inv.Equipements {
		fmt.Printf("%d. %s (%s) HP:%d DMG:%d Prix:%d capsules\n", i+1, eq.Name, eq.Type, eq.BonusHP, eq.BonusDmg, eq.Price)
	}

	fmt.Println("\n=== Items ===")
	for i, it := range inv.Items {
		fmt.Printf("%d. %s (%s) Valeur:%d Prix:%d capsules\n", i+1, it.Name, it.Effect, it.Value, it.Price)
	}
}

// HasEquipment : vérifie si un équipement est déjà possédé
// renvoie true si oui, false sinon
func (inv *Inventory) HasEquipment(name string) bool {
	for _, eq := range inv.Equipements {
		if eq.Name == name { // compare le nom
			return true
		}
	}
	return false
}

// HasItem : vérifie si un item est déjà possédé
// renvoie true si l'objet est présent dans l'inventaire
func (inv *Inventory) HasItem(name string) bool {
	for _, it := range inv.Items {
		if it.Name == name { // compare le nom
			return true
		}
	}
	return false
}
