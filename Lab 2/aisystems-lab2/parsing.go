package main

import (
	"slices"
	"strconv"
	"strings"
)

var parameters = map[string][]string{
	"primary_stat":           {"str", "dex", "cns", "int", "wsd", "chr"},
	"primary_fighting_style": {"melee", "ranged", "caster", "support", "swordlemage"},
	"difficulty":             {}}

var rules = []string{"beginner_friendly", "challenging", "militant", "magic", "body", "soul"}

func FormQuery(inputSplit []string) (string, bool) {
	positiveParameters := make(map[string][]string)
	negativeParameters := make(map[string][]string)
	var positiveRules []string
	var negativeRules []string
	difficulties := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	if ok := parseInputIntoEntries(inputSplit, positiveParameters, negativeParameters,
		&positiveRules, &negativeRules, &difficulties); !ok {
		return "", false
	}
	query := formEntriesIntoQuery(positiveParameters, negativeParameters, positiveRules, negativeRules, difficulties)
	return query + ".", true
}

func parseInputIntoEntries(inputSplit []string, positiveParameters, negativeParameters map[string][]string,
	positiveRules, negativeRules *[]string, difficulties *[]int) bool {
	for i := 0; i < len(inputSplit); {
		negative := false
		if strings.HasPrefix(inputSplit[i], "!") {
			negative = true
			inputSplit[i] = inputSplit[i][1:]
		}
		if strings.HasPrefix(inputSplit[i], "-") {
			if ok := parseRule(inputSplit[i], negative, positiveRules, negativeRules); !ok {
				return false
			}
			i++
			continue
		}
		if !strings.HasSuffix(inputSplit[i], ":") || i == len(inputSplit)-1 {
			return false
		}
		param, value := inputSplit[i], inputSplit[i+1]
		if ok := parseParameter(param, value, negative, positiveParameters, negativeParameters, difficulties); !ok {
			return false
		}
		i += 2
	}
	return true
}

func formEntriesIntoQuery(positiveParameters, negativeParameters map[string][]string,
	positiveRules, negativeRules []string, difficulties []int) string {
	query := ""
	for _, rule := range positiveRules {
		query = appendRule(query, rule) + ","
	}
	for k, v := range positiveParameters {
		query = appendParameter(query, k, v) + ","
	}
	if len(negativeRules) > 0 || len(negativeParameters) > 0 {
		if query == "" {
			query += "class(Class), "
		}
		for _, rule := range negativeRules {
			query = appendRule(query+" (\\+ ", rule) + "),"
		}
		for k, v := range negativeParameters {
			query = appendParameter(query+" (\\+ ", k, v) + "),"
		}
	}
	query = appendDifficulties(query, difficulties)
	if strings.HasSuffix(query, ",") {
		query = query[:len(query)-1]
	}
	return query
}

func parseRule(word string, negative bool, positiveRules, negativeRules *[]string) bool {
	rule := word[1:]
	if !slices.Contains(rules, rule) {
		return false
	}
	if negative {
		*negativeRules = append(*negativeRules, rule)
	} else {
		*positiveRules = append(*positiveRules, rule)
	}
	return true
}

func parseParameter(param, values string, negative bool,
	positiveParameters, negativeParameters map[string][]string, difficulties *[]int) bool {
	parameter := param[:len(param)-1]
	if _, ok := parameters[parameter]; !ok {
		return false
	}
	valuesSplit := strings.Split(values, ",")
	if parameter == "difficulty" {
		return parseDifficulty(valuesSplit, negative, difficulties)
	}
	for _, value := range valuesSplit {
		if !slices.Contains(parameters[parameter], value) {
			return false
		}
	}
	if negative {
		negativeParameters[parameter] = valuesSplit
	} else {
		positiveParameters[parameter] = valuesSplit
	}
	return true
}

func parseDifficulty(values []string, negative bool, difficulties *[]int) bool {
	for _, value := range values {
		if len(value) == 1 {
			return false
		}
		sign := string(value[0])
		number := value[1:]
		n, err := strconv.Atoi(number)
		if err != nil || n > 10 || n < 1 {
			return false
		}
		var matches []int
		switch sign {
		case "<":
			matches = ModifyWithConstraint(*difficulties, n, LessThan)
		case ">":
			matches = ModifyWithConstraint(*difficulties, n, GreaterThan)
		case "=":
			matches = ModifyWithConstraint(*difficulties, n, Equals)
		default:
			return false
		}
		if negative {
			*difficulties = removeMatches(*difficulties, matches)
		} else {
			*difficulties = matches
		}
	}
	return true
}

func appendRule(query string, rule string) string {
	if query == "" {
		return rule + "(Class)"
	}
	return query + rule + "(Class)"
}

func appendParameter(query string, param string, values []string) string {
	fact := "(" + param + "(Class, '" + values[0] + "')"
	if len(values) > 1 {
		for _, v := range values[1:] {
			fact += ";" + param + "(Class, '" + v + "')"
		}
	}
	fact += ")"
	if query == "" {
		return fact
	}
	return query + fact
}

func appendDifficulties(query string, difficulties []int) string {
	if len(difficulties) == 10 {
		return query
	}
	diffs := "difficulty(Class, Diff),("
	if len(difficulties) == 0 {
		diffs += "Diff < 0)"
	} else {
		for i := range difficulties {
			diffs += "Diff=" + strconv.Itoa(difficulties[i])
			if i != len(difficulties)-1 {
				diffs += ";"
			}
		}
		diffs += ")"
	}
	if query == "" {
		return diffs
	}
	query += diffs
	return query
}
