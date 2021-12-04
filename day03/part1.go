package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parse_line(line string) []int {
	var columns []int
	for _, value := range line {
		numValue, err := strconv.Atoi(string(value))
		if err != err {
			panic(err)
		}
		columns = append(columns, numValue)
	}
	return columns
}
func calculate_gamma(totals []int, threshold int) int64 {
	gamma := ""
	for _, value := range totals {
		if value > threshold {
			gamma += "1"
		} else {
			gamma += "0"
		}
	}
	out, _ := strconv.ParseInt(gamma, 2, 64)
	return out
}

func calculate_epsilon(totals []int, threshold int) int64 {
	epsilon := ""
	for _, value := range totals {
		if value >= threshold {
			epsilon += "0"
		} else {
			epsilon += "1"
		}
	}
	out, _ := strconv.ParseInt(epsilon, 2, 64)
	return out
}

func main() {
	if len(os.Args) < 2 {
		panic("Not enough command line arguments\n")
	}
	file_path := os.Args[1]
	lines, err := os.ReadFile(file_path)
	if err != nil {
		panic(err)
	}
	var power_readings [][]int
	for _, line := range strings.Split(string(lines), "\n") {
		power_readings = append(power_readings, parse_line(line))
	}

	var totals []int
	for i, row := range power_readings {
		for j := range row {
			if j >= len(totals) {
				totals = append(totals, 0)
			}
			totals[j] += power_readings[i][j]
		}
	}
	gamma := calculate_gamma(totals, len(power_readings)/2)
	epsilon := calculate_epsilon(totals, len(power_readings)/2)

	fmt.Printf("%d x %d\n", gamma, epsilon)
	fmt.Println(gamma * epsilon)
}
