package day4

import (
	"fmt"

	"github.com/marceljk/advent-of-code/pkg/utils"
)

type Field string
type Row []Field
type Diagram []Row

const (
	PaperRoll Field = "@"
	Empty     Field = "."
	Remove    Field = "X"
)

func (d Diagram) GetNeighbors(cell, row int) ([]Field, error) {
	cellsLength := len(d)
	if row < 0 || cell < 0 {
		return nil, fmt.Errorf("invalid input. negativ input")
	}
	if cell >= cellsLength {
		return nil, fmt.Errorf("invalid input. cell value is too big")
	}
	rowLength := len(d[cell]) // Assuming all rows have the same length
	if row >= rowLength {
		return nil, fmt.Errorf("invalid input. row value is too big")
	}
	var result []Field
cellLoop:
	for x := -1; x < 2; x++ {
	rowLoop:
		for y := -1; y < 2; y++ {
			cellIdx := cell + x
			if cellIdx < 0 || cellIdx >= cellsLength {
				continue cellLoop
			}
			rowIdx := row + y
			if rowIdx < 0 || rowIdx >= rowLength {
				continue rowLoop
			}
			// skip itself
			if x == 0 && y == 0 {
				continue
			}
			result = append(result, d[cellIdx][rowIdx])
		}
	}
	return result, nil
}

func (d Diagram) canRoleBeRemoved(cell, row int) (bool, error) {
	if d[cell][row] != PaperRoll {
		return false, nil
	}
	neightbors, err := d.GetNeighbors(cell, row)
	if err != nil {
		return false, fmt.Errorf("can not find neighbors: %w", err)
	}
	paperRolls := utils.CountOccurrences(neightbors, PaperRoll)
	return paperRolls < 4, nil

}

func (d Diagram) deepCopy() Diagram {
	var newDiagram Diagram = make([]Row, len(d))
	for cellIdx := range d {
		newDiagram[cellIdx] = make(Row, len(d[cellIdx]))
		copy(newDiagram[cellIdx], d[cellIdx])
	}
	return newDiagram
}

func (d Diagram) print() {
	for _, row := range d {
		fmt.Printf("%+v\n", row)
	}
}

func (d Diagram) RemovableRolls() (removableRolls int, newDiagram Diagram, err error) {
	newDiagram = d.deepCopy()
	for cellIdx := range d {
		for rowIdx := range d[cellIdx] {
			var ok bool
			if ok, err = d.canRoleBeRemoved(cellIdx, rowIdx); err != nil {
				return
			} else if ok {
				removableRolls++
				newDiagram[cellIdx][rowIdx] = Remove
			}
		}
	}
	return
}
