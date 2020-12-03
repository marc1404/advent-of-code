package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	lines := readAllLines("./day03/input.txt")

	log.Println("Day 03 Part 01")
	partOne(lines)

	log.Println()

	log.Println("Day 03 Part 02")
	partTwo(lines)
}

func partOne(lines []string) {
	treeCount := countTrees(lines, 3, 1)

	log.Println("Counted", treeCount, "trees")
}

type Slope struct {
	modX int
	modY int
}

func partTwo(lines []string) {
	treeProduct := 1
	slopes := []Slope{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}

	for _, slope := range slopes {
		treeCount := countTrees(lines, slope.modX, slope.modY)
		treeProduct = treeProduct * treeCount
	}

	log.Println("Tree product:", treeProduct)
}

func countTrees(lines []string, modX int, modY int) int {
	x := 0
	y := 0
	treeCount := 0

	for {
		square := readCoordinate(lines, x, y)

		if square == "" {
			break
		}

		if square == "#" {
			treeCount = treeCount + 1
		}

		x = x + modX
		y = y + modY
	}

	return treeCount
}

func readCoordinate(lines []string, x, y int) string {
	if y >= len(lines) {
		return ""
	}

	line := lines[y]
	x = clamp(x, len(line))

	return string(line[x])
}

func clamp(x, max int) int {
	if x < max || x == 0 {
		return x
	}

	return clamp(x-max, max)
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
