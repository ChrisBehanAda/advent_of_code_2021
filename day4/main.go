package day4

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Solve() {
	input := getInputString("./day4/input.txt")
	nums, boards := numbersAndBoards(input)
	boardsCpy := copyInt3d(boards)
	score := bingo(nums, boards)
	fmt.Printf("The winning score is: %v\n", score)
	lastScore := lastWinningBoardsScore(nums, boardsCpy)
	fmt.Printf("The last winning score is: %v\n", lastScore)
}

func copyInt3d(src [][][]int64) [][][]int64 {
	dst := make([][][]int64, len(src))
	for i := range src {
		dst[i] = copyInt2d(src[i])
	}
	return dst
}

func copyInt2d(src [][]int64) [][]int64 {
	dst := make([][]int64, len(src))
	for i := range src {
		dst[i] = copyInt1d(src[i])
	}
	return dst
}

func copyInt1d(src []int64) []int64 {
	dst := make([]int64, len(src))
	copy(dst, src)
	return dst
}

func getInputString(path string) string {
	input, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	inputString := strings.Trim(string(input), "\n")
	return inputString
}

func numbersAndBoards(input string) ([]int64, [][][]int64) {
	numsAndBoards := strings.SplitN(input, "\n", 2)
	numString := numsAndBoards[0]
	numStringSlice := strings.Split(numString, ",")
	nums := []int64{}
	for _, str := range numStringSlice {
		n, err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			panic(err)
		}
		nums = append(nums, n)
	}
	boardString := strings.Trim(numsAndBoards[1], "\n")
	boardSlice := strings.Split(boardString, "\n")
	boards := buildBoards(boardSlice)
	return nums, boards
}

func buildBoards(input []string) [][][]int64 {
	noEmptylines := []string{}
	for _, line := range input {
		if len(line) < 2 {
			continue
		}
		noEmptylines = append(noEmptylines, line)
	}

	boards := [][][]int64{}
	board := [][]int64{}
	for i, line := range noEmptylines {
		if i%5 == 0 && i != 0 {
			boards = append(boards, board)
			board = [][]int64{}
		}

		// convert string of ints to int slice
		chars := strings.Split(line, " ")
		// Need to remove whitespace because of input formatting
		chars = removeEmpties(chars)
		numLine := []int64{}
		for _, char := range chars {
			num, err := strconv.ParseInt(char, 10, 64)
			if err != nil {
				panic(err)
			}
			numLine = append(numLine, num)
		}
		board = append(board, numLine)
	}
	// need to add final board
	boards = append(boards, board)
	return boards
}

func removeEmpties(s []string) []string {
	nonEmpty := []string{}
	for _, val := range s {
		if val != "" {
			nonEmpty = append(nonEmpty, val)
		}
	}
	return nonEmpty
}

func bingo(nums []int64, boards [][][]int64) int64 {
	for _, num := range nums {
		markBoards(num, boards)
		winner, board, _ := checkWinners(boards)
		if winner {
			score := calculateScore(num, board)
			return score
		}
	}
	return -1
}

func markBoards(num int64, boards [][][]int64) {
	for _, board := range boards {
		if len(board) == 0 {
			continue
		}
		for row := 0; row < 5; row++ {
			for col := 0; col < 5; col++ {
				if board[row][col] == num {
					board[row][col] = -1
				}
			}
		}
	}
}

func checkWinners(boards [][][]int64) (bool, [][]int64, int) {
	for idx, board := range boards {
		// Check each row
		for row := 0; row < 5; row++ {
			markCounter := 0
			for col := 0; col < 5; col++ {
				if board[row][col] == -1 {
					markCounter++
				}
			}
			if markCounter == 5 {
				return true, board, idx
			}
		}

		// Check each column
		for col := 0; col < 5; col++ {
			markCounter := 0
			for row := 0; row < 5; row++ {
				if board[row][col] == -1 {
					markCounter++
				}
			}
			if markCounter == 5 {
				return true, board, idx
			}
		}
	}
	return false, [][]int64{}, -1
}

func calculateScore(num int64, board [][]int64) int64 {
	sum := int64(0)
	for row := 0; row < 5; row++ {
		for col := 0; col < 5; col++ {
			if board[row][col] != -1 {
				sum += board[row][col]
			}
		}
	}
	return sum * num
}

func lastWinningBoardsScore(nums []int64, boards [][][]int64) int64 {
	for _, n := range nums {
		markBoards(n, boards)
		if len(boards) > 1 {
			winner, _, idx := checkWinners(boards)
			for winner {
				boards = remove(boards, idx)
				winner, _, idx = checkWinners(boards)
			}
		} else {
			w, board, _ := checkWinners(boards)
			if w {
				return calculateScore(n, board)
			}
		}
	}
	return -1
}

func remove(boards [][][]int64, idx int) [][][]int64 {
	copy(boards[idx:], boards[idx+1:])
	boards[len(boards)-1] = [][]int64{}
	boards = boards[:len(boards)-1]
	return boards
}
