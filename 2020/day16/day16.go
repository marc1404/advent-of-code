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
	testLines1 := readAllLines("./day16/test_input_1.txt")
	testLines2 := readAllLines("./day16/test_input_2.txt")
	lines := readAllLines("./day16/input.txt")

	testInput1 := parseInput(testLines1)
	testInput2 := parseInput(testLines2)
	input := parseInput(lines)

	log.Println("Day 16 Part 01")
	testPartOne(testInput1)
	partOne(input)

	log.Println()

	log.Println("Day 16 Part 02")
	testPartTwo(testInput2)
	partTwo(input)
}

func testPartOne(input Input) {
	log.Println("expected: 71, actual:", determineTicketScanningErrorRate(input))
}

func partOne(input Input) {
	log.Println("Ticket scanning error rate:", determineTicketScanningErrorRate(input))
}

func testPartTwo(input Input) {
	input.solvePartTwo()
}

func partTwo(input Input) {
	input.solvePartTwo()
}

func (input Input) solvePartTwo() {
	height := len(input.yourTicket.fields)
	width := len(input.nearbyTickets) + 1
	columns := make([]Column, height)

	for i, value := range input.yourTicket.fields {
		column := Column{i, make([]int, width), make([]string, 0)}
		column.values[0] = value
		columns[i] = column
	}

	for x, ticket := range input.nearbyTickets {
		for y, value := range ticket.fields {
			column := columns[y]
			column.values[x+1] = value
		}
	}

	for i, column := range columns {
		for _, field := range input.fields {
			valid := true

			for _, value := range column.values {
				if !field.isValueValid(value) {
					valid = false
					break
				}
			}

			if !valid {
				continue
			}

			column.possibleFields = append(column.possibleFields, field.name)
			columns[i] = column
		}
	}

	cols := make([]Column, 0)

	for _, column := range columns {
		if len(column.possibleFields) > 0 {
			cols = append(cols, column)
		}
	}

	return
}

func determineTicketScanningErrorRate(input Input) int {
	ticketScanningErrorRate := 0

	for _, ticket := range input.nearbyTickets {
		invalidValues := ticket.getInvalidValues(input.fields)

		for _, invalidValue := range invalidValues {
			ticketScanningErrorRate += invalidValue
		}
	}

	return ticketScanningErrorRate
}

type Column struct {
	index          int
	values         []int
	possibleFields []string
}

type Span struct {
	min int
	max int
}

type Field struct {
	name  string
	spans []Span
}

type Ticket struct {
	fields []int
}

type Input struct {
	fields        []Field
	yourTicket    Ticket
	nearbyTickets []Ticket
}

func (span Span) isValid(value int) bool {
	return value >= span.min && value <= span.max
}

func (field Field) isValueValid(value int) bool {
	for _, span := range field.spans {
		if span.isValid(value) {
			return true
		}
	}

	return false
}

func (ticket Ticket) getInvalidValues(fields []Field) []int {
	var invalidValues []int

	for _, value := range ticket.fields {
		isValueValid := false

		for _, field := range fields {
			if field.isValueValid(value) {
				isValueValid = true
			}
		}

		if !isValueValid {
			invalidValues = append(invalidValues, value)
		}
	}

	return invalidValues
}

func parseInput(lines []string) Input {
	mode := "fields"
	var fields []Field
	var yourTicket Ticket
	var nearbyTickets []Ticket

	for _, line := range lines {
		switch line {
		case "your ticket:":
			mode = "your_ticket"
			continue
		case "nearby tickets:":
			mode = "nearby_tickets"
			continue
		}

		switch mode {
		case "fields":
			fields = append(fields, parseField(line))
		case "your_ticket":
			yourTicket = parseTicket(line)
		case "nearby_tickets":
			nearbyTickets = append(nearbyTickets, parseTicket(line))
		}
	}

	return Input{fields, yourTicket, nearbyTickets}
}

func parseSpan(min, max string) Span {
	return Span{
		min: parseInt(min),
		max: parseInt(max),
	}
}

func parseField(line string) Field {
	pattern := regexp.MustCompile("^(.+): (\\d+)-(\\d+) or (\\d+)-(\\d+)$")
	matches := pattern.FindAllStringSubmatch(line, -1)
	match := matches[0]

	if len(match) != 6 {
		panic(fmt.Sprintf("Could not parse field: %v!", line))
	}

	name := match[1]
	spans := []Span{
		parseSpan(match[2], match[3]),
		parseSpan(match[4], match[5]),
	}

	return Field{name, spans}
}

func parseTicket(line string) Ticket {
	parts := strings.Split(line, ",")
	fields := make([]int, len(parts))

	for i, part := range parts {
		fields[i] = parseInt(part)
	}

	return Ticket{fields}
}

func parseInt(input string) int {
	output, _ := strconv.Atoi(input)

	return output
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
