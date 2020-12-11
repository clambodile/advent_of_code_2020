package day3

import (
	"container/ring"
	"github.com/clambodile/advent_of_code_2020/util/io"
)

const (
	open = '.'
	tree = '#'
)

func Challenge1() (int, error) {
	dX := 3
	dY := 1

	//read grid from file
	filename := "./day3/d3.txt"
	lines, err := io.ReadLines(filename)
	if err != nil {
		return 0, err
	}
	//turn into array of rings
	var rings []*ring.Ring
	for _, line := range lines {
		rings = append(rings, runeRing(line))
	}

	collisions, err := countCollisions(rings, dX, dY)
	return collisions, err
}

func Challenge2() (int, error) {

	slopes := [][2]int{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}
	//read grid from file
	filename := "./day3/d3.txt"
	lines, err := io.ReadLines(filename)
	if err != nil {
		return 0, err
	}

	collisionsProduct := 1
	for _, slope := range slopes {
		//turn into array of rings
		var rings []*ring.Ring
		for _, line := range lines {
			rings = append(rings, runeRing(line))
		}

		dX := slope[0]
		dY := slope[1]
		collisions, err := countCollisions(rings, dX, dY)
		if err != nil {
			return 0, err
		}
		collisionsProduct *= collisions
	}
	return collisionsProduct, nil
}

func countCollisions(rings []*ring.Ring, dX, dY int) (int, error) {
	treesEncountered := 0
	//traverse array of rings
	for i := 0; i < len(rings); i += dY {
		r := rings[i]
		for j := 0; j < dX*(i/dY); j++ {
			r = r.Next()
		}
		if r.Value == tree {
			treesEncountered++
		}
	}
	return treesEncountered, nil
}

func runeRing(s string) *ring.Ring {
	r := ring.New(len(s))
	for _, rn := range s {
		r.Value = rn
		r = r.Next()
	}
	return r
}
