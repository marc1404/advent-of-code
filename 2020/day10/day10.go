package main

import (
	"bufio"
	"container/list"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	testFewAdapters := []int{
		16,
		10,
		15,
		5,
		1,
		11,
		7,
		19,
		6,
		12,
		4,
	}

	testManyAdapters := []int{
		28,
		33,
		18,
		42,
		31,
		14,
		46,
		20,
		48,
		47,
		24,
		23,
		49,
		45,
		19,
		38,
		39,
		11,
		1,
		32,
		25,
		35,
		8,
		17,
		7,
		9,
		4,
		2,
		34,
		10,
		3,
	}

	adapters := readLinesAsAdapters("./day10/input.txt")

	log.Println("Day 10 Part 01")
	testPartOne(testFewAdapters, testManyAdapters)
	partOne(adapters)

	log.Println()

	log.Println("Day 10 Part 02")
	testPartTwo(testFewAdapters, testManyAdapters)
	partTwo(adapters)
}

func testPartOne(fewAdapters, manyAdapters []int) {
	printJoltageRatingAndDifferences(fewAdapters)
	printJoltageRatingAndDifferences(manyAdapters)
}

func partOne(adapters []int) {
	_, differenceToCount := getJoltageRatingAndDifferences(adapters)
	product := differenceToCount[1] * differenceToCount[3]

	log.Println("Product of 1-jolt and 3-jolt differences:", product)
}

func testPartTwo(fewAdapters, manyAdapters []int) {
	log.Println("Arrangement count for few adapters:", len(findArrangements(fewAdapters)))
	log.Println("Arrangement count for many adapters:", len(findArrangements(manyAdapters)))
}

func partTwo(adapters []int) {
	log.Println("Arrangement count for input adapters:", len(findArrangements(adapters)))
}

func getJoltageRatingAndDifferences(adapters []int) (deviceJoltageRating int, differenceToCount map[int]int) {
	sortAscending(adapters)

	deviceJoltageRating = determineDeviceJoltageRating(adapters)
	differenceToCount = collectJoltageDifferences(adapters, deviceJoltageRating, make(map[int]int))

	return deviceJoltageRating, differenceToCount
}

func printJoltageRatingAndDifferences(adapters []int) {
	deviceJoltageRating, differenceToCount := getJoltageRatingAndDifferences(adapters)

	log.Println("Device joltage rating:", deviceJoltageRating)
	log.Println("Joltage differences:", differenceToCount)
}

func sortAscending(values []int) {
	sort.Slice(values, func(i, j int) bool {
		return values[i] < values[j]
	})
}

func determineDeviceJoltageRating(adapters []int) int {
	highestAdapter := 0

	for _, adapter := range adapters {
		if adapter > highestAdapter {
			highestAdapter = adapter
		}
	}

	return highestAdapter + 3
}

func collectJoltageDifferences(adapters []int, targetJoltage int, differenceToCount map[int]int) map[int]int {
	adapters = append(adapters, targetJoltage)
	sourceJoltage := 0

	for _, adapter := range adapters {
		difference := adapter - sourceJoltage
		differenceToCount[difference]++
		sourceJoltage = adapter
	}

	return differenceToCount
}

func findArrangements(adapters []int) map[string]bool {
	sortAscending(adapters)

	deviceJoltageRating := determineDeviceJoltageRating(adapters)
	adapters = append([]int{0}, adapters...)
	adapters = append(adapters, deviceJoltageRating)
	adapterList := list.New()

	for _, adapter := range adapters {
		adapterList.PushBack(adapter)
	}

	arrangements := make(map[string]bool)
	arrangements[getArrangementKey(adapterList)] = true
	arrangements = findSubArrangements(adapterList, arrangements)

	return arrangements
}

func findSubArrangements(adapters *list.List, arrangements map[string]bool) map[string]bool {
	for adapter := adapters.Front().Next(); adapter.Next() != nil; adapter = adapter.Next() {
		left := adapter.Prev()
		right := adapter.Next()
		difference := right.Value.(int) - left.Value.(int)

		if difference > 3 {
			continue
		}

		subArrangement := cloneListWithout(adapters, adapter)
		subArrangementKey := getArrangementKey(subArrangement)

		if arrangements[subArrangementKey] {
			continue
		}

		arrangements[subArrangementKey] = true
		arrangements = findSubArrangements(subArrangement, arrangements)
	}

	return arrangements
}

func cloneListWithout(original *list.List, elementToRemove *list.Element) *list.List {
	clone := list.New()

	for element := original.Front(); element != nil; element = element.Next() {
		if element == elementToRemove {
			continue
		}

		clone.PushBack(element.Value)
	}

	return clone
}

func getArrangementKey(adapters *list.List) string {
	var builder strings.Builder

	for adapter := adapters.Front(); adapter != nil; adapter = adapter.Next() {
		builder.WriteString(fmt.Sprintf("%v;", adapter.Value))
	}

	return builder.String()
}

func readLinesAsAdapters(filePath string) []int {
	file, err := os.Open(filePath)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	adapters := make([]int, 0)

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			continue
		}

		adapter, _ := strconv.Atoi(line)
		adapters = append(adapters, adapter)
	}

	return adapters
}
