package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Floor struct {
	Number   int
	Monsters []Monster
	Boss     *Monster
	Chest    *Chest
}

type Chest struct {
	RewardType string // "argent" ou "item"
	Amount     int
}

type Tower struct {
	Floors    []Floor
	MaxFloor  int
}

// Générer la tour avec mobs, sous-boss, boss et coffres
func genererTour(maxFloor int) Tower {
	rand.Seed(time.Now().UnixNano())
	tower := Tower{MaxFloor: maxFloor}

	// Sous-boss
	sousBoss := []Monster{
		{Name: "Garde vétéran", PVMax: 70, PVActuels: 70, DegatsBase: 12, Initiative: 12},
		{Name: "Prisonnier fou", PVMax: 80, PVActuels: 80, DegatsBase: 14, Initiative: 11},
		{Name: "Chimère carcérale", PVMax: 90, PVActuels: 90, DegatsBase: 16, Initiative: 10},
		{Name: "Évadé mutant", PVMax: 85, PVActuels: 85, DegatsBase: 15, Initiative: 13},
	}

	for i := 1; i <= maxFloor; i++ {
		floor := Floor{Number: i}

		// Boss final à l'étage 20
		if i == 20 {
			boss := Monster{Name: "Directeur suprême", PVMax: 200, PVActuels: 200, DegatsBase: 25, Initiative: 15}
			floor.Boss = &boss
		} else if i == 5 || i == 10 || i == 15 || i == 18 { // Sous-boss aux étages clés
			idx := (i / 5) - 1
			sb := sousBoss[idx]
			floor.Boss = &sb
		} else if i%3 == 0 { // Étages spéciaux avec coffre
			chest := Chest{}
			if rand.Intn(2) == 0 {
				chest.RewardType = "argent"
				chest.Amount = 20 + rand.Intn(31) // 20 à 50 capsules
			} else {
				chest.RewardType = "item"
				chest.Amount = rand.Intn(len(ItemsList))
			}
			floor.Chest = &chest
		} else { // Monstres normaux
			nbMonstres := rand.Intn(3) + 1
			for j := 0; j < nbMonstres; j++ {
				m := genererMonstreTour()
				floor.Monsters = append(floor.Monsters, m)

			}
		}

		tower.Floors = append(tower.Floors, floor)
	}

	return tower
}

// Afficher l'étage et ses occupants
func (t *Tower) afficherEtage(num int, joueur *Character) {
	if num < 1 || num > t.MaxFloor {
		fmt.Println("❌ Étape invalide")
		return
	}
	floor := t.Floors[num-1]

	if floor.Boss != nil {
		if num == 20 {
			fmt.Printf("👑 Boss final : %s (%d PV, %d dégâts)\n", floor.Boss.Name, floor.Boss.PVActuels, floor.Boss.DegatsBase)
		} else {
			fmt.Printf("👹 Sous-boss : %s (%d PV, %d dégâts)\n", floor.Boss.Name, floor.Boss.PVActuels, floor.Boss.DegatsBase)
		}
	} else if floor.Chest != nil {
		fmt.Println("🎁 Un coffre est ici !")
		if floor.Chest.RewardType == "argent" {
			fmt.Printf("💰 Tu trouves %d capsules !\n", floor.Chest.Amount)
			joueur.Argent += floor.Chest.Amount
		} else {
			item := ItemsList[floor.Chest.Amount]
			fmt.Printf("🎁 Tu trouves un item : %s\n", item.Name)
			joueur.Inventaire = append(joueur.Inventaire, item.Name)
		}
	} else {
		fmt.Println("👾 Monstres présents :")
		for _, m := range floor.Monsters {
			fmt.Printf("- %s (%d PV, %d dégâts)\n", m.Name, m.PVActuels, m.DegatsBase)
		}
	}
}

// Récupérer un monstre pour le combat
func (t *Tower) getMonsterForCombat(num int) *Monster {
	if num < 1 || num > t.MaxFloor {
		return nil
	}
	floor := t.Floors[num-1]
	if floor.Boss != nil {
		return floor.Boss
	} else if len(floor.Monsters) > 0 {
		rand.Seed(time.Now().UnixNano())
		idx := rand.Intn(len(floor.Monsters))
		return &floor.Monsters[idx]
	}
	return nil
}

