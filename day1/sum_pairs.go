package day1

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

func SumTrips(nums[]int, target int) (first, second, third int, err error) {
	var singles []int
	pairSums := map[int][]int{}
	for _, n := range nums {
		rem := target - n
		companions, ok := pairSums[rem]
		if ok {
			return companions[0], companions[1], n, nil
		}
		for _, single := range singles {
			pairSums[single + n] = []int{single, n}
		}
		singles = append(singles, n)
	}
	return 0, 0, 0, errors.New("no inputs match target")
}