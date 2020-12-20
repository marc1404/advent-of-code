package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	testLines := readAllLines("./day20/test_input.txt")
	lines := readAllLines("./day20/input.txt")

	testTiles := parseTiles(testLines)
	tiles := parseTiles(lines)

	findNeighbors(testTiles)
	findNeighbors(tiles)

	log.Println("Day 20 Part 01")
	testPartOne(testTiles)
	partOne(tiles)

	log.Println()

	log.Println("Day 20 Part 02")
	testPartTwo()
	partTwo()
}

func testPartOne(tiles []*Tile) {
	log.Println("expected: 20899048083289, actual:", calculateProductOfCornerTileIds(tiles))
}

func partOne(tiles []*Tile) {
	log.Println("What do you get if you multiply together the IDs of the four corner tiles?")
	log.Println("Answer:", calculateProductOfCornerTileIds(tiles))
}

func testPartTwo() {
}

func partTwo() {
}

func parseTiles(lines []string) []*Tile {
	var tiles []*Tile
	tile := createEmptyTile()
	mode := "id"

	for _, line := range lines {
		if line == "" {
			if tile.id != -1 {
				tile.initializeBorders()

				tiles = append(tiles, tile)
			}

			tile = createEmptyTile()
			mode = "id"
			continue
		}

		switch mode {
		case "id":
			idAsString := strings.TrimPrefix(line, "Tile ")
			idAsString = strings.TrimSuffix(idAsString, ":")
			id, _ := strconv.Atoi(idAsString)
			tile.id = id
			mode = "image"
		case "image":
			tile.image = append(tile.image, line)
		default:
			panic(fmt.Sprintf("Unexpected mode: %v!", mode))
		}
	}

	if tile.id != -1 {
		tile.initializeBorders()

		tiles = append(tiles, tile)
	}

	return tiles
}

func findNeighbors(tiles []*Tile) {
	for _, tile := range tiles {
		tile.findNeighbors(tiles)
	}
}

func calculateProductOfCornerTileIds(tiles []*Tile) int {
	checksumProduct := 1

	for _, tile := range tiles {
		if tile.isCorner() {
			checksumProduct *= tile.id
		}
	}

	return checksumProduct
}

type Tile struct {
	id      int
	image   []string
	borders []*Border
}

func createEmptyTile() *Tile {
	return &Tile{-1, make([]string, 0), make([]*Border, 0)}
}

func (tile *Tile) initializeBorders() {
	tile.borders = make([]*Border, 4)
	leftSequenceBuilder := strings.Builder{}
	rightSequenceBuilder := strings.Builder{}

	for _, row := range tile.image {
		leftSequenceCharacter := string(row[0])
		rightSequenceCharacter := string(row[len(row)-1])

		leftSequenceBuilder.WriteString(leftSequenceCharacter)
		rightSequenceBuilder.WriteString(rightSequenceCharacter)
	}

	tile.borders[0] = createBorder(tile.image[0], tile)                 // top
	tile.borders[1] = createBorder(rightSequenceBuilder.String(), tile) // right
	tile.borders[2] = createBorder(tile.image[len(tile.image)-1], tile) // bottom
	tile.borders[3] = createBorder(leftSequenceBuilder.String(), tile)  // left
}

func (tile *Tile) findNeighbors(tiles []*Tile) {
	for _, border := range tile.borders {
		border.findNeighbors(tiles)
	}
}

func (tile *Tile) isCorner() bool {
	edgeCount := 0

	for _, border := range tile.borders {
		if len(border.neighbors) == 0 {
			edgeCount++
		}
	}

	return edgeCount == 2
}

type Border struct {
	sequence  string
	tile      *Tile
	neighbors []*Border
}

func createBorder(sequence string, tile *Tile) *Border {
	return &Border{sequence, tile, make([]*Border, 0)}
}

func (border *Border) findNeighbors(tiles []*Tile) {
	for _, neighbor := range tiles {
		if neighbor.id == border.tile.id {
			continue
		}

		for _, neighborBorder := range neighbor.borders {
			if border.checkAlignment(neighborBorder) {
				border.neighbors = append(border.neighbors, neighborBorder)
			}
		}
	}
}

func (border *Border) checkAlignment(neighbor *Border) bool {
	if border.sequence == neighbor.sequence {
		return true
	}

	if reverseString(border.sequence) == neighbor.sequence {
		return true
	}

	return false
}

func reverseString(input string) string {
	output := make([]rune, len(input))

	for inputIndex, character := range input {
		outputIndex := len(output) - 1 - inputIndex
		output[outputIndex] = character
	}

	return string(output)
}

func readAllLines(filePath string) (lines []string) {
	file, err := os.Open(filePath)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	return lines
}
