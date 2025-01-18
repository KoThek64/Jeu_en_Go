package character

import (
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/floor"
)

// Init met en place un personnage. Pour le moment
// cela consiste simplement à initialiser une variable
// responsable de définir l'étape d'animation courante.
func (c *Character) Init(f floor.Floor, floorWidth, floorHeight int) {
	c.animationStep = 1

	c.X = floorWidth / 2
	c.Y = floorHeight / 2

	// Vérifie si le point de départ est sur une case d'eau
	if configuration.Global.AvoidWater {
		if c.Y < len(f.Content) && c.X < len(f.Content[c.Y]) {
			if f.Content[c.Y][c.X] == 4 {
				// Si c'est le cas, déplace le personnage vers une case non bloquante
				c.X = floorWidth/2 + 1
				c.Y = floorHeight / 2
			}
		}
	}
}
