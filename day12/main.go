package day12

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

func Solve() {
	input := getInput("./day12/input.txt")
	adjList := adjList(input)
	fmt.Printf("Adjacency List: %v\n", adjList)
	paths := countPaths(adjList)
	fmt.Printf("There are %v unique paths.\n", paths)
	paths2 := countPaths2(adjList)
	fmt.Printf("There are %v unique paths when visiting one small cave twice.\n", paths2)
}

func getInput(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	input := strings.Trim(string(data), "\n")
	return input
}

func adjList(input string) map[string][]string {
	adjList := make(map[string][]string)
	edges := strings.Split(input, "\n")
	for _, e := range edges {
		nodes := strings.Split(e, "-")
		from, to := nodes[0], nodes[1]
		adjList[from] = append(adjList[from], to)
		adjList[to] = append(adjList[to], from)
	}
	return adjList
}

func countPaths(adjList map[string][]string) int {
	count := 0
	visited := make(map[string]bool)
	visited["start"] = true
	dfs("start", adjList, visited, &count)
	return count
}

func dfs(node string, adjList map[string][]string, visited map[string]bool, count *int) {
	if node == "end" {
		*count++
		return
	}
	// check if small cave
	if unicode.IsLower(rune(node[0])) {
		visited[node] = true
	}

	for _, cave := range adjList[node] {
		if !visited[cave] {
			v := copyMap(visited)
			dfs(cave, adjList, v, count)
		}
	}
}

func countPaths2(adjList map[string][]string) int {
	count := 0
	visited := make(map[string]int)
	for k := range visited {
		visited[string(k)] = 0
	}
	visited["start"] = 2
	dfs2("start", adjList, visited, &count)
	return count
}

func dfs2(node string, adjList map[string][]string, visited map[string]int, count *int) {
	if node == "end" {
		*count++
		return
	}
	// check if small cave
	if unicode.IsLower(rune(node[0])) {
		visited[node]++
	}

	for _, cave := range adjList[node] {
		if visited[cave] == 0 || visited[cave] == 1 && !hasTwo(visited) {
			v := copyMap2(visited)
			dfs2(cave, adjList, v, count)
		}
	}
}

func copyMap(m map[string]bool) map[string]bool {
	copy := make(map[string]bool)
	for key, value := range m {
		copy[key] = value
	}
	return copy
}

func copyMap2(m map[string]int) map[string]int {
	copy := make(map[string]int)
	for key, value := range m {
		copy[key] = value
	}
	return copy
}

func hasTwo(m map[string]int) bool {
	for _, val := range m {
		if val == 2 {
			return true
		}
	}
	return false
}
