package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	testInstructions := readAllLines("./day12/test_input.txt")
	instructions := readAllLines("./day12/input.txt")

	log.Println("Day 12 Part 01")
	testPartOne(testInstructions)
	partOne(instructions)

	log.Println()

	log.Println("Day 12 Part 02")
	testPartTwo(testInstructions)
	partTwo(instructions)
}

func testPartOne(instructions []string) {
	log.Println(distanceAfterNavigation(instructions))
}

func partOne(instructions []string) {
	log.Println(distanceAfterNavigation(instructions))
}

func testPartTwo(instructions []string) {
	log.Println(distanceAfterWaypointNavigation(instructions))
}

func partTwo(instructions []string) {
	log.Println(distanceAfterWaypointNavigation(instructions))
}

func distanceAfterWaypointNavigation(instructions []string) int {
	var shipEast, shipNorth int
	waypointEast := 10
	waypointNorth := 1

	for _, instruction := range instructions {
		action, value := parseInstruction(instruction)

		switch action {
		case "N":
			waypointEast, waypointNorth = move("N", waypointEast, waypointNorth, value)
		case "S":
			waypointEast, waypointNorth = move("S", waypointEast, waypointNorth, value)
		case "E":
			waypointEast, waypointNorth = move("E", waypointEast, waypointNorth, value)
		case "W":
			waypointEast, waypointNorth = move("W", waypointEast, waypointNorth, value)
		case "L":
			waypointEast, waypointNorth = rotate(waypointEast, waypointNorth, rotateLeft, value)
		case "R":
			waypointEast, waypointNorth = rotate(waypointEast, waypointNorth, rotateRight, value)
		case "F":
			shipEast, shipNorth = moveToWaypoint(shipEast, shipNorth, waypointEast, waypointNorth, value)
		default:
			panic(fmt.Sprintf("Unexpected action: %v!", action))
		}
	}

	return abs(shipEast) + abs(shipNorth)
}

func distanceAfterNavigation(instructions []string) int {
	var east, north int
	orientation := "E"

	for _, instruction := range instructions {
		action, value := parseInstruction(instruction)

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

func parseInstruction(instruction string) (string, int) {
	action := string(instruction[0])
	value, _ := strconv.Atoi(instruction[1:])

	return action, value
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

func moveToWaypoint(shipEast, shipNorth, waypointEast, waypointNorth, distance int) (int, int) {
	for i := 0; i < distance; i++ {
		shipEast += waypointEast
		shipNorth += waypointNorth
	}

	return shipEast, shipNorth
}

type RotateFunc func(int, int) (int, int)

func rotate(east, north int, rotateFunc RotateFunc, degrees int) (int, int) {
	turns := degrees / 90

	for i := 0; i < turns; i++ {
		east, north = rotateFunc(east, north)
	}

	return east, north
}

func rotateLeft(oldEast, oldNorth int) (newEast, newNorth int) {
	newEast = -oldNorth
	newNorth = oldEast

	return newEast, newNorth
}

func rotateRight(oldEast, oldNorth int) (newEast, newNorth int) {
	newEast = oldNorth
	newNorth = -oldEast

	return newEast, newNorth
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
