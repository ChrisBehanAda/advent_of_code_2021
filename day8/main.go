package day8

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var segments = map[int]string{
	0: "abcefg",
	1: "cf",
	2: "acdeg",
	3: "acdfg",
	4: "bcdf",
	5: "abdfg",
	6: "abdefg",
	7: "acf",
	8: "abcdefg",
	9: "abcdfg",
}

func Solve() {
	input := getInputString("./day8/input.txt")
	signalPatterns, outputVals := parseInput(input)
	fmt.Printf("SIGNAL PATTERNS: %v\n", signalPatterns)
	fmt.Printf("OUTPUT VALUES: %v\n", outputVals)
	easyLengths := easySegmentLengths()
	easyCount := easyCount(outputVals, easyLengths)
	fmt.Printf("the easy lengths are: %v\n", easyLengths)
	fmt.Printf("There are %v occurrences of 1,4, 7 or 8.\n", easyCount)
	sortByLengthAndContents(signalPatterns)
	sortContents(outputVals)
	fmt.Printf("SORTED SIGNAL PATTERNS: %v\n", signalPatterns)
	fmt.Printf("SORTED OUTPUT VAULES: %v\n", outputVals)
	decodedOutputVals := decodedOutputVals(signalPatterns, outputVals)
	fmt.Printf("Decoded output values are: %v\n", decodedOutputVals)
	sum := sumDecodedVals(decodedOutputVals)
	fmt.Printf("Sum of decoded output values is: %v\n", sum)
}

func getInputString(path string) string {
	input, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	trimmedInput := strings.Trim(string(input), "\n")
	return trimmedInput
}

func parseInput(input string) ([][]string, [][]string) {
	lines := strings.Split(input, "\n")
	signalPatterns := [][]string{}
	outputVals := [][]string{}
	for _, line := range lines {
		splitLine := strings.Split(line, "|")
		signalPatternLine := strings.Trim(splitLine[0], " ")
		pattern := strings.Split(signalPatternLine, " ")
		signalPatterns = append(signalPatterns, pattern)
		outputValueLine := strings.Trim(splitLine[1], " ")
		val := strings.Split(outputValueLine, " ")
		outputVals = append(outputVals, val)
	}
	return signalPatterns, outputVals
}

func easySegmentLengths() []int {
	easyLengths := []int{}
	easySegKeys := []int{1, 4, 7, 8}
	for _, key := range easySegKeys {
		easyLengths = append(easyLengths, len(segments[key]))
	}
	return easyLengths
}

func easyCount(outputVals [][]string, easyLengths []int) int {
	count := 0
	for _, output := range outputVals {
		for _, val := range output {
			for _, e := range easyLengths {
				if len(val) == e {
					count++
				}
			}
		}
	}
	return count
}

// Sorts a slice of string slices by length of the strings
// and sorts the characters within each string lexicographically.
func sortByLengthAndContents(strs [][]string) {
	for _, strSlice := range strs {
		// sorts string slice by length of strs
		sort.Slice(strSlice, func(i, j int) bool {
			return len(strSlice[i]) <= len(strSlice[j])
		})
	}
	sortContents(strs)
}

func sortContents(strs [][]string) {
	for a, strSlice := range strs {
		// sorts contents of strings lexicographically
		for b, str := range strSlice {
			runes := []rune(str)
			sort.Slice(runes, func(i, j int) bool {
				return runes[i] <= runes[j]
			})
			s := string(runes)
			strs[a][b] = s
		}
	}
}

func decodedOutputVals(signals [][]string, output [][]string) [][]int {
	patternMap := make(map[string]int)
	digitMap := make(map[int]string)
	// Create 2d slice of dimensions output len X 4
	decodedVals := make([][]int, len(output))
	for i := range decodedVals {
		decodedVals[i] = make([]int, 4)
	}

	for eIdx, entry := range signals {
		for _, pattern := range entry {
			if len(pattern) == 2 {
				patternMap[pattern] = 1
				digitMap[1] = pattern
			} else if len(pattern) == 3 {
				patternMap[pattern] = 7
				digitMap[7] = pattern
			} else if len(pattern) == 4 {
				patternMap[pattern] = 4
				digitMap[4] = pattern
			} else if len(pattern) == 7 {
				patternMap[pattern] = 8
				digitMap[8] = pattern
			} else if len(pattern) == 5 {
				if differingChars(digitMap[7], pattern) == 0 {
					patternMap[pattern] = 3
					digitMap[3] = pattern
				} else if differingChars(pattern, digitMap[4]) == 3 {
					patternMap[pattern] = 2
					digitMap[2] = pattern
				} else {
					patternMap[pattern] = 5
					digitMap[5] = pattern
				}
			} else if len(pattern) == 6 {
				if differingChars(digitMap[4], pattern) == 0 {
					patternMap[pattern] = 9
					digitMap[9] = pattern
				} else if differingChars(digitMap[7], pattern) == 0 {
					patternMap[pattern] = 0
					digitMap[0] = pattern
				} else {
					patternMap[pattern] = 6
					digitMap[6] = pattern
				}
			}
		}
		outputVals := output[eIdx]
		for i, vals := range outputVals {
			decodedVals[eIdx][i] = patternMap[vals]
		}
	}
	return decodedVals
}

// Returns the number of characters in s1 that are not in s2.
func differingChars(s1, s2 string) int {
	count := len(s1)
	for _, c1 := range s1 {
		found := false
		for _, c2 := range s2 {
			if c1 == c2 {
				found = true
				break
			}
		}
		if found {
			count--
		}
	}
	return count
}

func sumDecodedVals(vals [][]int) int {
	sum := 0
	for _, v := range vals {
		str := ""
		for _, i := range v {
			str += strconv.Itoa(i)
		}
		intVal, _ := strconv.Atoi(str)
		sum += intVal
	}
	return sum
}
