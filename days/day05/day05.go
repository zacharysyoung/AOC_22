package day03

import (
	"strings"
)

type (
	stack  int
	crate  string
	crates []crate
)
type stacks map[stack]crates

type move struct {
	count    int
	from, to stack
}
type moves []move

func readInput(input string) (stacks_ stacks, moves_ moves) {
	var (
		rawRows  []string
		rowCount = 0

		stackCount = 0

		rawRow  string
		stack_  stack
		crate_  crate
		crates_ crates
	)

	stacks_ = make(stacks)
	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			continue
		}

		if line[0:2] == " 1" {
			rowCount = len(rawRows)

			// " 1   2   3 " → ["1","2","3"] → 3 (stackCount)
			line = strings.TrimSpace(line)
			strStackNums := strings.Split(line, "   ")
			stackCount = len(strStackNums)

			// add crates for each row, from bottom-row-up
			for i := rowCount; i > 0; i-- {
				rawRow = rawRows[i-1]
				for j := 0; j < stackCount; j++ {
					crate_ = crate(rawRow[j*4+1])
					if crate_ == " " {
						continue
					}
					stack_ = stack(j + 1)
					crates_ = stacks_[stack_]
					stacks_[stack_] = append(crates_, crate_)
				}
			}
		}

		if rowCount == 0 {
			rawRows = append(rawRows, line)
			continue
		}
	}

	return
}
