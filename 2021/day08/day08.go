package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
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

func get_value(signal string) int {
	if len(signal) == 2 {
		return 1
	} else if len(signal) == 4 {
		return 4
	} else if len(signal) == 3 {
		return 7
	} else if len(signal) == 7 {
		return 8
	}
	return -1
}

func subtract_set(set string, sub string) string {
	for _, i := range sub {
		index_of_i := strings.Index(set, string(i))
		if index_of_i >= 0 {
			set = set[0:index_of_i] + set[index_of_i+1:]
		}
	}
	return set
}

func sort_string(in string) string {
	components := strings.Split(in, "")
	sort.Strings(components)
	return strings.Join(components, "")
}

func find_value(signal string, mapping []string) int {
	for i, candidate := range mapping {
		if sort_string(signal) == sort_string(candidate) {
			return i
		}
	}
	return -1
}

func main() {
	if len(os.Args) < 2 {
		panic("Not enough command line arguments\n")
	}
	file_path := os.Args[1]
	lines := get_input(file_path)

	total := 0
	for _, line := range lines {
		parts := strings.Split(line, "|")
		testing_singals := strings.Split(strings.Trim(parts[0], " "), " ")
		output_values := strings.Split(strings.Trim(parts[1], " "), " ")
		var mapping = make([]string, 10)

		// one pass to get the key-values
		for _, signal := range testing_singals {
			value := get_value(signal)
			if value < 0 {
				continue
			}
			mapping[value] = signal
		}

		// another pass to map the rest
		for _, signal := range testing_singals {
			value := get_value(signal)
			if value > 0 {
				continue
			}
			if len(signal) == 6 {
				// could be 0, 6 or 9
				zero_six_test_set := subtract_set(subtract_set(mapping[8], mapping[7]), mapping[4])
				zero_six_test := subtract_set(signal, zero_six_test_set)
				if len(zero_six_test) == 4 {
					//could be 0 or 6
					sub_one := subtract_set(zero_six_test, mapping[1])
					if len(sub_one) == 2 {
						// is 0
						mapping[0] = signal
					} else if len(sub_one) == 3 {
						// is 6
						mapping[6] = signal
					}

				} else {
					// is 9
					mapping[9] = signal
				}
			} else if len(signal) == 5 {
				sub_one := subtract_set(signal, mapping[1])
				if len(sub_one) == 3 {
					// it's a 3
					mapping[3] = signal
				} else {
					// actually same as the zero_six_test_set
					five_test_set := subtract_set(subtract_set(mapping[8], mapping[7]), mapping[4])
					five_test := subtract_set(signal, five_test_set)
					if len(five_test) == 4 {
						// is 5
						mapping[5] = signal
					} else {
						mapping[2] = signal
					}
				}
			}
		}
		output_str := ""
		for _, signal := range output_values {
			value := find_value(signal, mapping)
			value_str := strconv.FormatInt(int64(value), 10)
			output_str += value_str
		}
		output_val, _ := strconv.Atoi(output_str)
		total += output_val
	}
	fmt.Println(total)
}
