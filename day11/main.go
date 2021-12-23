package day11

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type point struct {
	row int
	col int
}

func Solve() {
	input := getInput("./day11/input.txt")
	grid := buildGrid(input)
	flashes := flashes(grid, 100)
	fmt.Printf("There were %v flashes.\n", flashes)
	grid2 := buildGrid(input)
	stepsToSync := flashesSync(grid2)
	fmt.Printf("Octopi sync after %v steps.\n", stepsToSync)

}

func getInput(path string) string {
	input, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	trimmedInput := strings.Trim(string(input), "\n")
	return trimmedInput
}

func buildGrid(input string) [][]int {
	grid := [][]int{}
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		row := []int{}
		for _, r := range line {
			s := string(r)
			num, _ := strconv.Atoi(s)
			row = append(row, num)
		}
		grid = append(grid, row)
	}
	return grid
}

func flashes(grid [][]int, steps int) int {
	count := 0
	for i := 0; i < steps; i++ {
		increaseEnergy(grid)
		handleFlashes(grid)
		c := countFlashes(grid)
		count += c
	}
	return count
}

func flashesSync(grid [][]int) int {
	step := 1
	for {

		increaseEnergy(grid)
		handleFlashes(grid)
		c := countFlashes(grid)
		if c == len(grid)*len(grid[0]) {
			return step
		}
		step++
	}
}

func increaseEnergy(grid [][]int) {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			grid[i][j]++
		}
	}
}

func handleFlashes(grid [][]int) {
	flashed := []point{}
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[0]); col++ {
			p := point{row: row, col: col}
			if grid[p.row][p.col] > 9 && !contains(flashed, p) {
				flash(grid, p, &flashed)
			}
		}
	}
}

func contains(points []point, p point) bool {
	for _, i := range points {
		if i == p {
			return true
		}
	}
	return false
}

func flash(grid [][]int, p point, flashed *[]point) {
	*flashed = append(*flashed, p)
	ul := point{row: p.row - 1, col: p.col - 1}
	u := point{row: p.row - 1, col: p.col}
	ur := point{row: p.row - 1, col: p.col + 1}
	l := point{row: p.row, col: p.col - 1}
	r := point{row: p.row, col: p.col + 1}
	bl := point{row: p.row + 1, col: p.col - 1}
	b := point{row: p.row + 1, col: p.col}
	br := point{row: p.row + 1, col: p.col + 1}
	adj := []point{ul, u, ur, l, r, bl, b, br}
	for _, a := range adj {
		if a.row >= 0 && a.row < len(grid) && a.col >= 0 && a.col < len(grid[0]) {
			grid[a.row][a.col]++
			if grid[a.row][a.col] > 9 && !contains(*flashed, a) {
				flash(grid, a, flashed)
			}
		}
	}
}

func countFlashes(grid [][]int) int {
	count := 0
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[0]); col++ {
			if grid[row][col] > 9 {
				count++
				grid[row][col] = 0
			}
		}
	}
	return count
}
