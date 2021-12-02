package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func lower_bound(i int) int {
	return i
}

func upper_bound(i int, l int) int {
	if i+3 < l {
		return i + 3
	}
	return l
}

func sum(nums []int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

func main() {
	data, err := os.ReadFile("inputs.txt")
	if err != nil {
		panic(err)
	}

	var nums []int
	for _, line := range strings.Split(string(data), "\n") {
		n, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		nums = append(nums, n)
	}

	previous := -1
	num_increased := 0
	for i := range nums {
		cur := sum(nums[lower_bound(i):upper_bound(i, len(nums))])
		if previous != -1 && previous < cur {
			num_increased += 1
		}
		previous = cur

	}
	fmt.Println(num_increased)

}
