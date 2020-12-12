package main

import (
	"fmt"
	"github.com/clambodile/advent_of_code_2020/day11"
)

func main() {
	//logSolution(1, 1, day1.Challenge1)
	//logSolution(1, 2, day1.Challenge2)
	//logSolution(2, 1, day2.Challenge1)
	//logSolution(2, 2, day2.Challenge2)
	//logSolution(3, 1, day3.Challenge1)
	//logSolution(3, 2, day3.Challenge2)
	//logSolution(4, 1, day4.Challenge1)
	//logSolution(4, 2, day4.Challenge2)
	//logSolution(5, 1, day5.Challenge1)
	//logSolution(5, 2, day5.Challenge2)
	//logSolution(6, 1, day6.Challenge1)
	//logSolution(6, 2, day6.Challenge2)
	//logSolution(7, 1, day7.Challenge1)
	//logSolution(7, 2, day7.Challenge2)
	//logSolution(8, 1, day8.Challenge1)
	//logSolution(8, 2, day8.Challenge2)
	//logSolution(9, 1, day9.Challenge1)
	//logSolution(9, 2, day9.Challenge2)
	//logSolution(10, 1, day10.Challenge1)
	//logSolution(10, 2, day10.Challenge2)
	//logSolution(11, 1, day11.Challenge1)
	logSolution(11, 2, day11.Challenge2)
}

func logSolution(dayN, cN int, f func() (int, error)) {
	solution, err := f()
	if err != nil {
		fmt.Printf("Error in day %d challenge %d: %s\n", dayN, cN, err)
	} else {
		fmt.Printf("Day %d challenge %d solution: %d\n", dayN, cN, solution)
	}
}
