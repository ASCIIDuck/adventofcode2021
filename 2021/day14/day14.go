package main

import (
	"fmt"
	"io/ioutil"
	"os"
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

func expand_template(polymer string, rules map[string]string) string {
	output := string(polymer[0])
	for i := range polymer {
		if i+1 >= len(polymer) {
			break
		}
		output += rules[polymer[i:i+2]] + string(polymer[i+1])
	}

	return output
}

func expand_pairs(polymer string, rules map[string]string, depth int) map[string]int {
	pairs := string_to_pairs(polymer)

	for i := 0; i < depth; i++ {
		var new_pairs = make(map[string]int)
		for pair, count := range pairs {
			expanded_pairs := string_to_pairs(string(pair[0]) + rules[pair] + string(pair[1]))
			for p, c := range expanded_pairs {
				new_pairs[p] += count * c
			}
		}
		pairs = new_pairs
	}

	return pairs
}

func sum_pairs(a map[string]int, b map[string]int) map[string]int {
	var out = make(map[string]int)
	for p, v := range a {
		out[p] = v
	}
	for p, v := range b {
		out[p] += v
	}
	return out
}

func string_to_pairs(polymer string) map[string]int {
	var output = make(map[string]int)
	for i := range polymer {
		if i == len(polymer)-1 {
			break
		}
		output[polymer[i:i+2]]++
	}
	return output
}

func calculate_score(pairs map[string]int) int {
	min := 99999999999999999
	max := 0
	var chars = make(map[string]int)
	for k, v := range pairs {
		chars[string(k[0])] += v
		chars[string(k[1])] += v
	}
	for _, v := range chars {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}

	return (max - min) / 2
}

func main() {
	if len(os.Args) < 2 {
		panic("Not enough command line arguments\n")
	}
	file_path := os.Args[1]
	lines := get_input(file_path)

	polymer := lines[0]
	var rules = make(map[string]string)
	lines = lines[2:]
	for _, line := range lines {
		parts := strings.Split(line, " -> ")
		rules[parts[0]] = parts[1]
	}

	seed_pair := string(polymer[0]) + string(polymer[len(polymer)-1])
	part1_counts := expand_pairs(polymer, rules, 10)
	part1_counts[seed_pair]++
	part2_counts := expand_pairs(polymer, rules, 40)
	part2_counts[seed_pair]++

	fmt.Printf("Score for part 1: %d\n", calculate_score(part1_counts))
	fmt.Printf("Score for part 2: %d\n", calculate_score(part2_counts))

}
