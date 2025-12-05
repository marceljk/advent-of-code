package day5

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/marceljk/advent-of-code/pkg/utils/globalflags"
	"github.com/spf13/cobra"
)

type FreshIngredient struct {
	Start int
	End   int
}

type AvailableIngredient = int

func NewCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "5",
		Short: "Solution for 2025 day 5",
		Long:  "Solution for 2025 day 5.",
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
	freshIds, availableIds, err := parseIngredientsList(model.FileContent)
	if err != nil {
		return fmt.Errorf("can not parse list: %w", err)
	}
	var amountFresh int
	for _, availId := range availableIds {
		for _, idRange := range freshIds {
			if idRange.Start <= availId && idRange.End >= availId {
				amountFresh++
				break
			}
		}
	}
	fmt.Printf("Solution 1: %d\n", amountFresh)
	return nil
}

func task2(model globalflags.Model) error {
	freshIds, _, err := parseIngredientsList(model.FileContent)
	if err != nil {
		return fmt.Errorf("can not parse list: %w", err)
	}
	sort.Slice(freshIds, func(i, j int) bool {
		return freshIds[i].Start < freshIds[j].Start
	})
	optimizedIdsList := optimizeList(freshIds)
	totalIds := calculateIdsSum(optimizedIdsList)
	fmt.Printf("Solution 2: %d\n", totalIds)
	return nil
}

func calculateIdsSum(allFreshIngred []FreshIngredient) int {
	totalSum := 0
	for _, ingredient := range allFreshIngred {
		// to get all ids within a range, end - start + 1 must be calculated. for example range 3-5, has 3, 4 and 5 (= 3 ids). (5-3)+1 = 3
		diff := (ingredient.End - ingredient.Start) + 1
		totalSum = totalSum + diff
	}
	return totalSum
}

func optimizeList(freshIds []FreshIngredient) []FreshIngredient {
	optimizedIdsList := []FreshIngredient{}
	for _, idRange := range freshIds {
		if len(optimizedIdsList) == 0 {
			optimizedIdsList = append(optimizedIdsList, idRange)
			continue
		}
		changed := false
		for optimizedIdx := range optimizedIdsList {
			if optimizedIdsList[optimizedIdx].End >= idRange.Start && optimizedIdsList[optimizedIdx].Start <= idRange.Start {
				changed = true
				if optimizedIdsList[optimizedIdx].End < idRange.End {
					optimizedIdsList[optimizedIdx].End = idRange.End
				}
			}
			if optimizedIdsList[optimizedIdx].Start <= idRange.End && optimizedIdsList[optimizedIdx].End >= idRange.End {
				changed = true
				if optimizedIdsList[optimizedIdx].Start > idRange.Start {
					optimizedIdsList[optimizedIdx].Start = idRange.Start
				}
			}
		}
		if !changed {
			optimizedIdsList = append(optimizedIdsList, idRange)
		}
	}
	return optimizedIdsList
}

func parseIngredientsList(input string) (fresh []FreshIngredient, available []AvailableIngredient, err error) {
	parsingFreshDone := false
	lines := strings.Split(input, "\n")
	// parse ranges
	for _, line := range lines {
		if line == "" {
			parsingFreshDone = true
			continue
		}
		if !parsingFreshDone {
			freshIng, err := parseFreshIngredient(line)
			if err != nil {
				return nil, nil, fmt.Errorf("can not parse fresh ingredient: %w", err)
			}
			fresh = append(fresh, *freshIng)
		} else {
			availId, err := strconv.Atoi(line)
			if err != nil {
				return nil, nil, fmt.Errorf("can not parse available ingredient: %w", err)
			}
			available = append(available, availId)
		}
	}
	return
}

func parseFreshIngredient(input string) (*FreshIngredient, error) {
	values := strings.Split(input, "-")
	if len(values) != 2 {
		return nil, fmt.Errorf("invalid input, can not parse line: %q", input)
	}
	firstValue, err := strconv.Atoi(values[0])
	if err != nil {
		return nil, fmt.Errorf("invalid input, can not parse value: %q: %w", values[0], err)
	}
	secondValue, err := strconv.Atoi(values[1])
	if err != nil {
		return nil, fmt.Errorf("invalid input, can not parse value: %q: %w", values[1], err)
	}
	if secondValue < firstValue {
		firstValue, secondValue = secondValue, firstValue
	}
	return &FreshIngredient{
		Start: firstValue,
		End:   secondValue,
	}, nil
}
