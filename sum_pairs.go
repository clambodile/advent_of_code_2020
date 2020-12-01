package advent_of_code_2020

import "errors"

func SumPairs(nums []int, target int) (first, second int, err error) {

	//entries of nums which were previously checked for companionship
	checked := map[int]bool{}

	for _, num := range nums {
		//"companion" num which, added to first num, yields target num
		companion := target - num
		_, ok := checked[companion]
		if ok {
			return num, companion, nil
		} else {
			checked[num] = true
		}
	}
	return 0, 0, errors.New("no inputs match target")
}
