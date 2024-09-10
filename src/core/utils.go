package core

import "math"

// Map manipulates a slice and transforms it to a slice of another type
//
// Source: https://github.com/samber/lo/blob/master/slice.go#L25
func Map[T any, R any](collection []T, iteratee func(item T, index int) R) []R {
	result := make([]R, len(collection))

	for i, item := range collection {
		result[i] = iteratee(item, i)
	}

	return result
}

// Fract returns the fractional portion of a float32.
// For example, 3.1415 returns 0.1415
func Fract(f float32) float32 {
	_, fractional := math.Modf(float64(f))
	return float32(fractional)
}
