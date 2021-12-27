package day13

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Solve() {
	input := getInput("./day13/input.txt")
	pointsInput, foldsInput := pointsAndFold(input)
	points := points(pointsInput)
	grid := buildGrid(points)
	c0 := count(grid)
	fmt.Printf("Points before first fold: %v\n", c0)
	folds := folds(foldsInput)
	f := applyFold(grid, folds[0])
	c1 := count(f)
	fmt.Printf("Points after first fold: %v\n", c1)
	allFolds := applyFolds(grid, folds)
	printPaper(allFolds)
}

type point struct {
	row int
	col int
}

func getInput(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	input := strings.Trim(string(data), "\n")
	return input
}

func pointsAndFold(input string) (string, string) {
	pointsAndFold := strings.Split(input, "\n\n")
	return pointsAndFold[0], pointsAndFold[1]
}

func points(input string) []point {
	lines := strings.Split(input, "\n")
	points := []point{}
	for _, line := range lines {
		ps := strings.Split(line, ",")
		col, _ := strconv.Atoi(ps[0])
		row, _ := strconv.Atoi(ps[1])
		p := point{row: row, col: col}
		points = append(points, p)
	}
	return points
}

func buildGrid(points []point) [][]int {
	x, y := -1, -1
	for _, p := range points {
		if p.col > x {
			x = p.col
		}
		if p.row > y {
			y = p.row
		}
	}

	grid := [][]int{}
	for row := 0; row <= y; row++ {
		line := make([]int, x+1)
		grid = append(grid, line)
	}

	for _, p := range points {
		grid[p.row][p.col] = 1
	}

	return grid
}

func folds(input string) []point {
	instructions := strings.Split(input, "\n")
	folds := []point{}
	for _, i := range instructions {
		axisAndCoordinate := strings.Split(i, "=")
		axis := axisAndCoordinate[0]
		axis = axis[len(axis)-1:]
		co, _ := strconv.Atoi(axisAndCoordinate[1])
		p := point{}
		if []rune(axis)[0] == 'x' {
			p = point{row: 0, col: co}
		} else {
			p = point{row: co, col: 0}
		}
		folds = append(folds, p)
	}
	return folds
}

func applyFold(grid [][]int, fold point) [][]int {
	// horizontal fold
	if fold.col == 0 {
		for row := fold.row + 1; row < len(grid); row++ {
			for col := 0; col < len(grid[0]); col++ {
				if grid[row][col] == 1 {
					newRow := fold.row - (row - fold.row)
					grid[newRow][col] = 1
				}
			}
		}
		return grid[:fold.row]
	}
	// vertical fold
	if fold.row == 0 {
		for row := 0; row < len(grid); row++ {
			for col := fold.col + 1; col < len(grid[0]); col++ {
				if grid[row][col] == 1 {
					newCol := fold.col - (col - fold.col)
					grid[row][newCol] = 1
				}
			}
		}
		newGrid := [][]int{}
		for row := 0; row < len(grid); row++ {
			newGrid = append(newGrid, grid[row][:fold.col])
		}
		return newGrid
	}
	return [][]int{}
}

func count(grid [][]int) int {
	count := 0
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[0]); col++ {
			if grid[row][col] == 1 {
				count++
			}
		}
	}
	return count
}

func applyFolds(grid [][]int, folds []point) [][]int {
	folded := grid
	for _, f := range folds {
		folded = applyFold(folded, f)
	}
	return folded
}

func printPaper(grid [][]int) {
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[0]); col++ {
			fmt.Printf("%v", grid[row][col])
		}
		fmt.Print("\n")
	}
}
