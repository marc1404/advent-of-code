package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	lines := readAllLines("./day12/input.txt")

	log.Println("Day 12 Part 01")
	testPartOne()
	partOne(lines)

	log.Println()

	log.Println("Day 12 Part 02")
	testPartTwo()
	partTwo()
}

func testPartOne() {
	instructions := []string{
		"F10",
		"N3",
		"F7",
		"R90",
		"F11",
	}

	log.Println(distanceAfterNavigation(instructions))
}

func partOne(instructions []string) {
	log.Println(distanceAfterNavigation(instructions))
}

func testPartTwo() {

}

func partTwo() {

}

func distanceAfterNavigation(instructions []string) int {
	var east, north int
	orientation := "E"

	for _, instruction := range instructions {
		action := string(instruction[0])
		value, _ := strconv.Atoi(instruction[1:])

		switch action {
		case "N":
			east, north = move("N", east, north, value)
		case "S":
			east, north = move("S", east, north, value)
		case "E":
			east, north = move("E", east, north, value)
		case "W":
			east, north = move("W", east, north, value)
		case "L":
			orientation = turn(orientation, turnLeft, value)
		case "R":
			orientation = turn(orientation, turnRight, value)
		case "F":
			east, north = move(orientation, east, north, value)
		default:
			panic(fmt.Sprintf("Unexpected action: %v!", action))
		}
	}

	return abs(east) + abs(north)
}

func move(direction string, east, north, distance int) (int, int) {
	switch direction {
	case "N":
		north += distance
	case "S":
		north -= distance
	case "E":
		east += distance
	case "W":
		east -= distance
	default:
		panic(fmt.Sprintf("Unexpected direction: %v!", direction))
	}

	return east, north
}

func turn(orientation string, turnFunc TurnFunc, degrees int) string {
	turns := degrees / 90

	for i := 0; i < turns; i++ {
		orientation = turnFunc(orientation)
	}

	return orientation
}

type TurnFunc func(string) string

func turnLeft(orientation string) string {
	switch orientation {
	case "N":
		return "W"
	case "S":
		return "E"
	case "E":
		return "N"
	case "W":
		return "S"
	default:
		panic(fmt.Sprintf("Unexpected orientation: %v!", orientation))
	}
}

func turnRight(orientation string) string {
	switch orientation {
	case "N":
		return "E"
	case "S":
		return "W"
	case "E":
		return "S"
	case "W":
		return "N"
	default:
		panic(fmt.Sprintf("Unexpected orientation: %v!", orientation))
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
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
