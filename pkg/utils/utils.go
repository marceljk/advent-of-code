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
