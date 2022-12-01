package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

type Dot struct {
	x int
	y int
}

type Instruction struct {
	axis string
	pos  int
}

func get_input(file_path string) []string {
	input, err := ioutil.ReadFile(file_path)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(strings.TrimSpace(string(input)), "\n")

	return lines
}

func getField(d Dot, field string) int {
	r := reflect.ValueOf(d)
	f := reflect.Indirect(r).FieldByName(field)
	return int(f.Int())
}

func main() {
	if len(os.Args) < 2 {
		panic("Not enough command line arguments\n")
	}
	file_path := os.Args[1]
	lines := get_input(file_path)

	var paper = make(map[Dot]int)
	var insts []Instruction
	for _, line := range lines {
		if line == "" {
			continue
		}
		parts := strings.Split(line, ",")
		if len(parts) < 2 {
			r := regexp.MustCompile(`fold along (?P<axis>[xy])=(?P<pos>\d+)`)
			res := r.FindStringSubmatch(line)
			axis := string(res[1])
			pos, _ := strconv.Atoi(string(res[2]))

			insts = append(insts, Instruction{axis, pos})
		} else {
			x, _ := strconv.Atoi(parts[0])
			y, _ := strconv.Atoi(parts[1])
			paper[Dot{x, y}] = 1
		}
	}

	for i, inst := range insts {
		for d := range paper {
			x := d.x
			y := d.y
			if inst.axis == "x" {
				if x > inst.pos {
					new_x := inst.pos - (x - inst.pos)
					delete(paper, Dot{x, y})
					paper[Dot{new_x, y}] = 1
				}
			} else {
				if y > inst.pos {
					new_y := inst.pos - (y - inst.pos)
					delete(paper, Dot{x, y})
					paper[Dot{x, new_y}] = 1
				}
			}
		}

		num_dots := 0
		for _, i = range paper {
			if i == 1 {
				num_dots++
			}
		}
		fmt.Printf("The number of visible dots is %d\n", num_dots)
	}

	max_x := 0
	max_y := 0
	for d := range paper {
		if d.x > max_x {
			max_x = d.x
		}
		if d.y > max_y {
			max_y = d.y
		}
	}
	var output = make([][]rune, max_x+1)
	for i := range output {
		output[i] = make([]rune, max_y+1)
	}
	for d := range paper {
		output[d.x][d.y] = 'X'
	}

	for i := range output {
		for j := range output[i] {
			if output[i][j] == 0 {
				fmt.Print(".")
			} else {
				fmt.Print("X")
			}
		}
		fmt.Println("")
	}
}
