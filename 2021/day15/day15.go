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

func get_candidates(x int, y int) []Point {
	candidates := []Point{}
	if x >= 1 {
		candidates = append(candidates, Point{x - 1, y})
	}
	if y >= 1 {
		candidates = append(candidates, Point{x, y - 1})
	}
	return candidates

}

func min_risk(candidates []Point, riskmap [][]int) Point {
	min_risk := -1
	min_risk_point := Point{}
	for _, c := range candidates {
		if riskmap[c.x][c.y] < min_risk || min_risk == -1 {
			min_risk = riskmap[c.x][c.y]
			min_risk_point = c
		}
	}
	return min_risk_point
}

func makeMatrix(wei int, hei int, def int) [][]int {
	out := make([][]int, wei)
	for i := range out {
		out[i] = make([]int, hei)
		for j := range out[i] {
			out[i][j] = def
		}
	}
	return out
}

func getNeighbors(p Point, n int) []Point {
	dirs := []Point{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	}

	output := []Point{}
	for _, d := range dirs {
		c := Point{p.x + d.x, p.y + d.y}
		if c.x >= 0 && c.y >= 0 && c.x < n && c.y < n {
			output = append(output, c)
		}

	}
	return output
}

func find_lowest(queue []Point, pri [][]int) int {
	lowest := -1
	lowest_index := 0
	for i, cur := range queue {
		risk := pri[cur.x][cur.y]
		if (risk >= 0 && risk < lowest) || lowest == -1 {
			lowest = pri[cur.x][cur.y]
			lowest_index = i
		}
	}

	return lowest_index
}

func solve(riskmap [][]int) int {
	n := len(riskmap)
	total_risk := makeMatrix(n, n, -1)
	total_risk[0][0] = 0
	queue := []Point{{0, 0}}

	var cur Point
	for len(queue) > 0 {
		i := find_lowest(queue, total_risk)
		cur = queue[i]
		queue = append(queue[0:i], queue[i+1:]...)
		neighbors := getNeighbors(cur, n)
		for _, n := range neighbors {
			if total_risk[n.x][n.y] == -1 {
				//unvisited
				total_risk[n.x][n.y] = total_risk[cur.x][cur.y] + riskmap[n.x][n.y]
				queue = append(queue, n)
			} else if total_risk[cur.x][cur.y]+riskmap[n.x][n.y] < total_risk[n.x][n.y] {
				total_risk[n.x][n.y] = total_risk[cur.x][cur.y] + riskmap[n.x][n.y]
			}
		}
		if len(queue) > n*n {
			panic("Fuck this")
		}
	}

	return total_risk[n-1][n-1]
}

func main() {
	if len(os.Args) < 2 {
		panic("Not enough command line arguments\n")
	}
	file_path := os.Args[1]
	lines := get_input(file_path)

	risk := make([][]int, len(lines))
	for i := range lines {
		risk[i] = convert_to_int(strings.Split(lines[i], ""))
	}
	n := len(risk)

	fmt.Printf("Lowest risk part1 %d\n", solve(risk))

	tile_factor := 5
	tiled_risk := make([][]int, tile_factor*n)
	for oi := 0; oi < tile_factor; oi++ {
		for oj := 0; oj < tile_factor; oj++ {
			for i := range risk {
				if oj == 0 {
					tiled_risk[i+oi*n] = make([]int, n*tile_factor)
				}
				for j := range risk[i] {
					risk := risk[i][j]
					adjusted_risk := (risk + oi + oj) % 10
					if adjusted_risk < risk+oi+oj {
						adjusted_risk += 1
					}
					if adjusted_risk > 9 || adjusted_risk < 1 {
						panic("Fuck this")
					}
					tiled_risk[i+oi*n][j+oj*n] = adjusted_risk
				}
			}
		}
	}
	fmt.Printf("Lowest risk part2 %d\n", solve(tiled_risk))
}
