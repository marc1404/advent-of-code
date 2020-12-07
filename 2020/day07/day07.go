package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	lines := readAllLines("./day07/input.txt")
	nameToBag := parseBags(lines)
	shinyGoldBag := nameToBag["shiny gold"]

	log.Println("Day 07 Part 01")
	testPartOne()
	partOne(shinyGoldBag)

	log.Println()

	log.Println("Day 07 Part 02")
	testPartTwo()
	partTwo(shinyGoldBag)
}

func testPartOne() {
	lines := []string{
		"light red bags contain 1 bright white bag, 2 muted yellow bags.",
		"dark orange bags contain 3 bright white bags, 4 muted yellow bags.",
		"bright white bags contain 1 shiny gold bag.",
		"muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.",
		"shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.",
		"dark olive bags contain 3 faded blue bags, 4 dotted black bags.",
		"vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.",
		"faded blue bags contain no other bags.",
		"dotted black bags contain no other bags.",
	}

	nameToBag := parseBags(lines)
	shinyGoldBag := nameToBag["shiny gold"]

	log.Println(countUniqueParentBags(shinyGoldBag))
}

func partOne(shinyGoldBag *Bag) {
	log.Println("Bags which eventually contain a shiny gold bag:", countUniqueParentBags(shinyGoldBag))
}

func testPartTwo() {
	lines := []string{
		"light red bags contain 1 bright white bag, 2 muted yellow bags.",
		"dark orange bags contain 3 bright white bags, 4 muted yellow bags.",
		"bright white bags contain 1 shiny gold bag.",
		"muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.",
		"shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.",
		"dark olive bags contain 3 faded blue bags, 4 dotted black bags.",
		"vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.",
		"faded blue bags contain no other bags.",
		"dotted black bags contain no other bags.",
	}

	nameToBag := parseBags(lines)
	shinyGoldBag := nameToBag["shiny gold"]

	log.Println(countChildrenBags(shinyGoldBag))
}

func partTwo(shinyGoldBag *Bag) {
	log.Println("Bags contained in a shiny gold bag:", countChildrenBags(shinyGoldBag))
}

type Bag struct {
	name        string
	parents     []*Bag
	children    []*Bag
	bagToAmount map[string]int
}

func countUniqueParentBags(bag *Bag) int {
	parentBags := collectParentBags(bag, make(map[string]*Bag))

	return len(parentBags)
}

func collectParentBags(bag *Bag, parentBagMap map[string]*Bag) map[string]*Bag {
	for _, parentBag := range bag.parents {
		parentBagMap[parentBag.name] = parentBag
		parentBagMap = collectParentBags(parentBag, parentBagMap)
	}

	return parentBagMap
}

func countChildrenBags(bag *Bag) int {
	childrenCount := 0

	for _, childBag := range bag.children {
		childAmount := bag.bagToAmount[childBag.name]
		childrenCount += childAmount + countChildrenBags(childBag)*childAmount
	}

	return childrenCount
}

func parseBags(lines []string) (nameToBag map[string]*Bag) {
	nameToBag = make(map[string]*Bag)

	for _, line := range lines {
		line = strings.TrimSuffix(line, ".")
		parentParts := strings.Split(line, " bags contain ")
		parentBagName := parentParts[0]
		tail := parentParts[1]
		parentBag := getBag(nameToBag, parentBagName)
		containedBags := strings.Split(tail, ", ")

		for _, containedBag := range containedBags {
			if containedBag == "no other bags" {
				continue
			}

			childParts := strings.Fields(containedBag)
			amount, _ := strconv.Atoi(childParts[0])
			childBagName := strings.Join(childParts[1:], " ")
			childBagName = normalizeBagName(childBagName)
			childBag := getBag(nameToBag, childBagName)

			parentBag.children = append(parentBag.children, childBag)
			childBag.parents = append(childBag.parents, parentBag)

			parentBag.bagToAmount[childBagName] = amount
		}
	}

	return nameToBag
}

func getBag(nameToBag map[string]*Bag, bagName string) *Bag {
	bag, hasBag := nameToBag[bagName]

	if !hasBag {
		bag = &Bag{name: bagName, bagToAmount: make(map[string]int)}
		nameToBag[bagName] = bag
	}

	return bag
}

func normalizeBagName(bagName string) string {
	bagName = strings.TrimSuffix(bagName, " bags")
	bagName = strings.TrimSuffix(bagName, " bag")

	return bagName
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
