package quadtree

func makeFromArrayRecursive(floorContent [][]int, topLeftX, topLeftY, width, height int) *node {

	content := floorContent[topLeftY][topLeftX]
	is_leaf := true

	check_leaf: for x := topLeftX; x < topLeftX + width; x ++ {
		for y := topLeftY; y < topLeftY + height; y ++ {
			if floorContent[y][x] != content {
				is_leaf = false
				break check_leaf
			}

		}
	}

	returnNode := node{
		topLeftX: topLeftX,
		topLeftY: topLeftY,
		width: width,
		height: height,
		isLeaf: is_leaf,
	}


	if is_leaf {
		returnNode.content = content
	}else {
		middleWidth := width / 2
		middleHeight := height / 2

		returnNode.topLeftNode = makeFromArrayRecursive(floorContent, topLeftX, topLeftY, middleWidth, middleHeight)
		returnNode.topRightNode = makeFromArrayRecursive(floorContent, topLeftX + middleWidth, topLeftY, width - middleWidth, middleHeight)
		returnNode.bottomLeftNode = makeFromArrayRecursive(floorContent, topLeftX, topLeftY + middleHeight, middleWidth, height - middleHeight)
		returnNode.bottomRightNode = makeFromArrayRecursive(floorContent, topLeftX + middleWidth, topLeftY + middleHeight, width - middleWidth, height - middleHeight)
	}

	return &returnNode
}

// MakeFromArray construit un quadtree représentant un terrain
// étant donné un tableau représentant ce terrain.
func MakeFromArray(floorContent [][]int) Quadtree {
	
	width := len(floorContent[0])
	height := len(floorContent)

	return Quadtree{
		width: width,
		height: height,
		root: makeFromArrayRecursive(floorContent, 0, 0, width, height),
	}
}
