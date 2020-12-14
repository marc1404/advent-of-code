package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	testInstructions := readAllLines("./day14/test_input.txt")
	instructions := readAllLines("./day14/input.txt")

	log.Println("Day 14 Part 01")
	testPartOne(testInstructions)
	partOne(instructions)

	log.Println()

	log.Println("Day 14 Part 02")
	testPartTwo()
	partTwo(instructions)
}

func testPartOne(instructions []string) {
	log.Println(initializeAndSumFerryDockingProgram(instructions))
}

func partOne(instructions []string) {
	log.Println(initializeAndSumFerryDockingProgram(instructions))
}

func testPartTwo() {

}

func partTwo(instructions []string) {

}

func initializeAndSumFerryDockingProgram(instructions []string) int {
	memory := initializeFerryDockingProgram(instructions)
	sum := 0

	for _, value := range memory {
		sum += value
	}

	return sum
}

func initializeFerryDockingProgram(instructions []string) map[int]int {
	memory := make(map[int]int)
	var mask string

	for _, instruction := range instructions {
		target, value := parseInstruction(instruction)

		switch target {
		case "mask":
			mask = value
		default:
			memory = writeToMemoryThroughMask(target, value, mask, memory)
		}
	}

	return memory
}

func parseInstruction(instruction string) (target, value string) {
	parts := strings.Split(instruction, " = ")
	target = parts[0]
	value = parts[1]

	return target, value
}

func writeToMemoryThroughMask(target string, valueAsString string, mask string, memory map[int]int) map[int]int {
	memoryAddress := parseMemoryAddress(target)
	value, _ := strconv.Atoi(valueAsString)
	valueAsBinary := strconv.FormatInt(int64(value), 2)
	valueAsBinary = fmt.Sprintf("%036v", valueAsBinary)
	resultBits := make([]string, 36)

	for i := 0; i < 36; i++ {
		bit := string(valueAsBinary[i])
		maskBit := string(mask[i])
		resultBits[i] = applyMask(bit, maskBit)
	}

	resultAsBinary := strings.Join(resultBits, "")
	result, _ := strconv.ParseInt(resultAsBinary, 2, 64)
	memory[memoryAddress] = int(result)

	return memory
}

func applyMask(bit string, mask string) string {
	switch mask {
	case "0":
		return "0"
	case "1":
		return "1"
	case "X":
		return bit
	default:
		panic(fmt.Sprintf("Unexpected mask: %v!", mask))
	}
}

func parseMemoryAddress(target string) int {
	pattern := regexp.MustCompile("mem\\[([0-9]+)]")
	matches := pattern.FindAllStringSubmatch(target, -1)
	match := matches[0]
	memoryAddress, _ := strconv.Atoi(match[1])

	return memoryAddress
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
