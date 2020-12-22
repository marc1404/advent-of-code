package main

import (
	"bufio"
	"fmt"
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

	testDeck1 := testDecks[0]
	testDeck2 := testDecks[1]

	deck1 := decks[0]
	deck2 := decks[1]

	log.Println("Day 22 Part 01")
	testPartOne(testDeck1, testDeck2)
	partOne(deck1, deck2)

	log.Println()

	log.Println("Day 22 Part 02")
	testPartTwo(testDeck1, testDeck2)
	partTwo(deck1, deck2)
}

func testPartOne(deck1, deck2 []int) {
	deck := playCombat(deck1, deck2)

	log.Println("expected: 306, actual:", calculateScore(deck))
}

func partOne(deck1, deck2 []int) {
	deck := playCombat(deck1, deck2)

	log.Println("What is the winning player's score?")
	log.Println("Answer:", calculateScore(deck))
}

func testPartTwo(deck1, deck2 []int) {
	_, deck := playRecursiveCombat(deck1, deck2)

	log.Println("expected: 291, actual:", calculateScore(deck))
}

func partTwo(deck1, deck2 []int) {
	_, deck := playRecursiveCombat(deck1, deck2)

	log.Println("What is the winning player's score?")
	log.Println("Answer:", calculateScore(deck))
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

func playRecursiveCombat(deck1, deck2 []int) (int, []int) {
	infinityPreventionMemory := make(map[string]bool)

	for {
		infinityPreventionKey := fmt.Sprintf("%v%v", deck1, deck2)
		hasSameConfiguration, _ := infinityPreventionMemory[infinityPreventionKey]

		if hasSameConfiguration {
			return 1, deck1
		}

		infinityPreventionMemory[infinityPreventionKey] = true

		card1 := deck1[0]
		card2 := deck2[0]

		deck1 = deck1[1:]
		deck2 = deck2[1:]

		shouldRecurse := len(deck1) >= card1 && len(deck2) >= card2
		winner := 0

		switch shouldRecurse {
		case true:
			subDeck1 := append([]int{}, deck1[0:card1]...)
			subDeck2 := append([]int{}, deck2[0:card2]...)
			winner, _ = playRecursiveCombat(subDeck1, subDeck2)
		case false:
			if card1 > card2 {
				winner = 1
			}

			if card2 > card1 {
				winner = 2
			}
		}

		switch winner {
		case 1:
			deck1 = append(deck1, card1, card2)
		case 2:
			deck2 = append(deck2, card2, card1)
		default:
			panic(fmt.Sprintf("Unexpected winner: %v!", winner))
		}

		if len(deck1) == 0 {
			return 2, deck2
		}

		if len(deck2) == 0 {
			return 1, deck1
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
