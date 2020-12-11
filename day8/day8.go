package day8

import (
	"errors"
	"github.com/clambodile/advent_of_code_2020/util/io"
	"strconv"
	"strings"
)

type Instruction struct {
	Op  string
	Arg int
}

var Accumulator = 0
var SeenIndices = map[int]bool{}

func Challenge1() (int, error) {
	filename := "./day8/d8.txt"
	lines, err := io.ReadLines(filename)
	if err != nil {
		return 0, err
	}
	instrs := parseInstructions(lines)
	i := 0
	for !SeenIndices[i] {
		SeenIndices[i] = true
		i, err = executeInstruction(instrs, i)
		if err != nil {
			return 0, err
		}
	}
	return Accumulator, nil
}

func Challenge2() (int, error) {
	filename := "./day8/d8.txt"
	lines, err := io.ReadLines(filename)
	if err != nil {
		return 0, err
	}
	for i := range lines {
		Accumulator = 0
		instrs := parseInstructions(lines)
		if instrs[i].Op == "jmp" {
			instrs[i].Op = "nop"
		} else if instrs[i].Op == "nop" {
			instrs[i].Op = "jmp"
		} else {
			continue
		}
		terminates := untilLoop(instrs)
		if terminates {
			return Accumulator, nil
		}
	}

	return 0, errors.New("solution not found")
}

func parseInstructions(lines []string) []Instruction {
	instrs := make([]Instruction, len(lines))
	for i, line := range lines {
		instrs[i] = parseInstruction(line)
	}
	return instrs
}
func untilLoop(instrs []Instruction) bool {
	seen := map[int]bool{}
	var i int
	for i = 0; (i > -1) && !seen[i]; {
		seen[i] = true
		i, _ = executeInstruction(instrs, i)
	}
	if i == -1 {
		return true
	}
	return false
}
func executeInstruction(instrs []Instruction, i int) (int, error) {
	if i == len(instrs) {
		return -1, nil
	}
	instr := instrs[i]
	switch instr.Op {
	case "acc":
		Accumulator += instr.Arg
		return i + 1, nil
	case "jmp":
		return i + instr.Arg, nil
	case "nop":
		return i + 1, nil
	}
	return -9999, errors.New("no valid instruction recognized")
}

func parseInstruction(instr string) Instruction {
	opArg := strings.Split(instr, " ")
	op := opArg[0]
	arg, _ := strconv.Atoi(opArg[1])
	return Instruction{
		Op:  op,
		Arg: arg,
	}
}
