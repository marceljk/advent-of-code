package day6

import "fmt"

type Operation string

type Problem struct {
	Operation Operation
	Numbers   []int
}

func (p Problem) calcAdd() (total int) {
	for _, num := range p.Numbers {
		total += num
	}
	return
}

func (p Problem) calcMul() (total int) {
	total = 1
	for _, num := range p.Numbers {
		total *= num
	}
	return
}

func (p Problem) Calculate() (int, error) {
	switch p.Operation {
	case AddOp:
		return p.calcAdd(), nil
	case MulOp:
		return p.calcMul(), nil
	default:
		return 0, fmt.Errorf("no operation set")
	}
}

const (
	AddOp Operation = "+"
	MulOp Operation = "*"
)
