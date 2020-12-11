package day11

import (
	"fmt"
	"github.com/clambodile/advent_of_code_2020/util/io"
)

func Challenge1() (int, error) {
	filename := "./day11/d11.txt"
	lines, err := io.ReadLines(filename)
	if err != nil {
		return 0, err
	}
	fp := parseFloorPlan(lines)
	fmt.Println(fp.totalFilled())
	next := fp.applyRules()
	for !fp.equals(next) {
		fp = next
		fmt.Println(fp.totalFilled())
		next = fp.applyRules()
	}
	return fp.totalFilled(), nil
}

const (
	floor      = 0
	emptySeat  = 1
	filledSeat = 2
)

type floorPlan [][]int

func (fp *floorPlan) countFilledAdjacent(y, x int) int {
	total := 0
	height := len(*fp)
	width := len((*fp)[0])
	for i := y - 1; i <= y+1; i++ {
		for j := x - 1; j <= x+1; j++ {
			if i == y && j == x {
				continue
			}
			if i >= height {
				continue
			}
			if j >= width {
				continue
			}
			if i < 0 || j < 0 {
				continue
			}
			if (*fp)[i][j] == filledSeat {
				total++
			}
		}
	}
	return total
}

func (fp *floorPlan) applyRules() *floorPlan {
	newFloorPlan := make(floorPlan, len(*fp))
	for i, line := range *fp {
		newLine := make([]int, len(line))
		newFloorPlan[i] = newLine
		for j, cell := range line {
			if cell == floor {
				newFloorPlan[i][j] = floor
				continue
			}
			count := fp.countFilledAdjacent(i, j)
			if cell == emptySeat && count == 0 {
				newFloorPlan[i][j] = filledSeat
				continue
			}
			if cell == filledSeat && count >= 4 {
				newFloorPlan[i][j] = emptySeat
				continue
			}
			newFloorPlan[i][j] = cell
		}
	}
	return &newFloorPlan
}

func (fp *floorPlan) equals(fp2 *floorPlan) bool {
	for i, line := range *fp {
		for j, cell := range line {
			if cell != (*fp2)[i][j] {
				return false
			}
		}
	}
	return true
}

func (fp *floorPlan) totalFilled() int {
	total := 0
	for _, line := range *fp {
		for _, cell := range line {
			if cell == filledSeat {
				total++
			}
		}
	}
	return total
}

func parseFloorPlan(lines []string) *floorPlan {
	floorPlan := make(floorPlan, len(lines))
	for i, line := range lines {
		floorPlanLine := make([]int, len(line))
		for j, char := range line {
			switch char {
			case '.':
				floorPlanLine[j] = 0
			case 'L':
				floorPlanLine[j] = 1
			case '#':
				floorPlanLine[j] = 2
			}
		}
		floorPlan[i] = floorPlanLine
	}
	return &floorPlan
}
