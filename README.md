# projet-jeux-go

# 🏚️ COCHLEA – RPG en terminal

**COCHLEA** est un **RPG en terminal** développé en **Go**, dans un univers sombre et oppressant.  
Le joueur incarne un prisonnier piégé dans une **tour labyrinthique** où chaque étage est gardé par des créatures de plus en plus puissantes.  
Votre objectif : **survivre, progresser et vaincre le maître de la tour**.  

---

## 🎯 Objectif du jeu
Le but de **COCHLEA** est de **gravir les 20 étages de la tour**.  

- Gérez votre **faim** et votre **fatigue** pour rester en vie.  
- Progressez **étage par étage**, combattez des monstres, des sous-boss et un boss final.  
- Utilisez vos **sorts**, vos **objets** et la **boutique** pour améliorer vos chances.  
- Atteignez le dernier étage et remportez la victoire.  

👉 Si vous échouez : **Game Over**.  
👉 Si vous réussissez : **vous êtes libre**.  

---

## 🧍 Classes disponibles
Au début du jeu, vous choisissez une **classe** qui définit vos bonus et malus :  

- 🔪 **Meurtrier** : +20 PV max, mais +20% de fatigue par étage.  
- 🥷 **Voleur** : +5 initiative et bonus de vol, mais -20 PV max.  
- 💻 **Hacker** : commence avec un sort bonus (*Piratage*), mais moins d’équipement.  
- 👹 **Psychopathe** : +10 dégâts de base, mais faim/fatigue augmentent 2× plus vite.  
- 🛠️ **Admin (debug)** : accès spécial pour tests.  

---

## ⚔️ Gameplay
- **Combat au tour par tour** : attaques, sorts, inventaire, fuite.  
- **Gestion de la survie** : faim et fatigue influencent vos statistiques.  
- **Inventaire** : potions, nourriture, armes, armures, livres de sorts.  
- **Économie** : dépensez vos capsules pour acheter du meilleur équipement.  
- **Progression** : monstres aléatoires, sous-boss stratégiques, boss final.  

---

## 🛒 Boutique
Bienvenue dans la boutique !  
Dépensez vos capsules pour acheter armes, protections et objets essentiels.  

### 🗡️ Armes
- 🔪 **Couteau artisanal** : +15 dégâts | 20 capsules  
- 🪵 **Bâton en bois** : +8 dégâts | 10 capsules  
- ⚒️ **Barre de fer** : +30 dégâts | 30 capsules  

### 🛡️ Armures
- 👕 **Gilet pare-balles léger** : +20 PV | 40 capsules  
- 🧥 **Gilet pare-balles lourd** : +50 PV | 80 capsules  

### ⚗️ Consommables
- 🍷 **Potion de soin** : soigne +30 PV | 25 capsules  
- 🍞 **Pain sec** : réduit la faim (+5) | 10 capsules  
- ☠️ **Potion de poison** : inflige +20 dégâts empoisonnés | 20 capsules  
- 📘 **Livre de sort – Boule de Feu** : apprend un sort | 50 capsules  

---

## 🔧 Stack technique  

Le projet est développé en **Go** et fonctionne entièrement dans le **terminal**.  
Chaque composant a un rôle précis pour organiser le code et faciliter la maintenance :  

### 🌐 Langage et outils
- **Langage** : Go (Golang)  
- **Interface** : terminal (console)  
- **Gestion des dépendances** : modules Go
  
---
## 🗂️ Structure du projet

```text
prison/
├── main.go              // Boucle principale du jeu et menu principal
├── character.go         // Gestion du personnage : création, stats, compétences
├── inventory.go         // Gestion de l’inventaire, objets et équipements
├── shop.go              // Boutique : achat et utilisation d’objets
├── equipment.go         // Définition et gestion des armes et armures
├── combat.go            // Système de combat tour par tour
├── monster.go           // Création et gestion des monstres et boss
├── dungeon.go           // Progression dans les étages, coffres et monstres
├── training.go          // Zone d’entraînement ou combats optionnels
└── utils.go             // Fonctions utilitaires (input, conversion, affichage)
```
---

## 🚀 Installation & Lancement

```bash
# Cloner le projet
git clone https://github.com/Yakolacoco/Cochlea-.git
cd src

# Lancer le jeu (via Git Bash)
go run *.go


