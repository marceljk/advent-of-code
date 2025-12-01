package main

import (
	_ "embed"
	"fmt"
	"log"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func importToArray(input string) ([][]int, error) {
	lines := strings.Split(input, "\n")
	arr := make([][]int, len(lines))

	for idx, val := range lines {
		rawInput := strings.Split(val, " ")
		arr[idx] = make([]int, len(rawInput))
		for j, val := range rawInput {
			num, err := strconv.Atoi(val)
			if err != nil {
				return nil, fmt.Errorf("could not parse %q to int: %w", val, err)
			}
			arr[idx][j] = num
		}
	}
	return arr, nil
}

func isSafePartOne(arr []int) bool {
	if len(arr) == 0 {
		return true
	}

	prevNum := arr[0]
	nextNum := arr[1]

	if prevNum == nextNum {
		return false
	}

	increasing := prevNum < nextNum

	for idx, val := range arr {
		if idx == 0 {
			continue
		}

		diff := val - prevNum
		prevNum = val

		if increasing && diff <= 3 && diff > 0 {
			continue
		}
		if !increasing && diff >= -3 && diff < 0 {
			continue
		}

		return false
	}

	return true
}

func isSafePartTwo(arr []int) bool {
	if len(arr) == 0 {
		return true
	}

	prevNum := arr[0]
	nextNum := arr[1]

	increasing := prevNum < nextNum

	for idx, val := range arr {
		if idx == 0 {
			continue
		}

		diff := val - prevNum
		prevNum = val

		if increasing && diff <= 3 && diff > 0 {
			continue
		}
		if !increasing && diff >= -3 && diff < 0 {
			continue
		}

		return isSafeWithOneRemovedLevel(arr) 
	}

	return true
}

func remove(slice []int, s int) []int {
	tmp := make([]int, len(slice))
	copy(tmp, slice)
	return append(tmp[:s], tmp[s+1:]...)
}

func isSafeWithOneRemovedLevel(arr []int) bool {
	for idx := range arr {
		withoutVal := remove(arr, idx)
		fmt.Printf("orig: %+v, subSet: %+v\n", arr, withoutVal)
		if isSafePartOne(withoutVal) {
			fmt.Printf("found %+v\n", withoutVal)
			return true
		}
	}
	return false
}

func main() {
	arr, err := importToArray(input)
	if err != nil {
		log.Fatal(err)
	}

	// Begin Part 1
	sum := 0
	for _, subArr := range arr {
		if isSafePartOne(subArr) {
			sum++
		}
	}
	fmt.Println(sum)
	// End Part 1

	// Begin Part 2
	sum = 0
	for _, subArr := range arr {
		if isSafePartTwo(subArr) {
			sum++
		}
	}
	fmt.Println(sum)
	// End Part 2
}
