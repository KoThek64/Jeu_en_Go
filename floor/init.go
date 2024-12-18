package floor

import (
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/quadtree"
	"bufio"
	"strconv"
	"os"
)

// Init initialise les structures de données internes de f.
func (f *Floor) Init() {
	f.content = make([][]int, configuration.Global.NumTileY)
	for y := 0; y < len(f.content); y++ {
		f.content[y] = make([]int, configuration.Global.NumTileX)
	}

	switch configuration.Global.FloorKind {
	case FromFileFloor:
		f.fullContent = readFloorFromFile(configuration.Global.FloorFile)
	case QuadTreeFloor:
		f.quadtreeContent = quadtree.MakeFromArray(readFloorFromFile(configuration.Global.FloorFile))
	}
}

// lecture du contenu d'un fichier représentant un terrain
// pour le stocker dans un tableau
func readFloorFromFile(fileName string) (floorContent [][]int) {
	file, err := os.Open(fileName)
    if err != nil { panic(err) }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        txt := scanner.Text()

        ligne := make([]int, len(txt))

        for carIndex :=  0; carIndex < len(txt); carIndex ++ {
			
            n, err := strconv.Atoi(txt[carIndex:carIndex+1])
            if err != nil { panic(err) }

            ligne[carIndex] = n

        }

        floorContent = append(floorContent, ligne)
    }

    return
}
