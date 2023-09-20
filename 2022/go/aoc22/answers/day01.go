package answers

import (
	"aoc22/registry"
	"github.com/golang/glog"
	"slices"
	"strconv"
	"strings"
)

func processInput(input string) []int64 {
	elves := strings.Split(input, "\n\n")
	total_calories := make([]int64, 0, len(elves))
	for _, elf := range elves {
		var total int64 = 0
		for _, line := range strings.Split(elf, "\n") {
			i, err := strconv.ParseInt(line, 10, 64)
			if err != nil {
				glog.Fatal("Failed to parse %s into an int: %s", line, err)
			}
			total = total + i
		}
		total_calories = append(total_calories, total)
	}
	slices.Sort(total_calories)
	return total_calories
}

func Day01PartA(input string) string {
	total_calories := processInput(input)
	return strconv.FormatInt(total_calories[len(total_calories)-1], 10)
}

func Day01PartB(input string) string {
	total_calories := processInput(input)
	length := len(total_calories)
	var top3 int64 = 0
	for i := 1; i <= 3; i++ {
		top3 = top3 + total_calories[length-i]
	}
	return strconv.FormatInt(top3, 10)
}

func init() {
	registry.RegisterAnswer("day01", Day01PartA, Day01PartB)
}
