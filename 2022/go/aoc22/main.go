package main

import (
	_ "aoc22/answers" // Ignore the "unused" import, I must import this to get the answers to register
	"aoc22/aocutils"
	"aoc22/registry"
	"flag"
	"fmt"
	"github.com/golang/glog"
	"os"
)

var SESSION_ID = "53616c7465645f5f7d24e96769086bafac28042b55c77674553f2bfe0b77258e797fc8193b97cd19005c1686052d63d4abce595edaa25634226f2ee98ed06125"

const SESSION_ENV_VAR = "AOC_SESSION_ID"

var day int
var part string

func init() {
	flag.IntVar(&day, "day", 0, "Int between 1 and 31 indicating which day's answers should be run")
	flag.StringVar(&part, "part", "*", "One of three accepted values : A, B, and *")
	flag.Parse()
	if day == 0 || day > 31 {
		glog.Fatalf("You must provide a valid value for --day that is between 1 and 31, inclusive")
	}
}
func main() {
	fmt.Println("Registered answers: ", registry.ListAnswers())
	sessionID := os.Getenv(SESSION_ENV_VAR)
	if sessionID == "" {
		glog.Fatalf("No session ID found in %s. This environment variable is required.", SESSION_ENV_VAR)
	}

	input := aocutils.GetInput(sessionID, 2022, day)
	answers := registry.GetAnswerStruct(fmt.Sprintf("day%02d", day))

	if part == "A" || part == "*" {
		fmt.Println("Running Part A")
		ans := answers.PartA(input)
		fmt.Printf("Answer to Part A: %s\n", ans)
	}
	if part == "B" || part == "*" {
		fmt.Println("Running Part B")
		ans := answers.PartB(input)
		fmt.Printf("Answer to Part B: %s\n", ans)
	}
}
