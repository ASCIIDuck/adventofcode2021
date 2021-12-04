package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parse_line(line string) []string {
	return strings.Split(line, "")
}

func most_common(readings [][]string, column_no int) string {
	total := 0
	for i := range readings {
		num, _ := strconv.Atoi(readings[i][column_no])
		total += num
	}
	if float64(total) >= float64(len(readings))/2 {
		return "1"
	} else {
		return "0"
	}
}

func least_common(readings [][]string, column_no int) string {
	if most_common(readings, column_no) == "1" {
		return "0"
	} else {
		return "1"
	}
}

func yeet(matrix [][]string, row_no int) [][]string {
	return append(matrix[0:row_no], matrix[row_no+1:]...)
}

func find_reading(readings [][]string, chooser func([][]string, int) string) []string {
	column := 0
	for len(readings) > 1 {
		chosen_value := chooser(readings, column)
		i := 0
		for i < len(readings) && len(readings) > 1 {
			if readings[i][column] != chosen_value {
				readings = yeet(readings, i)
			} else {
				i += 1
			}
		}
		column += 1
	}
	return readings[0]
}

func string_to_binary_int(in []string) int64 {
	n, _ := strconv.ParseInt(strings.Join(in, ""), 2, 64)
	return n
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
	var readings [][]string
	for _, line := range strings.Split(string(lines), "\n") {
		readings = append(readings, parse_line(line))
	}

	copy_readings := make([][]string, len(readings))
	copy(copy_readings, readings)
	o2_reading := string_to_binary_int(find_reading(copy_readings, most_common))
	copy(copy_readings, readings)
	co2_reading := string_to_binary_int(find_reading(copy_readings, least_common))
	fmt.Println(o2_reading * co2_reading)
}
