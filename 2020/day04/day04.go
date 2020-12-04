package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	lines := readAllLines("./day04/input.txt")
	passports := parsePassports(lines)

	log.Println("Day 04 Part 01")
	testPartOne()
	partOne(passports)

	log.Println("")

	log.Println("Day 04 Part 02")
	testPartTwo()
	partTwo(passports)
}

func testPartOne() {
	passports := parsePassports([]string{
		"ecl:gry pid:860033327 eyr:2020 hcl:#fffffd",
		"byr:1937 iyr:2017 cid:147 hgt:183cm",
		"",
		"iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884",
		"hcl:#cfa07d byr:1929",
		"",
		"hcl:#ae17e1 iyr:2013",
		"eyr:2024",
		"ecl:brn pid:760753108 byr:1931",
		"hgt:179cm",
		"",
		"hcl:#cfa07d eyr:2025 pid:166559648",
		"iyr:2011 ecl:brn hgt:59in",
	})

	for _, passport := range passports {
		isValid := validatePassport(passport, false)

		log.Println(passport, isValid)
	}
}

func partOne(passports []Passport) {
	log.Println("Valid passports:", countValidPassports(passports, false))
}

func testPartTwo() {
	log.Println("expect byr:2002 to be true:", validateField("byr", "2002"))
	log.Println("expect byr:2003 to be true:", validateField("byr", "2003"))

	log.Println("expect hgt:60in to be true:", validateField("hgt", "60in"))
	log.Println("expect hgt:190cm to be true:", validateField("hgt", "190cm"))
	log.Println("expect hgt:190in to be false:", validateField("hgt", "190in"))
	log.Println("expect hgt:190 to be false:", validateField("hgt", "190"))

	log.Println("expect hcl:#123abc to be true:", validateField("hcl", "#123abc"))
	log.Println("expect hcl:#123abz to be false:", validateField("hcl", "#123abz"))
	log.Println("expect hcl:123abc to be false:", validateField("hcl", "123abc"))

	log.Println("expect ecl:brn to be true:", validateField("ecl", "brn"))
	log.Println("expect ecl:wat to be false:", validateField("ecl", "wat"))

	log.Println("expect pid:000000001 to be true:", validateField("pid", "000000001"))
	log.Println("expect pid:0123456789 to be false:", validateField("pid", "0123456789"))
}

func partTwo(passports []Passport) {
	log.Println("Valid passports:", countValidPassports(passports, true))
}

type Passport struct {
	Fields []string
}

func countValidPassports(passports []Passport, beStrict bool) int {
	var validCount int

	for _, passport := range passports {
		isValid := validatePassport(passport, beStrict)

		if isValid {
			validCount++
		}
	}

	return validCount
}

func validatePassport(passport Passport, beStrict bool) bool {
	requiredFields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	actualFieldMap := make(map[string]string)

	for _, actualField := range passport.Fields {
		fieldParts := strings.Split(actualField, ":")
		fieldName := fieldParts[0]
		fieldValue := fieldParts[1]
		actualFieldMap[fieldName] = fieldValue
	}

	for _, requiredField := range requiredFields {
		fieldValue := actualFieldMap[requiredField]

		if fieldValue == "" {
			return false
		}

		if !beStrict {
			continue
		}

		if !validateField(requiredField, fieldValue) {
			return false
		}
	}

	return true
}

func validateField(fieldName string, fieldValue string) bool {
	if fieldName == "byr" {
		byr, err := strconv.Atoi(fieldValue)

		return err == nil && byr >= 1920 && byr <= 2002
	}

	if fieldName == "iyr" {
		iyr, err := strconv.Atoi(fieldValue)

		return err == nil && iyr >= 2010 && iyr <= 2020
	}

	if fieldName == "eyr" {
		eyr, err := strconv.Atoi(fieldValue)

		return err == nil && eyr >= 2020 && eyr <= 2030
	}

	if fieldName == "hgt" {
		pattern := regexp.MustCompile("([0-9]+)(cm|in)")
		matches := pattern.FindAllStringSubmatch(fieldValue, -1)

		if len(matches) != 1 {
			return false
		}

		match := matches[0]

		if len(match) != 3 {
			return false
		}

		height, err := strconv.Atoi(match[1])

		if err != nil {
			return false
		}

		unit := match[2]

		if unit == "cm" {
			return height >= 150 && height <= 193
		}

		if unit == "in" {
			return height >= 59 && height <= 76
		}

		return false
	}

	if fieldName == "hcl" {
		matched, err := regexp.MatchString("#[0-9a-f]{6}", fieldValue)

		return err == nil && matched
	}

	if fieldName == "ecl" {
		matched, err := regexp.MatchString("(amb|blu|brn|gry|grn|hzl|oth)", fieldValue)

		return err == nil && matched
	}

	if fieldName == "pid" {
		matched, err := regexp.MatchString(`^(\d){9}$`, fieldValue)

		return err == nil && matched
	}

	if fieldName == "cid" {
		return true
	}

	return false
}

func parsePassports(lines []string) []Passport {
	var passports []Passport
	var passport Passport

	for _, line := range lines {
		if line == "" {
			passports = append(passports, passport)
			passport = Passport{}
			continue
		}

		passport.Fields = append(passport.Fields, strings.Split(line, " ")...)
	}

	return passports
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
		lines = append(lines, line)
	}

	return lines
}
