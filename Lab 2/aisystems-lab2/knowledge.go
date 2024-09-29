package main

import (
	"github.com/ichiban/prolog"
)

func InitiateInterpreter() *prolog.Interpreter {
	pro := prolog.New(nil, nil)
	return pro
}

func InitiateKnowledgeBase(pro *prolog.Interpreter, path string) {
	knowledgeBase := ReadKnowledgeBaseFromFile(path)
	if err := pro.Exec(knowledgeBase); err != nil {
		panic(err)
	}
}

const incorrectQuery = "Incorrect query. Please use the specified querying rules and parameters."
const noAnswer = "No answer. There are no classes that fit your query."

func HandleQuery(pro *prolog.Interpreter, inputSplit []string) string {
	query, ok := inputToQuery(inputSplit)
	if !ok {
		return incorrectQuery
	}
	answer := executeQuery(pro, query)
	return answer
}

func inputToQuery(inputSplit []string) (string, bool) {
	positiveParameters := make(map[string][]string)
	negativeParameters := make(map[string][]string)
	var positiveRules []string
	var negativeRules []string
	difficulties := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	query, ok := FormQuery(inputSplit, positiveParameters, negativeParameters,
		&positiveRules, &negativeRules, &difficulties)
	return query, ok
}

func executeQuery(pro *prolog.Interpreter, query string) string {
	sols, err := pro.Query(query)
	if err != nil {
		panic(err)
	}
	answer := ""
	for sols.Next() {
		var s struct {
			Class string
		}
		if err := sols.Scan(&s); err != nil {
			panic(err)
		}
		answer += s.Class + " "
	}
	if answer == "" {
		answer = noAnswer
	}
	return "Matching classes: " + answer
}
