package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Floor : représente un étage de la tour
// Number = numéro de l'étage
// Monsters = liste des monstres normaux sur cet étage
// Boss = pointeur vers un boss ou sous-boss (nil si pas de boss)
// Chest = pointeur vers un coffre (nil si pas de coffre)
type Floor struct {
	Number   int
	Monsters []Monster
	Boss     *Monster
	Chest    *Chest
}

// Chest : représente un coffre avec une récompense
// RewardType = "argent" ou "item"
// Amount = soit la quantité d'argent, soit index dans ItemsList
type Chest struct {
	RewardType string
	Amount     int
}

// Tower : structure principale de la tour
// Floors = liste des étages
// MaxFloor = nombre maximum d'étages
type Tower struct {
	Floors   []Floor
	MaxFloor int
}

// genererTour : crée la tour avec les monstres, sous-boss, boss et coffres
// maxFloor = nombre d'étages à générer
func genererTour(maxFloor int) Tower {
	rand.Seed(time.Now().UnixNano()) // initialisation du random
	tower := Tower{MaxFloor: maxFloor}

	// liste des sous-boss
	sousBoss := []Monster{
		{Name: "Erwan le viking", PVMax: 230, PVActuels: 230, DegatsBase: 20, Initiative: 12},
		{Name: "Sathyan le pakpak", PVMax: 150, PVActuels: 150, DegatsBase: 27, Initiative: 15},
		{Name: "Victor le mangeur de concombre", PVMax: 120, PVActuels: 120, DegatsBase: 17, Initiative: 14},
		{Name: "Mac", PVMax: 300, PVActuels: 300, DegatsBase: 10, Initiative: 13},
		{Name: "Hossam le karateka", PVMax: 160, PVActuels: 160, DegatsBase: 32, Initiative: 15},
		{Name: "Lucas le radin", PVMax: 50, PVActuels: 50, DegatsBase: 10, Initiative: 15},

	}

	// boucle pour chaque étage
	for i := 1; i <= maxFloor; i++ {
		floor := Floor{Number: i}

		switch {
		case i == 20: // Boss
			boss := Monster{Name: "Gabriel le mec gentil", PVMax: 280, PVActuels: 280, DegatsBase: 30, Initiative: 20}
			floor.Boss = &boss
		
		case i == 25: // boss final ultime
			ultimeBoss := Monster{
				Name:       "Nael l'épicier: Il lance des épices",
				PVMax:      300,
				PVActuels:  300,
				DegatsBase: 25,
				Initiative: 20,
				IsUltimate: true,
			}
			floor.Boss = &ultimeBoss


		case i == 5 || i == 10 || i == 15 || i == 18: // sous-boss
			sb := sousBoss[rand.Intn(len(sousBoss))]
			floor.Boss = &sb

		case i%3 == 0: // étage avec coffre
			chest := Chest{}
			if rand.Intn(2) == 0 {
				chest.RewardType = "argent"
				chest.Amount = 20 + rand.Intn(31) // 20 à 50 capsules
			} else {
				chest.RewardType = "item"
				chest.Amount = rand.Intn(len(ItemsList))
			}
			floor.Chest = &chest

		default: // monstres normaux
			nbMonstres := rand.Intn(3) + 1 // 1 à 3 monstres
			for j := 0; j < nbMonstres; j++ {
				m := genererMonstreTour() // génération aléatoire
				floor.Monsters = append(floor.Monsters, m)
			}
		}

		tower.Floors = append(tower.Floors, floor)
	}

	return tower
}


// afficherEtage : affiche le contenu d'un étage et gère les coffres
// num = numéro de l'étage
// joueur = pointeur vers le joueur pour ajouter récompenses
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
	} else if floor.Chest != nil { // si il y a un coffre
		fmt.Println("🎁 Un coffre est ici !")
		if floor.Chest.RewardType == "argent" {
			fmt.Printf("💰 Tu trouves %d capsules !\n", floor.Chest.Amount)
			joueur.Argent += floor.Chest.Amount
		} else {
			item := ItemsList[floor.Chest.Amount] // récupère l'item via l'index
			fmt.Printf("🎁 Tu trouves un item : %s\n", item.Name)
			joueur.Inventaire = append(joueur.Inventaire, item.Name)
		}
	} else { // monstres normaux
		fmt.Println("👾 Monstres présents :")
		for _, m := range floor.Monsters {
			fmt.Printf("- %s (%d PV, %d dégâts)\n", m.Name, m.PVActuels, m.DegatsBase)
		}
	}
}

// getMonsterForCombat : récupère un monstre à combattre pour cet étage
// renvoie nil si pas de monstre
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
func genererMonstreTour() Monster {
    mobs := []Monster{
        {Name: "Rat géant", PVMax: 40, PVActuels: 40, DegatsBase: 7, Initiative: 8},
        {Name: "Zombie affamé", PVMax: 60, PVActuels: 60, DegatsBase: 10, Initiative: 9},
        {Name: "Garde armé", PVMax: 100, PVActuels: 100, DegatsBase: 13, Initiative: 10},
        {Name: "Prisonnier fou", PVMax: 60, PVActuels: 35, DegatsBase: 20, Initiative: 12},
        {Name: "Chien dressé", PVMax: 40, PVActuels: 40, DegatsBase: 10, Initiative: 12},
    }
    rand.Seed(time.Now().UnixNano())
    return mobs[rand.Intn(len(mobs))]
}

