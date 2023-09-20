package answers

import (
	"aoc22/registry"
	"github.com/golang/glog"
	"strconv"
	"strings"
)

const (
	ROCK     = 1
	PAPER    = 2
	SCISSORS = 3
)

var winnerMap = map[int]int{
	ROCK:     PAPER,
	PAPER:    SCISSORS,
	SCISSORS: ROCK,
}

func relativeValue(in string) int {
	inRune := in[0]
	if inRune >= byte('X') {
		return int(inRune-byte('X')) + 1
	}
	return int(inRune-byte('A')) + 1
}

func playCmp(them string, me string) int {
	relativeThem := relativeValue(them)
	relativeMe := relativeValue(me)
	if relativeThem == relativeMe {
		return 0
	} else if winnerMap[relativeThem] == relativeMe {
		return 1
	} else {
		return -1
	}
}

func reverseLookup(in int) int {
	for key, value := range winnerMap {
		if value == in {
			return key
		}
	}
	return 0
}

func parseInput(input string, scoreRound func(string, string) int) int {
	score := 0
	for _, round := range strings.Split(input, "\n") {
		parts := strings.SplitN(round, " ", 2)
		them := parts[0]
		me := parts[1]
		roundScore := scoreRound(them, me)
		glog.Infof("Round '%s' scored at '%d'\n", round, roundScore)
		score = score + roundScore
	}
	return score
}

func Day02PartA(input string) string {
	partARoundScore := func(them string, me string) int {
		roundScore := 0
		switch playCmp(them, me) {
		case -1:
			roundScore = roundScore + 0
		case 0:
			roundScore = roundScore + 3
		case 1:
			roundScore = roundScore + 6
		}
		roundScore = roundScore + relativeValue(me)
		return roundScore
	}
	score := parseInput(input, partARoundScore)
	return strconv.Itoa(score)
}

func Day02PartB(input string) string {
	partBRoundScore := func(them string, outcome string) int {
		roundScore := 0
		switch playCmp("Y", outcome) {
		case -1:
			roundScore = roundScore + 0 + reverseLookup(relativeValue(them))
		case 0:
			roundScore = roundScore + 3 + relativeValue(them)
		case 1:
			roundScore = roundScore + 6 + winnerMap[relativeValue(them)]
		}
		return roundScore
	}
	score := parseInput(input, partBRoundScore)
	return strconv.Itoa(score)
}

func init() {
	registry.RegisterAnswer("day02", Day02PartA, Day02PartB)
}
