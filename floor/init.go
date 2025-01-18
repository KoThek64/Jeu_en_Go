package floor

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"math/rand"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/quadtree"
)

// Init initialise les structures de données internes de f.
func (f *Floor) Init() {
	f.Content = make([][]int, configuration.Global.NumTileY)
	for y := 0; y < len(f.Content); y++ {
		f.Content[y] = make([]int, configuration.Global.NumTileX)
	}

	switch configuration.Global.FloorKind {
	case FromFileFloor:
		f.FullContent = readFloorFromFile(configuration.Global.FloorFile)
	case QuadTreeFloor:
		f.QuadtreeContent = quadtree.MakeFromArray(readFloorFromFile(configuration.Global.FloorFile))
	}
}

// lecture du contenu d'un fichier représentant un terrain
// pour le stocker dans un tableau
func readFloorFromFile(fileName string) (floorContent [][]int) {
	if configuration.Global.RandomGeneration {
		fileName = "../floor-files/random"
		RandomMapInFile(fileName)
	}

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


func RandomMapInFile(nomFichier string) error {
	// Vider le fichier avant d'écrire de nouvelles données
	err := os.Truncate(nomFichier, 0)
	if err != nil {
		return fmt.Errorf("erreur lors de l'effacement du fichier: %v", err)
	}

	// Ouvrir le fichier en mode ajout
	file, err := os.OpenFile(nomFichier, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("erreur lors de l'ouverture du fichier: %v", err)
	}
	defer file.Close()

	// Écrire chaque sous-tableau dans le fichier
	for i:=0; i < configuration.Global.RandomMapDimensions[0]; i++ {
		for j:=0; j < configuration.Global.RandomMapDimensions[1]; j++ {
			_, err := file.WriteString(fmt.Sprintf("%d", rand.Intn(5)))
			if err != nil {
				return fmt.Errorf("erreur lors de l'écriture dans le fichier: %v", err)
			}
		}
		_, err = file.WriteString("\n") // Ajouter une nouvelle ligne après chaque sous-tableau
		if err != nil {
			return fmt.Errorf("erreur lors de l'écriture de la nouvelle ligne: %v", err)
		}
	}
	return nil
}
