package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

type Data [][]rune

func parseInput(inputString string) *Data {
	lines := strings.Split(inputString, "\n")
	var result Data = make([][]rune, len(lines))
	for idx := range lines {
		result[idx] = []rune(lines[idx])
	}
	return &result
}

func checkHorizontal(i int, j int, data *Data) (sum int) {
	size := len((*data)[i])
	if j+3 < size && (*data)[i][j] == 'X' && (*data)[i][j+1] == 'M' && (*data)[i][j+2] == 'A' && (*data)[i][j+3] == 'S' {
		sum++
	}
	if j-3 >= 0 && (*data)[i][j] == 'X' && (*data)[i][j-1] == 'M' && (*data)[i][j-2] == 'A' && (*data)[i][j-3] == 'S' {
		sum++
	}
	return
}

func checkVertical(i int, j int, data *Data) (sum int) {
	size := len(*data)
	if i+3 < size && (*data)[i][j] == 'X' && (*data)[i+1][j] == 'M' && (*data)[i+2][j] == 'A' && (*data)[i+3][j] == 'S' {
		sum++
	}
	if i-3 >= 0 && (*data)[i][j] == 'X' && (*data)[i-1][j] == 'M' && (*data)[i-2][j] == 'A' && (*data)[i-3][j] == 'S' {
		sum++
	}
	return
}

func checkDiagonal(i int, j int, data *Data) (sum int) {
	iSize := len(*data)
	jSize := len((*data)[i])
	if i+3 < iSize && j+3 < jSize && (*data)[i][j] == 'X' && (*data)[i+1][j+1] == 'M' && (*data)[i+2][j+2] == 'A' && (*data)[i+3][j+3] == 'S' {
		sum++
	}
	if i-3 >= 0 && j+3 < jSize && (*data)[i][j] == 'X' && (*data)[i-1][j+1] == 'M' && (*data)[i-2][j+2] == 'A' && (*data)[i-3][j+3] == 'S' {
		sum++
	}
	if i+3 < iSize && j-3 >= 0 && (*data)[i][j] == 'X' && (*data)[i+1][j-1] == 'M' && (*data)[i+2][j-2] == 'A' && (*data)[i+3][j-3] == 'S' {
		sum++
	}
	if i-3 >= 0 && j-3 >= 0 && (*data)[i][j] == 'X' && (*data)[i-1][j-1] == 'M' && (*data)[i-2][j-2] == 'A' && (*data)[i-3][j-3] == 'S' {
		sum++
	}
	return
}

func getAllXmasEntries(data *Data) (sum int) {
	for i, iVal := range *data {
		for j, jVal := range iVal {
			if jVal != 'X' {
				continue
			}
			sum = sum + checkHorizontal(i, j, data)
			sum = sum + checkVertical(i, j, data)
			sum = sum + checkDiagonal(i, j, data)
		}
	}
	return sum
}

func getAllXmasEntriesPart2(data Data) (sum int) {
	iSize := len(data)
	for i, iVal := range data {
		jSize := len(iVal)
		for j, jVal := range iVal {
			if jVal != 'A' {
				continue
			}
			if i-1 < 0 || i+1 >= iSize {
				continue
			}
			if j-1 < 0 || j+1 >= jSize {
				continue
			}

			aboveLeft := data[i-1][j-1]
			belowLeft := data[i+1][j-1]
			aboveRight := data[i-1][j+1]
			belowRight := data[i+1][j+1]

			if (aboveLeft == 'M' && belowRight == 'S') || (aboveLeft == 'S' && belowRight == 'M') {
				if (aboveRight == 'M' && belowLeft == 'S') || (aboveRight == 'S' && belowLeft == 'M') {
					sum++
				}
			}
		}
	}
	return sum
}

func main() {
	//Begin part 1
	data := parseInput(input)
	sum := getAllXmasEntries(data)
	fmt.Println(sum)
	//End part 1

	//Begin part 2
	data = parseInput(input)
	sum = getAllXmasEntriesPart2(*data)
	fmt.Println(sum)
	//End part 2
}
