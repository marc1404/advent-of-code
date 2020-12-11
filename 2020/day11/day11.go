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

}

func partOne(lines []string) {

}

func testPartTwo() {

}

func partTwo(lines []string) {

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
