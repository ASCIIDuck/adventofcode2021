package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

func get_input(file_path string) []string {
	input, err := ioutil.ReadFile(file_path)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(strings.TrimSpace(string(input)), "\n")

	return lines
}

func get_keys(m map[string]string) []string {
	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	return keys
}

func get_values(m map[string]string) []string {
	var values []string
	for _, value := range m {
		values = append(values, value)
	}
	return values
}

func in(s string, arr []string) bool {
	for _, cur := range arr {
		if cur == s {
			return true
		}
	}
	return false
}

func main() {
	if len(os.Args) < 2 {
		panic("Not enough command line arguments\n")
	}
	file_path := os.Args[1]
	lines := get_input(file_path)
	corrupt_points := map[string]int{
		")": 3,
		"]": 57,
		"}": 1197,
		">": 25137,
	}
	pairs := map[string]string{
		"(": ")",
		"[": "]",
		"{": "}",
		"<": ">",
	}
	opens := get_keys(pairs)
	closes := get_values(pairs)
	fmt.Println(opens)
	fmt.Println(closes)

	corrupt_total := 0
	var complete_lines []string
	for _, chunk := range lines {
		var queue []string
		corrupt_value := 0
		for _, char := range chunk {
			if in(string(char), opens) {
				queue = append(queue, string(char))
			} else if in(string(char), closes) {
				opener := queue[len(queue)-1]
				queue = queue[0 : len(queue)-1]
				if pairs[opener] != string(char) {
					corrupt_value = corrupt_points[string(char)]
					break
				}
			}
		}
		if corrupt_value == 0 {
			complete_lines = append(complete_lines, chunk)
		}
		corrupt_total += corrupt_value
	}

	complete_points := map[string]int{
		")": 1,
		"]": 2,
		"}": 3,
		">": 4,
	}
	var complete_totals = make([]int, len(complete_lines))
	for i, chunk := range complete_lines {
		var queue []string

		for _, char := range chunk {
			if in(string(char), opens) {
				queue = append(queue, string(char))
			} else if in(string(char), closes) {
				queue = queue[0 : len(queue)-1]
			}
		}
		for j := len(queue) - 1; j >= 0; j-- {
			complete_totals[i] = complete_totals[i]*5 + complete_points[pairs[queue[j]]]
		}
	}
	sort.Ints(complete_totals)
	fmt.Println(complete_totals)
	middle := len(complete_totals) / 2
	fmt.Printf("Total for p1 %d\n", corrupt_total)
	fmt.Printf("Total for p2 %d\n", complete_totals[middle])
}
