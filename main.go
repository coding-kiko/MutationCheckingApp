package main

import (
	"fmt"
	"strings"
)

type adn struct {
	Size     int
	Sequence []string
}

func main() {
	sample := adn{}
	//sample.Sequence = []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"} // 3
	//sample.Sequence = []string{"ATGCGA", "CAGTGC", "TTATGT", "AGTAGG", "CTCCTA", "TCACTG"} // 3
	sample.Sequence = []string{"TTGCGA", "CAGTGC", "TTATAT", "AGTAGG", "CGCCTA", "TCACTG"} // 0
	sample.Size = len(sample.Sequence)
	fmt.Println(sample.isMutant())
}

func (a *adn) isMutant() bool {
	sequences := 0

	for i := 0; i < a.Size; i++ {
		for j := 0; j < a.Size; j++ {
			if string(a.Sequence[i][j]) == "X" {
				continue
			}
			if a.checkDown(i, j) {
				sequences++
				continue
			}
			if a.checkRight(i, j) {
				sequences++
				continue
			}
			if a.checkDiagonalLeft(i, j) {
				sequences++
				continue
			}
			if a.checkDiagonalRight(i, j) {
				sequences++
				continue
			}
		}
	}

	fmt.Println(a.Sequence)
	fmt.Println(sequences)
	return sequences > 1
}

func (a *adn) checkDown(i, j int) bool {
	// check if a 4 sized sequence even fits
	if i > a.Size-4 {
		return false
	}

	compare := string(a.Sequence[i][j])
	for n := i + 1; n < i+4; n++ {
		if string(a.Sequence[n][j]) != compare {
			return false
		}
	}

	// invalidate used nitrogen base (avoid sharing between sequences)
	for n := i; n < i+4; n++ {
		a.Sequence[n] = a.Sequence[n][:j] + "X" + a.Sequence[n][j+1:]
	}

	return true
}

func (a *adn) checkRight(i, j int) bool {
	// check if a 4 sized sequence even fits
	if j > a.Size-4 {
		return false
	}

	compare := string(a.Sequence[i][j])
	for n := j + 1; n < j+4; n++ {
		if string(a.Sequence[i][n]) != compare {
			return false
		}
	}

	// invalidate used nitrogen base (avoid sharing between sequences)
	a.Sequence[i] = strings.Repeat("X", 4) + a.Sequence[i][j+4:]

	return true
}

func (a *adn) checkDiagonalRight(i, j int) bool {
	// check if a 4 sized sequence even fits
	if i > a.Size-4 || j > a.Size-4 {
		return false
	}

	compare := string(a.Sequence[i][j])
	for n, m := i+1, j+1; n < j+4; n, m = n+1, m+1 {
		if string(a.Sequence[n][m]) != compare {
			return false
		}
	}

	// invalidate used nitrogen base (avoid sharing between sequences)
	for n, m := i, j; n < i+4; n, m = n+1, m+1 {
		a.Sequence[n] = a.Sequence[n][:m] + "X" + a.Sequence[n][m+1:]
	}

	return true
}

func (a *adn) checkDiagonalLeft(i, j int) bool {
	// check if a 4 sized sequence even fits
	if i > a.Size-4 || j < 3 {
		return false
	}

	compare := string(a.Sequence[i][j])
	for n, m := i+1, j-1; n < i+4; n, m = n+1, m-1 {
		if string(a.Sequence[n][m]) != compare {
			return false
		}
	}

	// invalidate used nitrogen base (avoid sharing between sequences)
	for n, m := i+1, j-1; n < i+4; n, m = n+1, m-1 {
		a.Sequence[n] = a.Sequence[n][:m] + "X" + a.Sequence[n][m+1:]
	}

	return true
}
