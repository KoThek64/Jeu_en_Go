package floor

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/configuration"
	"gitlab.univ-nantes.fr/jezequel-l/quadtree/quadtree"
)

// Init initialise les structures de données internes de f.
func (f *Floor) Init() {
	f.AlrRegistered = false
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

	if configuration.Global.LogicMapGeneration {
		fileName = "../floor-files/logic"
		writeFloorToFile(diamondSquared(configuration.Global.LogicMapSize))
	}

	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		txt := scanner.Text()

		ligne := make([]int, len(txt))

		for carIndex := 0; carIndex < len(txt); carIndex++ {

			n, err := strconv.Atoi(txt[carIndex : carIndex+1])
			if err != nil {
				panic(err)
			}

			ligne[carIndex] = n

		}

		floorContent = append(floorContent, ligne)
	}

	return
}

// writeFloorToFile prend un tableau 2D représentant une carte et l'écrit dans le fichier ../floor-files/logic
func writeFloorToFile(grid [][]int) error {
	// Ouvrir le fichier en mode écriture. Si le fichier n'existe pas, il sera créé.
	file, err := os.OpenFile("../floor-files/logic", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("erreur lors de l'ouverture du fichier: %v", err)
	}
	defer file.Close()

	// Créer un writer tamponné pour améliorer les performances d'écriture
	writer := bufio.NewWriter(file)

	// Parcourir le tableau et écrire chaque ligne dans le fichier
	for i := 0; i < len(grid); i++ { // i correspond à la ligne
		for j := 0; j < len(grid[i]); j++ { // j correspond à la colonne
			_, err := writer.WriteString(fmt.Sprintf("%d", grid[i][j])) // Pas d'espace entre les valeurs
			if err != nil {
				return fmt.Errorf("erreur lors de l'écriture de la cellule (i=%d, j=%d): %v", i, j, err)
			}
		}
		// Ajouter une nouvelle ligne après chaque rangée
		_, err := writer.WriteString("\n")
		if err != nil {
			return fmt.Errorf("erreur lors de l'écriture de la nouvelle ligne (i=%d): %v", i, err)
		}
	}

	// S'assurer que tous les tampons sont écrits dans le fichier
	err = writer.Flush()
	if err != nil {
		return fmt.Errorf("erreur lors de la finalisation de l'écriture: %v", err)
	}

	return nil
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
	for i := 0; i < configuration.Global.RandomMapDimensions[0]; i++ {
		for j := 0; j < configuration.Global.RandomMapDimensions[1]; j++ {
			if i == configuration.Global.RandomMapDimensions[0]/2 && j == configuration.Global.RandomMapDimensions[1]/2 {
				_, err := file.WriteString(fmt.Sprintf("%d", 2))
				if err != nil {
					return fmt.Errorf("erreur lors de l'écriture dans le fichier: %v", err)
				}
			}else {
				_, err := file.WriteString(fmt.Sprintf("%d", rand.Intn(5)))
				if err != nil {
					return fmt.Errorf("erreur lors de l'écriture dans le fichier: %v", err)
				}
			}

		}
		_, err = file.WriteString("\n") // Ajouter une nouvelle ligne après chaque sous-tableau
		if err != nil {
			return fmt.Errorf("erreur lors de l'écriture de la nouvelle ligne: %v", err)
		}
	}
	return nil
}

// diamondSquared génère un terrain en utilisant l'algorithme Diamond-Square
func diamondSquared(size int) [][]int {
	if (size-1)&(size-2) != 0 || size < 2 {
		panic("La taille doit être une puissance de 2 plus 1 (ex: 3, 5, 9, 17, ...)")
	}

	// Créer une grille vide
	floorContent := make([][]int, size)
	for i := range floorContent {
		floorContent[i] = make([]int, size)
	}

	// Initialiser les coins aléatoirement
	floorContent[0][0] = rand.Intn(5)
	floorContent[0][size-1] = rand.Intn(5)
	floorContent[size-1][0] = rand.Intn(5)
	floorContent[size-1][size-1] = rand.Intn(5)

	step := size - 1 // Taille du pas initial

	// Diamond-Square Algorithm
	for step > 1 {
		halfStep := step / 2

		// Étape 1 : Diamond Step
		for x := halfStep; x < size-1; x += step {
			for y := halfStep; y < size-1; y += step {
				// Moyenne des quatre coins
				avg := (
					floorContent[x-halfStep][y-halfStep] +
					floorContent[x-halfStep][y+halfStep] +
					floorContent[x+halfStep][y-halfStep] +
					floorContent[x+halfStep][y+halfStep]) / 4

				// Ajouter une valeur aléatoire
				n0 := avg + 10
				if n0 < 0 {
					floorContent[x][y] = 1 // Océan
				} else if n0 == 0 {
					floorContent[x][y] = 4 // sable
				} else if n0 == 1 {
					floorContent[x][y] = 0 // feuilles
				} else if n0 == 2 {
					floorContent[x][y] = 2 // Cailloux
				} else if n0 == 3 {
					floorContent[x][y] = 3 // Bois
				}
			}
		}

		// Étape 2 : Square Step
		for x := 0; x < size; x += halfStep {
			for y := (x + halfStep) % step; y < size; y += step {
				// Moyenne des voisins adjacents (sans sortir des limites)
				count := 0
				sum := 0
				if x-halfStep >= 0 {
					sum += floorContent[x-halfStep][y]
					count++
				}
				if x+halfStep < size {
					sum += floorContent[x+halfStep][y]
					count++
				}
				if y-halfStep >= 0 {
					sum += floorContent[x][y-halfStep]
					count++
				}
				if y+halfStep < size {
					sum += floorContent[x][y+halfStep]
					count++
				}
				n := (sum / count) + rand.Intn(4) -1
				if n < 0 {
					floorContent[x][y] = 1 // Océan
				} else if n == 0 {
					floorContent[x][y] = 4 // Feuilles
				} else if n == 1 {
					floorContent[x][y] = 0 // Sable
				} else if n == 2 {
					floorContent[x][y] = 2 // Cailloux
				} else if n == 3 {
					floorContent[x][y] = 3 // Bois
				}
			}
		}

		step /= 2
	}
	floorContent[size/2][size/2] = 2

	return floorContent
}
