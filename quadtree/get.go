package quadtree

func nextNode(j int, i int, noeud node) (next_node node) {
	if j < noeud.bottomLeftNode.topLeftY {
		//la case est dans la partie superieur du quadtree
		if i < noeud.topRightNode.topLeftX {
			return *noeud.topLeftNode
		} else {
			return *noeud.topRightNode
		}
	} else {
		//la case est dans la partie inferieur du quadtree
		if i < noeud.bottomRightNode.topLeftX {
			return *noeud.bottomLeftNode
		} else {
			return *noeud.bottomRightNode
		}
	}
}

// GetContent remplit le tableau contentHolder (qui reprÃ©sente
// un terrain dont la case le plus en haut Ã  gauche a pour coordonnÃ©es
// (topLeftX, topLeftY)) Ã  partir du qadtree q.
func (q Quadtree) GetContent(topLeftX, topLeftY int, contentHolder [][]int) { //Simon
	// recuperation du noeud racine
	var racine node = *q.root

	//parcours des cases du tableau
	for i := 0; i < len(contentHolder); i++ {
		for j := 0; j < len(contentHolder[i]); j++ {
			//current_node : noeud sur lequel on travaille
			var current_node node = racine

			//Y, X: conversion de la position realtive dans contentHolder Ã  une position absolue dans le quadtree
			var Y int = i + topLeftY
			var X int = j + topLeftX

			//verification que la case relative ne soit pas hors du quadtree
			var out_of_map bool = (Y < current_node.topLeftY ||
				Y > (current_node.topLeftY+current_node.height-1) ||
				X < current_node.topLeftX ||
				X > (current_node.topLeftX+current_node.width-1))

			if out_of_map {
				contentHolder[i][j] = -1
			} else {

				//recherche du leaf dans lequel se trouve la case
				for !current_node.isLeaf {
					current_node = nextNode(Y, X, current_node)
				}
				contentHolder[i][j] = current_node.content
			}
		}
	}
}
