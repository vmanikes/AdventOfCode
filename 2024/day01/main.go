package main

import (
	"fmt"
	"log/slog"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("2024/day01/input.txt")
	if err != nil {
		slog.Error("unable to read file", slog.Any("err", err))
		os.Exit(1)
	}

	part01(data)
	part02(data)
}

func part01(data []byte) {
	var listA, listB []int

	// Parse input data
	for _, line := range strings.Split(string(data), "\n") {
		numbers := strings.Fields(line) // Handles any whitespace separator
		if len(numbers) < 2 {
			fmt.Printf("Skipping malformed line: %s\n", line)
			continue
		}

		for idx, number := range numbers[:2] { // Only process the first two numbers
			val, err := strconv.Atoi(number)
			if err != nil {
				fmt.Printf("Error converting '%s' to integer: %v\n", number, err)
				continue
			}

			switch idx {
			case 0:
				listA = append(listA, val)
			case 1:
				listB = append(listB, val)
			}
		}
	}

	// Sort both lists
	sort.Ints(listA)
	sort.Ints(listB)

	// Ensure lists are the same length
	if len(listA) != len(listB) {
		fmt.Println("Error: Lists are of unequal length.")
		return
	}

	// Calculate distance
	var distance float64
	for i := range listA {
		distance += math.Abs(float64(listA[i] - listB[i]))
	}

	fmt.Printf("Distance: %.2f\n", distance)
}

func part02(data []byte) {
	var (
		listA        []int
		frequencyMap = make(map[int]int)
	)

	// Parse input data
	for _, line := range strings.Split(string(data), "\n") {
		numbers := strings.Fields(line) // Handles any whitespace separator
		for idx, number := range numbers {
			val, err := strconv.Atoi(number)
			if err != nil {
				fmt.Printf("Error converting '%s' to integer: %v\n", number, err)
				continue
			}

			switch idx {
			case 0:
				listA = append(listA, val)
			case 1:
				frequencyMap[val]++
			}
		}
	}

	// Calculate distance
	distance := 0
	for _, number := range listA {
		if count, exists := frequencyMap[number]; exists {
			distance += number * count
		}
	}

	fmt.Println("Distance:", distance)
}
