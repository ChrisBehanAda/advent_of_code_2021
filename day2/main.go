package day2

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Solve() {
	input := getInputString("./day2/input.txt")
	instructions := strings.Split(input, "\n")
	x, y := finalDistances(instructions)
	x2, y2 := finalDistances2(instructions)
	fmt.Println("DAY 2")
	fmt.Println("Part 1")
	fmt.Printf("Final horizontal distance is %v.\nFinal vertical distance is %v.\n", x, y)
	fmt.Printf("Answer is %v\n", x*y)
	fmt.Println("Part 2")
	fmt.Printf("Final horizontal distance is %v.\nFinal vertical distance is %v.\n", x2, y2)
	fmt.Printf("Answer is %v\n", x2*y2)
}

func getInputString(path string) string {
	input, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(input)
}

// Part 1
func finalDistances(instructions []string) (int, int) {
	horizontal := 0
	vertical := 0
	for _, instruction := range instructions {
		if instruction == "" {
			continue
		}
		dirAndDistance := strings.Split(instruction, " ")
		dir := dirAndDistance[0]
		distance, err := strconv.Atoi(dirAndDistance[1])
		if err != nil {
			panic(err)
		}

		if dir == "forward" {
			horizontal += distance
		} else if dir == "up" {
			vertical -= distance
		} else {
			vertical += distance
		}
	}
	return horizontal, vertical
}

// Part 2
func finalDistances2(instructions []string) (int, int) {
	aim := 0
	horizontal := 0
	vertical := 0
	for _, instruction := range instructions {
		if instruction == "" {
			continue
		}
		dirAndDistance := strings.Split(instruction, " ")
		dir := dirAndDistance[0]
		distance, err := strconv.Atoi(dirAndDistance[1])
		if err != nil {
			panic(err)
		}

		if dir == "forward" {
			horizontal += distance
			vertical += aim * distance
		} else if dir == "up" {
			aim -= distance
		} else {
			aim += distance
		}
	}
	return horizontal, vertical

}
