package day9

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/marceljk/advent-of-code/pkg/utils/globalflags"
	"github.com/marceljk/advent-of-code/pkg/utils/runner"
	"github.com/spf13/cobra"
)

func NewCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "9",
		Short: "Solution for 2025 day 9",
		Long:  "Solution for 2025 day 9.",
		Args:  cobra.NoArgs,
		RunE:  runner.GenericRunE(task1, task2),
	}
	globalflags.ConfigureFlags(cmd)
	return cmd
}

type Point struct {
	X, Y int
}

func (p Point) area(another Point) int64 {
	xDiff := math.Abs(float64(p.X-another.X)) + 1
	yDiff := math.Abs(float64(p.Y-another.Y)) + 1
	return int64(xDiff * yDiff)
}

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}

func parseInput(input string) []Point {
	points := []Point{}
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		pointValues := strings.Split(line, ",")
		x, err := strconv.Atoi(pointValues[0])
		handleErr(err)
		y, err := strconv.Atoi(pointValues[1])
		handleErr(err)
		points = append(points, Point{X: x, Y: y})
	}
	return points
}

func task1(input string) error {
	points := parseInput(input)
	lenPoints := len(points)
	var biggestArea int64 = math.MinInt64
	for i := range points {
		for j := i + 1; j < lenPoints; j++ {
			area := points[i].area(points[j])
			if area > biggestArea {
				biggestArea = area
			}
		}
	}
	fmt.Printf("Solution 1: %d\n", biggestArea)
	return nil
}

func task2(input string) error {
	// points := parseInput(input)
	// fmt.Printf("Solution 2: %d\n", biggestArea)
	return nil
}
