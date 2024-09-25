package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func GetPath() (string, error) {
	args := os.Args[1:]
	if len(args) < 1 {
		return "", errors.New(
			"usage: go run aisystems-lab2 <location/name.pl>")
	}
	path := args[0]
	return path, nil
}

func Greet() {
	hello := "AI Systems Lab 2: Knowledge-based decision-making system"
	fmt.Println(strings.Repeat("-", len(hello)))
	fmt.Println(hello)
	fmt.Println(strings.Repeat("-", len(hello)))
}

func Goodbye() {
	fmt.Println("Closing application. Goodbye!")
}

// ShowHelpOptions TODO
func ShowHelpOptions() {

}

// ShowQueryingHelp TODO
func ShowQueryingHelp() {

}

// ShowTokensHelp TODO
func ShowTokensHelp() {

}

var scanner = bufio.NewScanner(os.Stdin)

func ReadInput() string {
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}

func ReadKnowledgeBaseFromFile(path string) string {
	f, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	knowledgeBase := string(f)
	return knowledgeBase
}
