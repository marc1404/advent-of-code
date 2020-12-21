package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func main() {
	testLines := readAllLines("./day21/test_input.txt")
	lines := readAllLines("./day21/input.txt")

	log.Println("Day 21 Part 01")
	testPartOne(testLines)
	partOne(lines)

	log.Println()

	log.Println("Day 21 Part 02")
	testPartTwo()
	partTwo()
}

func testPartOne(lines []string) {
	input := parseInput(lines)
	ingredients := input.findAllergenFreeIngredients()

	log.Println("expected: 5, actual:", input.countOccurrencesOfIngredients(ingredients))
}

func partOne(lines []string) {
	input := parseInput(lines)
	ingredients := input.findAllergenFreeIngredients()

	log.Println("How many times do allergen free ingredients appear?")
	log.Println("Answer:", input.countOccurrencesOfIngredients(ingredients))
}

func testPartTwo() {
}

func partTwo() {
}

func parseInput(lines []string) Input {
	foods := make([]Food, len(lines))
	nameToIngredient := make(map[string]*Ingredient)
	nameToAllergen := make(map[string]*Allergen)

	for index, line := range lines {
		parts := strings.Split(line, " (contains ")

		ingredientNames := parts[0]
		allergenNames := strings.TrimSuffix(parts[1], ")")

		ingredients := parseIngredients(ingredientNames, nameToIngredient)
		allergens := parseAllergens(allergenNames, nameToAllergen)

		for _, ingredient := range ingredients {
			for _, allergen := range allergens {
				ingredient.nameToPossibleAllergens[allergen.name] = allergen
				allergen.nameToPossibleIngredients[ingredient.name] = ingredient
			}
		}

		foods[index] = createFood(ingredients, allergens)
	}

	return Input{
		foods,
		nameToIngredient,
		nameToAllergen,
	}
}

func parseIngredients(line string, nameToIngredient map[string]*Ingredient) []*Ingredient {
	names := strings.Split(line, " ")
	ingredients := make([]*Ingredient, len(names))

	for index, name := range names {
		ingredient, hasIngredient := nameToIngredient[name]

		if !hasIngredient {
			ingredient = &Ingredient{name, make(map[string]*Allergen)}
			nameToIngredient[name] = ingredient
		}

		ingredients[index] = ingredient
	}

	return ingredients
}

func parseAllergens(line string, nameToAllergen map[string]*Allergen) []*Allergen {
	names := strings.Split(line, ", ")
	allergens := make([]*Allergen, len(names))

	for index, name := range names {
		allergen, hasAllergen := nameToAllergen[name]

		if !hasAllergen {
			allergen = &Allergen{name, make(map[string]*Ingredient)}
			nameToAllergen[name] = allergen
		}

		allergens[index] = allergen
	}

	return allergens
}

type Input struct {
	foods            []Food
	nameToIngredient map[string]*Ingredient
	nameToAllergen   map[string]*Allergen
}

func (input Input) findAllergenFreeIngredients() []*Ingredient {
	ingredients := make([]*Ingredient, 0)

	for _, ingredient := range input.nameToIngredient {
		if ingredient.isAllergenFree(input.foods) {
			ingredients = append(ingredients, ingredient)
		}
	}

	return ingredients
}

func (input Input) countOccurrencesOfIngredients(ingredients []*Ingredient) int {
	occurrenceCount := 0

	for _, ingredient := range ingredients {
		occurrenceCount += input.countOccurrencesOfIngredient(ingredient)
	}

	return occurrenceCount
}

func (input Input) countOccurrencesOfIngredient(ingredient *Ingredient) int {
	occurrenceCount := 0

	for _, food := range input.foods {
		if food.hasIngredient(ingredient) {
			occurrenceCount++
		}
	}

	return occurrenceCount
}

type Food struct {
	nameToIngredient map[string]*Ingredient
	nameToAllergen   map[string]*Allergen
}

func (food Food) hasIngredient(ingredient *Ingredient) bool {
	_, contains := food.nameToIngredient[ingredient.name]

	return contains
}

func (food Food) hasAllergen(allergen *Allergen) bool {
	_, contains := food.nameToAllergen[allergen.name]

	return contains
}

func createFood(ingredients []*Ingredient, allergens []*Allergen) Food {
	food := Food{make(map[string]*Ingredient), make(map[string]*Allergen)}

	for _, ingredient := range ingredients {
		food.nameToIngredient[ingredient.name] = ingredient
	}

	for _, allergen := range allergens {
		food.nameToAllergen[allergen.name] = allergen
	}

	return food
}

type Ingredient struct {
	name                    string
	nameToPossibleAllergens map[string]*Allergen
}

func (ingredient *Ingredient) isAllergenFree(foods []Food) bool {
	for _, allergen := range ingredient.nameToPossibleAllergens {
		isFreeOfCurrentAllergen := false

		for _, food := range foods {
			if !food.hasAllergen(allergen) {
				continue
			}

			if !food.hasIngredient(ingredient) {
				isFreeOfCurrentAllergen = true
				break
			}
		}

		if !isFreeOfCurrentAllergen {
			return false
		}
	}

	return true
}

type Allergen struct {
	name                      string
	nameToPossibleIngredients map[string]*Ingredient
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
