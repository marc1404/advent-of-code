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
	testInput := []TestInput{
		{"1 + 2 * 3 + 4 * 5 + 6", 71, 231},
		{"1 + (2 * 3) + (4 * (5 + 6))", 51, 51},
		{"2 * 3 + (4 * 5)", 26, 46},
		{"5 + (8 * 3 + 9 + 3 * 4 * 3)", 437, 1445},
		{"5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))", 12240, 669060},
		{"((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2", 13632, 23340},
	}
	input := readAllLines("./day18/input.txt")

	log.Println("Day 18 Part 01")
	testPartOne(testInput)
	partOne(input)

	log.Println()

	log.Println("Day 18 Part 02")
	testPartTwo(testInput)
	partTwo(input)
}

func testPartOne(testInput []TestInput) {
	test(testInput, false)
}

func partOne(expressions []string) {
	log.Println("Sum of resulting values:", sumOfExpressions(expressions, false))
}

func testPartTwo(testInput []TestInput) {
	test(testInput, true)
}

func partTwo(expressions []string) {
	log.Println("Sum of resulting values:", sumOfExpressions(expressions, true))
}

func sumOfExpressions(expressions []string, useAdvancedMath bool) int {
	sum := 0

	for _, expression := range expressions {
		sum += calculateExpression(expression, useAdvancedMath)
	}

	return sum
}

func calculateExpression(expression string, useAdvancedMath bool) int {
	expression = strings.ReplaceAll(expression, " ", "")
	pattern := regexp.MustCompile(`\([\d\+\*]+\)`)

	for {
		matches := pattern.FindAllStringSubmatch(expression, -1)

		if len(matches) == 0 {
			break
		}

		for _, match := range matches {
			wrappedExpression := match[0]
			subExpression := strings.TrimPrefix(wrappedExpression, "(")
			subExpression = strings.TrimSuffix(subExpression, ")")
			result := calculateSimpleExpression(subExpression, useAdvancedMath)
			expression = strings.Replace(expression, wrappedExpression, strconv.Itoa(result), 1)
		}
	}

	return calculateSimpleExpression(expression, useAdvancedMath)
}

func calculateSimpleExpression(expression string, useAdvancedMath bool) int {
	expression = strings.ReplaceAll(expression, " ", "")

	if useAdvancedMath {
		advancedExpression := rewriteToAdvancedMath(expression)

		if advancedExpression != expression {
			return calculateExpression(advancedExpression, true)
		}
	}

	result := 0
	expression = expression + "+"
	characters := strings.Split(expression, "")
	numberBuilder := strings.Builder{}
	operator := "+"

	for _, character := range characters {
		if character != "+" && character != "*" {
			numberBuilder.WriteString(character)
			continue
		}

		numberAsString := numberBuilder.String()

		numberBuilder.Reset()

		number, _ := strconv.Atoi(numberAsString)

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

func rewriteToAdvancedMath(expression string) string {
	if !strings.ContainsRune(expression, '*') {
		return expression
	}

	pattern := regexp.MustCompile(`(\d+(\+\d+))`)

	return pattern.ReplaceAllString(expression, `($1)`)
}

func test(testInput []TestInput, useAdvancedMath bool) {
	for _, test := range testInput {
		var expected int

		switch useAdvancedMath {
		case false:
			expected = test.expectedSimpleMath
		case true:
			expected = test.expectedAdvancedMath
		}

		actual := calculateExpression(test.expression, useAdvancedMath)

		log.Println(fmt.Sprintf("expected: %v, actual: %v", expected, actual))
	}
}

type TestInput struct {
	expression           string
	expectedSimpleMath   int
	expectedAdvancedMath int
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
