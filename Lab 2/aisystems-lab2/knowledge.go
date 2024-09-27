package main

import (
	"fmt"
	"github.com/ichiban/prolog"
	"slices"
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

var parameters = map[string][]string{
	"primary_stat":           {"STR", "DEX", "CNS", "INT", "WSD", "CHR"},
	"primary_fighting_style": {"melee", "ranged", "caster", "support", "swordlemage"},
	"difficulty":             {}}

var rules = []string{"beginner_friendly", "challenging", "militant", "magic", "body", "soul"}

func HandleQuery(pro *prolog.Interpreter, inputSplit []string) string {
	answer := ""
	query, ok := formFact(inputSplit[0][:len(inputSplit[0])-1], inputSplit[1])
	if !ok {
		return "Incorrect query."
	}
	answer = executeQuery(pro, query)
	return answer
}

func executeQuery(pro *prolog.Interpreter, query string) string {
	fmt.Println(query)
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
	return answer
}

func formRule(s string) (string, bool) {
	if slices.Contains(rules, s) {
		return s + "(Class).", true
	}
	return "", false
}

func formFact(param string, value string) (string, bool) {
	if _, ok := parameters[param]; ok {
		if param == "difficulty" {
			return handleDifficulty(value)
		} else if slices.Contains(parameters[param], value) {
			return param + "(Class, '" + value + "').", true
		}
	}
	return "", false
}

func handleDifficulty(value string) (string, bool) {
	return "", false
}

//q := `class(Class), \+ (beginner_friendly(Class); primary_fighting_style(Class, 'caster')), difficulty(Class, Diff), Diff < 8.`
//sols, err := pro.Query(q)
//if err != nil {
//panic(err)
//}
//
//for sols.Next() {
//var s struct {
//Class string
//}
//if err := sols.Scan(&s); err != nil {
//panic(err)
//}
//fmt.Printf("%s", s.Class)
//}
//
//fmt.Printf("\n")
//if err := sols.Close(); err != nil {
//panic(err)
//}
