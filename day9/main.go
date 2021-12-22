package day9

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type coordinate struct {
	row int
	col int
}

func Solve() {
	input := getInputString("./day9/input.txt")
	heightMap := parseHeightMap(input)
	lowPoints := findLowPoints(heightMap)
	riskLevel := riskLevel(heightMap, lowPoints)
	fmt.Printf("Risk level is: %v\n", riskLevel)
	threeLargest := largestBasins(heightMap, lowPoints)
	ans := 1
	for _, s := range threeLargest {
		ans *= s
	}
	fmt.Printf("Three largest basins multiplied together equal: %v\n", ans)
}

func getInputString(path string) string {
	input, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	trimmedInput := strings.Trim(string(input), "\n")
	return trimmedInput
}

func parseHeightMap(input string) [][]int {
	lines := strings.Split(input, "\n")
	heightMap := make([][]int, len(lines))
	for i, line := range lines {
		nums := []int{}
		for _, r := range line {
			n, err := strconv.Atoi(string(r))
			if err != nil {
				panic(err)
			}
			nums = append(nums, n)
		}
		heightMap[i] = nums
	}
	return heightMap
}

func findLowPoints(heightMap [][]int) []coordinate {
	lowPoints := []coordinate{}
	for row := 0; row < len(heightMap); row++ {
		for col := 0; col < len(heightMap[0]); col++ {
			co := coordinate{row: row, col: col}
			if lowpoint(co, heightMap) {
				lowPoints = append(lowPoints, co)
			}
		}
	}
	return lowPoints
}

func lowpoint(co coordinate, heightMap [][]int) bool {
	adjacent := adjacentPoints(co)
	isLowpoint := true
	for _, point := range adjacent {
		if point.row >= 0 && point.row < len(heightMap) && point.col >= 0 && point.col < len(heightMap[0]) {
			if heightMap[point.row][point.col] <= heightMap[co.row][co.col] {
				isLowpoint = false
				break
			}
		}
	}
	return isLowpoint
}

func adjacentPoints(co coordinate) []coordinate {
	up := coordinate{row: co.row - 1, col: co.col}
	down := coordinate{row: co.row + 1, col: co.col}
	left := coordinate{row: co.row, col: co.col - 1}
	right := coordinate{row: co.row, col: co.col + 1}
	return []coordinate{up, down, left, right}
}

func riskLevel(heightMap [][]int, points []coordinate) int {
	level := 0
	for _, p := range points {
		level += (1 + heightMap[p.row][p.col])
	}
	return level
}

func largestBasins(heightMap [][]int, lowpoints []coordinate) []int {
	basinSizes := []int{}
	for _, p := range lowpoints {
		size := basinSize(heightMap, p)
		basinSizes = append(basinSizes, size)
	}
	sort.Ints(basinSizes)
	return basinSizes[len(basinSizes)-3:]
}

func basinSize(heightMap [][]int, co coordinate) int {
	visited := []coordinate{co}
	q := []coordinate{co}
	size := 0
	for len(q) > 0 {
		// set point to first elem in queue
		point := q[0]
		// pop from queue
		q = q[1:]
		size++
		// visited = append(visited, point)
		adjPoints := adjacentPoints(point)
		for _, p := range adjPoints {
			// check we haven't already visited this point
			alreadyVisited := false
			for _, v := range visited {
				if p == v {
					alreadyVisited = true
					break
				}
			}
			if alreadyVisited {
				continue
			}

			// Check adj point is in bounds and not 9
			if p.row >= 0 && p.row < len(heightMap) && p.col >= 0 && p.col < len(heightMap[0]) && heightMap[p.row][p.col] != 9 {
				q = append(q, p)
				visited = append(visited, p)
			}
		}
	}
	return size
}
