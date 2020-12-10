package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	testNumbers := []int{
		35,
		20,
		15,
		25,
		47,
		40,
		62,
		55,
		65,
		95,
		102,
		117,
		150,
		182,
		127,
		219,
		299,
		277,
		309,
		576,
	}

	numbers := readLinesAsNumbers("./day09/input.txt")

	log.Println("Day 09 Part 01")
	testPartOne(testNumbers)
	partOne(numbers)

	log.Println()

	log.Println("Day 09 Part 02")
	testPartTwo(testNumbers)
	partTwo(numbers)
}

func testPartOne(numbers []int) {
	verifyExchangeMaskingAdditionSystem(numbers, 5)
}

func partOne(numbers []int) {
	verifyExchangeMaskingAdditionSystem(numbers, 25)
}

func testPartTwo(numbers []int) {
	findEncryptionWeakness(numbers, 127)
}

func partTwo(numbers []int) {
	findEncryptionWeakness(numbers, 20874512)
}

func verifyExchangeMaskingAdditionSystem(numbers []int, preambleLength int) {
	if len(numbers) <= preambleLength {
		return
	}

	preamble := numbers[:preambleLength]
	checksum := numbers[preambleLength]
	pairs := findAdditionPairs(preamble, checksum)

	if len(pairs) == 0 {
		log.Println("Invalid checksum:", checksum)
	}

	verifyExchangeMaskingAdditionSystem(numbers[1:], preambleLength)
}

func findAdditionPairs(preamble []int, checksum int) (additionPairs [][]int) {
	for i, a := range preamble {
		for j, b := range preamble {
			if i == j {
				continue
			}

			if (a + b) == checksum {
				additionPairs = append(additionPairs, []int{a, b})
			}
		}
	}

	return additionPairs
}

func findEncryptionWeakness(numbers []int, invalidChecksum int) {
	for i, _ := range numbers {
		sum := 0

		for j, number := range numbers[i:] {
			sum += number

			if sum > invalidChecksum {
				break
			}

			if j == 0 {
				continue
			}

			if sum == invalidChecksum {
				contiguousRange := numbers[i : i+j+1]
				min, max := findMinMax(contiguousRange)
				encryptionWeakness := min + max

				log.Printf("Contiguous range: %v, Smallest: %v, Largest: %v, Encryption weakness: %v", contiguousRange, min, max, encryptionWeakness)
			}
		}
	}
}

func findMinMax(values []int) (min, max int) {
	min = int(^uint(0) >> 1)

	for _, value := range values {
		if value < min {
			min = value
		}

		if value > max {
			max = value
		}
	}

	return min, max
}

func readLinesAsNumbers(filePath string) []int {
	file, err := os.Open(filePath)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	numbers := make([]int, 0)

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			continue
		}

		number, _ := strconv.Atoi(line)
		numbers = append(numbers, number)
	}

	return numbers
}
