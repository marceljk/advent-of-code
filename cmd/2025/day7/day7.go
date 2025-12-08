package day7

import (
	"fmt"
	"strings"

	"github.com/marceljk/advent-of-code/pkg/utils/globalflags"
	"github.com/marceljk/advent-of-code/pkg/utils/runner"
	"github.com/spf13/cobra"
)

type Field rune

const (
	StartField    Field = 'S'
	BeamField     Field = '|'
	SplitterField Field = '^'
	EmptyField    Field = '.'
)

type Diagram [][]Field

func (d Diagram) beam() (int, int) {
	var totalTimesSplitted int
	beamingRow := map[int]int{} // beamingRow stores how many paths(=value), can reach the beam on the index(=key)
	for rowIdx := range d {
		nextBeamingRow := map[int]int{}
		for colIdx := range d[rowIdx] {
			field := d[rowIdx][colIdx]
			if field == StartField {
				nextBeamingRow[colIdx] = 1
			}
			// If current field is a splitter field and beamingRow has a value > 0.
			if field == SplitterField && beamingRow[colIdx] > 0 {
				totalTimesSplitted++
				leftIdx := colIdx - 1
				if leftIdx >= 0 {
					nextBeamingRow[leftIdx] += beamingRow[colIdx]
				}
				rightIdx := colIdx + 1
				if rightIdx < len(d[rowIdx]) {
					nextBeamingRow[rightIdx] += beamingRow[colIdx]
				}
				delete(beamingRow, colIdx)
			}
		}
		for idx, nextBeaming := range nextBeamingRow {
			beamingRow[idx] += nextBeaming
		}
		fmt.Println(beamingRow)
	}
	possibleSolutions := 0
	for _, val := range beamingRow {
		possibleSolutions += val
	}
	return totalTimesSplitted, possibleSolutions
}

func NewCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "7",
		Short: "Solution for 2025 day 7",
		Long:  "Solution for 2025 day 7.",
		Args:  cobra.NoArgs,
		RunE:  runner.GenericRunE(task1, task2),
	}
	globalflags.ConfigureFlags(cmd)
	return cmd
}

func task1(input string) error {
	diagram, err := parseInput(input)
	if err != nil {
		return fmt.Errorf("can not parse input: %w", err)
	}
	totalTimesSplitted, _ := diagram.beam()
	fmt.Printf("Solution 1: %d\n", totalTimesSplitted)
	return nil
}

func task2(input string) error {
	diagram, err := parseInput(input)
	if err != nil {
		return fmt.Errorf("can not parse input: %w", err)
	}
	_, possiblePaths := diagram.beam()
	fmt.Printf("Solution 2: %d\n", possiblePaths)
	return nil
}

func parseInput(input string) (Diagram, error) {
	lines := strings.Split(input, "\n")
	diagram := Diagram{}
	for _, line := range lines {
		runeLine := []rune(line)
		fields := make([]Field, len(runeLine))
		for _, runeField := range runeLine {
			switch runeField {
			case rune(StartField):
				fields = append(fields, StartField)
			case rune(SplitterField):
				fields = append(fields, SplitterField)
			case rune(EmptyField):
				fields = append(fields, EmptyField)
			default:
				return nil, fmt.Errorf("invalid field %q", runeField)
			}
		}
		diagram = append(diagram, fields)
	}
	return diagram, nil
}
