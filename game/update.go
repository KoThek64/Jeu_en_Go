package game

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
)

// Update met à jour les données du jeu à chaque 1/60 de seconde.
// Il faut bien faire attention à l'ordre des mises-à-jour car elles
// dépendent les unes des autres (par exemple, pour le moment, la
// mise-à-jour de la caméra dépend de celle du personnage et la définition
// du terrain dépend de celle de la caméra).
func (g *Game) Update() error {
	blocking := g.floor.Blocking(g.character.X, g.character.Y, g.camera.X, g.camera.Y)
	g.character.Update(blocking, g.floor)
	g.camera.Update(g.character.X, g.character.Y)
	g.floor.Update(g.camera.X, g.camera.Y)

	if configuration.Global.Zoomable {
		g.handleZoom()
	}

	if !g.floor.AlrRegistered {
		g.handleMapSaving()
	}
	

	return nil
}

var zoomInPressed bool = false
var zoomOutPressed bool = false

func (g *Game) handleZoom() {
	if ebiten.IsKeyPressed(ebiten.KeyO) && !zoomInPressed {
		zoomInPressed = true
		configuration.Global.NumTileX -= 1
		configuration.Global.NumTileY -= 1
		if configuration.Global.NumTileX < 4 {
			configuration.Global.NumTileX = 4
		}
		if configuration.Global.NumTileY < 4 {
			configuration.Global.NumTileY = 4
		}
		configuration.Global.SetComputedFields()
	} else if !ebiten.IsKeyPressed(ebiten.KeyO) && zoomInPressed {
		zoomInPressed = false
	}

	if ebiten.IsKeyPressed(ebiten.KeyP) && !zoomOutPressed {
		zoomOutPressed = true
		configuration.Global.NumTileX += 1
		configuration.Global.NumTileY += 1
		configuration.Global.SetComputedFields()
	} else if !ebiten.IsKeyPressed(ebiten.KeyP) && zoomOutPressed {
		zoomOutPressed = false
	}
}

func (g *Game) handleMapSaving() {
	// Vérifier si la touche S est pressée
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		// Vérifier si une sauvegarde a déjà été enregistrée dans ce cycle

		// Fonction pour sauvegarder
		saveToFile := func() error {
			// Obtenir la date et l'heure actuelles
			now := time.Now()
			timestamp := now.Format("2006-01-19_15-04") // Format : AAAA-MM-JJ_HH-MM

			// Créer le chemin du fichier
			outputDir := "../floor-files/enregistrement"
			outputFile := fmt.Sprintf("%s/%s.txt", outputDir, timestamp) // Nom basé sur la date et l'heure

			// Ouvrir le fichier source (celui que l'on veut copier)
			sourceFile, err := os.Open("../floor-files/random")
			if err != nil {
				return fmt.Errorf("impossible d'ouvrir le fichier random : %v", err)
			}
			defer sourceFile.Close()

			// Créer le fichier de destination
			destFile, err := os.Create(outputFile)
			if err != nil {
				return fmt.Errorf("impossible de créer le fichier de destination : %v", err)
			}
			defer destFile.Close()

			// Copier le contenu du fichier source vers le fichier de destination
			_, err = io.Copy(destFile, sourceFile)
			if err != nil {
				return fmt.Errorf("impossible de copier le contenu du fichier : %v", err)
			}

			fmt.Printf("Map sauvegardé dans '%s'\n", outputFile)
			return nil
		}

		// Appeler la fonction pour sauvegarder
		err := saveToFile()
		if err != nil {
			fmt.Printf("Erreur lors de la sauvegarde : %v\n", err)
		}

		// Marquer comme déjà enregistré pour ce cycle
		g.floor.AlrRegistered = true
	}
}
