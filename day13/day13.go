package day13

import (
	"github.com/clambodile/advent_of_code_2020/util/io"
	mth "github.com/clambodile/advent_of_code_2020/util/math"
	"math"
	"strconv"
	"strings"
)

//first line is earliest time available to leave
//inputs are either blank or are a number
//the numbers are multiple
//which multiple can be multiplied by itself repeatedly to get the smallest gap above the target time?

func Challenge1() (int, error) {
	lines, err := io.ReadLines("./day13/d13.txt")
	if err != nil {
		return 0, err
	}
	target, err := strconv.Atoi(lines[0])
	if err != nil {
		return 0, err
	}
	intervals := make([]int, 0)
	chars := strings.Split(lines[1], ",")
	for _, char := range chars {
		interval, err := strconv.Atoi(char)
		if err != nil {
			continue
		}
		intervals = append(intervals, interval)
	}

	id := 0
	minWait := math.MaxInt64
	for _, interval := range intervals {
		i := 0
		for ; i < target; i += interval {
		}
		if i-target < minWait {
			minWait = i - target
			id = interval
		}
	}
	return id * minWait, nil
}

func Challenge2() (int, error) {
	lines, err := io.ReadLines("./day13/d13.txt")
	if err != nil {
		return 0, err
	}

	chars := strings.Split(lines[1], ",")
	intervals := make([]int, len(chars))
	for i, char := range chars {
		if char == "x" {
			intervals[i] = 1
		} else {
			num, err := strconv.Atoi(char)
			if err != nil {
				return 0, err
			}
			intervals[i] = num
		}
	}
	sequence := ascendingMultipliersSpecialI(intervals)
	return sequence[0], nil
}

//given a list of integers, what is the smallest set of multiples of those integers, such that
//the integers occur with a specified gap between each one?

func smallestMultipleAbove(target, multiplier float64) float64 {
	smallestMultiple := math.Ceil(target/multiplier) * multiplier
	if smallestMultiple == target {
		smallestMultiple += multiplier
	}
	return smallestMultiple
}

func smallestMultipleAboveI(target, multiplier int) int {
	return int(smallestMultipleAbove(float64(target), float64(multiplier)))
}

func ascendingMultipliers(multipliers []float64) []float64 {
	original := multipliers
Next:
	prev := multipliers[0]
	for i := 1; i < len(multipliers); i++ {
		num := original[i]
		num = smallestMultipleAbove(prev, num)
		if num-prev != 1 {
			multipliers[0] += original[0]
			goto Next
		}
		prev = num
		multipliers[i] = num

	}
	return multipliers
}
func ascendingMultipliersSpecial(multipliers []float64) []float64 {
	original := multipliers
	multipliers[0] = smallestMultipleAbove(100000000000000, multipliers[0])
Next:
	prev := multipliers[0]
	for i := 1; i < len(multipliers); i++ {
		num := original[i]
		num = smallestMultipleAbove(prev, num)
		if num-prev != 1 {
			multipliers[0] += original[0]
			goto Next
		}
		prev = num
		multipliers[i] = num

	}
	return multipliers
}

func ascendingMultipliersI(multipliers []int) []int {
	return mth.FloatsToInts(ascendingMultipliers(mth.IntsToFloats(multipliers)))
}

func ascendingMultipliersSpecialI(multipliers []int) []int {
	return mth.FloatsToInts(ascendingMultipliersSpecial(mth.IntsToFloats(multipliers)))
}

func ascendingMultipliersNoDiv(multipliers []int) []int {
	//create chains of factors
	type product struct {
		Value int
		Multiplier int
		Multiplicand int
	}
	products := map[int]product{}

}