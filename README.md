# projet-jeux-go

# 🏚️ Prison – RPG Prison 

**Prison** est un **RPG en terminal** développé en **Go**, dans un univers carcéral sombre et stratégique.  
Le joueur incarne un prisonnier qui doit **survivre**, **combattre** et tenter de **s’évader** en progressant à travers les étages de la prison.  

---

## 🎯 Objectif du jeu  
Le but de **Prison** est de **s’évader de la prison**.  

- Gérez votre **faim** et votre **fatigue** pour rester en vie.  
- Progressez **étage par étage**, affrontez des ennemis et des mini-boss.  
- Utilisez vos **compétences**, votre **inventaire** et le **marchand** pour améliorer vos chances.  
- Atteignez l’**étage final** et vainquez le **directeur de la prison** pour vous échapper.  

👉 Si vous échouez : **Game Over**.  
👉 Si vous réussissez : **vous êtes libre**.  

---

## 🧍 Personnages (Peines / Classes)  
Au début du jeu, vous choisissez une **peine** qui définit vos bonus et malus :  

- 🔪 **Meurtrier** : +20 PV max, mais +20% fatigue par étage.  
- 🥷 **Voleur** : +5 initiative et bonus de vol, mais -20 PV max.  
- 💻 **Hacker** : sort bonus *Pirater*, mais moins d’équipement.  
- 💊 **Dealer** : plus d’argent et meilleures potions, mais moins d’XP.  
- 🧨 **Terroriste** : +25 HP max, skill "Bombe artisanale", mais -30 argent.  
- 🧪 **Scientifique fou** : commence avec une seringue spéciale, mais fatigue +15.  
- 🕵️ **Prisonnier politique** : +50 argent, skill "Discours", mais -20 HP.  
- 🚔 **Ancien flic** : +5 initiative, skill "Menottes", mais détesté par les autres.  
- 👹 **Psychopathe** : +10 dégâts de base, mais faim/fatigue augmentent 2× plus vite.  

---

## ⚔️ Gameplay  
- **Tour par tour** : attaque, inventaire, compétences, fuite.  
- **Survie** : faim et fatigue influencent vos performances.  
- **Inventaire** : potions, nourriture, armes improvisées, talismans…  
- **Craft** : créez des objets (ex. gilet pare-balles en journaux).  
- **Statuts spéciaux** : bénédictions et malédictions qui changent vos capacités.  

---

## 🛒 Marchand (Contrebandeur)  
Achetez et vendez des objets grâce à la **monnaie de prison** : cigarettes ou pièces.  

---


## 📂 Structure du projet
```text

prison/
├── main.go              // Menu principal, boucle de jeu
├── character.go         // Structure et gestion du personnage
├── inventory.go         // Gestion de l’inventaire et des objets
├── shop.go              // Marchand
├── equipment.go         // Gestion des équipements
├── combat.go            // Combat tour par tour
├── monster.go           // Structure des monstres
├── training.go          // Combat d'entraînement
├── dungeon.go           // Progression dans les étages (à ajouter)
└── utils.go             // Fonctions utilitaires (ex : input, vérifications)
```
---

## 🔧 Stack technique  
- **Langage** : Go  
- **Interface** : en terminal  
- **Architecture** :  
  - `character.go` → gestion des personnages  
  - `combat.go` → système de combat  
  - `items.go` → objets et inventaire  
  - `dungeon.go` → progression et étages  
  - `shop.go` → marchand/économie  
  - `main.go` → boucle principale du jeu  

---

## 🚀 Installation & Lancement  
```bash
git clone https://github.com/Yakolacoco/projet-jeux-go-.git
cd projet-jeux-go-

# Lancer le jeu
go run main.go