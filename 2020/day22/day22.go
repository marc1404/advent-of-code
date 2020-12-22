package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	testLines := readAllLines("./day22/test_input.txt")
	lines := readAllLines("./day22/input.txt")

	testDecks := parseDecks(testLines)
	decks := parseDecks(lines)

	log.Println("Day 22 Part 01")
	testPartOne(testDecks)
	partOne(decks)

	log.Println()

	log.Println("Day 22 Part 02")
	testPartTwo()
	partTwo()
}

func testPartOne(decks [][]int) {
	deck := playCombat(decks[0], decks[1])

	log.Println("expected: 306, actual:", calculateScore(deck))
}

func partOne(decks [][]int) {
	deck := playCombat(decks[0], decks[1])

	log.Println("What is the winning player's score?")
	log.Println("Answer:", calculateScore(deck))
}

func testPartTwo() {
}

func partTwo() {
}

func playCombat(deck1, deck2 []int) []int {
	for {
		card1 := deck1[0]
		card2 := deck2[0]

		if card1 > card2 {
			deck1 = append(deck1[1:], card1, card2)
			deck2 = deck2[1:]
		}

		if card2 > card1 {
			deck1 = deck1[1:]
			deck2 = append(deck2[1:], card2, card1)
		}

		if len(deck1) == 0 {
			return deck2
		}

		if len(deck2) == 0 {
			return deck1
		}
	}
}

func parseDecks(lines []string) [][]int {
	var decks [][]int
	var deck []int

	for _, line := range lines {
		if strings.HasPrefix(line, "Player ") {
			continue
		}

		if line == "" {
			decks = append(decks, deck)
			deck = make([]int, 0)
			continue
		}

		card, _ := strconv.Atoi(line)
		deck = append(deck, card)
	}

	if len(deck) > 0 {
		decks = append(decks, deck)
	}

	return decks
}

func calculateScore(deck []int) int {
	score := 0
	cardCount := len(deck)

	for index, card := range deck {
		score += card * (cardCount - index)
	}

	return score
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
		lines = append(lines, line)
	}

	return lines
}
