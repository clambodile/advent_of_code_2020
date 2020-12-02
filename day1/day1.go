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

	//first, second, err := SumPairs(nums, 2020)
	answers, err := SumN(nums, 2, 2020)
	if err != nil {
		return 0, err
	}
	first := answers[0]
	second := answers[1]
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

	//first, second, third, err := SumTrips(nums, 2020)
	answers, err := SumN(nums, 3, 2020)
	first := answers[0]
	second := answers[1]
	third := answers[2]
	if err != nil {
		return 0, err
	}
	return first * second * third, nil
}

func ChallengeX(n int) (int, error) {
	nums, err := io.ReadIntsFile("./day1/d1c1.txt")
	if err != nil {
		return 0, err
	}

	answers, err := SumN(nums, n, 2020)
	product := 1
	for _, answer := range answers {
		product *= answer
	}
	if err != nil {
		return 0, err
	}
	return product, nil
}