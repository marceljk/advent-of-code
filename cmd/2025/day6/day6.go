package day6

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/marceljk/advent-of-code/pkg/utils/globalflags"
	"github.com/marceljk/advent-of-code/pkg/utils/runner"
	"github.com/spf13/cobra"
)

func NewCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "6",
		Short: "Solution for 2025 day 6",
		Long:  "Solution for 2025 day 6.",
		Args:  cobra.NoArgs,
		RunE:  runner.GenericRunE(task1, task2),
	}
	globalflags.ConfigureFlags(cmd)
	return cmd
}

func task1(input string) error {
	problems, err := parseInput1(input)
	if err != nil {
		return fmt.Errorf("parse input: %w", err)
	}
	var total int
	for _, problem := range problems {
		sum, err := problem.Calculate()
		if err != nil {
			return fmt.Errorf("can not calculate sum: %w", err)
		}
		total += sum
	}
	fmt.Printf("Solution 1: %d\n", total)
	return nil
}

func task2(input string) error {
	problems, err := parseInput2(input)
	if err != nil {
		return fmt.Errorf("parse input: %w", err)
	}
	var total int
	for _, problem := range problems {
		sum, err := problem.Calculate()
		if err != nil {
			return fmt.Errorf("can not calculate sum: %w", err)
		}
		total += sum
	}
	fmt.Printf("Solution 2: %d\n", total)
	return nil
}

func parseInput1(input string) ([]Problem, error) {
	var problems []Problem
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		// Trim space and replace it with seperator ","
		line = strings.Trim(line, " ")
		line = regexp.MustCompile(`\s+`).ReplaceAllLiteralString(line, ",")
		fields := strings.Split(line, ",")

		// Init problems slice
		if problems == nil {
			problems = make([]Problem, len(fields))
		}
		for idx, field := range fields {
			switch field {
			case string(AddOp):
				problems[idx].Operation = AddOp
				continue
			case string(MulOp):
				problems[idx].Operation = MulOp
				continue
			}
			digit, err := strconv.Atoi(field)
			if err != nil {
				return nil, err
			}
			problems[idx].Numbers = append(problems[idx].Numbers, digit)
		}
	}
	return problems, nil
}

func parseInput2(input string) ([]Problem, error) {
	var readFromTopToBottom [][]string
	problems := []Problem{}
	lines := strings.Split(input, "\n")
	/*
	 * flip the input to read each lines from top to bottom, instead from left to right
	 */
	for _, line := range lines {
		lineChars := strings.Split(line, "")
		for charIdx := range lineChars {
			if readFromTopToBottom == nil {
				readFromTopToBottom = make([][]string, len(lineChars))
			}
			if readFromTopToBottom[charIdx] == nil {
				readFromTopToBottom[charIdx] = []string{}
			}
			readFromTopToBottom[charIdx] = append(readFromTopToBottom[charIdx], lineChars[charIdx])
		}
	}

	for _, line := range readFromTopToBottom {
		line := strings.Join(line, "")
		line = regexp.MustCompile(`\s+`).ReplaceAllLiteralString(line, "")
		if strings.HasSuffix(line, string(AddOp)) {
			problems = append(problems, Problem{
				Numbers:   []int{},
				Operation: AddOp,
			})
			// remove suffix
			line = line[:len(line)-1]
		} else if strings.HasSuffix(line, string(MulOp)) {
			problems = append(problems, Problem{
				Numbers:   []int{},
				Operation: MulOp,
			})
			// remove suffix
			line = line[:len(line)-1]
		} else if line == "" {
			continue
		}
		number, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}
		currentProblem := &problems[len(problems)-1]
		currentProblem.Numbers = append(currentProblem.Numbers, number)
	}
	return problems, nil
}
