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
