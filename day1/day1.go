package day1

import (
	"github.com/clambodile/advent_of_code_2020/io"
)

//given an input list of integers, find a pair which sum to 2020 and return the product
func Challenge1() (int, error) {
	nums, err := io.ReadIntsFile("./day1/d1c1.txt")
	if err != nil {
		return 0, err
	}

	first, second, err := SumPairs(nums, 2020)
	if err != nil {
		return 0, err
	}
	return first * second, nil
}

//find a triad which sum to 2020 and return the product
func Challenge2() (int, error) {
	nums, err := io.ReadIntsFile("./day1/d1c1.txt")
	if err != nil {
		return 0, err
	}

	first, second, third, err := SumTrips(nums, 2020)
	if err != nil {
		return 0, err
	}
	return first * second * third, nil
}