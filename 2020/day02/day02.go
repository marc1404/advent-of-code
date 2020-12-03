package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	lines := readAllLines("./day02/input.txt")

	log.Println("Day 02 Part 01")
	testPartOne()
	partOne(lines)

	log.Println()

	log.Println("Day 02 Part 02")
	testPartTwo()
	partTwo(lines)
}

func testPartOne() {
	lines := []string{
		"1-3 a: abcde",
		"1-3 b: cdefg",
		"2-9 c: ccccccccc",
	}

	for _, policyAndPassword := range lines {
		log.Println(policyAndPassword, isPasswordValidSledRental(policyAndPassword))
	}
}

func partOne(lines []string) {
	validCount := 0

	for _, line := range lines {
		if line == "" {
			continue
		}

		isValid := isPasswordValidSledRental(line)

		if isValid {
			validCount = validCount + 1
		}
	}

	log.Println("Valid passwords", validCount)
}

func testPartTwo() {
	lines := []string{
		"1-3 a: abcde",
		"1-3 b: cdefg",
		"2-9 c: ccccccccc",
	}

	for _, policyAndPassword := range lines {
		log.Println(policyAndPassword, isPasswordValidTobogganCorporatePolicy(policyAndPassword))
	}
}

func partTwo(lines []string) {
	validCount := 0

	for _, line := range lines {
		if line == "" {
			continue
		}

		isValid := isPasswordValidTobogganCorporatePolicy(line)

		if isValid {
			validCount = validCount + 1
		}
	}

	log.Println("Valid passwords", validCount)
}

func isPasswordValidSledRental(policyAndPassword string) bool {
	policyMin, policyMax, policyLetter, password := parsePolicyAndPassword(policyAndPassword)
	occurrences := strings.Count(password, policyLetter)

	if occurrences < policyMin {
		return false
	}

	if occurrences > policyMax {
		return false
	}

	return true
}

func isPasswordValidTobogganCorporatePolicy(policyAndPassword string) bool {
	policyLeft, policyRight, policyLetter, password := parsePolicyAndPassword(policyAndPassword)
	policyLeft = policyLeft - 1
	policyRight = policyRight - 1
	hasLeft := string(password[policyLeft]) == policyLetter
	hasRight := string(password[policyRight]) == policyLetter

	return hasLeft != hasRight
}

func parsePolicyAndPassword(policyAndPassword string) (policyMin int, policyMax int, policyLetter string, password string) {
	parts := strings.Split(policyAndPassword, ": ")
	policy := parts[0]
	password = parts[1]
	policyMin, policyMax, policyLetter = parsePolicy(policy)

	return policyMin, policyMax, policyLetter, password
}

func parsePolicy(policy string) (policyMin int, policyMax int, policyLetter string) {
	parts := strings.Split(policy, " ")
	policyMinMax := parts[0]
	policyLetter = parts[1]
	minMaxParts := strings.Split(policyMinMax, "-")
	policyMin, _ = strconv.Atoi(minMaxParts[0])
	policyMax, _ = strconv.Atoi(minMaxParts[1])

	return policyMin, policyMax, policyLetter
}

func readAllLines(filePath string) []string {
	file, err := os.Open(filePath)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := make([]string, 1)

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	return lines
}
