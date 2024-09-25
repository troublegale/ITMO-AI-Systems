package main

import (
	"fmt"
)

func main() {

	path, err := GetPath()
	if err != nil {
		fmt.Println(err)
		return
	}

	Greet()
	defer Goodbye()

	fmt.Println("Initializing knowledge base...")
	pro := InitiateInterpreter()
	InitiateKnowledgeBase(pro, path)
	fmt.Println("Knowledge base initialized. You can start writing queries!")

}
