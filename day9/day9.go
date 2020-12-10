package day9

import (
	"github.com/clambodile/advent_of_code_2020/contiguous_sums"
	"github.com/clambodile/advent_of_code_2020/io"
	"github.com/clambodile/advent_of_code_2020/preamble_sums"
	"math"
)

func Challenge1() (int, error) {
	filename := "./day9/d9.txt"
	nums, err := io.ReadIntsFile(filename)
	if err != nil {
		return 0, err
	}
	preambleLength := 25
	invalids := preamble_sums.ValidatePreambleSums(nums, preambleLength)
	firstInvalid := nums[invalids[0]]
	return firstInvalid, nil
}

func Challenge2() (int, error) {
	filename := "./day9/d9.txt"
	nums, err := io.ReadIntsFile(filename)
	if err != nil {
		return 0, err
	}
	target := 18272118
	summands, err := contiguous_sums.FindContiguousSummands(nums, target)
	if err != nil {
		return 0, err
	}
	min := math.MaxInt64
	max := 0
	for _, summand := range summands {
		if summand < min {
			min = summand
		}
		if summand > max {
			max = summand
		}
	}
	return min + max, nil
}