package main

import (
    "bufio"
    "fmt"
    "strings"
    "unicode"
    "strconv"
)

type Character struct {
    Nom              string
    Classe           string
    Niveau           int
    PVMax            int
    PVActuels        int
    Inventaire       []string
    Argent           int
    DegatsBase       int
    Initiative       int
    Passif           string
    Faim             int
    Fatigue          int
    ÉquipementArme   *Equipment
    ÉquipementArmure *Equipment
    Skills           []string
    XP         int // XP actuelle
    XPNext     int // XP nécessaire pour le prochain niveau

}

func characterCreation(scanner *bufio.Scanner) Character {
    var nom, classe string
    fmt.Print("Quel est ton nom ? ")
    scanner.Scan()
    nom = formatNom(scanner.Text())

    for {
        fmt.Print("Choisis ta peine (Meurtrier, Voleur, Hacker, Psychopathe) : ")
        scanner.Scan()
        classe = strings.Title(strings.ToLower(scanner.Text()))
        if classe == "Meurtrier" || classe == "Voleur" || classe == "Hacker" || classe == "Psychopathe" || classe == "Admin" {
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
        Inventaire: []string{},
        DegatsBase: 15,    
        Initiative: 10,
        Argent:     20,    
        Faim:       20,
        Fatigue:    20,
        ÉquipementArme:   nil,
        ÉquipementArmure: nil,
        Skills: []string{"Coup de poing"},
        XP: 0,
        XPNext: 10,

    }

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
        c.Passif = "+10 dégâts de base, mais faim/fatigue augmentent 2× plus vite. 50% de chance de faire x2 dégâts."
        c.DegatsBase += 10
    }

    c.PVActuels = c.PVMax 
    return c
}

func apresCombat(c *Character) {
    perte := 3
    if c.Classe == "Psychopathe" {
        perte *= 2
    }

    c.Faim -= perte
    c.Fatigue -= perte

    if c.Faim < 10 || c.Fatigue < 10 {
        if c.Initiative > 0 {
            c.Initiative--
            fmt.Println("⚠️ Tu te sens faible... ton initiative diminue.")
        }
    }

    if c.Faim < 0 {
        c.Faim = 0
    }
    if c.Fatigue < 0 {
        c.Fatigue = 0
    }
}

func manger(c *Character) {
    for i, item := range c.Inventaire {
        if item == "Pain sec" {
            fmt.Println("🍞 Tu manges un pain sec. +5 faim.")
            c.Faim += 5
            if c.Faim > 20 {
                c.Faim = 20
            }
            c.Inventaire = append(c.Inventaire[:i], c.Inventaire[i+1:]...)
            return
        }
    }
    fmt.Println("❌ Tu n’as rien à manger.")
}

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

func accessInventory(c *Character, scanner *bufio.Scanner) {
    for {
        fmt.Println("\n--- Inventaire ---")
        if len(c.Inventaire) == 0 {
            fmt.Println("Inventaire vide.")
            return
        }

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

        equip := getEquipmentByName(itemName)
        if equip != nil {
            equiperItem(c, equip, index)
        } else {
            useConsumable(c, itemName, index)
        }
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

func getEquipmentByName(name string) *Equipment {
    for i := range Equipments {
        if Equipments[i].Name == name {
            return &Equipments[i]
        }
    }
    return nil
}

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

    c.Inventaire = append(c.Inventaire[:index], c.Inventaire[index+1:]...)
}

func useConsumable(c *Character, name string, index int) {
    switch name {
    case "Potion de soin":
        c.PVActuels += 30
        if c.PVActuels > c.PVMax {
            c.PVActuels = c.PVMax
        }
        fmt.Println("💊 Tu utilises une potion de soin :", c.PVActuels, "/", c.PVMax)
    case "Pain sec":
        fmt.Println("🍞 Tu manges du pain sec. +5 faim.")
        c.Faim += 5
        if c.Faim > 20 {
            c.Faim = 20
        }
    case "Potion de poison":
        fmt.Println("☠️ Tu prépares une potion de poison. Elle pourra être utilisée en combat.")
        // Tu peux ajouter une logique spécifique si tu veux l'utiliser directement en combat
    case "Livre de Sort : Boule de Feu":
        spellBook(c) // Apprentissage du sort
    default:
        fmt.Println("❌ Cet item ne peut pas être utilisé.")
    }

    // Supprime l'item utilisé
    c.Inventaire = append(c.Inventaire[:index], c.Inventaire[index+1:]...)
}
// Ajoute le sort Boule de Feu au personnage
func spellBook(c *Character) {
    for _, s := range c.Skills {
        if s == "Boule de Feu" {
            fmt.Println("❌ Vous connaissez déjà ce sort !")
            return
        }
    }
    c.Skills = append(c.Skills, "Boule de Feu")
    fmt.Println("✨ Vous avez appris le sort Boule de Feu !")

}

func gagnerXP(c *Character, xpGagne int) {
	fmt.Printf("⭐ Tu gagnes %d XP !\n", xpGagne)
	c.XP += xpGagne

	for c.XP >= c.XPNext {
		c.XP -= c.XPNext
		c.Niveau++
		fmt.Printf("🔼 Félicitations ! Tu passes niveau %d !\n", c.Niveau)
		levelUp(c)
		// Augmente XPNext pour le prochain niveau (exponentiel)
		c.XPNext += 5 + c.Niveau*2
	}
}

func levelUp(c *Character) {
	c.PVMax += 5
	c.PVActuels = c.PVMax // soigner complètement au niveau up
	c.DegatsBase += 2      // arrondi à +2 dégâts
	c.Initiative += 1      // +1 initiative
	fmt.Println("✨ Stats améliorées : +5 PV, +2 dégâts, +1 initiative")
}





func parseInt(s string) int {
    val, err := strconv.Atoi(s)
    if err != nil {
        return 0
    }
    return val
}

