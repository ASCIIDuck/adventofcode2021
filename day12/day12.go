package main

import (
	"fmt"
	"io/ioutil"
	"os"
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

func pop(queue [][]string) ([]string, [][]string) {
	head := queue[0]
	queue = queue[1:]
	return head, queue
}

func add_to_graph(graph map[string][]string, nodeA string, nodeB string) map[string][]string {
	if graph[nodeA] == nil {
		graph[nodeA] = []string{nodeB}
	} else if in(graph[nodeA], nodeB) == 0 {
		graph[nodeA] = append(graph[nodeA], nodeB)
	}

	//now do it the other way
	if graph[nodeB] == nil {
		graph[nodeB] = []string{nodeA}
	} else if in(graph[nodeB], nodeA) == 0 {
		graph[nodeB] = append(graph[nodeB], nodeA)
	}
	return graph
}

func in(list []string, maybe_member string) int {
	total := 0
	for _, member := range list {
		if member == maybe_member {
			total++
		}
	}
	return total
}

func step1_validator(path []string) bool {
	last := path[len(path)-1]
	if (strings.ToLower(last) == last && in(path, last) == 1) || strings.ToUpper(last) == last {
		return true
	}
	return false
}

func step2_validator(path []string) bool {
	lower_counts := make(map[string]int)
	for _, node := range path {
		if strings.ToUpper(node) == node {
			continue
		}
		lower_counts[node]++
	}

	distinct_lower := 0
	total_lower := 0
	for node, count := range lower_counts {
		if (node == "start" || node == "end") && count > 1 {
			return false
		}
		distinct_lower++
		total_lower += count
	}
	return distinct_lower+1 >= total_lower
}

func count_paths(graph map[string][]string, validator func([]string) bool) int {
	var queue [][]string
	queue = append(queue, []string{"start"})
	path_count := 0
	for len(queue) > 0 {
		var cur = make([]string, len(queue[0]))
		copy(cur, queue[0])
		queue = queue[1:]
		latest_node := cur[len(cur)-1]
		for _, adj := range graph[latest_node] {
			if adj == "end" {
				path_count++
			} else if validator(append(cur, adj)) {
				queue = append(queue, append(cur, adj))
			}
		}
	}
	return path_count
}

func main() {
	if len(os.Args) < 2 {
		panic("Not enough command line arguments\n")
	}
	file_path := os.Args[1]
	lines := get_input(file_path)

	graph := make(map[string][]string)
	for _, line := range lines {
		parts := strings.Split(line, "-")
		name := parts[0]
		adj := parts[1]
		graph = add_to_graph(graph, name, adj)
	}

	step1_path_count := count_paths(graph, step1_validator)
	step2_path_count := count_paths(graph, step2_validator)

	fmt.Printf("Number of paths found in step 1 %d\n", step1_path_count)
	fmt.Printf("Number of paths found in step 2 %d\n", step2_path_count)

}
