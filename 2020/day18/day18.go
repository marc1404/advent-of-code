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
	lines := readAllLines("./day18/input.txt")

	log.Println("Day 18 Part 01")
	testPartOne()
	partOne(lines)

	log.Println()

	log.Println("Day 18 Part 02")
	testPartTwo()
	partTwo()
}

func testPartOne() {
	log.Println("expected: 71, actual:", calculateExpression("1 + 2 * 3 + 4 * 5 + 6"))
	log.Println("expected: 51, actual:", calculateExpression("1 + (2 * 3) + (4 * (5 + 6))"))
	log.Println("expected: 26, actual:", calculateExpression("2 * 3 + (4 * 5)"))
	log.Println("expected: 437, actual:", calculateExpression("5 + (8 * 3 + 9 + 3 * 4 * 3)"))
	log.Println("expected: 12240, actual:", calculateExpression("5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))"))
	log.Println("expected: 13632, actual:", calculateExpression("((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2"))

}

func partOne(lines []string) {
	sum := 0

	for _, line := range lines {
		sum += calculateExpression(line)
	}

	log.Println("Sum of resulting values:", sum)
}

func testPartTwo() {

}

func partTwo() {

}

func calculateExpression(expression string) int {
	pattern := regexp.MustCompile("\\([\\d\\s\\+\\*]+\\)")

	for {
		matches := pattern.FindAllStringSubmatch(expression, -1)

		if len(matches) == 0 {
			break
		}

		for _, match := range matches {
			wrappedExpression := match[0]
			subExpression := strings.TrimPrefix(wrappedExpression, "(")
			subExpression = strings.TrimSuffix(subExpression, ")")
			result := calculateSimpleExpression(subExpression)
			expression = strings.Replace(expression, wrappedExpression, strconv.Itoa(result), 1)
		}
	}

	return calculateSimpleExpression(expression)
}

func calculateSimpleExpression(expression string) int {
	result := 0
	expression = strings.ReplaceAll(expression, " ", "")
	expression = expression + "+"
	characters := strings.Split(expression, "")
	numberBuilder := strings.Builder{}
	operator := "+"
	number := 0

	for _, character := range characters {
		if character != "+" && character != "*" {
			numberBuilder.WriteString(character)
			continue
		}

		numberAsString := numberBuilder.String()

		numberBuilder.Reset()

		number, _ = strconv.Atoi(numberAsString)

		switch operator {
		case "+":
			result += number
		case "*":
			result *= number
		default:
			panic(fmt.Sprintf("Unexpected operator: %v!", operator))
		}

		operator = character
	}

	return result
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
