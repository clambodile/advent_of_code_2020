package day12

import (
	"github.com/clambodile/advent_of_code_2020/util/io"
	"math"
	"regexp"
	"strconv"
)

type move struct {
	Dir int
	Amt int
}

const (
	north   = 0
	east    = 1
	south   = 2
	west    = 3
	left    = 4
	right   = 5
	forward = 6
)

type Coord struct {
	X int
	Y int
}

type part1State struct {
	Facing int
	Coord
}

type part2State struct {
	Waypoint Coord
	Coord
}

type vector [2]int8

func Challenge1() (int, error) {
	lines, err := io.ReadLines("./day12/d12.txt")
	if err != nil {
		return 0, err
	}
	moves := parseMoves(lines)
	endState := executeMoves1(moves)
	return int(math.Abs(float64(endState.Coord.X)) + math.Abs(float64(endState.Y))), nil
}

func Challenge2() (int, error) {
	lines, err := io.ReadLines("./day12/d12.txt")
	if err != nil {
		return 0, err
	}
	moves := parseMoves(lines)
	endState := executeMoves2(moves)
	return int(math.Abs(float64(endState.Coord.X)) + math.Abs(float64(endState.Y))), nil
}

func executeMoves1(moves []move) part1State {
	st := part1State{
		Facing: east,
		Coord: Coord{0,0},
	}
	for _, mv := range moves {
		st = executeMove1(st, mv)
	}
	return st
}

func executeMoves2(moves []move) part2State {
	st := part2State{
		Waypoint: Coord{10, 1},
		Coord:    Coord{0, 0},
	}
	for _, mv := range moves {
		st = executeMove2(st, mv)
	}
	return st
}

func executeMove1(st part1State, mv move) part1State {
	switch mv.Dir {
	case north:
		st.Y += mv.Amt
	case south:
		st.Y -= mv.Amt
	case east:
		st.X += mv.Amt
	case west:
		st.X -= mv.Amt
	case left:
		times := mv.Amt / 90
		for i := 0; i < times; i++ {
			st.Facing = quarterCounterClockwise(st.Facing)
		}
	case right:
		times := mv.Amt / 90
		for i := 0; i < times; i++ {
			st.Facing = quarterClockwise(st.Facing)
		}
	case forward:
		st = executeMove1(st, move{Dir: st.Facing, Amt: mv.Amt})
	}
	return st
}

func executeMove2(st part2State, mv move) part2State {
	switch mv.Dir {
	case north:
		st.Waypoint.Y += mv.Amt
	case south:
		st.Waypoint.Y -= mv.Amt
	case east:
		st.Waypoint.X += mv.Amt
	case west:
		st.Waypoint.X -= mv.Amt
	case left:
		times := mv.Amt / 90
		for i := 0; i < times; i++ {
			quarterCounterClockwise2(&st)
		}
	case right:
		times := mv.Amt / 90
		for i := 0; i < times; i++ {
			quarterClockwise2(&st)
		}
	case forward:
		for i := 0; i < mv.Amt; i++ {
			st.X += st.Waypoint.X
			st.Y += st.Waypoint.Y
		}
	}
	return st
}
func quarterClockwise(d int) int {
	return (d + 1) % 4
}

func quarterClockwise2(st *part2State) *part2State {
	prevX := st.Waypoint.X
	prevY := st.Waypoint.Y
	st.Waypoint.X = prevY
	st.Waypoint.Y = -prevX
	return st
}
func quarterCounterClockwise(d int) int {
	return (d + 3) % 4
}

func quarterCounterClockwise2(st *part2State) *part2State {
	prevX := st.Waypoint.X
	prevY := st.Waypoint.Y
	st.Waypoint.X = -prevY
	st.Waypoint.Y = prevX
	return st
}
var vecs = [4]vector{
	{1, 0},
	{0, 1},
	{-1, 0},
	{0, -1},
}

func vec(d int) vector {
	return vecs[d]
}

func parseMoves(lines []string) []move {
	moves := make([]move, 0)
	charToDir := map[string]int{
		"N": north,
		"S": south,
		"E": east,
		"W": west,
		"L": left,
		"R": right,
		"F": forward,
	}
	moveRegex := regexp.MustCompile(`([NSEWLRF])(\d+)`)
	for _, line := range lines {
		matches := moveRegex.FindAllStringSubmatch(line, -1)
		dirChar := matches[0][1]
		amountStr := matches[0][2]
		dir := charToDir[dirChar]
		amount, _ := strconv.Atoi(amountStr)
		mv := move{
			Dir: dir,
			Amt: amount,
		}
		moves = append(moves, mv)
	}
	return moves
}
