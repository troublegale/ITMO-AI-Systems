package main

import (
	"github.com/ichiban/prolog"
	"slices"
	"strconv"
	"strings"
	"unicode"
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

func Classes() []string {
	return []string{"bard", "barbarian", "fighter", "wizard", "druid", "cleric", "artificer", "warlock",
		"monk", "paladin", "rogue", "ranger", "sorcerer", "alchemist", "warlord", "jaeger", "stargazer",
		"bloodhunter", "runekeeper", "shaman"}
}

func HandleQuery(pro *prolog.Interpreter, inputSplit []string) string {
	query, ok := inputToQuery(inputSplit)
	if !ok {
		return incorrectQuery
	}
	if slices.Contains(Classes(), query) {
		return getClassInfo(pro, query)
	}
	answer := executeQuery(pro, query)
	return answer
}

func inputToQuery(inputSplit []string) (string, bool) {
	query, ok := FormQuery(inputSplit)
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

func getClassInfo(pro *prolog.Interpreter, class string) string {
	runes := []rune(class)
	runes[0] = unicode.ToUpper(runes[0])
	class = string(runes)
	query := "primary_stat('" + class + "', Stat), primary_fighting_style('" + class + "', Style), " +
		"difficulty('" + class + "', Diff)."
	sols, err := pro.Query(query)
	if err != nil {
		panic(err)
	}
	answer := "-----The " + class + " Class-----\n"
	sols.Next()
	var s struct {
		Stat  string
		Style string
		Diff  int
	}
	if err := sols.Scan(&s); err != nil {
		panic(err)
	}
	answer += "Primary stat: " + strings.ToUpper(s.Stat) + "\n"
	answer += "Primary fighting style: " + s.Style + "\n"
	answer += "Difficulty: " + strconv.Itoa(s.Diff)
	return answer
}
