package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	groups := readGroups("./day06/input.txt")

	log.Println("Day 06 Part 01")
	testPartOne()
	partOne(groups)

	log.Println()

	log.Println("Day 06 Part 02")
	partTwo()
}

func testPartOne() {
	log.Println(countYesInGroup([]string{
		"abcx",
		"abcy",
		"abcz",
	}))

	log.Println(countYesInGroups([][]string{
		{"abc"},
		{"a", "b", "c"},
		{"ab", "ac"},
		{"a", "a", "a", "a"},
		{"b"},
	}))
}

func partOne(groups [][]string) {
	log.Println("Questions answered with yes:", countYesInGroups(groups))
}

func partTwo() {

}

func countYesInGroups(groups [][]string) int {
	yesCount := 0

	for _, answers := range groups {
		yesCount += countYesInGroup(answers)
	}

	return yesCount
}

func countYesInGroup(answers []string) int {
	questionMap := make(map[string]bool)

	for _, answer := range answers {
		for _, letterAsByte := range answer {
			letter := string(letterAsByte)
			questionMap[letter] = true
		}
	}

	return len(questionMap)
}

func readGroups(filePath string) [][]string {
	file, err := os.Open(filePath)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var groups [][]string
	var answers []string

	for scanner.Scan() {
		line := scanner.Text()
		log.Println(line)
		if line == "" {
			groups = append(groups, answers)
			answers = []string{}
			continue
		}

		answers = append(answers, line)
	}

	groups = append(groups, answers)

	return groups
}
