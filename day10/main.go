package day10

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

var points = map[rune]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

var completionPoints = map[rune]int{
	')': 1,
	']': 2,
	'}': 3,
	'>': 4,
}

var open = map[rune]bool{
	'(': true,
	'[': true,
	'{': true,
	'<': true,
}

var pairs = map[rune]rune{
	')': '(',
	']': '[',
	'}': '{',
	'>': '<',
}

var closer = map[rune]rune{
	'(': ')',
	'[': ']',
	'{': '}',
	'<': '>',
}

func Solve() {
	input := getInput("./day10/input.txt")
	lines := strings.Split(input, "\n")
	fmt.Printf("total lines: %v\n", len(lines))
	score := syntaxScore(lines)
	fmt.Printf("Syntax score is %v\n", score)
	incompleteLines := filterCorruptLines(lines)
	completionChars := closingChars(incompleteLines)
	completionScores := completionScores(completionChars)
	medianScore := medianScore(completionScores)
	fmt.Printf("Median score is %v\n", medianScore)

}

func getInput(path string) string {
	input, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	trimmedInput := strings.Trim(string(input), "\n")
	return trimmedInput
}

func syntaxScore(lines []string) int {
	score := 0
	stack := []rune{}
	for _, line := range lines {
		for _, r := range line {
			if open[r] {
				stack = append(stack, r)
			} else {
				if len(stack) > 0 {
					n := len(stack) - 1
					matching := stack[n]
					stack = stack[:n]

					if pairs[r] != matching {
						score += points[r]
						break
					}
				} else {
					score += points[r]
					break
				}
			}
		}
	}
	return score
}

func filterCorruptLines(lines []string) []string {
	incomplete := []string{}
	stack := []rune{}
	for _, line := range lines {
		corrupt := false
		for _, r := range line {
			if open[r] {
				stack = append(stack, r)
			} else {
				if len(stack) > 0 {
					n := len(stack) - 1
					matching := stack[n]
					stack = stack[:n]

					if pairs[r] != matching {
						corrupt = true
						break
					}
				} else {
					corrupt = true
					break
				}
			}
		}
		if !corrupt {
			incomplete = append(incomplete, line)
		}
	}
	return incomplete
}

func closingChars(lines []string) [][]rune {
	completionChunks := [][]rune{}
	for _, line := range lines {
		stack := []rune{}
		for _, r := range line {
			if open[r] {
				stack = append(stack, r)
			} else {
				stack = stack[:len(stack)-1]
			}
		}
		chunk := []rune{}
		for i := len(stack) - 1; i >= 0; i-- {
			completer := closer[rune(stack[i])]
			chunk = append(chunk, completer)
		}
		completionChunks = append(completionChunks, chunk)
	}
	return completionChunks
}

func completionScores(input [][]rune) []int {
	scores := []int{}
	for _, runes := range input {
		total := 0
		for _, r := range runes {
			total = (total * 5) + completionPoints[r]
		}
		scores = append(scores, total)
	}
	return scores
}

func medianScore(scores []int) int {
	sort.Ints(scores)
	mid := len(scores) / 2
	return scores[mid]
}
