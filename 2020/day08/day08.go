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
	lines := readAllLines("./day08/input.txt")
	stateMachine := createStateMachine(lines)

	log.Println("Day 08 Part 01")
	testPartOne()
	partOne(stateMachine)

	log.Println()

	log.Println("Day 08 Part 02")
	partTwo(lines)
}

func testPartOne() {
	instructions := []string{
		"nop +0",
		"acc +1",
		"jmp +4",
		"acc +3",
		"jmp -3",
		"acc -99",
		"acc +1",
		"jmp -4",
		"acc +6",
	}

	stateMachine := createStateMachine(instructions)

	stateMachine.runStateMachine(true)
	log.Println(stateMachine.accumulator)
}

func partOne(stateMachine *StateMachine) {
	stateMachine.runStateMachine(true)
	log.Println("Accumulator:", stateMachine.accumulator)
}

func partTwo(lines []string) {
	for pointer, line := range lines {
		var replaceOld, replaceNew string

		if strings.HasPrefix(line, "jmp") {
			replaceOld = "jmp"
			replaceNew = "nop"
		}

		if strings.HasPrefix(line, "nop") {
			replaceOld = "nop"
			replaceNew = "jmp"
		}

		if replaceOld == "" || replaceNew == "" {
			continue
		}

		instructions := make([]string, len(lines))

		copy(instructions, lines)

		instructions[pointer] = strings.Replace(instructions[pointer], replaceOld, replaceNew, 1)
		stateMachine := createStateMachine(instructions)
		exitCode := stateMachine.runStateMachine(true)

		if exitCode == 0 {
			log.Println("Accumulator:", stateMachine.accumulator)
		}
	}
}

type StateMachine struct {
	instructions       []string
	accumulator        int
	pointer            int
	instructionToCalls map[int]int
}

func createStateMachine(instructions []string) *StateMachine {
	return &StateMachine{
		instructions,
		0,
		0,
		make(map[int]int),
	}
}

func (stateMachine *StateMachine) runStateMachine(preventRepeatedInstruction bool) int {
	pointer := stateMachine.pointer

	if pointer == len(stateMachine.instructions) {
		return 0
	}

	if preventRepeatedInstruction && stateMachine.instructionToCalls[pointer] == 1 {
		return -1
	}

	stateMachine.instructionToCalls[pointer]++
	instruction := stateMachine.instructions[stateMachine.pointer]
	operation, argument := parseInstruction(instruction)

	switch operation {
	case "acc":
		stateMachine.acc(argument)
	case "jmp":
		stateMachine.jmp(argument)
	case "nop":
		stateMachine.nop()
	default:
		panic(fmt.Sprintf("Unknown operation: %v!", operation))
	}

	return stateMachine.runStateMachine(preventRepeatedInstruction)
}

func (stateMachine *StateMachine) acc(argument int) {
	stateMachine.accumulator += argument
	stateMachine.pointer++
}

func (stateMachine *StateMachine) jmp(argument int) {
	stateMachine.pointer += argument
}

func (stateMachine *StateMachine) nop() {
	stateMachine.pointer++
}

func parseInstruction(instruction string) (operation string, argument int) {
	parts := strings.Split(instruction, " ")
	operation = parts[0]
	argument, _ = strconv.Atoi(parts[1])

	return operation, argument
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
