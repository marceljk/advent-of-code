package main

import (
	_ "embed"
	"fmt"
	"log"
	"strconv"
	"strings"
)

//go:embed input.txt
var inputFile string

type rawRule struct {
	before int
	after  int
}

type data [][]int

type inputModel struct {
	rules rule
	data  data
}

type rule map[int][]int

func parseRules(rawRules string) (rule, error) {
	ruleEntries := strings.Split(rawRules, "\n")
	var rules []rawRule
	for _, rawEntry := range ruleEntries {
		entry := strings.Split(rawEntry, "|")
		if len(entry) != 2 {
			return nil, fmt.Errorf("input does not have the expected format in rules: %q", rawEntry)
		}

		rawBefore := entry[0]
		rawAfter := entry[1]

		before, err := strconv.Atoi(rawBefore)
		if err != nil {
			return nil, fmt.Errorf("could not parse rule %q, value %q: %w", rawEntry, rawBefore, err)
		}
		after, err := strconv.Atoi(rawAfter)
		if err != nil {
			return nil, fmt.Errorf("could not parse rule %q, value %q: %w", rawEntry, rawAfter, err)
		}

		rules = append(rules, rawRule{
			before: before,
			after:  after,
		})
	}
	result := optimizeRule(rules)
	return result, nil
}

func parseData(rawData string) (data, error) {
	dataEntries := strings.Split(rawData, "\n")
	result := make(data, len(dataEntries))
	for idx, rawData := range dataEntries {
		entry := strings.Split(rawData, ",")
		data := []int{}
		for _, val := range entry {
			parsedVal, err := strconv.Atoi(val)
			if err != nil {
				return nil, fmt.Errorf("could not parse val %q in dataRow %v: %w", val, rawData, err)
			}
			data = append(data, parsedVal)
		}
		result[idx] = data
	}
	return result, nil
}

func parseInput(input string) (*inputModel, error) {
	dataSplitted := strings.Split(input, "\n\n")
	if len(dataSplitted) != 2 {
		return nil, fmt.Errorf("input does not have the expected format")
	}
	rules, err := parseRules(dataSplitted[0])
	if err != nil {
		return nil, fmt.Errorf("could not parse rules: %w", err)
	}

	data, err := parseData(dataSplitted[1])
	if err != nil {
		return nil, fmt.Errorf("could not parse data: %w", err)
	}

	return &inputModel{
		rules,
		data,
	}, nil
}

func optimizeRule(inputRule []rawRule) rule {
	result := make(rule, 0)
	for _, rule := range inputRule {
		if result[rule.before] == nil {
			result[rule.before] = make([]int, 0)
		}
		result[rule.before] = append(result[rule.before], rule.after)
	}
	return result
}

func isStatementValid(stmt []int, rules rule) bool {
	for idx, val := range stmt {
		if idx == 0 {
			continue
		}
		if !isRuleValid(idx, stmt[:idx], rules[val]) {
			return false
		}
	}
	return true
}

func removeValidData(input inputModel) data {
	resultData := make(data, 0)
	data := input.data
	for _, stmt := range data {
		if isStatementValid(stmt, input.rules) {
			continue
		}
		resultData = append(resultData, stmt)
	}
	return resultData
}

func removeInvalidData(input inputModel) data {
	var result data
	data := input.data
	for _, stmt := range data {
		if isStatementValid(stmt, input.rules) {
			result = append(result, stmt)
		}
	}
	return result
}

func sumMiddleValue(input data) int {
	var result int
	for idx, stmt := range input {
		middleIdx := ((len(stmt) + 1) / 2) - 1
		if middleIdx < 0 {
			log.Fatalf("middleIdx: %v, stmt: %v, idx: %v", middleIdx, stmt, idx)
		}
		result += stmt[middleIdx]
	}
	return result
}

func isRuleValid(idx int, entry []int, rule []int) bool {
	check := entry[:idx]
	for _, afterRule := range rule {
		for _, beforeValue := range check {
			if afterRule == beforeValue {
				return false
			}
		}
	}
	return true
}

func fixInvalidData(rules rule, input data) data {
	result := make(data, len(input))
	// var selectedRules rule
	for entryIdx, entry := range input {
		// resultEntry := make([]int, len(entry))
		// copy(entry, resultEntry)

		// for _, val := range entry {
		// 	if selectedRules[val] == nil {
		// 		selectedRules[val] = make([]int, 0)
		// 	}
		// 	selectedRules[val] = rules[val]
		// }
		idx := 0
		hasChanged := false
		for {
			if idx == len(entry) && !hasChanged {
				break
			}
			idx = idx % len(entry)
			if idx == 0 {
				hasChanged = false
			}
			val := entry[idx]

			if !isRuleValid(idx, entry, rules[val]) {
				hasChanged = true
				a := entry[idx-1]
				b := entry[idx]

				entry[idx] = a
				entry[idx-1] = b
				idx = 0

				continue
			}

			idx++
		}

		result[entryIdx] = entry
	}
	return result
}

func main() {
	// Begin Part 1
	input, err := parseInput(inputFile)
	if err != nil {
		log.Fatalln(err)
	}
	filteredData := removeInvalidData(*input)
	result := sumMiddleValue(filteredData)
	fmt.Println(result)
	// End Part 1

	// Begin Part 2
	filteredData = removeValidData(*input)
	orderedData := fixInvalidData(input.rules, filteredData)
	result = sumMiddleValue(orderedData)
	fmt.Println(result)
	// End Part 2
}
