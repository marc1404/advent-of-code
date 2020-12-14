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
	log.Println(initializeAndSumFerryDockingProgram(instructions, writeToMemoryThroughMask))
}

func partOne(instructions []string) {
	log.Println(initializeAndSumFerryDockingProgram(instructions, writeToMemoryThroughMask))
}

func testPartTwo() {
	instructions := []string{
		"mask = 000000000000000000000000000000X1001X",
		"mem[42] = 100",
		"mask = 00000000000000000000000000000000X0XX",
		"mem[26] = 1",
	}

	log.Println(initializeAndSumFerryDockingProgram(instructions, writeToMemoryWithFloatingBits))
}

func partTwo(instructions []string) {
	log.Println(initializeAndSumFerryDockingProgram(instructions, writeToMemoryWithFloatingBits))
}

func initializeAndSumFerryDockingProgram(instructions []string, writeToMemory WriteToMemory) int {
	memory := initializeFerryDockingProgram(instructions, writeToMemory)
	sum := 0

	for _, value := range memory {
		sum += value
	}

	return sum
}

type WriteToMemory func(target string, valueAsString string, mask string, memory map[int]int) map[int]int

func initializeFerryDockingProgram(instructions []string, writeToMemory WriteToMemory) map[int]int {
	memory := make(map[int]int)
	var mask string

	for _, instruction := range instructions {
		target, value := parseInstruction(instruction)

		switch target {
		case "mask":
			mask = value
		default:
			memory = writeToMemory(target, value, mask, memory)
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
	valueAsBinary := intToBinary(value)
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

func writeToMemoryWithFloatingBits(target string, valueAsString string, mask string, memory map[int]int) map[int]int {
	value, _ := strconv.Atoi(valueAsString)
	memoryAddress := parseMemoryAddress(target)
	memoryAsBinary := intToBinary(memoryAddress)

	for i := 0; i < 36; i++ {
		bit := string(memoryAsBinary[i])
		maskBit := string(mask[i])
		newBit := applyMaskWithFloatingBits(bit, maskBit)
		memoryAsBinary = replaceAtIndex(memoryAsBinary, []rune(newBit)[0], i)
	}

	permutations := getMemoryAddressPermutations(memoryAsBinary, []string{})

	for _, permutation := range permutations {
		address, _ := strconv.ParseInt(permutation, 2, 64)
		memory[int(address)] = value
	}

	return memory
}

func getMemoryAddressPermutations(memoryAsBinary string, permutations []string) []string {
	index := strings.Index(memoryAsBinary, "X")

	if index == -1 {
		permutations = append(permutations, memoryAsBinary)
		return permutations
	}

	zeroVariant := replaceAtIndex(memoryAsBinary, '0', index)
	oneVariant := replaceAtIndex(memoryAsBinary, '1', index)

	permutations = getMemoryAddressPermutations(zeroVariant, permutations)
	permutations = getMemoryAddressPermutations(oneVariant, permutations)

	return permutations
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

func applyMaskWithFloatingBits(bit string, mask string) string {
	switch mask {
	case "0":
		return bit
	case "1":
		return "1"
	case "X":
		return "X"
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

func intToBinary(value int) string {
	binary := strconv.FormatInt(int64(value), 2)

	return fmt.Sprintf("%036v", binary)
}

func replaceAtIndex(in string, r rune, i int) string {
	out := []rune(in)
	out[i] = r

	return string(out)
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
