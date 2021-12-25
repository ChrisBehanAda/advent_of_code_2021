package day25

import (
	"fmt"
	"os"
	"strings"
)

type point struct {
	row int
	col int
}

func Solve() {
	input := getInput("./day25/input.txt")
	grid := buildGrid(input)
	stopAfter := stopAfter(grid)
	fmt.Printf("Stopped after: %v", stopAfter)
}

func getInput(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	input := strings.Trim(string(data), "\n")
	return input
}

func buildGrid(input string) [][]rune {
	grid := [][]rune{}
	rows := strings.Split(input, "\n")
	for _, r := range rows {
		col := []rune(r)
		grid = append(grid, col)
	}
	return grid
}

func stopAfter(grid [][]rune) int {
	steps := 0
	for {
		e := moveEast(grid)
		s := moveSouth(grid)
		steps++
		if !e && !s {
			return steps
		}
	}
}

func moveEast(grid [][]rune) bool {
	toMove := []point{}
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[0]); col++ {
			if grid[row][col] == '>' {
				if col == len(grid[0])-1 {
					if grid[row][0] == '.' {
						p := point{row: row, col: col}
						toMove = append(toMove, p)
					}
				} else if grid[row][col+1] == '.' {
					p := point{row: row, col: col}
					toMove = append(toMove, p)
				}
			}
		}
	}

	for _, p := range toMove {
		if p.col == len(grid[0])-1 {
			grid[p.row][p.col] = '.'
			grid[p.row][0] = '>'
		} else {
			grid[p.row][p.col] = '.'
			grid[p.row][p.col+1] = '>'
		}
	}
	return len(toMove) > 0
}

func moveSouth(grid [][]rune) bool {
	toMove := []point{}
	for col := 0; col < len(grid[0]); col++ {
		for row := 0; row < len(grid); row++ {
			if grid[row][col] == 'v' {
				if row == len(grid)-1 {
					if grid[0][col] == '.' {
						p := point{row: row, col: col}
						toMove = append(toMove, p)
					}
				} else if grid[row+1][col] == '.' {
					p := point{row: row, col: col}
					toMove = append(toMove, p)
				}
			}
		}
	}

	for _, p := range toMove {
		if p.row == len(grid)-1 {
			grid[p.row][p.col] = '.'
			grid[0][p.col] = 'v'
		} else {
			grid[p.row][p.col] = '.'
			grid[p.row+1][p.col] = 'v'
		}
	}
	return len(toMove) > 0
}
