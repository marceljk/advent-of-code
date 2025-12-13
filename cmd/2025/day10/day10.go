package day10

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/marceljk/advent-of-code/pkg/utils/globalflags"
	"github.com/marceljk/advent-of-code/pkg/utils/runner"
	"github.com/spf13/cobra"
)

// stores the light diagram as int. Input is interpreted as binary representation. E.g. [.##.] => 0110 (=binary) => 6 (=decimal)
type LightDiagram = int

// stores a button as int. Input is interpreted as binary representation. E.g. LightDiagram has 4 digits in binary (=0110)
// and button is (0,2,3) (=input) => 1011 (=binary) => 7 (=decimal)
type Button = int

type Manual struct {
	Buttons     []Button
	WantedState LightDiagram
	// IsState     LightDiagram
}

func NewCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "10",
		Short: "Solution for 2025 day 10",
		Long:  "Solution for 2025 day 10.",
		Args:  cobra.NoArgs,
		RunE:  runner.GenericRunE(task1, task2),
	}
	globalflags.ConfigureFlags(cmd)
	return cmd
}

func task1(input string) error {
	manuels, err := parseInput(input)
	if err != nil {
		return err
	}
	depths := findAllFewestButtonPresses(manuels)
	var result int
	for _, val := range depths {
		result += val
	}
	fmt.Printf("Solution 1: %d\n", result)
	return nil
}

func task2(input string) error {
	return nil
}

func findAllFewestButtonPresses(manuels []Manual) []int {
	depths := []int{}
	for _, manuel := range manuels {
		result := findFewestButtonPress(0, 1, manuel.WantedState, manuel.Buttons)
		depths = append(depths, result)
	}
	return depths
}

func findFewestButtonPress(currentVal, depth int, wantedState LightDiagram, values []Button) int {
	resultDepth := math.MaxInt
	for idx, val := range values {
		result := currentVal ^ val
		if result == wantedState {
			return depth
		}
		// fmt.Printf("depth %d val %d result %d\n", depth, val, result)
		findDepth := findFewestButtonPress(result, depth+1, wantedState, values[idx+1:])
		if findDepth < resultDepth {
			resultDepth = findDepth
		}
		findDepth = findFewestButtonPress(currentVal, depth+1, wantedState, values[idx+1:])
		if findDepth < resultDepth {
			resultDepth = findDepth
		}
	}
	return resultDepth
}

func parseInput(input string) ([]Manual, error) {
	result := []Manual{}
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		// parse light diag
		lightDiagramReg := regexp.MustCompile(`(\.|#)+`)
		lightDiagramStr := string(lightDiagramReg.Find([]byte(line)))
		var wantedState LightDiagram
		var size int
		for _, val := range lightDiagramStr {
			size++
			wantedState = wantedState << 1
			switch val {
			case '.':
				wantedState += 0
			case '#':
				wantedState += 1
			default:
				return nil, fmt.Errorf("unexpected light state, got: %q", val)
			}
		}

		// parse buttons
		buttonsReg := regexp.MustCompile(`(\((\d|,)+\))+`)
		buttonsByteArr := buttonsReg.FindAll([]byte(line), -1)
		var buttons []Button
		for _, buttonVal := range buttonsByteArr {
			buttonStr := string(buttonVal)
			buttonStr = strings.TrimPrefix(buttonStr, "(")
			buttonStr = strings.TrimSuffix(buttonStr, ")")
			digits := strings.Split(buttonStr, ",")
			var button Button
			for _, digitStr := range digits {
				digit, err := strconv.Atoi(digitStr)
				if err != nil {
					return nil, err
				}
				factor := size - digit - 1
				button = button | (1 << factor)
				// fmt.Printf("size: %d factor: %d digit: %d button: %d\n", size, factor, digit, button)
			}
			// fmt.Println()
			buttons = append(buttons, button)
		}
		// fmt.Println()
		man := Manual{
			WantedState: wantedState,
			// IsState:     isState,
			Buttons: buttons,
		}
		result = append(result, man)
	}
	return result, nil
}
