package main

import (
	"fmt"
	"log/slog"
	"os"
	"strconv"
	"strings"
)

func main() {
	partOne()
	partTwo()
}

func partTwo() {
	data, err := os.ReadFile("2024/day02/input.txt")
	if err != nil {
		slog.Error("unable to read file", slog.Any("err", err))
		os.Exit(1)
	}

	// Split the input into lines
	reports := strings.Split(string(data), "\n")

	// Count the number of safe reports
	safeCount := 0
	for _, report := range reports {
		levels := parseLevels(report)
		if isSafeWithDampener(levels) {
			safeCount++
		}
	}

	fmt.Printf("Number of safe reports ( with damp ): %d\n", safeCount)
}

func partOne() {
	data, err := os.ReadFile("2024/day02/input.txt")
	if err != nil {
		slog.Error("unable to read file", slog.Any("err", err))
		os.Exit(1)
	}

	// Split the input into lines
	reports := strings.Split(string(data), "\n")

	// Count the number of safe reports
	safeCount := 0
	for _, report := range reports {
		levels := parseLevels(report)
		if isSafe(levels) {
			safeCount++
		}
	}

	fmt.Printf("Number of safe reports: %d\n", safeCount)
}

// parseLevels parses a single line of input into a slice of integers
func parseLevels(report string) []int {
	parts := strings.Fields(report)
	levels := make([]int, len(parts))
	for i, part := range parts {
		num, err := strconv.Atoi(part)
		if err != nil {
			panic(fmt.Sprintf("Invalid number in input: %s", part))
		}
		levels[i] = num
	}
	return levels
}

// isSafe checks if a report is safe according to the given rules
func isSafe(levels []int) bool {
	if len(levels) < 2 {
		return true // A single level is always safe
	}

	increasing := levels[1] > levels[0]
	decreasing := levels[1] < levels[0]

	// Iterate through the levels to validate the rules
	for i := 1; i < len(levels); i++ {
		diff := levels[i] - levels[i-1]
		if diff == 0 || diff < -3 || diff > 3 {
			return false // Rule 2: Adjacent levels differ by at least 1 and at most 3
		}
		if increasing && levels[i] <= levels[i-1] {
			return false // Rule 1: All levels must be increasing
		}
		if decreasing && levels[i] >= levels[i-1] {
			return false // Rule 1: All levels must be decreasing
		}
	}

	return true
}

func isSafeWithDampener(levels []int) bool {
	if isSafe(levels) {
		return true
	}

	for i := 0; i < len(levels); i++ {
		modified := append([]int{}, levels[:i]...)   // Copy levels before the removed index
		modified = append(modified, levels[i+1:]...) // Append levels after the removed index
		if isSafe(modified) {
			return true
		}
	}

	return false
}
