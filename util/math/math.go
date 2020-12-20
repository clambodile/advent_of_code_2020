package math

import "math"

func MaxInt(nums []int) int {
	max := math.MinInt64
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
}

func FloatsToInts(floats []float64) []int {
	ints := make([]int, len(floats))
	for i, f := range floats {
		ints[i] = int(f)
	}
	return ints
}

func IntsToFloats(ints []int) []float64 {
	floats := make([]float64, len(ints))
	for i, n := range ints {
		floats[i] = float64(n)
	}
	return floats
}
