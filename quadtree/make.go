package quadtree

// makeFromArrayRecursive construit un nœud de quadtree de manière récursive à partir d'un tableau 2D donné.
// - floorContent : le tableau 2D représentant le terrain.
// - topLeftX, topLeftY : coordonnées du coin supérieur gauche de la zone actuelle du nœud.
// - width, height : dimensions de la zone actuelle du nœud.
func makeFromArrayRecursive(floorContent [][]int, topLeftX, topLeftY, width, height int) *node {
	// Supposer initialement que la zone actuelle a un contenu uniforme
	content := floorContent[topLeftY][topLeftX]
	is_leaf := true

	// Vérifier si la zone n'est pas uniforme
check_leaf:
	for x := topLeftX; x < topLeftX+width; x++ {
		for y := topLeftY; y < topLeftY+height; y++ {
			if floorContent[y][x] != content {
				is_leaf = false
				break check_leaf
			}
		}
	}

	// Créer un nouveau nœud de quadtree
	returnNode := node{
		topLeftX: topLeftX,
		topLeftY: topLeftY,
		width:    width,
		height:   height,
		isLeaf:   is_leaf,
	}

	// Si la zone est uniforme, assigner le contenu
	if is_leaf {
		returnNode.content = content
	} else {
		// Diviser la zone en quatre quadrants et créer les nœuds enfants
		middleWidth := width / 2
		middleHeight := height / 2

		returnNode.topLeftNode = makeFromArrayRecursive(floorContent, topLeftX, topLeftY, middleWidth, middleHeight)
		returnNode.topRightNode = makeFromArrayRecursive(floorContent, topLeftX+middleWidth, topLeftY, width-middleWidth, middleHeight)
		returnNode.bottomLeftNode = makeFromArrayRecursive(floorContent, topLeftX, topLeftY+middleHeight, middleWidth, height-middleHeight)
		returnNode.bottomRightNode = makeFromArrayRecursive(floorContent, topLeftX+middleWidth, topLeftY+middleHeight, width-middleWidth, height-middleHeight)
	}

	return &returnNode
}

// MakeFromArray construit un quadtree à partir d'un tableau 2D représentant un terrain.
// - floorContent : le tableau 2D à convertir en quadtree.
func MakeFromArray(floorContent [][]int) Quadtree {
	

	width := len(floorContent[0]) // Obtenir la largeur du terrain
	height := len(floorContent)   // Obtenir la hauteur du terrain

	// Retourner le quadtree avec son nœud racine initialisé récursivement
	return Quadtree{
		width:  width,
		height: height,
		root:   makeFromArrayRecursive(floorContent, 0, 0, width, height),
	}
}
