package day10

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/clambodile/advent_of_code_2020/util/io"
	mbth "github.com/clambodile/advent_of_code_2020/util/math"
	"math"
	"reflect"
	"sort"
)

func Challenge1() (int, error) {
	filename := "./day10/d10.txt"
	nums, err := io.ReadIntsFile(filename)
	nums = append(nums, 0)
	if err != nil {
		return 0, err
	}
	sort.Ints(nums)
	oneGaps := 0
	threeGaps := 1
	for i := 1; i < len(nums); i++ {
		num := nums[i]
		prev := nums[i-1]
		gap := int(math.Abs(float64(num - prev)))
		if gap == 1 {
			oneGaps++
		}
		if gap == 3 {
			threeGaps++
		}
	}
	return oneGaps * threeGaps, nil
}

func Challenge2() (int, error) {
	filename := "./day10/d10.txt"
	nums, err := io.ReadIntsFile(filename)
	if err != nil {
		return 0, err
	}
	max := mbth.MaxInt(nums)
	outlet := max + 3
	nums = append(nums, 0, outlet)
	sort.Ints(nums)
	//how many valid orderings?
	validSequences, err := countValidSequences(nums)
	if err != nil {
		return 0, err
	}
	fmt.Printf("%+v", reflect.TypeOf(countValidSequences))
	return validSequences, nil
}

var memo = map[string]int{}

func countValidSequences(nums []int) (int, error) {
	keyBytes, _ := json.Marshal(nums)
	key := string(keyBytes)
	answer, exists := memo[key]
	if exists {
		return answer, nil
	}
	if len(nums) == 1 {
		return 1, nil
	}
	total := 0
	first := nums[0]
	//a sequence is valid if the max gap between adjacent elements is 3
	for i := 1; i <= 3 && i < len(nums); i++ {
		nth := nums[i]
		gap := nth - first
		if gap < 0 {
			return 0, errors.New("list not sorted")
		}
		if gap <= 3 {
			subsequences, err := countValidSequences(nums[i:])
			if err != nil {
				return 0, err
			}
			total += subsequences
		}
	}
	memo[key] = total
	return total, nil
}
