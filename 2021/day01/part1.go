package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("inputs.txt")

	if err != nil {
		panic(err)
	}

	previous := -1
	num_increased := 0
	for _, line := range strings.Split(string(data), "\n") {
		num, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		if previous != -1 && previous < num {
			num_increased += 1
		}
		previous = num
	}
	fmt.Println(num_increased)

}
