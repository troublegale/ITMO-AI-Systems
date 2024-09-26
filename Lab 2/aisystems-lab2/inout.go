package main

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/ichiban/prolog"
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
	fmt.Println()
	fmt.Println("Closing application. Goodbye!")
}

func ShowHelpOptions() {
	fmt.Println("To see query options use 'help query'")
	fmt.Println("To see parameter options use 'help param'")
	fmt.Println()
}

func ShowQueryingHelp() {
	fmt.Println("% Simple query")
	fmt.Println("Usage: 'parameter: value'")
	fmt.Println("Example: 'primary_stat: INT'")
	fmt.Println()
	fmt.Println("% Negative query")
	fmt.Println("Usage: '!parameter: value'")
	fmt.Println("Example: '!primary_stat: STR'")
	fmt.Println()
	fmt.Println("% Query with multiple options (OR)")
	fmt.Println("Usage: 'parameter: (value1 value2 ...)'")
	fmt.Println("Example: 'primary_stat: (INT WSD)'")
	fmt.Println()
	fmt.Println("% Query with multiple options (AND)")
	fmt.Println("Usage: 'parameter1: value parameter2: value'")
	fmt.Println("Example: 'primary_stat: DEX difficulty: 4'")
	fmt.Println()
	fmt.Println("% Query with a rule")
	fmt.Println("Usage: '-rule'")
	fmt.Println("Example: '-beginner_friendly'")
	fmt.Println()
	fmt.Println("% All combined together")
	fmt.Println("primary_stat: (DEX CHR) !-beginner_friendly !primary_fighting_style: caster difficulty: <9")
	fmt.Println()
}

func ShowParametersHelp() {
	fmt.Println("primary_stat: STR DEX CNS INT WSD CHR")
	fmt.Println("primary_fighting_style: melee ranged caster support swordlemage")
	fmt.Println("difficulty: [1-10] || [< > <= >=][1-10]")
	fmt.Println("rules: -beginner_friendly -challenging -militant -magic -body -soul")
	fmt.Println()
}

func HandleUserInput(pro *prolog.Interpreter) {
	for {
		fmt.Print("$ ")
		input := readInput()
		if input != "" {
			inputSplit := strings.Fields(input)
			if strings.EqualFold(inputSplit[0], "help") {
				handleHelp(inputSplit)
			} else {
				HandleQuery(pro, inputSplit)
			}
		}
	}
}

func handleHelp(inputSplit []string) {
	if len(inputSplit) == 1 {
		ShowHelpOptions()
		return
	}
	switch strings.ToLower(inputSplit[1]) {
	case "query":
		ShowQueryingHelp()
	case "param":
		ShowParametersHelp()
	default:
		ShowHelpOptions()
	}
}

var scanner = bufio.NewScanner(os.Stdin)

func readInput() string {
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
