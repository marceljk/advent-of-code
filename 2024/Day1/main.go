package main

import (
	_ "embed"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Arrays struct {
	firstArr  []int
	secondArr []int
}

//go:embed input.txt
var input string

func importToArray(input string) (*Arrays, error) {
	arr := &Arrays{}
	lines := strings.Split(input, "\n")
	for _, val := range lines {
		splittedLine := strings.Split(val, " ")
		firstVal, err := strconv.Atoi(splittedLine[0])
		if err != nil {
			return nil, fmt.Errorf("invalid value in input.txt: %q. %w", splittedLine[0], err)
		}
		arr.firstArr = append(arr.firstArr, firstVal)
		secondVal, err := strconv.Atoi(splittedLine[len(splittedLine)-1])
		if err != nil {
			return nil, fmt.Errorf("invalid value in input.txt: %q. %w", splittedLine[len(splittedLine)-1], err)
		}
		arr.secondArr = append(arr.secondArr, secondVal)
	}
	return arr, nil
}

func sortArray(input []int, size int) []int {
	if size == 1 {
		return input
	}

	for i := 0; i < size-1; i++ {
		if input[i] > input[i+1] {
			input[i], input[i+1] = input[i+1], input[i]
		}
	}

	sortArray(input, size-1)

	return input
}

func containsNumTimes(num int, arr []int) int {
	sum := 0
	for _, val := range arr {
		if val == num {
			sum++
		}
	}
	return sum
}

func main() {
	arr, err := importToArray(input)
	if err != nil {
		log.Fatal(err)
	}
	arr.firstArr = sortArray(arr.firstArr, len(arr.firstArr))
	arr.secondArr = sortArray(arr.secondArr, len(arr.secondArr))

	
	// Begin Part 1
	sum := 0
	for idx := range arr.firstArr {
		diff := arr.firstArr[idx] - arr.secondArr[idx]
		if diff < 0 {
			diff = diff * -1
		}
		sum = sum + diff
	}
	fmt.Println(sum)
	// End Part 1

	// Begin Part 2
	sum = 0
	for _, val := range arr.firstArr {
		contains := containsNumTimes(val, arr.secondArr)
		sum = sum + (val * contains)
	}
	fmt.Println(sum)
	// End Part 2
}
