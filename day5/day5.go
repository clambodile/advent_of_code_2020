package day5

import (
	"errors"
	"fmt"
	"github.com/clambodile/advent_of_code_2020/io"
	"math"
)

func Challenge1() (int, error) {
	lines, err := io.ReadLines("./day5/d5.txt")
	if err != nil {
		return 0, err
	}
	max := 0
	for _, line := range lines {
		seatID, err := readSeatCode(line)
		if err != nil {
			return 0, err
		}
		if seatID > max {
			max = seatID
		}
	}
	return max, nil
}

func Challenge2() (int, error) {

	lines, err := io.ReadLines("./day5/d5.txt")
	if err != nil {
		return 0, err
	}
	seats := [843]bool{}
	for _, line := range lines {
		seatID, err := readSeatCode(line)
		if err != nil {
			return 0, err
		}
		seats[seatID] = true
	}
	for i := 0; i < len(seats) - 2; i++ {
		if (seats[i] == true) && (seats[i + 1] != true) && (seats[i + 2] == true) {
			return i + 1, nil
		}
	}
	return 0, errors.New("solution not found")
}

func readSeatCode(code string) (int, error) {
	rowBits := make([]int, 7)
	colBits := make([]int, 3)
	for i, rn := range code {
		if rn == 'F' {
			rowBits[i] = 0
		} else if rn == 'B' {
			rowBits[i] = 1
		} else if rn == 'L' {
			colBits[i - 7] = 0
		} else if rn == 'R' {
			colBits[i - 7] = 1
		} else {
			return 0, fmt.Errorf("unrecognized token: %c", rn)
		}
	}
	row := readBinaryPartition(rowBits, 128)
	col := readBinaryPartition(colBits, 8)
	seatID := row * 8 + col
	return seatID, nil
}

func readBinaryPartition(bin []int, n int) int {
	acc := 0
	for i, bit := range bin {
		if bit == 1 {
			acc += n / int(math.Pow(2, float64(i+1)))
		}
	}
	return acc
}