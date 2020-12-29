package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	log.Println("Day 23 Part 01")
	testPartOne()
	partOne()

	log.Println()

	log.Println("Day 23 Part 02")
	testPartTwo()
	partTwo()
}

func testPartOne() {
}

func partOne() {
}

func testPartTwo() {
}

func partTwo() {
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
