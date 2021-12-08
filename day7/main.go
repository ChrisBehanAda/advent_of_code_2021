package day7

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func Solve() {
	input := getInputString("./day7/input.txt")
	crabPositions := parseInts(input)
	max := findMax(crabPositions)
	counts := make([]int, max+1)
	pos, cost := mostEfficientPosition(counts, crabPositions, false)
	fmt.Printf("Most efficient position is %v, with a total fuel cost of %v\n", pos, cost)
	counts2 := make([]int, max+1)
	pos2, cost2 := mostEfficientPosition(counts2, crabPositions, true)
	fmt.Printf("Most efficient position with expensive fuel is %v, with a total fuel cost of %v\n", pos2, cost2)
}

func getInputString(path string) string {
	input, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	trimmedInput := strings.Trim(string(input), "\n")
	return trimmedInput
}

func parseInts(input string) []int {
	nums := []int{}
	numStrings := strings.Split(input, ",")
	for _, str := range numStrings {
		n, err := strconv.Atoi(str)
		if err != nil {
			panic(err)
		}
		nums = append(nums, n)
	}
	return nums
}

func findMax(nums []int) int {
	max := -1
	for _, n := range nums {
		if n > max {
			max = n
		}
	}
	return max
}

func mostEfficientPosition(counts []int, positions []int, expensive bool) (int, int) {
	for c := 0; c < len(counts); c++ {
		count := 0
		for _, pos := range positions {
			distance := abs(c - pos)
			if expensive {
				count += increasedFuelCost(distance)
			} else {
				count += abs(c - pos)
			}
		}
		counts[c] = count
	}
	minPos, minCost := math.MaxInt, math.MaxInt
	for pos, cost := range counts {
		if cost < minCost {
			minCost = cost
			minPos = pos
		}
	}
	return minPos, minCost
}

func increasedFuelCost(distance int) int {
	cost := 0
	for i := 1; i <= distance; i++ {
		cost += i
	}
	return cost
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
