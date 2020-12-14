package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	earliest, busses := readEarliestAndBusses("./day13/input.txt")

	log.Println("Day 13 Part 01")
	testPartOne()
	partOne(earliest, busses)

	log.Println()

	log.Println("Day 13 Part 02")
	testPartTwo()
	partTwo()
}

func testPartOne() {
	earliest := 939
	busses := []int{7, 13, 59, 31, 19}
	bus, departure := getEarliestBus(earliest, busses)

	log.Println(hash(departure, earliest, bus))
}

func partOne(earliest int, busses []int) {
	bus, departure := getEarliestBus(earliest, busses)

	log.Println(hash(departure, earliest, bus))
}

func testPartTwo() {

}

func partTwo() {

}

func hash(departure, earliest, bus int) int {
	return (departure - earliest) * bus
}

func getEarliestBus(earliest int, busses []int) (earliestBus, earliestDeparture int) {
	earliestDeparture = int(^uint(0) >> 1)

	for _, bus := range busses {
		departure := getDepartureTimestampOnBus(earliest, bus)

		if departure < earliestDeparture {
			earliestBus = bus
			earliestDeparture = departure
		}
	}

	return earliestBus, earliestDeparture
}

func getDepartureTimestampOnBus(earliest, bus int) int {
	roundsFloor := float64(earliest) / float64(bus)
	roundsCeil := int(math.Ceil(roundsFloor))

	return roundsCeil * bus
}

func readEarliestAndBusses(filePath string) (earliest int, busses []int) {
	file, err := os.Open(filePath)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()

	earliestAsString := scanner.Text()
	earliest, _ = strconv.Atoi(earliestAsString)

	scanner.Scan()

	bussesAsString := scanner.Text()
	bussesParts := strings.Split(bussesAsString, ",")

	for _, busAsString := range bussesParts {
		if busAsString == "x" {
			continue
		}

		bus, _ := strconv.Atoi(busAsString)
		busses = append(busses, bus)
	}

	return earliest, busses
}
