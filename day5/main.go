package day5

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type line struct {
	x1, x2, y1, y2 int
}

func Solve() {
	input := getInputString("./day5/input.txt")
	allLines := parseInput(input)
	hVLines := horizontalAndVertical(allLines)
	x, y := maxXY(allLines)
	grid1 := buildGrid(x, y)
	grid2 := buildGrid(x, y)
	drawLines(grid1, hVLines)
	drawLines2(grid2, allLines)
	ans1 := dangerZones(grid1)
	ans2 := dangerZones(grid2)
	fmt.Printf("The number of danger zones with horizontal and vertical lines is: %v\n", ans1)
	fmt.Printf("The number of danger zones with all lines is: %v\n", ans2)
}

func getInputString(path string) string {
	input, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	trimmedInput := strings.Trim(string(input), "\n")
	return trimmedInput
}

func parseInput(s string) []line {
	lines := []line{}
	lineStrings := strings.Split(s, "\n")
	for _, ls := range lineStrings {
		p1AndP2String := strings.Split(ls, " -> ")
		p1XAndY := strings.Split(p1AndP2String[0], ",")
		p2XAndY := strings.Split(p1AndP2String[1], ",")
		p1X, _ := strconv.Atoi(p1XAndY[0])
		p1Y, _ := strconv.Atoi(p1XAndY[1])
		p2X, _ := strconv.Atoi(p2XAndY[0])
		p2Y, _ := strconv.Atoi(p2XAndY[1])
		line := line{x1: p1X, x2: p2X, y1: p1Y, y2: p2Y}
		lines = append(lines, line)
	}
	return lines
}

func horizontalAndVertical(lines []line) []line {
	filteredLines := []line{}
	for _, l := range lines {
		if l.x1 == l.x2 || l.y1 == l.y2 {
			filteredLines = append(filteredLines, l)
		}
	}
	return filteredLines
}

func maxXY(lines []line) (int, int) {
	maxX, maxY := 0, 0
	for _, l := range lines {
		if l.x1 > maxX || l.x2 > maxX {
			newXMax := 0
			if l.x1 >= l.x2 {
				newXMax = l.x1
			} else {
				newXMax = l.x2
			}
			maxX = newXMax
		}
		if l.y1 > maxY || l.y2 > maxY {
			newYMax := 0
			if l.y1 >= l.y2 {
				newYMax = l.y1
			} else {
				newYMax = l.y2
			}
			maxY = newYMax
		}
	}
	return maxX, maxY
}

func buildGrid(x, y int) [][]int {
	grid := [][]int{}
	for i := 0; i <= y; i++ {
		row := []int{}
		for j := 0; j <= x; j++ {
			row = append(row, 0)
		}
		grid = append(grid, row)
	}
	return grid
}

func drawLines(grid [][]int, lines []line) {
	for _, l := range lines {
		if l.x1 == l.x2 {
			length := math.Abs(float64(l.y1 - l.y2))
			start := min(l.y1, l.y2)
			for i := start; i <= start+int(length); i++ {
				grid[i][l.x1]++
			}
		} else if l.y1 == l.y2 {
			length := math.Abs(float64(l.x1 - l.x2))
			start := min(l.x1, l.x2)
			for i := start; i <= start+int(length); i++ {
				grid[l.y1][i]++
			}
		}
	}
}

func drawLines2(grid [][]int, lines []line) {
	for _, l := range lines {
		if l.x1 == l.x2 {
			length := math.Abs(float64(l.y1 - l.y2))
			start := min(l.y1, l.y2)
			for i := start; i <= start+int(length); i++ {
				grid[i][l.x1]++
			}
		} else if l.y1 == l.y2 {
			length := math.Abs(float64(l.x1 - l.x2))
			start := min(l.x1, l.x2)
			for i := start; i <= start+int(length); i++ {
				grid[l.y1][i]++
			}
		} else {
			var lowX, lowY, hiX int
			if l.y1 < l.y2 {
				lowY = l.y1
				lowX = l.x1
				hiX = l.x2
			} else {
				lowY = l.y2
				lowX = l.x2
				hiX = l.x1
			}
			if lowX < hiX {
				for i := 0; i <= hiX-lowX; i++ {
					grid[lowY+i][lowX+i]++
				}
			} else {
				for i := 0; i <= lowX-hiX; i++ {
					grid[lowY+i][lowX-i]++
				}
			}
		}
	}
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func dangerZones(grid [][]int) int {
	count := 0
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[0]); col++ {
			if grid[row][col] > 1 {
				count++
			}
		}
	}
	return count
}
