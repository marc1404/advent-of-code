package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

func main() {
	lines := readAllLines("./day05/input.txt")

	log.Println("Day 05 Part 01")
	testPartOne()
	partOne(lines)

	log.Println()

	log.Println("Day 05 Part 02")
	partTwo(lines)
}

func testPartOne() {
	log.Println(parseBoardingPass("FBFBBFFRLR"))
	log.Println(parseBoardingPass("BFFFBBFRRR"))
	log.Println(parseBoardingPass("FFFBBBFRRR"))
	log.Println(parseBoardingPass("BBFFBBFRLL"))
}

func partOne(boardingPasses []string) {
	var highestSeatId int

	for _, boardingPass := range boardingPasses {
		_, _, seatId := parseBoardingPass(boardingPass)

		if seatId > highestSeatId {
			highestSeatId = seatId
		}
	}

	log.Println("Highest seat ID on a boarding pass is:", highestSeatId)
}

func partTwo(boardingPasses []string) {
	seatIdMap := make(map[int]bool)
	min := int(^uint(0) >> 1)
	max := 0

	for _, boardingPass := range boardingPasses {
		_, _, seatId := parseBoardingPass(boardingPass)
		seatIdMap[seatId] = true

		if seatId < min {
			min = seatId
		}

		if seatId > max {
			max = seatId
		}
	}

	for seatId := min + 1; seatId < max; seatId++ {
		if seatIdMap[seatId-1] && !seatIdMap[seatId] && seatIdMap[seatId+1] {
			log.Println("My seat ID:", seatId)
		}
	}
}

func parseBoardingPass(boardingPass string) (row, column, seatId int) {
	rowPartition := boardingPass[0:7]
	columnPartition := boardingPass[7:10]

	row = doBinarySpacePartitioning(rowPartition, PartitionLetters{"F", "B"}, Span{0, 127})
	column = doBinarySpacePartitioning(columnPartition, PartitionLetters{"L", "R"}, Span{0, 7})

	seatId = row*8 + column

	return row, column, seatId
}

func doBinarySpacePartitioning(partition string, partitionLetters PartitionLetters, span Span) int {
	for _, letterAsByte := range partition {
		letter := string(letterAsByte)
		delta := int(math.Round(float64((span.max - span.min + 1) / 2)))

		if letter == partitionLetters.lower {
			span.max -= delta
			continue
		}

		if letter == partitionLetters.upper {
			span.min += delta
			continue
		}

		panic("Unexpected letter: " + letter)
	}

	if span.min != span.max {
		panic(fmt.Sprintf("Span min %v != max %v", span.min, span.max))
	}

	return span.min
}

type PartitionLetters struct {
	lower string
	upper string
}

type Span struct {
	min int
	max int
}

func readAllLines(filePath string) []string {
	file, err := os.Open(filePath)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := make([]string, 0)

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			continue
		}

		lines = append(lines, line)
	}

	return lines
}
