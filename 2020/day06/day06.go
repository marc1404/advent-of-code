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
	partTwo(groups)
}

func testPartOne() {
	log.Println(countYesInGroup([]string{
		"abcx",
		"abcy",
		"abcz",
	}))

	log.Println(countInGroups([][]string{
		{"abc"},
		{"a", "b", "c"},
		{"ab", "ac"},
		{"a", "a", "a", "a"},
		{"b"},
	}, countYesInGroup))
}

func partOne(groups [][]string) {
	log.Println("Questions answered with yes:", countInGroups(groups, countYesInGroup))
}

func partTwo(groups [][]string) {
	log.Println("Questions answered with yes in unison:", countInGroups(groups, countUnisonYesInGroup))
}

type CountInGroup func([]string) int

func countInGroups(groups [][]string, countInGroup CountInGroup) int {
	yesCount := 0

	for _, answers := range groups {
		yesCount += countInGroup(answers)
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

func countUnisonYesInGroup(answers []string) int {
	unisonYesCount := 0
	questionToYesCount := make(map[string]int)
	personCount := len(answers)

	for _, answer := range answers {
		for _, letterAsByte := range answer {
			letter := string(letterAsByte)
			questionToYesCount[letter]++

			if questionToYesCount[letter] == personCount {
				unisonYesCount++
			}
		}
	}

	return unisonYesCount
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
