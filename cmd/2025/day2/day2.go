package day2

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/marceljk/advent-of-code/pkg/utils"
	"github.com/marceljk/advent-of-code/pkg/utils/globalflags"
	"github.com/spf13/cobra"
)

type IdRange struct {
	start int
	end   int
}

func NewCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "2",
		Short: "Solution for 2025 day 2",
		Long:  "Solution for 2025 day 2.",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			model, err := globalflags.ParseInput(cmd)
			if err != nil {
				return fmt.Errorf("failec parsing model: %w", err)
			}

			switch model.Task {
			case 2:
				return task2(*model)
			default:
				return task1(*model)
			}
		},
	}
	globalflags.ConfigureFlags(cmd)
	return cmd
}

func task1(model globalflags.Model) error {
	list, err := parseIdRanges(model.FileContent)
	if err != nil {
		return fmt.Errorf("can not parse input: %w", err)
	}
	invalidIds := findInvalidIds1(list)
	sum := sumIds(invalidIds)
	fmt.Printf("Solution 1: %d\n", sum)
	return nil
}

func task2(model globalflags.Model) error {
	list, err := parseIdRanges(model.FileContent)
	if err != nil {
		return fmt.Errorf("can not parse input: %w", err)
	}
	invalidIds := findInvalidIds2(list)
	sum := sumIds(invalidIds)
	fmt.Printf("Solution 2: %d\n", sum)
	return nil
}

func findInvalidIds1(list []IdRange) []int {
	result := []int{}
	for _, idRange := range list {
		start, end := idRange.start, idRange.end
		for currentNumber := start; currentNumber <= end; currentNumber++ {
			valueString := strconv.Itoa(currentNumber)
			lenValueString := len(valueString)
			// If the number has an uneven number of digits, it can be skipped because there can not be a sequence twice
			if lenValueString%2 != 0 {
				continue
			}

			firstHalf := valueString[:(lenValueString / 2)]
			secondHalf := valueString[(lenValueString / 2):]
			if firstHalf == secondHalf {
				result = append(result, currentNumber)
			}
		}
	}
	return result
}

func findInvalidIds2(list []IdRange) []int {
	result := []int{}
	for _, idRange := range list {
		start, end := idRange.start, idRange.end
		invalidIdOfRange := map[int]bool{}
		for currentNumber := start; currentNumber <= end; currentNumber++ {
			valueString := strconv.Itoa(currentNumber)
			lenIdString := len(valueString)

			// Split the number in sub sequences and check if they appear multiple times
			// Set upperLimitLength to half of length + 1. Avoid unnecessary checks of too big subsequences
			upperLimitLength := (lenIdString / 2) + 1
		subsequence:
			for i := 1; i < upperLimitLength; i++ {
				// If subsequence can not split the string evenly, it can be skipped
				if lenIdString%i != 0 {
					continue
				}
				// Convert num with string to rune slice
				idSlice := make([]rune, lenIdString)
				for idx, char := range valueString {
					idSlice[idx] = char
				}

				// Split number in chunks and compare them
				var firstChunk []rune
				for chunk := range slices.Chunk(idSlice[:], i) {
					if firstChunk == nil {
						firstChunk = chunk
						continue
					}
					if !utils.IsSliceEqual(firstChunk, chunk) {
						continue subsequence
					}
				}
				invalidIdOfRange[currentNumber] = true
			}
		}
		invalidIds := utils.GetMapKeys(invalidIdOfRange)
		result = append(result, invalidIds...)
	}
	return result
}

func sumIds(ids []int) (sum int) {
	for _, num := range ids {
		sum += num
	}
	return
}

func parseIdRanges(input string) ([]IdRange, error) {
	result := []IdRange{}

	const separator = ","
	splittedInput := strings.Split(input, separator)
	for _, idRange := range splittedInput {
		splittedRange := strings.Split(idRange, "-")
		if len(splittedRange) != 2 {
			return nil, fmt.Errorf("unexpected product id range. expected format %q, got: %q", "202-300", idRange)
		}
		first, err := strconv.Atoi(splittedRange[0])
		if err != nil {
			return nil, fmt.Errorf("can not parse first value of product id range. expected int, got: %v. err: %w", splittedRange[0], err)
		}
		second, err := strconv.Atoi(splittedRange[1])
		if err != nil {
			return nil, fmt.Errorf("can not parse second value of product id range. expected int, got: %v. err: %w", splittedRange[1], err)
		}
		result = append(result, IdRange{
			start: first,
			end:   second,
		})
	}
	return result, nil
}
