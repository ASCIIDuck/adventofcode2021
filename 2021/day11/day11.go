package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
}

type Octopus struct {
	power int
	flash bool
}

func get_input(file_path string) []string {
	input, err := ioutil.ReadFile(file_path)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(strings.TrimSpace(string(input)), "\n")

	return lines
}

func convert_to_octo(in []string) []Octopus {
	var out = make([]Octopus, len(in))
	for i := range in {
		x, _ := strconv.Atoi(in[i])
		out[i] = Octopus{x, false}
	}
	return out
}

func adjacent_points(i int, j int, wdth int, hei int) []Point {
	candidate_points := []Point{
		{i, j - 1},
		{i, j + 1},
		{i - 1, j - 1},
		{i - 1, j},
		{i - 1, j + 1},
		{i + 1, j - 1},
		{i + 1, j},
		{i + 1, j + 1},
	}

	var adjacent_points []Point
	for _, p := range candidate_points {
		if p.x < 0 || p.x >= wdth || p.y < 0 || p.y >= hei {
			///////////////fmt.Printf("Rejecting point %dx%d\n", p.x, p.y)
			continue
		}
		adjacent_points = append(adjacent_points, p)
	}
	return adjacent_points
}

func charge(octopuses [][]Octopus, cur Point) [][]Octopus {
	octopuses[cur.x][cur.y].power++
	if octopuses[cur.x][cur.y].power > 9 && octopuses[cur.x][cur.y].flash != true {
		octopuses[cur.x][cur.y].flash = true
		for _, p := range adjacent_points(cur.x, cur.y, len(octopuses), len(octopuses[cur.x])) {
			octopuses = charge(octopuses, p)
		}
	}
	return octopuses
}

func do_step(octopuses [][]Octopus) (int, [][]Octopus) {
	flash_count := 0

	// increment
	for i := range octopuses {
		for j := range octopuses[i] {
			octopuses = charge(octopuses, Point{i, j})
		}
	}
	// find flashes and count
	for i := range octopuses {
		for j := range octopuses[i] {
			if octopuses[i][j].power > 9 {
				octopuses[i][j].power = 0
				octopuses[i][j].flash = false
				flash_count++
			}
		}
	}

	return flash_count, octopuses
}

func print_grid(octopuses [][]Octopus) {
	colorRed := "\033[31m"
	colorReset := "\033[0m"
	for _, line := range octopuses {
		for _, o := range line {
			if o.power == 0 {
				fmt.Print(string(colorRed), o.power, string(colorReset))
			} else {
				fmt.Print(o.power)
			}
		}
		fmt.Print("\n")
	}
}

func check_all(octopuses [][]Octopus) bool {
	for _, line := range octopuses {
		for _, o := range line {
			if o.power != 0 {
				return false
			}
		}
	}
	return true
}

func main() {
	if len(os.Args) < 2 {
		panic("Not enough command line arguments\n")
	}
	file_path := os.Args[1]
	lines := get_input(file_path)
	var octopuses = make([][]Octopus, len(lines))
	for i := range octopuses {
		octopuses[i] = convert_to_octo(strings.Split(lines[i], ""))
	}
	print_grid(octopuses)
	fmt.Println("-----")

	total_flash_count := 0
	all_flashed_together := check_all(octopuses)
	step := 0
	for all_flashed_together == false {
		step++
		flash_count := 0
		flash_count, octopuses = do_step(octopuses)
		total_flash_count += flash_count
		if step == 100 {
			fmt.Printf("Part 1 flash count %d\n", total_flash_count)
		}
		all_flashed_together = check_all(octopuses)
	}
	fmt.Printf("Part 2 all flash step %d\n", step)

}
