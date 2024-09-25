package main

import "fmt"

func main() {
	path, err := GetPath()
	if err != nil {
		fmt.Println(err)
		return
	}
	Greet()
	pro := InitiateInterpreter()
	InitiateKnowledgeBase(pro, path)
}
