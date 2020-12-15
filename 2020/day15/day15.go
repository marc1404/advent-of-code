package main

import "log"

func main() {
	log.Println("Day 15 Part 01")
	testPartOne()
	partOne()

	log.Println()

	log.Println("Day 15 Part 02")
	testPartTwo()
	partTwo()
}

func testPartOne() {
	log.Println(playMemoryUntilTurn([]int{0, 3, 6}, 2020))
	log.Println(playMemoryUntilTurn([]int{1, 3, 2}, 2020))
	log.Println(playMemoryUntilTurn([]int{2, 1, 3}, 2020))
	log.Println(playMemoryUntilTurn([]int{1, 2, 3}, 2020))
	log.Println(playMemoryUntilTurn([]int{2, 3, 1}, 2020))
	log.Println(playMemoryUntilTurn([]int{3, 2, 1}, 2020))
	log.Println(playMemoryUntilTurn([]int{3, 1, 2}, 2020))
}

func partOne() {
	log.Println("The 2020th number spoken is:", playMemoryUntilTurn([]int{2, 15, 0, 9, 1, 20}, 2020))
}

func testPartTwo() {
	log.Println(playMemoryUntilTurn([]int{0, 3, 6}, 30000000))
	log.Println(playMemoryUntilTurn([]int{1, 3, 2}, 30000000))
	log.Println(playMemoryUntilTurn([]int{2, 1, 3}, 30000000))
	log.Println(playMemoryUntilTurn([]int{1, 2, 3}, 30000000))
	log.Println(playMemoryUntilTurn([]int{2, 3, 1}, 30000000))
	log.Println(playMemoryUntilTurn([]int{3, 2, 1}, 30000000))
	log.Println(playMemoryUntilTurn([]int{3, 1, 2}, 30000000))
}

func partTwo() {
	log.Println("The 30000000th number spoken is:", playMemoryUntilTurn([]int{2, 15, 0, 9, 1, 20}, 30000000))
}

func playMemoryUntilTurn(numbers []int, stopAfterTurn int) int {
	numberToMemory := make(map[int][]int)

	for i, number := range numbers {
		updateMemory(numberToMemory, number, i+1)
	}

	turn := len(numbers) + 1

	for {
		lastNumber := numbers[turn-2]
		memory := numberToMemory[lastNumber]
		number := determineSpokenNumber(memory)
		numbers = append(numbers, number)

		updateMemory(numberToMemory, number, turn)

		turn++

		if turn > stopAfterTurn {
			return number
		}
	}
}

func determineSpokenNumber(memory []int) int {
	memoryLength := len(memory)

	if memoryLength == 1 {
		return 0
	}

	return memory[memoryLength-1] - memory[memoryLength-2]
}

func updateMemory(numberToMemory map[int][]int, number, turn int) {
	numberToMemory[number] = append(numberToMemory[number], turn)
}
