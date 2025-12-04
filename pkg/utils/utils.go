package utils

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

// ConcateUint concatenates the input uint values to one value
func ConcateUint[T uint8 | uint16 | uint | uint64](input ...T) uint64 {
	var result uint64
	for _, val := range input {
		result = result*10 + uint64(val)
	}
	return result
}

// CountOccurrences counts how often a specifc entry occurce in a slice
func CountOccurrences[T comparable](input []T, searchValue T) (count int) {
	for idx := range input {
		if input[idx] == searchValue {
			count++
		}
	}
	return
}
