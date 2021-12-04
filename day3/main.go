package day3

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Solve() {
	input := getInputString("./day3/input.txt")
	readings := strings.Split(input, "\n")
	gammaSum := sum(readings, false)
	epsilonSum := sum(readings, true)
	fmt.Println("DAY 3")
	fmt.Printf("Gamma sum is: %v\n", gammaSum)
	fmt.Printf("Epsilon sum is: %v\n", epsilonSum)
	fmt.Printf("Power Consumption is: %v\n", gammaSum*epsilonSum)
	oxygen := oxygenRating(readings)
	co2 := co2Rating(readings)
	fmt.Printf("Oxygen rating is %v\nCO2 rating is %v\n", oxygen, co2)
	fmt.Printf("Life support rating is %v\n", oxygen*co2)
}

func getInputString(path string) string {
	input, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	inputString := strings.Trim(string(input), "\n")
	return inputString
}

// Part 1
func sum(readings []string, leastCommon bool) int64 {
	bits := commonBits(readings, leastCommon)
	bitString := ""
	for _, bit := range bits {
		bitString = bitString + strconv.Itoa(bit)
	}

	sum, err := strconv.ParseInt(bitString, 2, 64)
	if err != nil {
		panic(err)
	}
	return sum
}

func commonBits(readings []string, leastCommon bool) [12]int {
	bitCounts := bitCounts(readings)
	commonBits := [12]int{}
	for idx, count := range bitCounts {
		if leastCommon {
			if count < 0 {
				commonBits[idx] = 1
			}
		} else {
			if count > 0 {
				commonBits[idx] = 1
			}
		}
	}
	return commonBits
}

func bitCounts(readings []string) [12]int {
	bitCounts := [12]int{}
	for _, bits := range readings {
		for idx, bit := range bits {
			if bit == '1' {
				bitCounts[idx]++
			} else {
				bitCounts[idx]--
			}
		}
	}
	return bitCounts
}

func oxygenRating(readings []string) int64 {
	remaining := readings
	idx := 0
	for len(remaining) != 1 {
		bitCounts := bitCounts(remaining)
		matches := []string{}
		for _, reading := range remaining {
			if bitCounts[idx] > 0 && reading[idx] == '1' {
				matches = append(matches, reading)
			} else if bitCounts[idx] < 0 && reading[idx] == '0' {
				matches = append(matches, reading)
			} else if bitCounts[idx] == 0 && reading[idx] == '1' {
				matches = append(matches, reading)
			}
		}
		remaining = matches
		idx++
	}
	rating, err := strconv.ParseInt(remaining[0], 2, 64)
	if err != nil {
		panic(err)
	}
	return rating
}

func co2Rating(readings []string) int64 {
	remaining := readings
	idx := 0
	for len(remaining) != 1 {
		bitCounts := bitCounts(remaining)
		matches := []string{}
		for _, reading := range remaining {
			if bitCounts[idx] > 0 && reading[idx] == '0' {
				matches = append(matches, reading)
			} else if bitCounts[idx] < 0 && reading[idx] == '1' {
				matches = append(matches, reading)
			} else if bitCounts[idx] == 0 && reading[idx] == '0' {
				matches = append(matches, reading)
			}
		}
		remaining = matches
		idx++
	}
	rating, err := strconv.ParseInt(remaining[0], 2, 64)
	if err != nil {
		panic(err)
	}
	return rating
}
