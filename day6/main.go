package day6

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Solve() {
	input := getInputString("./day6/input.txt")
	fishDays := parseFishDays(input)
	fishDaysCpy := make([]int, len(fishDays))
	copy(fishDaysCpy, fishDays)
	eightyDays := simulateBirths(fishDays, 80)
	fmt.Printf("After 80 days there are %v lanternfish.\n", len(eightyDays))
	twoHundredFiftySixDays := simulateBirths2(fishDaysCpy, 256)
	fmt.Printf("After 256 days there are %v lanternfish.\n", twoHundredFiftySixDays)
}

func getInputString(path string) string {
	input, err := os.ReadFile(path)
	if err != nil {
		panic(input)
	}
	trimmedInput := strings.Trim(string(input), "\n")
	return trimmedInput
}

func parseFishDays(input string) []int {
	days := strings.Split(input, ",")
	fishDays := []int{}
	for _, s := range days {
		day, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		fishDays = append(fishDays, day)
	}
	return fishDays
}

func simulateBirths(fishDays []int, days int) []int {
	for i := 0; i < days; i++ {
		for j, day := range fishDays {
			if day == 0 {
				fishDays = append(fishDays, 8)
				(fishDays)[j] = 6
			} else {
				(fishDays)[j]--
			}
		}
	}
	return fishDays
}

func simulateBirths2(fishDays []int, days int) int {
	counts := make(map[int]int)
	for i := 0; i <= 8; i++ {
		counts[i] = 0
	}
	for _, d := range fishDays {
		counts[d]++
	}

	for d := 0; d < days; d++ {
		val := counts[len(counts)-1]
		for c := len(counts) - 1; c >= 0; c-- {
			if c == 0 {
				zeros := val
				counts[len(counts)-1] = zeros
				counts[6] += zeros
			} else {
				tmp := counts[c-1]
				counts[c-1] = val
				val = tmp
			}
		}
	}

	total := 0
	for _, val := range counts {
		total += val
	}
	return total
}
