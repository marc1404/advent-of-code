package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	lines := readAllLines("./day11/input.txt")

	log.Println("Day 11 Part 01")
	testPartOne()
	partOne(lines)

	log.Println()

	log.Println("Day 11 Part 02")
	testPartTwo()
	partTwo(lines)
}

func testPartOne() {
	layout := readAllLines("./day11/test_input.txt")

	log.Println(countOccupiedWhenPredictionStabilized(layout))
}

func partOne(layout [][]rune) {
	log.Println(countOccupiedWhenPredictionStabilized(layout))
}

func testPartTwo() {

}

func partTwo(lines [][]rune) {

}

func countOccupiedWhenPredictionStabilized(layout [][]rune) int {
	prediction := layout
	var stable bool

	for {
		prediction, stable = predictSeatLayout(prediction)

		if stable {
			break
		}
	}

	return countOccupiedSeats(prediction)
}

func predictSeatLayout(layout [][]rune) ([][]rune, bool) {
	prediction := make([][]rune, len(layout))
	stable := true

	for y, row := range layout {
		prediction[y] = make([]rune, len(row))

		for x, seatValue := range row {
			seat := string(seatValue)
			predictedSeat := seatValue
			occupiedCount := countAdjacentOccupiedSeats(layout, y, x)

			if seat == "L" && occupiedCount == 0 {
				predictedSeat = []rune("#")[0]
				stable = false
			}

			if seat == "#" && occupiedCount >= 4 {
				predictedSeat = []rune("L")[0]
				stable = false
			}

			prediction[y][x] = predictedSeat
		}
	}

	return prediction, stable
}

func countAdjacentOccupiedSeats(layout [][]rune, y, x int) int {
	occupiedCount := 0
	modifiers := [][]int{
		{-1, -1},
		{-1, 0},
		{-1, 1},
		{0, -1},
		{0, 1},
		{1, -1},
		{1, 0},
		{1, 1},
	}

	for _, modifier := range modifiers {
		modY := modifier[0]
		modX := modifier[1]

		if isSeatOccupied(layout, y+modY, x+modX) {
			occupiedCount++
		}
	}

	return occupiedCount
}

func isSeatOccupied(layout [][]rune, y, x int) bool {
	if y < 0 || y >= len(layout) {
		return false
	}

	row := layout[y]

	if x < 0 || x >= len(row) {
		return false
	}

	seat := string(row[x])

	return seat == "#"
}

func countOccupiedSeats(layout [][]rune) int {
	occupiedCount := 0

	for _, row := range layout {
		for _, seatValue := range row {
			if string(seatValue) == "#" {
				occupiedCount++
			}
		}
	}

	return occupiedCount
}

func readAllLines(filePath string) [][]rune {
	file, err := os.Open(filePath)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := make([][]rune, 0)

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			continue
		}

		lines = append(lines, []rune(line))
	}

	return lines
}
