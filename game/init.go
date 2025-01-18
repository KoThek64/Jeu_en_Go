package game

// Init initialise les données d'un jeu. Il faut bien
// faire attention à l'ordre des initialisation car elles
// pourraient dépendre les unes des autres.
func (g *Game) Init() {
	g.floor.Init()
	floorWidth := g.floor.GetWidth()
	floorHeight := g.floor.GetHeight()
	g.character.Init(g.floor, floorWidth, floorHeight)
	g.camera.Init(g.character.X, g.character.Y)
}
