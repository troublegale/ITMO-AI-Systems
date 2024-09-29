package main

import (
	"github.com/ichiban/prolog"
	"slices"
	"strconv"
	"strings"
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
	"primary_stat":           {"str", "dex", "cns", "int", "wsd", "chr"},
	"primary_fighting_style": {"melee", "ranged", "caster", "support", "swordlemage"},
	"difficulty":             {}}

var rules = []string{"beginner_friendly", "challenging", "militant", "magic", "body", "soul"}

const incorrectQuery = "Incorrect query. Please use the specified querying rules and parameters."
const noAnswer = "No answer. There are no classes that fit your query."

func HandleQuery(pro *prolog.Interpreter, inputSplit []string) string {
	query, ok := formQuery(inputSplit)
	if !ok {
		return incorrectQuery
	}
	answer := executeQuery(pro, query)
	return answer
}

func formQuery(inputSplit []string) (string, bool) {
	positiveParameters := make(map[string][]string)
	negativeParameters := make(map[string][]string)
	var positiveRules []string
	var negativeRules []string
	difficulties := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := 0; i < len(inputSplit); {
		negative := false
		word := inputSplit[i]
		if strings.HasPrefix(word, "!") {
			negative = true
			word = word[1:]
		}
		if strings.HasPrefix(word, "-") {
			rule := word[1:]
			if !slices.Contains(rules, rule) {
				return "", false
			}
			if negative {
				negativeRules = append(negativeRules, rule)
			} else {
				positiveRules = append(positiveRules, rule)
			}
			i++
			continue
		}
		if !strings.HasSuffix(word, ":") {
			return "", false
		}
		parameter := word[:len(word)-1]
		if _, ok := parameters[parameter]; !ok {
			return "", false
		}
		if i == len(inputSplit)-1 {
			return "", false
		}
		values := strings.Split(inputSplit[i+1], ",")
		i += 2
		if parameter == "difficulty" {
			for _, value := range values {
				if len(value) == 1 {
					return "", false
				}
				sign := string(value[0])
				number := value[1:]
				n, err := strconv.Atoi(number)
				if err != nil || n > 10 || n < 1 {
					return "", false
				}
				var matches []int
				switch sign {
				case "<":
					matches = ModifyWithConstraint(difficulties, n, LessThan)
				case ">":
					matches = ModifyWithConstraint(difficulties, n, GreaterThan)
				case "=":
					matches = ModifyWithConstraint(difficulties, n, Equals)
				default:
					return "", false
				}
				if negative {
					difficulties = removeMatches(difficulties, matches)
				} else {
					difficulties = matches
				}
			}
			continue
		}
		for _, value := range values {
			if !slices.Contains(parameters[parameter], value) {
				return "", false
			}
		}
		if negative {
			negativeParameters[parameter] = values
		} else {
			positiveParameters[parameter] = values
		}
	}
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
	return query + ".", true
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
