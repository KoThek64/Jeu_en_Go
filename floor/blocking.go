package floor

import "gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"

// Blocking retourne, étant donnée la position du personnage,
// un tableau de booléen indiquant si les cases au dessus (0),
// à droite (1), au dessous (2) et à gauche (3) du personnage
// sont bloquantes.
func (f Floor) Blocking(characterXPos, characterYPos, camXPos, camYPos int) (blocking [4]bool) {
	relativeXPos := characterXPos - camXPos + configuration.Global.ScreenCenterTileX
	relativeYPos := characterYPos - camYPos + configuration.Global.ScreenCenterTileY

	if relativeXPos < 0 || relativeXPos >= len(f.Content[0]) || relativeYPos < 0 || relativeYPos >= len(f.Content) {
		blocking[0] = true
		blocking[1] = true
		blocking[2] = true
		blocking[3] = true
		return
	}

	if configuration.Global.AvoidWater {
		blocking[0] = relativeYPos <= 0 || f.Content[relativeYPos-1][relativeXPos] == 4
		blocking[1] = relativeXPos >= configuration.Global.NumTileX-1 || f.Content[relativeYPos][relativeXPos+1] == 4
		blocking[2] = relativeYPos >= configuration.Global.NumTileY-1 || f.Content[relativeYPos+1][relativeXPos] == 4
		blocking[3] = relativeXPos <= 0 || f.Content[relativeYPos][relativeXPos-1] == 4
	} else {
		blocking[0] = relativeYPos <= 0
		blocking[1] = relativeXPos >= configuration.Global.NumTileX-1
		blocking[2] = relativeYPos >= configuration.Global.NumTileY-1
		blocking[3] = relativeXPos <= 0
	}

	return
}
