package day1

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Solve() {
	input := getInputString("./day1/input.txt")
	delim := "\n"
	nums := stringToIntSlice(input, delim)
	ans1 := countIncreases(nums)
	ans2 := countTripletIncreases(nums)
	fmt.Println("DAY 1")
	fmt.Printf("Depth increased %v times.\n", ans1)
	fmt.Printf("Triplet depth increased %v times.\n", ans2)
}

func getInputString(path string) string {
	input, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(input)
}

func stringToIntSlice(s string, delim string) []int {
	strs := strings.Split(s, delim)
	nums := []int{}
	for _, strNum := range strs {
		if strNum == "" {
			continue
		}
		n, err := strconv.Atoi(strNum)
		if err != nil {
			panic(err)
		}
		nums = append(nums, n)
	}
	return nums
}

// Part 1
func countIncreases(nums []int) int {
	var prev int
	increases := 0
	for idx, num := range nums {
		if idx == 0 {
			prev = num
			continue
		}

		if prev < num {
			increases++
		}
		prev = num
	}
	return increases
}

// Part 2
func countTripletIncreases(nums []int) int {
	prevTripletEndIdx := 2
	curTripletEndIdx := 3
	increases := 0
	for curTripletEndIdx < len(nums) {
		prevSum := sumTripletFromEndIdx(nums, prevTripletEndIdx)
		curSum := sumTripletFromEndIdx(nums, curTripletEndIdx)
		if curSum > prevSum {
			increases++
		}
		prevTripletEndIdx = curTripletEndIdx
		curTripletEndIdx++
	}
	return increases
}

func sumTripletFromEndIdx(nums []int, idx int) int {
	return nums[idx] + nums[idx-1] + nums[idx-2]
}
