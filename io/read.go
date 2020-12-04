package io

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func ReadIntsFile(filename string) ([]int, error) {
	var nums []int
	input, err := ioutil.ReadFile(filename)
	if err != nil {
		return nums, err
	}
	str := strings.TrimSpace(string(input))
	strs := strings.Split(str, "\n")
	for _, s := range strs {
		n, err := strconv.Atoi(s)
		if err != nil {
			return nums, err
		}
		nums = append(nums, n)
	}
	return nums, nil
}

func ReadLines(filename string) ([]string, error) {
	input, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	str := strings.TrimSpace(string(input))
	lines := strings.Split(str, "\n")
	return lines, nil
}

func ReadLineGroups(filename string) ([]string, error) {
	input, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	str := strings.TrimSpace(string(input))
	lines := strings.Split(str, "\n\n")
	return lines, nil
}