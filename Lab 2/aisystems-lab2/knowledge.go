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

var primaryAttribute = []string{"INT", "DEX", "CNS", "INT", "WSD", "CHR"}

var primaryFightingStyle = []string{"melee", "ranged", "caster", "support", "swordlemage"}

var rules = []string{"beginner_friendly", "challenging", "militant", "magic", "body", "soul"}

func HandleQuery(pro *prolog.Interpreter, inputSplit []string) {

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
