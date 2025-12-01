package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type Mul struct {
	x int
	y int
}

func parseInput(inputString string) []Mul {
	var res []Mul
	data := strings.Split(inputString, "mul(")
	data = data[1:]
	for _, val := range data {
		values := strings.Split(val, ")")[:1]
		if strings.ContainsAny(values[0], " ") {
			continue
		}
		values = strings.Split(strings.Join(values, ""), ",")
		if len(values) != 2 {
			continue
		}

		xParsed, err := strconv.Atoi(values[0])
		if err != nil {
			continue
		}
		yParsed, err := strconv.Atoi(values[1])
		if err != nil {
			continue
		}

		res = append(res, Mul{xParsed, yParsed})
	}
	return res
}

func filterValidInstructions(inputString string) string {
	splitDont := strings.Split(inputString, "don't()")
	result := splitDont[0]
	splitDont = splitDont[1:]
	for _, val := range splitDont {
		values := strings.Split(val, "do()")
		if len(values) <= 1 {
			continue
		}

		result = result + strings.Join(values[1:], "")
	}
	return result
}

func main() {
	//Begin part 1
	data := parseInput(input)
	sum := 0
	for _, val := range data {
		sum = sum + (val.x * val.y)
	}
	fmt.Println(sum)
	//End part 1

	//Begin part 2
	sum = 0
	newInput := filterValidInstructions(input)
	data = parseInput(newInput)
	for _, val := range data {
		sum = sum + (val.x * val.y)
	}
	fmt.Println(sum)
	//End part 2
}
