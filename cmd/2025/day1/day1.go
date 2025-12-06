package day1

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/marceljk/advent-of-code/pkg/utils/globalflags"
	"github.com/marceljk/advent-of-code/pkg/utils/runner"
	"github.com/spf13/cobra"
)

type rotation struct {
	op    direction
	moves int
}

type direction string

const (
	left  direction = "L"
	right direction = "R"

	startValue = 50
)

func NewCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "1",
		Short: "Solution for 2025 day 1",
		Long:  "Solution for 2025 day 1.",
		Args:  cobra.NoArgs,
		RunE:  runner.GenericRunE(task1, task2),
	}
	globalflags.ConfigureFlags(cmd)
	return cmd
}

func task1(input string) error {
	rotations, err := parseRotation(input)
	if err != nil {
		return fmt.Errorf("failed parsing rotation: %w", err)
	}
	zeroValues := calcPassword1(startValue, rotations)
	fmt.Printf("Solution 1: %d\n", zeroValues)
	return nil
}

func task2(input string) error {
	rotations, err := parseRotation(input)
	if err != nil {
		return fmt.Errorf("failed parsing rotation: %w", err)
	}
	zeroValues := calcPassword2(startValue, rotations)
	fmt.Printf("Solution 2: %d\n", zeroValues)
	return nil
}

func parseRotation(input string) ([]rotation, error) {
	rotations := []rotation{}
	lines := strings.Split(input, "\n")
	for idx, line := range lines {
		var op direction
		if line[0] == byte('L') {
			op = left
		} else if line[0] == byte('R') {
			op = right
		} else {
			return nil, fmt.Errorf("input is invalid. line %v, input: %+v", idx, line[0])
		}

		value, err := strconv.Atoi(line[1:])
		if err != nil {
			return nil, fmt.Errorf("input is invalid. line %v, input: %q", idx, line[1:])
		}
		rotations = append(rotations, rotation{
			op:    op,
			moves: value,
		})
	}
	return rotations, nil
}

func calcPassword1(startValue int, rotations []rotation) uint {
	var zeroValues uint
	for _, value := range rotations {
		if value.op == left {
			startValue = startValue - value.moves
		} else if value.op == right {
			startValue = startValue + value.moves
		}
		for startValue > 99 || startValue < 0 {
			if startValue > 99 {
				startValue -= 100
			}
			if startValue < 0 {
				startValue += 100
			}
		}
		if startValue == 0 {
			zeroValues++
		}
	}
	return zeroValues
}

func calcPassword2(startValue int, rotations []rotation) uint {
	var zeroValues uint
	previous, current := startValue, startValue
	for _, value := range rotations {
		for value.moves > 99 {
			value.moves = value.moves - 100
			zeroValues++
		}
		if value.moves == 0 {
			continue
		}

		if value.op == left {
			previous = current
			current = current - value.moves
		} else if value.op == right {
			previous = current
			current = current + value.moves
		}

		if current > 99 || current < 0 {
			if current > 99 {
				current -= 100
			}
			if current < 0 {
				current += 100
			}
			if previous != 0 {
				zeroValues++
			}
		} else if current == 0 {
			zeroValues++
		}

	}
	return zeroValues
}
