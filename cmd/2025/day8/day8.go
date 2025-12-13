package day8

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"

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
		Use:   "8",
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
	connections := findNearestConnections(boxes)
	sort.Slice(connections, func(i, j int) bool {
		return connections[i].distance < connections[j].distance
	})
	circuits := connect1(connections, 1000)
	sort.Slice(circuits, func(i, j int) bool {
		return len(circuits[i]) < len(circuits[j])
	})
	// sort.Slice(boxes, func(i, j int) bool {
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
	// connectionSizes := utils.GetMapValues(connections)
	// sort.Slice(connectionSizes, func(i, j int) bool {
	// 	return connectionSizes[i] < connectionSizes[j]
	// })
	// top3Values := connectionSizes[len(connectionSizes)-3:]
	// fmt.Println(top3Values)
	top3Values := circuits[len(circuits)-3:]
	result := 1
	for _, val := range top3Values {
		result *= len(val)
	}

	fmt.Printf("Solution 1: %d\n", result)
	// for _, val := range boxes {
	// 	info := ""
	// 	if val.Idx != nil {
	// 		info = fmt.Sprintf("Result Idx: %d", **val.Idx)
	// 	}
	// 	fmt.Printf("%d,%d,%d %s\n", val.X, val.Y, val.Z, info)
	// }
	return nil
}

func task2(input string) error {
	boxes, err := parseInput(input)
	if err != nil {
		return err
	}
	connections := findNearestConnections(boxes)
	sort.Slice(connections, func(i, j int) bool {
		return connections[i].distance < connections[j].distance
	})
	lastConnection := connect2(connections)

	fmt.Printf("Solution 2: %d\n", boxes[lastConnection.a].X*boxes[lastConnection.b].X)
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

func findNearestConnections(boxes []Box) []Connection {
	boxesLen := len(boxes)
	sum := []Connection{}
	for i := range boxes {
		for j := i; j < boxesLen; j++ {
			if i == j { // skip itself
				continue
			}
			distance := boxes[i].distance(boxes[j])
			sum = append(sum, Connection{i, j, distance})
		}
	}
	return sum
}

func connect1(connections []Connection, times int) []Circuit {
	circuits := []Circuit{}
	connectionsSize := len(connections)
	for idx := range times {
		if idx < connectionsSize {
			conn := connections[idx]
			a := -1
			b := -1
			for cIdx, circuit := range circuits {
				for _, boxIdx := range circuit {
					if boxIdx == conn.a {
						a = cIdx
					}
					if boxIdx == conn.b {
						b = cIdx
					}
				}
			}
			if a == b && a != -1 && b != -1 {
				continue
			}
			if a != -1 && b != -1 {
				// append b to a
				circuits[a] = append(circuits[a], circuits[b]...)
				// remove b
				circuits = append(circuits[:b], circuits[b+1:]...)
				continue
			}
			if a == -1 && b != -1 {
				// add a to b
				circuits[b] = append(circuits[b], conn.a)
				continue
			}
			if a != -1 && b == -1 {
				// add b to a
				circuits[a] = append(circuits[a], conn.b)
				continue
			}
			if a == -1 && b == -1 {
				// create new circuit
				circuits = append(circuits, Circuit{conn.a, conn.b})
				continue
			}
		}
	}
	return circuits
}

func connect2(connections []Connection) Connection {
	circuits := []Circuit{}
	lastConnection := Connection{}
	for idx := range connections {
		conn := connections[idx]
		a := -1
		b := -1
		for cIdx, circuit := range circuits {
			for _, boxIdx := range circuit {
				if boxIdx == conn.a {
					a = cIdx
				}
				if boxIdx == conn.b {
					b = cIdx
				}
			}
		}
		beforeSize := len(circuits)
		if a == b && a != -1 && b != -1 {
			continue
		}
		if a != -1 && b != -1 {
			// append b to a
			circuits[a] = append(circuits[a], circuits[b]...)
			// remove b
			circuits = append(circuits[:b], circuits[b+1:]...)
			afterSize := len(circuits)
			if beforeSize == 2 && afterSize == 1 {
				lastConnection = conn
			}
			continue
		}
		if a == -1 && b != -1 {
			// add a to b
			circuits[b] = append(circuits[b], conn.a)
			lastConnection = conn
			continue
		}
		if a != -1 && b == -1 {
			// add b to a
			circuits[a] = append(circuits[a], conn.b)
			lastConnection = conn
			continue
		}
		if a == -1 && b == -1 {
			circuits = append(circuits, Circuit{conn.a, conn.b})
			continue
		}
	}
	return lastConnection
}

// func findNearestConnections(boxes []Box, totalConnections int) map[int]int {
// 	sum := map[int]int{}
// 	totalConnections = totalConnections - 1 // must be calculated -1
// 	for idx := range totalConnections {
// 		var closestI int
// 		var closestJ int
// 		closestDistance := math.MaxFloat64
// 		for i := range boxes {
// 			for j := range boxes {
// 				if boxes[i].Idx != nil && boxes[j].Idx != nil && *boxes[i].Idx == *boxes[j].Idx {
// 					continue
// 				}
// 				if i == j {
// 					continue
// 				}
// 				if i == closestJ && j == closestI {
// 					continue
// 				}
// 				distance := boxes[i].distance(boxes[j])
// 				if distance < closestDistance {
// 					closestI, closestJ = i, j
// 					closestDistance = distance
// 				}
// 			}
// 		}
// 		firstBox := &boxes[closestI]
// 		secondBox := &boxes[closestJ]
// 		if firstBox.Idx != nil && secondBox.Idx != nil {
// 			sum[**firstBox.Idx] += sum[**secondBox.Idx]
// 			delete(sum, **secondBox.Idx)
// 			*secondBox.Idx = *firstBox.Idx
// 		} else if firstBox.Idx != nil {
// 			secondBox.Idx = firstBox.Idx
// 			sum[**firstBox.Idx]++
// 		} else if secondBox.Idx != nil {
// 			firstBox.Idx = secondBox.Idx
// 			sum[**firstBox.Idx]++
// 		} else {
// 			pointerIdx := &idx
// 			firstBox.Idx = &pointerIdx
// 			secondBox.Idx = &pointerIdx
// 			sum[**firstBox.Idx]++
// 			sum[**firstBox.Idx]++
// 		}
// 	}
// 	return sum
// }
