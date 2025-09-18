package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

// Character : structure principale du joueur
// contient toutes les infos du personnage : stats, inventaire, equipements, skills, XP...
type Character struct {
	Nom              string     // nom du joueur
	Classe           string     // classe choisie (Meurtrier, Voleur, Hacker, Psychopathe, Admin)
	Niveau           int        // niveau actuel
	PVMax            int        // pv maximum
	PVActuels        int        // pv actuels
	Inventaire       []string   // liste des items possédés
	Argent           int        // argent du joueur
	DegatsBase       int        // degats de base sans arme
	Initiative       int        // ordre d'action en combat
	Passif           string     // passif / bonus de la classe
	Faim             int        // faim (max 20)
	Fatigue          int        // fatigue (max 20)
	ÉquipementArme   *Equipment // arme équipée
	ÉquipementArmure *Equipment // armure équipée
	Skills           []string   // compétences et sorts connus
	XP               int        // xp actuelle
	XPNext           int        // xp nécessaire pour passer au niveau suivant
	PoisonNextAttack bool       // poison
	BurnTurns int				// nombre de tour brulure
    BurnDmg   int				// dgt brulure	
}

// characterCreation : crée un personnage à partir du scanner
// demande nom et classe à l'utilisateur et initialise le personnage via initCharacter
func characterCreation(scanner *bufio.Scanner) Character {
	var nom, classe string
	fmt.Print("Quel est ton nom ? ")
	scanner.Scan()
	nom = formatNom(scanner.Text()) // met première lettre en majuscule

	// boucle jusqu'à ce que l'utilisateur choisisse une classe valide
	for {
		fmt.Print("Choisis ta peine (Meurtrier, Voleur, Hacker, Psychopathe) : ")
		scanner.Scan()
		classe = strings.Title(strings.ToLower(scanner.Text()))
		if classe == "Meurtrier" || classe == "Voleur" || classe == "Hacker" || classe == "Psychopathe" || classe == "Admin" {
			break
		}
		fmt.Println("❌ Classe invalide.")
	}

	// retourne le personnage initialisé selon la classe
	return initCharacter(nom, classe)
}

// initCharacter : initialise un personnage selon sa classe
// applique les bonus/malus de la classe et les stats de base
func initCharacter(nom string, classe string) Character {
	c := Character{
		Nom:              nom,
		Classe:           classe,
		Niveau:           1,
		Inventaire:       []string{},
		DegatsBase:       15,
		Initiative:       10,
		Argent:           20,
		Faim:             20,
		Fatigue:          20,
		ÉquipementArme:   nil,
		ÉquipementArmure: nil,
		Skills:           []string{"Coup de poing"}, // compétence de base
		XP:               0,
		XPNext:           10,
	}

	// applique les bonus/malus selon la classe choisie
	switch classe {
	case "Admin":
		c.PVMax = 100000
		c.Passif = "Administrateur"
		c.DegatsBase += 10000
		c.Argent += 50000
		c.Initiative += 500
	case "Meurtrier":
		c.PVMax = 120
		c.Passif = "+20 PV Max, mais +10% fatigue par étage."
		c.DegatsBase += 10
	case "Voleur":
		c.PVMax = 100
		c.Passif = "+5 initiative et +100 or."
		c.Initiative += 5
		c.Argent += 50
	case "Hacker":
		c.PVMax = 80
		c.Passif = "Sort passif : Pirater (monstre confus 1/2 chance de rater)."
	case "Psychopathe":
		c.PVMax = 90
		c.Passif = "+10 dégâts, faim/fatigue augmentent 2× plus vite, 50% chance x2 dégâts."
		c.DegatsBase += 10
	}

	c.PVActuels = c.PVMax
	return c
}

// apresCombat : applique les pertes de faim/fatigue après un combat
// réduit l'initiative si trop faible
// utilise la classe pour modifier la perte (Psychopathe double la perte)
func apresCombat(c *Character) {
	perte := 3
	if c.Classe == "Psychopathe" {
		perte *= 2
	}

	c.Faim -= perte
	c.Fatigue -= perte

	// si trop faible, l'initiative diminue
	if c.Faim < 10 || c.Fatigue < 10 {
		if c.Initiative > 0 {
			c.Initiative--
			fmt.Println("⚠️ Tu te sens faible... ton initiative diminue.")
		}
	}

	// empêche valeurs négatives
	if c.Faim < 0 {
		c.Faim = 0
	}
	if c.Fatigue < 0 {
		c.Fatigue = 0
	}
}

// manger : permet de manger un item consommable
// recherche "Pain sec" dans l'inventaire, augmente faim et supprime l'item
func manger(c *Character) {
	for i, item := range c.Inventaire {
		if item == "Pain sec" {
			fmt.Println("🍞 Tu manges un pain sec. +5 faim.")
			c.Faim += 5
			if c.Faim > 20 {
				c.Faim = 20
			}
			c.Inventaire = append(c.Inventaire[:i], c.Inventaire[i+1:]...) // retire l'objet
			return
		}
	}
	fmt.Println("❌ Tu n’as rien à manger.")
}

// displayInfo : affiche toutes les infos du personnage
// utilise tous les champs du Character pour afficher stats, inventaire, passif et XP
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
	fmt.Printf("XP : %d / %d\n", c.XP, c.XPNext)
}

// accessInventory : permet au joueur de gérer son inventaire
// affiche items, choix d'utiliser ou équiper
// utilise getEquipmentByName, equiperItem et useConsumable
func accessInventory(c *Character, scanner *bufio.Scanner) {
	for {
		fmt.Println("\n--- Inventaire ---")
		if len(c.Inventaire) == 0 {
			fmt.Println("Inventaire vide.")
			return
		}

		// affiche chaque item
		for i, item := range c.Inventaire {
			fmt.Printf("%d. %s\n", i+1, item)
		}
		fmt.Println("0. Retour au menu principal")
		fmt.Print("Choisis un item : ")
		scanner.Scan()
		choix := scanner.Text()

		if choix == "0" {
			return
		}

		index := parseInt(choix) - 1
		if index < 0 || index >= len(c.Inventaire) {
			fmt.Println("❌ Choix invalide.")
			continue
		}

		itemName := c.Inventaire[index]

		// vérifie si c'est un équipement ou un consommable
		equip := getEquipmentByName(itemName)
		if equip != nil {
			equiperItem(c, equip, index)
		} else {
			useConsumable(c, itemName, index)
		}
	}
}

// formatNom : met la première lettre du nom en majuscule
// utile pour uniformiser l'affichage
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

// getEquipmentByName : récupère un équipement depuis son nom
// retourne un pointeur sur l'Equipment ou nil si non trouvé
func getEquipmentByName(name string) *Equipment {
	for i := range Equipments {
		if Equipments[i].Name == name {
			return &Equipments[i]
		}
	}
	return nil
}

// equiperItem : équipe une arme ou armure pour le joueur
// met à jour stats et supprime l'objet de l'inventaire
func equiperItem(c *Character, equip *Equipment, index int) {
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

	// supprime l'objet de l'inventaire après équipement
	c.Inventaire = append(c.Inventaire[:index], c.Inventaire[index+1:]...)
}

// useConsumable : utilise un objet consommable
// gère potions, pain sec et livres de sort
func useConsumable(c *Character, name string, index int) {
	switch name {
	case "Potion de soin":
		c.PVActuels += 30
		if c.PVActuels > c.PVMax {
			c.PVActuels = c.PVMax
		}
		fmt.Println("💊 Tu utilises une potion de soin :", c.PVActuels, "/", c.PVMax)

	case "Potion de soin majeure":
		c.PVActuels += 60
		if c.PVActuels > c.PVMax {
			c.PVActuels = c.PVMax
		}
		fmt.Println("💊 Tu utilises une potion de soin majeure :", c.PVActuels, "/", c.PVMax)

	case "Élixir de régénération":
		c.PVActuels = c.PVMax
		fmt.Println("✨ Tu es totalement régénéré :", c.PVActuels, "/", c.PVMax)

	case "Pain sec":
		c.Faim += 5
		if c.Faim > 20 {
			c.Faim = 20
		}
		fmt.Println("🍞 Tu manges un pain sec. Faim :", c.Faim, "/20")

	case "Sandwich frais":
		c.Faim += 10
		if c.Faim > 20 {
			c.Faim = 20
		}
		fmt.Println("🥪 Tu manges un sandwich frais. Faim :", c.Faim, "/20")

	case "Barre énergétique":
		c.Faim += 8
		if c.Faim > 20 {
			c.Faim = 20
		}
		fmt.Println("🍫 Tu manges une barre énergétique. Faim :", c.Faim, "/20")

	case "Potion de force":
		c.DegatsBase += 5
		fmt.Println("💪 Tu bois une potion de force ! Dégâts augmentés temporairement :", c.DegatsBase)

	case "Potion de rapidité":
		c.Initiative += 3
		fmt.Println("⚡ Tu bois une potion de rapidité ! Initiative augmentée temporairement :", c.Initiative)

	case "Potion de poison":
    c.PoisonNextAttack = true
    fmt.Println("☠️ Potion de poison prête pour le combat !")
		// Gérer le poison pendant le combat

	case "Livre de Sort : Boule de Feu":
		learnSpell(c, "Boule de Feu")

	case "Livre de Sort : Éclair":
		learnSpell(c, "Éclair")
	



	default:
		fmt.Println("❌ Cet item ne peut pas être utilisé.")
	}

	// Supprimer l’item de l’inventaire après usage
	c.Inventaire = append(c.Inventaire[:index], c.Inventaire[index+1:]...)
}


// spellBook : permet d'apprendre un nouveau sort (Boule de Feu)
// vérifie si le joueur ne le connaît pas déjà
func learnSpell(c *Character, spellName string) {
	for _, s := range c.Skills {
		if s == spellName {
			fmt.Println("❌ Vous connaissez déjà ce sort !")
			return
		}
	}
	c.Skills = append(c.Skills, spellName)
	fmt.Printf("✨ Vous avez appris le sort %s !\n", spellName)
}


// gagnerXP : ajoute de l'XP et gère le level up
// appelle levelUp si le joueur a assez d'XP
func gagnerXP(c *Character, xpGagne int) {
	fmt.Printf("⭐ Tu gagnes %d XP !\n", xpGagne)
	c.XP += xpGagne

	// vérifie si le joueur passe un niveau
	for c.XP >= c.XPNext {
		c.XP -= c.XPNext
		c.Niveau++
		fmt.Printf("🔼 Félicitations ! Tu passes niveau %d !\n", c.Niveau)
		levelUp(c)
		// augmente le coût d’XP pour le prochain niveau
		c.XPNext += 5 + c.Niveau*2
	}
}

// levelUp : augmente les stats du personnage après un niveau gagné
// augmente PV max, soin complet, degats et initiative
func levelUp(c *Character) {
	c.PVMax += 5
	c.PVActuels = c.PVMax
	c.DegatsBase += 2
	c.Initiative += 1
	fmt.Println("✨ Stats améliorées : +5 PV, +2 dégâts, +1 initiative")
}

// parseInt : convertit string en int, retourne 0 si erreur
func parseInt(s string) int {
	val, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return val
}
