package main

import (
	"fmt"
	"io/ioutil"
	"os"
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

func convert_to_int(in []string) []int {
	var out = make([]int, len(in))
	for i := range in {
		out[i], _ = strconv.Atoi(in[i])
	}
	return out
}

type Point struct {
	x int
	y int
}

func get_adjacents(p Point, queued *[][]int) []Point {
	width := len(*queued)
	height := len((*queued)[0])
	var ret []Point
	if p.x >= 1 {
		ret = append(ret, Point{p.x - 1, p.y})
	}
	if p.y >= 1 {
		ret = append(ret, Point{p.x, p.y - 1})
	}
	if p.x < width-1 {
		ret = append(ret, Point{p.x + 1, p.y})
	}
	if p.y < height-1 {
		ret = append(ret, Point{p.x, p.y + 1})
	}
	return filter_out_queued(ret, queued)
}

func filter_out_queued(in []Point, queued *[][]int) []Point {
	var out []Point
	for _, p := range in {
		if (*queued)[p.x][p.y] < 1 {
			out = append(out, p)
			(*queued)[p.x][p.y] = 1
		}
	}
	return out
}

func main() {
	if len(os.Args) < 2 {
		panic("Not enough command line arguments\n")
	}
	file_path := os.Args[1]
	lines := get_input(file_path)

	var heightmap = make([][]int, len(lines))
	for i := range heightmap {
		heightmap[i] = convert_to_int(strings.Split(lines[i], ""))
	}

	var low_points []Point
	for i := range heightmap {
		for j := range heightmap[i] {
			if i >= 1 && heightmap[i][j] >= heightmap[i-1][j] {
				continue
			}
			if j >= 1 && heightmap[i][j] >= heightmap[i][j-1] {
				continue
			}
			if i < len(heightmap)-1 && heightmap[i][j] >= heightmap[i+1][j] {
				continue
			}
			if j < len(heightmap[i])-1 && heightmap[i][j] >= heightmap[i][j+1] {
				continue
			}

			var p Point
			p.x = i
			p.y = j
			low_points = append(low_points, p)
		}
	}
	total_danger := 0
	for _, p := range low_points {
		total_danger += heightmap[p.x][p.y] + 1
	}
	fmt.Printf("Low points: %d. Total danger is %d\n", len(low_points), total_danger)

	total := 1
	var largest_basins = make([]int, 3)

	var queued = make([][]int, len(heightmap))
	for i := range heightmap {
		queued[i] = make([]int, len(heightmap[0]))
	}

	for _, p := range low_points {
		var queue []Point
		queued[p.x][p.y] = 1 // mark this space so we don't add it back in
		queue = append(queue, get_adjacents(p, &queued)...)
		basin_size := 1
		for len(queue) > 0 {
			if len(queue) > len(heightmap)*len(heightmap) {
				panic("No.")
			}
			m := queue[0]
			queue = queue[1:]
			if heightmap[m.x][m.y] == 9 {
				continue
			}
			basin_size++
			queue = append(queue, get_adjacents(m, &queued)...)
		}
		for i := range largest_basins {
			cur_basin_size := largest_basins[i]
			largest_basins = append(largest_basins[0:i], largest_basins[i+1:]...)
			if basin_size > cur_basin_size {
				largest_basins = append(largest_basins, basin_size)
				break
			} else {
				largest_basins = append(largest_basins, cur_basin_size)
			}
		}
	}
	for _, basin_size := range largest_basins {
		total *= basin_size
	}
	fmt.Printf("Total basin size product %d", total)

}
