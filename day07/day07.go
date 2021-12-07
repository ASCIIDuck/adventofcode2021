package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func convert_to_int(in []string) []int {
	var out = make([]int, len(in))
	for i := range in {
		out[i], _ = strconv.Atoi(in[i])
	}
	return out
}

func get_input(file_path string) []string {
	input, err := ioutil.ReadFile(file_path)
	if err != nil {
		panic(err)
	}

	space_re := regexp.MustCompile(`\s+`)
	lines := strings.Split(strings.TrimSpace(space_re.ReplaceAllString(string(input), " ")), "\n")

	return lines
}

func get_max(list []int) int {
	max := list[0]
	for _, n := range list {
		if n > max {
			max = n
		}
	}
	return max
}

func get_min(list []int) int {
	min := list[0]
	for _, n := range list {
		if n < min {
			min = n
		}
	}
	return min
}

func incremental_cost(num_of_steps int) int {
	return (num_of_steps * (num_of_steps + 1)) / 2
}

func linear_cost(num_of_steps int) int {
	return num_of_steps
}

func calculate_fuel_cost(depths []int, target int, cost_func func(int) int) int {
	cost := 0
	for _, d := range depths {
		number_of_steps := int(math.Abs(float64(target - d)))
		cost += cost_func(number_of_steps)
	}
	return cost
}

func find_minimum_fuel(depths []int, cost_func func(int) int) int {
	min := get_min(depths)
	max := get_max(depths)
	fuel_cost := calculate_fuel_cost(depths, min, cost_func)

	for i := min; i <= max; i += 1 {
		cur_fuel_cost := calculate_fuel_cost(depths, i, cost_func)
		if fuel_cost > cur_fuel_cost {
			fuel_cost = cur_fuel_cost
		}

	}
	return fuel_cost
}

func main() {
	if len(os.Args) < 2 {
		panic("Not enough command line arguments\n")
	}
	file_path := os.Args[1]
	lines := get_input(file_path)

	depths := convert_to_int(strings.Split(lines[0], ","))

	fmt.Printf("Fuel cost for part 1 is %d\n", find_minimum_fuel(depths, linear_cost))
	fmt.Printf("Fuel cost for part 2 at %d\n", find_minimum_fuel(depths, incremental_cost))
}
