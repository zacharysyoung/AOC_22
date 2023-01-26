package day03

import (
	"fmt"
	"strings"
)

type (
	stack  int
	crates []string
	stacks map[stack]crates
)

type move struct {
	count    int
	from, to stack
}
type moves []move

func readInput(input string) (stacks_ stacks, moves_ moves) {
	if input[0:1] == "\n" {
		input = input[1:]
	}

	// Split stacks and moves at double-newline divide
	parts := strings.Split(input, "\n\n")
	stacksStr := parts[0]
	movesStr := parts[1]

	// Split lines of stacked crates
	stacksLines := strings.Split(stacksStr, "\n")
	rawRows := stacksLines[:len(stacksLines)-1]

	// " 1   2   3 " → ["1","2","3"] → stackCount = 3
	stackNumsStr := stacksLines[len(stacksLines)-1]
	stackNumsStr = strings.TrimSpace(stackNumsStr)
	strStackNums := strings.Split(stackNumsStr, "   ")
	stackCount := len(strStackNums)

	// Start "stacking crates" from the bottom row, up
	stacks_ = make(stacks)
	var (
		rawRow  string // a row of fixed-width fields (crates), like `[Z] [M] [P]` or `    [D]    `
		crate_  string // an individual, parsed field, like `Z` or `D`
		stack_  stack  // positional "ID" of a stack
		crates_ crates // an individual stack of crates, in stacks_
	)

	for i := len(rawRows); i > 0; i-- {
		rawRow = rawRows[i-1]
		for j := 0; j < stackCount; j++ {
			crate_ = string(rawRow[j*4+1])
			if crate_ == " " {
				continue
			}
			stack_ = stack(j + 1)
			crates_ = stacks_[stack_]
			stacks_[stack_] = append(crates_, crate_)
		}
	}

	movesLines := strings.Split(strings.TrimSpace(movesStr), "\n")
	moves_ = make(moves, len(movesLines))
	var (
		count    int
		from, to stack
	)
	for i, moveLine := range movesLines {
		fmt.Sscanf(moveLine, "move %d from %d to %d", &count, &from, &to)
		moves_[i] = move{count, from, to}
	}

	return
}
