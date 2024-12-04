package main

import (
	"fmt"
	"log/slog"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	re     = regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	reDo   = regexp.MustCompile(`do\(\)`)
	reDont = regexp.MustCompile(`don't\(\)`)
)

func main() {
	data, err := os.ReadFile("2024/day03/input.txt")
	if err != nil {
		slog.Error("unable to read file", slog.Any("err", err))
		os.Exit(1)
	}

	partOne(data)
	partTwo(data)
}

func partOne(data []byte) {
	matches := re.FindAllString(string(data), -1)

	sum := 0

	for _, match := range matches {
		match = strings.ReplaceAll(match, "mul(", "")
		match = strings.ReplaceAll(match, ")", "")

		numbers := strings.Split(match, ",")

		val1, _ := strconv.Atoi(numbers[0])
		val2, _ := strconv.Atoi(numbers[1])

		sum += val1 * val2

	}

	fmt.Println("Sum before do and don't ", sum)
}

func partTwo(data []byte) {
	mulEnabled := true
	sum := 0

	matches := re.FindAllString(string(data), -1)

	for _, part := range matches {
		if reDo.MatchString(part) {
			mulEnabled = true
			continue
		}

		if reDont.MatchString(part) {
			mulEnabled = false
			continue
		}

		vard := re.FindStringSubmatch(part)
		if len(vard) > 0 && mulEnabled {
			num1, _ := strconv.Atoi(vard[1])
			num2, _ := strconv.Atoi(vard[2])

			sum += num1 * num2
		}
	}

	fmt.Println("Sum after do and don't", sum)
}
