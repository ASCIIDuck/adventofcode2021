package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Space struct {
	value  string
	marked bool
}

type Board struct {
	spaces     [5][5]Space
	lastCalled string
	alreadyWon bool
}

func (b *Board) init_board(inputs []string) {
	space_re := regexp.MustCompile(`\s+`)
	for i := range inputs {
		inputs[i] = space_re.ReplaceAllString(strings.TrimSpace(inputs[i]), " ")
		values := strings.Split(inputs[i], " ")
		for j := range values {
			b.spaces[i][j] = Space{value: values[j], marked: false}
		}
	}
}

func (b *Board) check_columns() bool {
	for j := 0; j < 5; j++ {
		numMarked := 0
		for i := 0; i < 5; i++ {
			if b.spaces[i][j].marked {
				numMarked++
			}
		}
		if numMarked == 5 {
			return true
		}
	}
	return false
}

func (b *Board) check_rows() bool {
	for i := 0; i < 5; i++ {
		numMarked := 0
		for j := 0; j < 5; j++ {
			if b.spaces[i][j].marked {
				numMarked++
			}
		}
		if numMarked == 5 {
			return true
		}
	}
	return false
}

func (b *Board) check_wins() bool {
	b.alreadyWon = b.check_columns() || b.check_rows()
	return b.alreadyWon
}

func (b *Board) mark_number(called string) {
	b.lastCalled = called
	for i := range b.spaces {
		for j := range b.spaces[i] {
			if b.spaces[i][j].value == called {
				b.spaces[i][j].marked = true
			}
		}
	}
}

func (b *Board) score() int {
	score := 0
	for _, row := range b.spaces {
		for _, space := range row {
			if !space.marked {
				numeric, err := strconv.Atoi(space.value)
				if err != nil {
					panic(err)
				}
				score += numeric
			}
		}
	}
	lastCalledNumeric, _ := strconv.Atoi(b.lastCalled)
	return lastCalledNumeric * score
}

func main() {
	if len(os.Args) < 2 {
		panic("No enough command line arguments\n")
	}
	file_path := os.Args[1]
	input_bytes, err := os.ReadFile(file_path)
	if err != nil {
		panic(err)
	}
	input := strings.Split(string(input_bytes), "\n")

	numbersCalled := strings.Split(input[0], ",")

	input = input[2:]
	var buffer []string
	var boards []Board

	for _, line := range input {
		if line != "" {
			buffer = append(buffer, line)
		} else {
			var b Board
			b.init_board(buffer)
			boards = append(boards, b)
			buffer = buffer[:0]
		}
	}
	if len(buffer) > 0 {
		var b Board
		b.init_board(buffer)
		boards = append(boards, b)
	}

	boardsToWin := len(boards)
	for _, called := range numbersCalled {
		for j := range boards {
			if boards[j].alreadyWon {
				continue
			}
			boards[j].mark_number(called)
			if boards[j].check_wins() {
				if boardsToWin == len(boards) {
					fmt.Printf("First winning board #%d scored: %d\n", j, boards[j].score())
				}
				boardsToWin--
				if boardsToWin == 0 {
					fmt.Printf("Last board to win #%d scored: %d \n", j, boards[j].score())
				}
			}
		}
	}
}
