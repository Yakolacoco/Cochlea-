package main

import "fmt"

// Inventaire
type Inventory struct {
	Equipements []Equipment
	Items       []Item
}

// Ajouter un équipement
func (inv *Inventory) AddEquipment(eq Equipment) {
	inv.Equipements = append(inv.Equipements, eq)
}

// Ajouter un item
func (inv *Inventory) AddItem(it Item) {
	inv.Items = append(inv.Items, it)
}

// Afficher l'inventaire
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

// Vérifier si un équipement est dans l’inventaire
func (inv *Inventory) HasEquipment(name string) bool {
	for _, eq := range inv.Equipements {
		if eq.Name == name {
			return true
		}
	}
	return false
}

// Vérifier si un item est dans l’inventaire
func (inv *Inventory) HasItem(name string) bool {
	for _, it := range inv.Items {
		if it.Name == name {
			return true
		}
	}
	return false
}
