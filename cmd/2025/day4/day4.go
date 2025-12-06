package day4

import (
	"fmt"
	"strings"

	"github.com/marceljk/advent-of-code/pkg/utils/globalflags"
	"github.com/marceljk/advent-of-code/pkg/utils/runner"
	"github.com/spf13/cobra"
)

func NewCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "4",
		Short: "Solution for 2025 day 4",
		Long:  "Solution for 2025 day 4.",
		Args:  cobra.NoArgs,
		RunE:  runner.GenericRunE(task1, task2),
	}
	globalflags.ConfigureFlags(cmd)
	return cmd
}

func task1(input string) error {
	diagram, err := parseDiagram(input)
	if err != nil {
		return fmt.Errorf("can not parse diagram: %w", err)
	}
	removableRolls, _, err := diagram.RemovableRolls()
	if err != nil {
		return err
	}
	fmt.Printf("Solution 1: %d\n", removableRolls)
	return nil
}

func task2(input string) error {
	diagram, err := parseDiagram(input)
	if err != nil {
		return fmt.Errorf("can not parse diagram: %w", err)
	}
	var totalRemovableRolls int
	removableRolls, newDiagram, err := diagram.RemovableRolls()
	if err != nil {
		return err
	}
	totalRemovableRolls += removableRolls
	for removableRolls > 0 {
		removableRolls, newDiagram, err = newDiagram.RemovableRolls()
		if err != nil {
			return err
		}
		totalRemovableRolls += removableRolls
	}
	fmt.Printf("Solution 2: %d\n", totalRemovableRolls)
	return nil
}

func parseDiagram(input string) (Diagram, error) {
	var diagram Diagram
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		var row Row
		chars := strings.Split(line, "")
		for _, char := range chars {
			switch char {
			case string(PaperRoll):
				row = append(row, PaperRoll)
			case string(Empty):
				row = append(row, Empty)
			default:
				return nil, fmt.Errorf("invalid input. expected %q or %q, got: %q", PaperRoll, Empty, char)
			}
		}
		diagram = append(diagram, row)
	}
	return diagram, nil
}
