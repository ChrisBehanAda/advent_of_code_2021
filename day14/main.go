package day14

import (
	"fmt"
	"math"
	"os"
	"strings"
)

type rule struct {
	left      rune
	right     rune
	insertion rune
}

func Solve() {
	input := getInput("./day14/input.txt")
	template, rules := templateAndRules(input)
	result := react(template, rules, 10)
	c1 := counts(result)
	min, max := maxAndMin(c1)
	ans := max - min
	fmt.Printf("Quantity of most common element minus least common element after 10 steps is %v\n", ans)
	result2 := react2(template, rules, 40)
	c2 := counts2(result2, template)
	min2, max2 := maxAndMin(c2)
	ans2 := max2 - min2
	fmt.Printf("Quantity of most common element minus least common element after 40 steps is %v\n", ans2)
}

func getInput(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	input := strings.Trim(string(data), "\n")
	return input
}

func templateAndRules(input string) (string, map[string]string) {
	inputs := strings.Split(input, "\n\n")
	template := inputs[0]
	rules := make(map[string]string)
	ruleInputs := strings.Split(inputs[1], "\n")
	for _, i := range ruleInputs {
		pairAndInsert := strings.Split(i, " -> ")
		pairInput := pairAndInsert[0]
		insert := pairAndInsert[1]
		rules[pairInput] = insert
	}
	return template, rules
}

func react(template string, rules map[string]string, steps int) string {
	for s := 0; s < steps; s++ {
		l, r := 0, 1
		for r < len(template) {
			ins := rules[string(template[l])+string(template[r])]
			template = template[:l+1] + ins + template[r:]
			l += 2
			r += 2
		}
	}
	return template
}

func counts(template string) map[rune]int {
	counts := make(map[rune]int)
	for _, r := range template {
		_, exists := counts[r]
		if exists {
			counts[r]++
		} else {
			counts[r] = 1
		}
	}
	return counts
}

func maxAndMin(counts map[rune]int) (int, int) {
	min, max := math.MaxInt, math.MinInt
	for _, val := range counts {
		if val < min {
			min = val
		}
		if val > max {
			max = val
		}
	}
	return min, max
}

func react2(template string, rules map[string]string, steps int) map[string]int {
	pairs := make(map[string]int)
	for i := 0; i < len(template)-1; i++ {
		r1, r2 := reactions(template[i:i+2], rules)
		pairs[r1]++
		pairs[r2]++
	}
	for i := 0; i < steps-1; i++ {
		newPairs := make(map[string]int)
		for pair, val := range pairs {
			r1, r2 := reactions(pair, rules)
			newPairs[r1] += val
			newPairs[r2] += val
		}
		pairs = newPairs
	}
	return pairs
}

func reactions(pair string, rules map[string]string) (string, string) {
	r := rules[pair]
	return string(pair[0]) + r, r + string(pair[1])
}

func test() {
	m := make(map[string]int)
	m["a"]++
	fmt.Print(m)
}

func counts2(m map[string]int, template string) map[rune]int {
	counts := make(map[rune]int)
	for pair, val := range m {
		counts[rune(pair[0])] += val
	}
	lastChar := rune(template[len(template)-1])
	counts[lastChar]++
	return counts
}
