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
	testLines := readAllLines("./day19/test_input.txt")
	lines := readAllLines("./day19/input.txt")

	log.Println("Day 19 Part 01")
	testPartOne(testLines)
	partOne(lines)

	log.Println()

	log.Println("Day 19 Part 02")
	testPartTwo()
	partTwo()
}

func testPartOne(lines []string) {
	rules, messages := parseRulesAndMessage(lines)

	log.Println("expected: 2, actual:", rules[0].countValidMessages(messages))
}

func partOne(lines []string) {
	rules, messages := parseRulesAndMessage(lines)

	log.Println("How many messages completely match rule 0?")
	log.Println("Answer:", rules[0].countValidMessages(messages))
}

func testPartTwo() {
}

func partTwo() {
}

func parseRulesAndMessage(lines []string) ([]*Rule, []string) {
	mode := "rules"
	var ruleLines []string
	var messages []string

	for _, line := range lines {
		if line == "" {
			switch mode {
			case "rules":
				mode = "messages"
				continue
			case "messages":
				break
			default:
				panic(fmt.Sprintf("Unexpected mode: %v!", mode))
			}
		}

		switch mode {
		case "rules":
			ruleLines = append(ruleLines, line)
		case "messages":
			messages = append(messages, line)
		default:
			panic(fmt.Sprintf("Unexpected mode: %v!", mode))
		}
	}

	rules := parseRules(ruleLines)

	return rules, messages
}

func parseRules(lines []string) []*Rule {
	rules := make([]*Rule, len(lines))

	for _, line := range lines {
		parts := strings.Split(line, ": ")
		index, _ := strconv.Atoi(parts[0])
		pattern := parts[1]
		rule := &Rule{index, pattern, nil}
		rules[index] = rule

		rule.preProcessing()
	}

	for {
		shouldContinue := false

		for _, rule := range rules {
			patternChanged := rule.replaceRules(rules)

			if patternChanged {
				shouldContinue = true
			}
		}

		if !shouldContinue {
			break
		}
	}

	for _, rule := range rules {
		rule.postProcessing()
	}

	return rules
}

type Rule struct {
	index   int
	pattern string
	regex   *regexp.Regexp
}

func (rule *Rule) preProcessing() {
	rule.pattern = strings.ReplaceAll(rule.pattern, `"`, "")
}

func (rule *Rule) replaceRules(rules []*Rule) bool {
	regex := regexp.MustCompile(`(\d+)`)
	newPattern := regex.ReplaceAllStringFunc(rule.pattern, func(match string) string {
		index, _ := strconv.Atoi(match)
		replacement := rules[index].pattern

		return fmt.Sprintf("(%v)", replacement)
	})

	if newPattern == rule.pattern {
		return false
	}

	rule.pattern = newPattern

	return true
}

func (rule *Rule) postProcessing() {
	rule.pattern = strings.ReplaceAll(rule.pattern, " ", "")
	rule.pattern = fmt.Sprintf(`^%v$`, rule.pattern)
	rule.regex = regexp.MustCompile(rule.pattern)
}

func (rule *Rule) countValidMessages(messages []string) int {
	validCount := 0

	for _, message := range messages {
		if rule.regex.MatchString(message) {
			validCount++
		}
	}

	return validCount
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
