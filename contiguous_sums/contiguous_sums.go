package contiguous_sums

import "errors"

func FindContiguousSummands(nums []int, target int) ([]int, error) {
	length := len(nums)
	for i, num := range nums {
		sum := num
		summands := []int{num}
		for j := i + 1; sum < target && j < length; j++ {
			nextNum := nums[j]
			sum += nextNum
			summands = append(summands, nextNum)
		}
		if sum == target {
			return summands, nil
		}
	}
	return nil, errors.New("no solution found")
}
