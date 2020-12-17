package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	testLines := readAllLines("./day17/test_input.txt")
	lines := readAllLines("./day17/input.txt")

	testInput := parseInput(testLines)
	input := parseInput(lines)

	log.Println("Day 17 Part 01")
	testPartOne(testInput)
	partOne(input)

	log.Println()

	log.Println("Day 17 Part 02")
	testPartTwo(testInput)
	partTwo(input)
}

func testPartOne(input Input) {
	input.simulateCycles(6)
	log.Println("expected: 112, actual:", input.countActiveCubes())
}

func partOne(input Input) {
	input.simulateCycles(6)
	log.Println("Active cubes after the sixth cycle:", input.countActiveCubes())
}

func testPartTwo(input Input) {

}

func partTwo(input Input) {

}

type ConwayCube struct {
	position Position
	active   bool
}

type Position struct {
	x, y, z int
}

type Input struct {
	positionToCube map[string]ConwayCube
}

func (position Position) relative(x, y, z int) Position {
	return Position{
		x: position.x + x,
		y: position.y + y,
		z: position.z + z,
	}
}

func (position Position) toString() string {
	return fmt.Sprintf("%v,%v,%v", position.x, position.y, position.z)
}

func (input Input) getCubeAt(position Position, initialize bool) ConwayCube {
	cube, hasCube := input.positionToCube[position.toString()]

	if !hasCube {
		cube = ConwayCube{position, false}

		if initialize {
			input.positionToCube[position.toString()] = cube
		}
	}

	return cube
}

func (input Input) getNeighborCubes(cube ConwayCube, initialize bool) []ConwayCube {
	cubes := make([]ConwayCube, 26)
	index := 0

	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			for z := -1; z <= 1; z++ {
				if x == 0 && y == 0 && z == 0 {
					continue
				}

				position := cube.position.relative(x, y, z)
				cubes[index] = input.getCubeAt(position, initialize)
				index++
			}
		}
	}

	return cubes
}

func (input Input) initializeNeighbors() {
	activeCubes := input.getActiveCubes()

	for _, cube := range activeCubes {
		input.getNeighborCubes(cube, true)
	}
}

func (input Input) simulateCycle() {
	input.initializeNeighbors()

	modifiedCubes := make([]ConwayCube, len(input.positionToCube))
	index := 0

	for _, cube := range input.positionToCube {
		neighborCubes := input.getNeighborCubes(cube, false)
		activeCount := 0

		for _, neighborCube := range neighborCubes {
			if neighborCube.active {
				activeCount++
			}
		}

		switch cube.active {
		case true:
			if activeCount != 2 && activeCount != 3 {
				cube.active = false
			}
		case false:
			if activeCount == 3 {
				cube.active = true
			}
		}

		modifiedCubes[index] = cube
		index++
	}

	for _, cube := range modifiedCubes {
		input.positionToCube[cube.position.toString()] = cube
	}
}

func (input Input) simulateCycles(cycles int) {
	for i := 0; i < cycles; i++ {
		input.simulateCycle()
	}
}

func (input Input) getActiveCubes() []ConwayCube {
	var activeCubes []ConwayCube

	for _, cube := range input.positionToCube {
		if cube.active {
			activeCubes = append(activeCubes, cube)
		}
	}

	return activeCubes
}

func (input Input) countActiveCubes() int {
	return len(input.getActiveCubes())
}

func (input Input) printSlice(z, min, max int) {
	log.Println(fmt.Sprintf("z: %v, min: %v, max: %v", z, min, max))

	for x := min; x <= max; x++ {
		line := strings.Builder{}

		for y := min; y <= max; y++ {
			position := Position{x, y, z}
			cube := input.getCubeAt(position, false)
			active := ""

			switch cube.active {
			case true:
				active = "#"
			case false:
				active = "."
			}

			line.WriteString(active)
		}

		log.Println(line.String())
	}
}

func parseInput(lines []string) Input {
	positionToCube := make(map[string]ConwayCube)
	input := Input{positionToCube}

	for x, line := range lines {
		for y, activeRune := range line {
			activeString := string(activeRune)
			active := false

			switch activeString {
			case ".":
				active = false
			case "#":
				active = true
			default:
				panic(fmt.Sprintf("Unexpected active rune: %v!", activeString))
			}

			position := Position{x, y, 0}
			cube := ConwayCube{position, active}
			positionToCube[position.toString()] = cube
		}
	}

	return input
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

		if line == "" {
			continue
		}

		lines = append(lines, line)
	}

	return lines
}
