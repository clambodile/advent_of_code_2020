package main

import (
	"fmt"
	"github.com/clambodile/advent_of_code_2020/day1"
	"github.com/clambodile/advent_of_code_2020/day2"
	"github.com/clambodile/advent_of_code_2020/day3"
	"github.com/clambodile/advent_of_code_2020/day4"
)

func main() {
	logSolution(1, 1, day1.Challenge1)
	logSolution(1, 2, day1.Challenge2)
	logSolution(2, 1, day2.Challenge1)
	logSolution(2, 2, day2.Challenge2)
	logSolution(3, 1, day3.Challenge1)
	logSolution(3, 2, day3.Challenge2)
	logSolution(4, 1, day4.Challenge1)
	logSolution(4, 2, day4.Challenge2)
}

func logSolution(dayN, cN int, f func() (int, error))  {
	solution, err := f()
	if err != nil {
		fmt.Printf("Error in day %d challenge %d: %s\n", dayN, cN,err)
	} else {
		fmt.Printf("Day %d challenge %d solution: %d\n", dayN, cN, solution)
	}
}