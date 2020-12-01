package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("./day01/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	numbers := make([]int, 1)

	for scanner.Scan() {
		line := scanner.Text()

		number, err := strconv.Atoi(line)

		if err != nil {
			log.Fatal(err)
		}

		numbers = append(numbers, number)
	}

	partOne(numbers)
	partTwo(numbers)
}

func partOne(numbers []int) {
	for i, a := range numbers {
		for j, b := range numbers {
			if i == j {
				continue
			}

			if a+b != 2020 {
				continue
			}

			fmt.Println("Part one:", a*b)
			return
		}
	}
}

func partTwo(numbers []int) {
	for i, a := range numbers {
		for j, b := range numbers {
			for h, c := range numbers {
				if i == j || j == h || i == h {
					continue
				}

				if a+b+c != 2020 {
					continue
				}

				product := a * b * c

				if product == 0 {
					continue
				}

				fmt.Println("Part two:", product)
				return
			}
		}
	}
}
