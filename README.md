# 🎮 Jeu en Quadtree

> Jeu 2D de navigation et d'exploration avec système de quadtree et téléportation.

<p align="center">
  <img src="https://img.shields.io/badge/Go-1.21-00ADD8?logo=go&logoColor=white" alt="Go 1.21">
  <img src="https://img.shields.io/badge/Ebitengine-v2.5-FF6B6B?logo=go&logoColor=white" alt="Ebitengine">
  <img src="https://img.shields.io/badge/License-GNU-green.svg" alt="License">
</p>

---

## 📋 Contexte du projet

**Le jeu en Quadtree** est un projet académique réalisé dans le cadre de la **SAÉ du Semestre 1** du BUT Informatique, combinant les ressources R1.01 (Introduction au développement) et SAÉ 1.01 (Implémentation d'un besoin client).

| | |
|---|---|
| 🎓 **Formation** | BUT Informatique |
| 👥 **Équipe** | 2 étudiants |
| 📅 **Année** | 2024-2025 |

### 👨‍💻 Membres de l'équipe

- LACHAISE Mattys
- DOUCET Axel

---

## 🎯 Présentation

**Le Jeu en Quadtree** est un jeu d'exploration 2D utilisant une structure de données quadtree pour optimiser le rendu et les collisions. Le joueur contrôle un personnage animé qui peut se déplacer dans un environnement généré aléatoirement ou chargé depuis un fichier, éviter des obstacles et utiliser des portails de téléportation.

### ✨ Fonctionnalités principales

#### 🕹️ Gameplay
- Contrôle fluide d'un personnage animé (4 animations x 5 frames)
- Déplacement dans un monde basé sur des tuiles (tile-based)
- Système de collision avec les zones d'eau et obstacles
- Portails de téléportation entre différentes zones
- Caméra dynamique suivant le personnage

#### 🗺️ Génération de carte
- Chargement de cartes prédéfinies depuis des fichiers
- Génération procédurale aléatoire de niveaux
- Génération logique avec algorithmes avancés
- Support de cartes de taille configurable (jusqu'à 65x65)
- Sauvegarde automatique des parties avec horodatage

#### ⚙️ Optimisation technique
- Structure quadtree pour l'optimisation spatiale
- Système de caméra avec plusieurs modes
- Mode debug avec affichage des informations techniques
- Configuration JSON flexible
- Zoom dynamique en temps réel

---

## 🛠️ Stack technique

| Composant | Technologie |
|-----------|-------------|
| **Langage** | Go 1.21 |
| **Moteur de jeu** | Ebitengine v2.5.9 |
| **Structure de données** | Quadtree (implémentation custom) |
| **Format de config** | JSON |
| **Gestion fenêtre** | GLFW |

---

## 📁 Architecture du projet

```
sae/
├── cmd/                    # Point d'entrée de l'application
│   ├── main.go            # Fichier principal
│   ├── config.json        # Configuration du jeu
│   └── doc.go
│
├── game/                   # Logique principale du jeu
│   ├── game.go            # Structure et données du jeu
│   ├── init.go            # Initialisation
│   ├── update.go          # Boucle de mise à jour
│   ├── draw.go            # Rendu graphique
│   └── layout.go          # Gestion de la disposition
│
├── character/              # Système de personnage
│   ├── character.go       # Structure du personnage
│   ├── init.go            # Initialisation
│   ├── update.go          # Logique de déplacement
│   ├── draw.go            # Rendu du personnage
│   └── portail.go         # Gestion des portails
│
├── floor/                  # Système de terrain
│   ├── floor.go           # Structure du niveau
│   ├── init.go            # Chargement/génération
│   ├── update.go          # Mise à jour du terrain
│   ├── draw.go            # Rendu des tuiles
│   └── blocking.go        # Gestion des collisions
│
├── quadtree/               # Structure de données quadtree
│   ├── quadtree.go        # Définition de la structure
│   ├── make.go            # Construction
│   ├── get.go             # Requêtes spatiales
│   ├── make_test.go       # Tests unitaires
│   └── get_test.go
│
├── camera/                 # Système de caméra
│   ├── camera.go          # Structure de la caméra
│   ├── init.go            # Initialisation
│   └── update.go          # Suivi du personnage
│
├── assets/                 # Ressources du jeu
│   ├── assets.go          # Chargement des assets
│   └── licence            # Licences des ressources
│
├── configuration/          # Gestion de la configuration
│   └── configuration.go   # Parseur JSON
│
├── floor-files/            # Fichiers de niveaux
│   ├── exemple            # Niveau d'exemple
│   ├── logic              # Niveau logique
│   ├── random             # Niveau aléatoire
│   └── enregistrement/    # Sauvegardes automatiques
│
└── go.mod                 # Dépendances Go
```

---

## 🚀 Installation et lancement

### Prérequis

- **Go 1.21** ou supérieur
- Système d'exploitation : Linux ou Windows
- Bibliothèques système pour GLFW (voir [documentation Ebitengine](https://ebitengine.org/en/documents/install.html))

### Installation

```bash
# 1. Cloner le projet
git clone https://github.com/KoThek64/Jeu_en_Go.git
cd sae

# 2. Installer les dépendances
go mod download

# 3. Compiler le jeu
cd cmd
go build -o main

# 4. Lancer le jeu
./main
```

### 🎮 Contrôles

- **Flèches directionnelles** : Déplacement du personnage
- **Zoom** : Molette de la souris (si activé dans la config)

---

## ⚙️ Configuration

Modifiez [cmd/config.json](cmd/config.json) pour personnaliser le jeu :

```json
{
    "DebugMode": false,                   // Affiche les informations de debug
    "NumTileX": 9,                        // Largeur de la carte en tuiles
    "NumTileY": 9,                        // Hauteur de la carte en tuiles
    "TileSize": 16,                       // Taille d'une tuile en pixels
    "CameraMode": 1,                      // Mode de caméra (0: fixe, 1: suiveuse)
    "FloorKind": 1,                       // Type de terrain
    "RandomGeneration": false,            // Génération aléatoire
    "LogicMapGeneration": true,           // Génération logique avancée
    "LogicMapSize": 65,                   // Taille de la carte logique
    "AvoidWater": false,                  // Activer les collisions avec l'eau
    "Zoomable": true,                     // Activer le zoom
    "TeleportationExtension": false,      // Activer les portails
    "FloorFile": "../floor-files/exemple" // Fichier de niveau à charger
}
```

---

## 🧪 Tests

```bash
# Lancer les tests du quadtree
cd quadtree
go test -v

# Lancer tous les tests
go test ./...
```

---

## 📝 Fonctionnalités avancées

- **Quadtree** : Permet une recherche spatiale optimisée des éléments du jeu
- **Génération procédurale** : Algorithmes de génération de niveaux jouables
- **Sauvegarde automatique** : Les parties sont enregistrées avec horodatage dans `floor-files/enregistrement/`
- **Multi-modes de caméra** : Caméra fixe ou suivant le personnage
- **Système de portails** : Extension permettant la téléportation (mode expérimental)

---

<p align="center">
  Projet réalisé avec ❤️ dans le cadre du BUT Informatique
</p>