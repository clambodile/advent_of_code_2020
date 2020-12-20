package product_runs

import (
	"fmt"
)

func NextGreaterProduct(target, multiplier int) (int, error) {
	result := multiplier
	if multiplier == 0 && multiplier < target {
		return 0, fmt.Errorf("no solution for target %d and input 0", target)
	}
	if multiplier < 0 {
		result = -result
	}
	for result <= target {
		result += multiplier
	}
	return result, nil
}
func NextGreaterOrEqualProduct(target, multiplier int) (int, error) {
	result := multiplier
	if multiplier == 0 && multiplier < target {
		return 0, fmt.Errorf("no solution for target %d and input 0", target)
	}
	if multiplier < 0 {
		result = -result
	}
	for result < target {
		result += multiplier
	}
	return result, nil
}

//find the period of gaps of two ints
func ProductGaps(x, y int) []int {
	if x == y {
		return []int{}
	}
	i := x
	j := y
	//start at the first non-negative gap
	j, _ = NextGreaterOrEqualProduct(i, y)
	gap := j - i
	gaps := []int{gap}
	for i = i + x;; i += x {
		if j < i {
			j, _ = NextGreaterOrEqualProduct(i, y)
		}
		gap = j - i
		if gap == gaps[0] {
			break
		}
		gaps = append(gaps, gap)
	}
	return gaps
}

func findInt(ints []int, target int) (index int) {
	for i, num := range ints {
		if num == target {
			return i
		}
	}
	return -1
}
//take a list of multipliers and multiply them until they are a strictly ascending list
//5, 3, 2
//5, 6, 8
//10, 12
//15, 18
//20, 21, 22
func ProductRuns(ints []int) []int {
	gapses := make([][]int, len(ints) - 1)
	for i := 0; i < len(ints)-1; i++ {
		x := ints[i]
		y := ints[i + 1]
		gaps := ProductGaps(x, y)
		gapses[i] = gaps
	}
	product := 1
	for _, gaps := range gapses {
		i := findInt(gaps, 1) + 1
		product *= i
	}
}