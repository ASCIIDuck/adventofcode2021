package registry

import (
	"log"
)

type Answer struct {
	PartA func(string) string
	PartB func(string) string
}

var ANSWER_REGISTRY = make(map[string]Answer)

func RegisterAnswer(name string, partA func(string) string, partB func(string) string) {
	ANSWER_REGISTRY[name] = Answer{partA, partB}
}

func GetAnswerStruct(name string) Answer {
	if answer, ok := ANSWER_REGISTRY[name]; ok {
		return answer
	}
	log.Fatalf("Problem '%s' was not found.", name)
	return Answer{}
}

func ListAnswers() []string {
	keys := make([]string, 0, len(ANSWER_REGISTRY))
	for k := range ANSWER_REGISTRY {
		keys = append(keys, k)
	}
	return keys
}
