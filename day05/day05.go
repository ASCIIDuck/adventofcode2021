package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func get_input() []string {
	input, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(strings.TrimSpace(string(input)), "\n")

	return lines
}

type Point struct {
	x int
	y int
}

func (p *Point) parse_string(in string) {
	parts := strings.Split(in, ",")
	p.x, _ = strconv.Atoi(strings.TrimSpace(parts[0]))
	p.y, _ = strconv.Atoi(strings.TrimSpace(parts[1]))
}

type Line struct {
	start Point
	end   Point
}

func (l *Line) parse_string(definition string) {
	parts := strings.Split(definition, "->")
	first_point := strings.TrimSpace(parts[0])
	second_point := strings.TrimSpace(parts[1])

	var p1 Point
	p1.parse_string(first_point)
	var p2 Point
	p2.parse_string(second_point)

	if p1.x <= p2.x {
		l.start = p1
		l.end = p2
	} else {
		l.start = p2
		l.end = p1
	}
}

type Grid struct {
	grid [][]int
}

func (g *Grid) set_size(x int, y int) {
	g.grid = make([][]int, y)
	for i := range g.grid {
		g.grid[i] = make([]int, x)
	}
}

func (g *Grid) draw_line(line Line) int {
	var incr Point
	if line.start.x == line.end.x {
		incr.x = 0
		incr.y = (line.end.y - line.start.y) / int(math.Abs(float64(line.end.y-line.start.y)))
	} else {
		incr.x = 1
		incr.y = (line.end.y - line.start.y) / (line.end.x - line.start.x)
	}

	var cur Point
	cur.x = line.start.x
	cur.y = line.start.y

	intersection_count := 0
	for true {
		g.grid[cur.y][cur.x]++
		if g.grid[cur.y][cur.x] == 2 {
			intersection_count++
		}

		if cur.x == line.end.x && cur.y == line.end.y {
			break
		}

		cur.x += incr.x
		cur.y += incr.y
	}
	return intersection_count
}

func (g *Grid) draw_simple_line(line Line) int {
	if line.start.x != line.end.x && line.start.y != line.end.y {
		return 0
	} else {
		return g.draw_line(line)
	}
}

func main() {
	line_strs := get_input()

	var lines []Line
	var max Point
	for _, line_str := range line_strs {
		var line Line
		line.parse_string(line_str)
		lines = append(lines, line)
		if line.start.x > max.x {
			max.x = line.start.x
		}
		if line.end.x > max.x {
			max.x = line.end.x
		}
		if line.start.y > max.y {
			max.y = line.start.y
		}
		if line.end.y > max.y {
			max.y = line.end.y
		}
	}
	var grid Grid
	grid.set_size(max.x+1, max.y+1)

	simple_intersection_count := 0
	for _, line := range lines {
		simple_intersection_count += grid.draw_simple_line(line)
	}
	fmt.Printf("I found %d simple intersections\n", simple_intersection_count)

	grid.set_size(max.x+1, max.y+1)
	intersection_count := 0
	for _, line := range lines {
		intersection_count += grid.draw_line(line)
	}
	fmt.Printf("I found %d intersections\n", intersection_count)

}
