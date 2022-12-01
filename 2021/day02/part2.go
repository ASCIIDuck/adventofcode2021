package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type submarine struct {
	aim      int
	distance int
	depth    int
}

func (s *submarine) forward(n int) {
	s.distance += n
	s.depth += n * s.aim
}

func (s *submarine) down(n int) {
	s.aim += n
}

func (s *submarine) up(n int) {
	s.aim -= n
}

func main() {
	if len(os.Args) < 2 {
		panic("No enough command line arguments\n")
	}
	file_path := os.Args[1]
	instructions, err := os.ReadFile(file_path)
	if err != nil {
		panic(err)
	}
	sub := submarine{0, 0, 0}
	for _, line := range strings.Split(string(instructions), "\n") {
		components := strings.Split(line, " ")
		if len(components) < 2 {
			fmt.Printf("Line '%s' contained too few components\n", string(line))
			continue
		}
		direction := components[0]
		magnitude, err := strconv.Atoi(components[1])
		if err != nil {
			fmt.Printf("Failed to convert magnitude. '%s' is not numeric\n", components[1])

		}

		if direction == "forward" {
			sub.forward(magnitude)
		}
		if direction == "down" {
			sub.down(magnitude)
		}
		if direction == "up" {
			sub.up(magnitude)
		}
	}

	fmt.Printf("Sub ended at %d, %d\n", sub.distance, sub.depth)
	fmt.Printf("Final answer: %d\n", sub.distance*sub.depth)
}
