package day8b

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/marceljk/advent-of-code/pkg/utils"
	"github.com/marceljk/advent-of-code/pkg/utils/globalflags"
	"github.com/marceljk/advent-of-code/pkg/utils/runner"
	"github.com/spf13/cobra"
)

type Circuit []int

type Box struct {
	X, Y, Z int
	Idx     **int
}

type Connection struct {
	a, b     int
	distance float64
}

func (b Box) parse(input string) (*Box, error) {
	coordinates := strings.Split(input, ",")
	x, err := strconv.Atoi(coordinates[0])
	if err != nil {
		return nil, err
	}
	y, err := strconv.Atoi(coordinates[1])
	if err != nil {
		return nil, err
	}
	z, err := strconv.Atoi(coordinates[2])
	if err != nil {
		return nil, err
	}
	return &Box{
		X: x,
		Y: y,
		Z: z,
	}, nil
}

func (from Box) distance(to Box) float64 {
	x := (from.X - to.X) * (from.X - to.X)
	y := (from.Y - to.Y) * (from.Y - to.Y)
	z := (from.Z - to.Z) * (from.Z - to.Z)
	return math.Sqrt(float64(x + y + z))
}

func NewCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "8b",
		Short: "Solution for 2025 day 8",
		Long:  "Solution for 2025 day 8.",
		Args:  cobra.NoArgs,
		RunE:  runner.GenericRunE(task1, task2),
	}
	globalflags.ConfigureFlags(cmd)
	return cmd
}

func task1(input string) error {
	boxes, err := parseInput(input)
	if err != nil {
		return err
	}
	connections := findNearestConnections(boxes, 1000)
	// sort.Slice(sum, func(i, j int) bool {
	// 	if boxes[i].Idx == nil && boxes[j].Idx != nil {
	// 		return true
	// 	}
	// 	if boxes[i].Idx != nil && boxes[j].Idx == nil {
	// 		return false
	// 	}
	// 	if boxes[i].Idx == nil && boxes[j].Idx == nil {
	// 		return true
	// 	}
	// 	if boxes[i].Idx != nil && boxes[j].Idx != nil {
	// 		return **boxes[i].Idx < **boxes[j].Idx
	// 	}
	// 	return false
	// })
	connectionSizes := utils.GetMapValues(connections)
	sort.Slice(connectionSizes, func(i, j int) bool {
		return connectionSizes[i] < connectionSizes[j]
	})
	top3Values := connectionSizes[len(connectionSizes)-3:]
	result := 1
	for _, val := range top3Values {
		result *= val
	}

	fmt.Printf("Solution 1: %d\n", result)
	return nil
}

func task2(input string) error {
	return nil
}

func parseInput(input string) ([]Box, error) {
	var boxes []Box
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		box, err := Box{}.parse(line)
		if err != nil {
			return nil, err
		}
		boxes = append(boxes, *box)
	}
	return boxes, nil
}

func findNearestConnections(boxes []Box, totalConnections int) map[int]int {
	sum := map[int]int{}
	totalConnections = totalConnections - 1 // must be calculated -1
	for idx := range totalConnections {
		var closestI int
		var closestJ int
		closestDistance := math.MaxFloat64
		for i := range boxes {
			for j := range boxes {
				if boxes[i].Idx != nil && boxes[j].Idx != nil && *boxes[i].Idx == *boxes[j].Idx {
					continue
				}
				if i == j {
					continue
				}
				if i == closestJ && j == closestI {
					continue
				}
				distance := boxes[i].distance(boxes[j])
				if distance < closestDistance {
					closestI, closestJ = i, j
					closestDistance = distance
				}
			}
		}
		firstBox := &boxes[closestI]
		secondBox := &boxes[closestJ]
		if firstBox.Idx != nil && secondBox.Idx != nil {
			sum[**firstBox.Idx] += sum[**secondBox.Idx]
			delete(sum, **secondBox.Idx)
			*secondBox.Idx = *firstBox.Idx
		} else if firstBox.Idx != nil {
			secondBox.Idx = firstBox.Idx
			sum[**firstBox.Idx]++
		} else if secondBox.Idx != nil {
			firstBox.Idx = secondBox.Idx
			sum[**firstBox.Idx]++
		} else {
			pointerIdx := &idx
			firstBox.Idx = &pointerIdx
			secondBox.Idx = &pointerIdx
			sum[**firstBox.Idx]++
			sum[**firstBox.Idx]++
		}
	}
	return sum
}
