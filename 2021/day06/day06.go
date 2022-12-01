package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func get_input(file_path string) []string {
	input, err := ioutil.ReadFile(file_path)
	if err != nil {
		panic(err)
	}

	space_re := regexp.MustCompile(`\s+`)
	lines := strings.Split(strings.TrimSpace(space_re.ReplaceAllString(string(input), " ")), "\n")

	return lines
}
func convert_to_int(in []string) []int {
	var out = make([]int, len(in))
	for i := range in {
		out[i], _ = strconv.Atoi(in[i])
	}
	return out
}

func sum(in []int) int {
	total := 0
	for _, x := range in {
		total += x
	}
	return total
}

func main() {
	if len(os.Args) < 2 {
		panic("Not enough command line arguments\n")
	}
	file_path := os.Args[1]
	lines := get_input(file_path)

	fish_lives := convert_to_int(strings.Split(lines[0], ","))
	var fish = make([]int, 10)
	for i := range fish_lives {
		fish[fish_lives[i]]++
	}

	var num_of_fish = make([]int, 256)

	for i := range num_of_fish {
		var new_fish = make([]int, len(fish))
		for j := range fish {
			if j == 0 {
				new_fish[8] = fish[j]
				new_fish[6] = fish[j]
			} else {
				new_fish[j-1] += fish[j]
			}

		}
		fish = new_fish
		num_of_fish[i] = sum(fish)
	}
	fmt.Printf("After 80 days there are %d fish\n", num_of_fish[79])
	fmt.Printf("After 256 days there are %d fish\n", num_of_fish[255])

}
