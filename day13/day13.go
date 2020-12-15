package day13

import (
	"github.com/clambodile/advent_of_code_2020/util/io"
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
		for ; i < target; i += interval {}
		if i - target < minWait {
			minWait = i - target
			id = interval
		}
	}
	return id * minWait, nil
}