package utils

import (
	"fmt"
	"strconv"
)

// IsSliceEqual compares two rune slices and returns true, if their content is equal
func IsSliceEqual[K comparable](x, y []K) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range len(x) {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}

// GetMapKeys returns a slice with all keys of the input map
func GetMapKeys[K comparable, V any](input map[K]V) []K {
	result := []K{}
	for k := range input {
		result = append(result, k)
	}
	return result
}

// FindLargestValue takes as input an uint slice and returns the largest digit in this slice with their value in this slice
func FindLargestValue[T ~[]uint](input T) (largestVal uint, largestIdx int) {
	for inputIdx, inputVal := range input {
		if largestVal < inputVal {
			largestVal, largestIdx = inputVal, inputIdx
		}
	}
	return
}

// ConcateNum concatenates the input uint values to one value
func ConcateNum[T uint | int](input ...T) (uint64, error) {
	var concat string
	for _, val := range input {
		concat += fmt.Sprintf("%d", val)
	}
	result, err := strconv.ParseUint(concat, 10, 64)
	if err != nil {
		return 0, err
	}
	return result, nil
}
