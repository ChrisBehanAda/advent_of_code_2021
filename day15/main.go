package day15

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type point struct {
	row int
	col int
}

func Solve() {
	input := getInput("./day15/input.txt")
	grid := buildGrid(input)
	lowestRiskSum := lowestRiskPathSum(grid)
	fmt.Printf("Sum of lowest risk path is %v\n", lowestRiskSum)
	fmt.Printf(" test mod: %v\n", 5/10)
	fullGrid := buildFullGrid(grid)
	fmt.Printf("Full Grid:\n%v", fullGrid)
	lowestRiskSumFull := lowestRiskPathSum(fullGrid)
	fmt.Printf("Sum of lowest risk path on full grid is %v\n", lowestRiskSumFull)
}

func getInput(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	input := strings.Trim(string(data), "\n")
	return input
}

func buildGrid(input string) [][]int {
	grid := [][]int{}
	rows := strings.Split(input, "\n")
	for _, row := range rows {
		rowNums := []int{}
		for _, col := range row {
			n, _ := strconv.Atoi(string(col))
			rowNums = append(rowNums, n)
		}
		grid = append(grid, rowNums)
	}
	return grid
}

func lowestRiskPathSum(grid [][]int) int {
	risk := riskLevels(len(grid), len(grid[0]))
	risk[0][0] = 0
	q := []point{}
	q = append(q, point{row: 0, col: 0})
	for len(q) > 0 {
		p := q[0]
		q = q[1:]
		neighbours := neighbours(grid, p)
		for _, n := range neighbours {
			if risk[p.row][p.col]+grid[n.row][n.col] < risk[n.row][n.col] {
				risk[n.row][n.col] = risk[p.row][p.col] + grid[n.row][n.col]
				q = append(q, n)
			}
		}
	}
	return risk[len(grid)-1][len(grid[0])-1]
}

func riskLevels(rows, cols int) [][]int {
	levels := [][]int{}
	for r := 0; r < rows; r++ {
		row := []int{}
		for c := 0; c < cols; c++ {
			row = append(row, math.MaxInt)
		}
		levels = append(levels, row)
	}
	return levels
}

func neighbours(grid [][]int, p point) []point {
	up := point{row: p.row - 1, col: p.col}
	down := point{row: p.row + 1, col: p.col}
	left := point{row: p.row, col: p.col - 1}
	right := point{row: p.row, col: p.col + 1}
	adj := []point{up, down, left, right}
	neighbours := []point{}
	for _, p := range adj {
		if p.row >= 0 && p.row < len(grid) && p.col >= 0 && p.col < len(grid[0]) {
			neighbours = append(neighbours, p)
		}
	}
	return neighbours
}

func buildFullGrid(grid [][]int) [][]int {
	full := [][]int{}
	tileHeight := len(grid)
	tileWidth := len(grid[0])
	rows, cols := tileHeight*5, tileHeight*5
	for r := 0; r < rows; r++ {
		originalRow := r % tileHeight
		fullRow := []int{}
		for c := 0; c < cols; c++ {
			addAmt := c/tileWidth + r/tileHeight
			originalCol := c % tileWidth
			val := grid[originalRow][originalCol] + addAmt
			fullRow = append(fullRow, val)
		}
		full = append(full, fullRow)
	}
	// replace values over 9 with val % 9
	for r := 0; r < len(full); r++ {
		for c := 0; c < len(full[0]); c++ {
			if full[r][c] > 9 {
				full[r][c] = full[r][c] % 9
			}
		}
	}

	return full
}
