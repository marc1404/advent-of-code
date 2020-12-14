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
	partTwo(busses)
}

func testPartOne() {
	earliest := 939
	busses := []int{7, 13, 59, 31, 19}
	bus, departure := getEarliestBus(earliest, busses)

	log.Println(hash(departure, earliest, bus))
}

func partOne(earliest int, busses []int) {
	bussesInService := []int{}

	for _, bus := range busses {
		if bus == 1 {
			continue
		}

		bussesInService = append(bussesInService, bus)
	}

	bus, departure := getEarliestBus(earliest, bussesInService)

	log.Println(hash(departure, earliest, bus))
}

func testPartTwo() {
	log.Println(getSuccessionTimestamp([]int{7, 13, 1, 1, 59, 1, 31, 19}))
	log.Println(getSuccessionTimestamp([]int{17, 1, 13, 19}))
	log.Println(getSuccessionTimestamp([]int{67, 7, 59, 61}))
	log.Println(getSuccessionTimestamp([]int{67, 1, 7, 59, 61}))
	log.Println(getSuccessionTimestamp([]int{67, 7, 1, 59, 61}))
	log.Println(getSuccessionTimestamp([]int{1789, 37, 47, 1889}))
}

func partTwo(busses []int) {
	log.Println(getSuccessionTimestamp(busses))
}

func getSuccessionTimestamp(busses []int) int {
	timestamp := 0
	step := 1

	for i, bus := range busses {
		if bus == 1 {
			continue
		}

		for (timestamp+i)%bus > 0 {
			timestamp += step
		}

		step *= bus
	}

	return timestamp
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
			busAsString = "1"
		}

		bus, _ := strconv.Atoi(busAsString)
		busses = append(busses, bus)
	}

	return earliest, busses
}
