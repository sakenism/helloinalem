package main

import (
	"github.com/01-edu/z01"
	student "github.com/alem-01/quad-checker/student"

	"github.com/alem-01/quad-checker/lib"
)

func quadE(x, y int) {
	if x < 0 || y < 0 {
		return
	}
	for i := 1; i <= y; i++ {
		for j := 1; j <= x; j++ {
			if (i == 1 && j == 1) || (i == y && j == x && i > 1 && j > 1) {
				z01.PrintRune('A')
			} else if (i == 1 && j == x) || (i == y && j == 1) {
				z01.PrintRune('C')
			} else if i == 1 || i == y || j == 1 || j == x {
				z01.PrintRune('B')
			} else {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}
}

func main() {
	// testing examples of subjects
	table := []int{
		5, 3,
		5, 1,
		1, 1,
		1, 5,
	}

	// testing special cases and one valid random case.
	table = append(table,
		0, 0,
		-1, 6,
		6, -1,
		lib.RandIntBetween(1, 20), lib.RandIntBetween(1, 20),
	)

	// Tests all possibilities including 0 0, -x y, x -y
	for i := 0; i < len(table); i += 2 {
		if i != len(table)-1 {
			lib.Challenge("QuadE", quadE, student.QuadE, table[i], table[i+1])
		}
	}
}
