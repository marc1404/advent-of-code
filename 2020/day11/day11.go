package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	testLayout := readAllLines("./day11/test_input.txt")
	lines := readAllLines("./day11/input.txt")

	log.Println("Day 11 Part 01")
	testPartOne(testLayout)
	partOne(lines)

	log.Println()

	log.Println("Day 11 Part 02")
	testPartTwo(testLayout)
	partTwo(lines)
}

func testPartOne(layout [][]rune) {
	log.Println(countOccupiedWhenPredictionStabilized(layout, countAdjacentOccupiedSeats, 4))
}

func partOne(layout [][]rune) {
	log.Println(countOccupiedWhenPredictionStabilized(layout, countAdjacentOccupiedSeats, 4))
}

func testPartTwo(layout [][]rune) {
	printLayout(layout)

	tempLayout := [][]rune{
		[]rune(".......#."),
		[]rune("...#....."),
		[]rune(".#......."),
		[]rune("........."),
		[]rune("..#L....#"),
		[]rune("....#...."),
		[]rune("........."),
		[]rune("#........"),
		[]rune("...#....."),
	}

	log.Println(countOccupiedSeatsLineOfSight(tempLayout, 4, 3))

	tempLayout = [][]rune{
		[]rune("............."),
		[]rune(".L.L.#.#.#.#."),
		[]rune("............."),
	}

	log.Println(countOccupiedSeatsLineOfSight(tempLayout, 1, 1))

	tempLayout = [][]rune{
		[]rune(".##.##."),
		[]rune("#.#.#.#"),
		[]rune("##...##"),
		[]rune("...L..."),
		[]rune("##...##"),
		[]rune("#.#.#.#"),
		[]rune(".##.##."),
	}

	log.Println(countOccupiedSeatsLineOfSight(tempLayout, 3, 3))

	tempLayout = [][]rune{
		[]rune("#.##.##.##"),
		[]rune("#######.##"),
		[]rune("#.#.#..#.."),
		[]rune("####.##.##"),
		[]rune("#.##.##.##"),
		[]rune("#.#####.##"),
		[]rune("..#.#....."),
		[]rune("##########"),
		[]rune("#.######.#"),
		[]rune("#.#####.##"),
	}

	prediction, _ := predictSeatLayout(tempLayout, countOccupiedSeatsLineOfSight, 5)

	printLayout(prediction)

	log.Println(countOccupiedWhenPredictionStabilized(layout, countOccupiedSeatsLineOfSight, 5))
}

func partTwo(layout [][]rune) {
	log.Println(countOccupiedWhenPredictionStabilized(layout, countOccupiedSeatsLineOfSight, 5))
}

var directions = [][]int{
	{-1, -1},
	{-1, 0},
	{-1, 1},
	{0, -1},
	{0, 1},
	{1, -1},
	{1, 0},
	{1, 1},
}

func countOccupiedWhenPredictionStabilized(layout [][]rune, countSurroundingOccupiedSeats CountSurroundingOccupiedSeats, leaveSeatThreshold int) int {
	prediction := layout
	var stable bool

	for {
		prediction, stable = predictSeatLayout(prediction, countSurroundingOccupiedSeats, leaveSeatThreshold)

		if stable {
			break
		}
	}

	return countOccupiedSeats(prediction)
}

type CountSurroundingOccupiedSeats = func(layout [][]rune, y, x int) int

func predictSeatLayout(layout [][]rune, countSurroundingOccupiedSeats CountSurroundingOccupiedSeats, leaveSeatThreshold int) ([][]rune, bool) {
	prediction := make([][]rune, len(layout))
	stable := true

	for y, row := range layout {
		prediction[y] = make([]rune, len(row))

		for x, seatValue := range row {
			seat := string(seatValue)
			predictedSeat := seatValue
			occupiedCount := countSurroundingOccupiedSeats(layout, y, x)

			if seat == "L" && occupiedCount == 0 {
				predictedSeat = []rune("#")[0]
				stable = false
			}

			if seat == "#" && occupiedCount >= leaveSeatThreshold {
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

	for _, direction := range directions {
		modY := direction[0]
		modX := direction[1]

		if isSeatOccupied(layout, y+modY, x+modX) {
			occupiedCount++
		}
	}

	return occupiedCount
}

func countOccupiedSeatsLineOfSight(layout [][]rune, y, x int) int {
	occupiedCount := 0

	for _, direction := range directions {
		distance := 1
		modY := direction[0]
		modX := direction[1]

		for {
			seat, outOfBounds := getSeat(layout, y+modY*distance, x+modX*distance)
			distance++
			isOccupied := seat == "#"

			if isOccupied {
				occupiedCount++
			}

			if outOfBounds || isOccupied || seat == "L" {
				break
			}
		}
	}

	return occupiedCount
}

func getSeat(layout [][]rune, y, x int) (seat string, outOfBounds bool) {
	if y < 0 || y >= len(layout) {
		return "", true
	}

	row := layout[y]

	if x < 0 || x >= len(row) {
		return "", true
	}

	seat = string(row[x])

	return seat, false
}

func isSeatOccupied(layout [][]rune, y, x int) bool {
	seat, _ := getSeat(layout, y, x)

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

func printLayout(layout [][]rune) {
	for _, row := range layout {
		log.Println(string(row))
	}
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
