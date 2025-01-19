package character

import (
	"github.com/hajimehoshi/ebiten/v2"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/floor"
)

// Update met à jour la position du personnage, son orientation
// et son étape d'animation (si nécessaire) à chaque pas
// de temps, c'est-à-dire tous les 1/60 secondes.
func (c *Character) Update(blocking [4]bool, f floor.Floor) {
	blocking = f.Blocking(c.X, c.Y, c.X, c.Y)
	camXPos := c.X
	camYPos := c.Y
	relativeXPos := c.X - camXPos + configuration.Global.ScreenCenterTileX
	relativeYPos := c.Y - camYPos + configuration.Global.ScreenCenterTileY

	if !c.moving {
		teleported := false
		if configuration.Global.TeleportationExtension {
			teleported = c.teleport()
		}
		if configuration.Global.AvoidWater && !teleported {
			if ebiten.IsKeyPressed(ebiten.KeyRight) {
				c.orientation = orientedRight
				if !blocking[1] && f.Content[relativeYPos][relativeXPos+1] != 4 && c.X+1 < f.GetWidth() {
					c.xInc = 1
					c.moving = true
				}
			} else if ebiten.IsKeyPressed(ebiten.KeyLeft) {
				c.orientation = orientedLeft
				if !blocking[3] && f.Content[relativeYPos][relativeXPos-1] != 4 && c.X-1 >= 0 {
					c.xInc = -1
					c.moving = true
				}
			} else if ebiten.IsKeyPressed(ebiten.KeyUp) {
				c.orientation = orientedUp
				if !blocking[0] && f.Content[relativeYPos-1][relativeXPos] != 4 && c.Y-1 >= 0 {
					c.yInc = -1
					c.moving = true
				}
			} else if ebiten.IsKeyPressed(ebiten.KeyDown) {
				c.orientation = orientedDown
				if !blocking[2] && f.Content[relativeYPos+1][relativeXPos] != 4 && c.Y+1 < f.GetHeight() {
					c.yInc = 1
					c.moving = true
				}
			}
		} else if !configuration.Global.AvoidWater && !teleported {
			if ebiten.IsKeyPressed(ebiten.KeyRight) {
				c.orientation = orientedRight
				if !blocking[1] && c.X+1 < f.GetWidth() {
					c.xInc = 1
					c.moving = true
				}
			} else if ebiten.IsKeyPressed(ebiten.KeyLeft) {
				c.orientation = orientedLeft
				if !blocking[3] && c.X-1 >= 0 {
					c.xInc = -1
					c.moving = true
				}
			} else if ebiten.IsKeyPressed(ebiten.KeyUp) {
				c.orientation = orientedUp
				if !blocking[0] && c.Y-1 >= 0 {
					c.yInc = -1
					c.moving = true
				}
			} else if ebiten.IsKeyPressed(ebiten.KeyDown) {
				c.orientation = orientedDown
				if !blocking[2] && c.Y+1 < f.GetHeight() {
					c.yInc = 1
					c.moving = true
				}
			}
		}
	} else {
		c.animationFrameCount++
		if c.animationFrameCount >= configuration.Global.NumFramePerCharacterAnimImage {
			c.animationFrameCount = 0
			shiftStep := configuration.Global.TileSize / configuration.Global.NumCharacterAnimImages
			c.shift += shiftStep
			c.animationStep = -c.animationStep
			if c.shift > configuration.Global.TileSize-shiftStep {
				c.shift = 0
				c.moving = false
				c.X += c.xInc
				c.Y += c.yInc
				c.xInc = 0
				c.yInc = 0
			}
		}
	}
}
