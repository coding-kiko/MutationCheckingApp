package mutant

import (
	"strings"
)

type DnaSample struct {
	Size   int
	Dna    []string
	Mutant bool
}

func NewDnaSample(dna []string) *DnaSample {
	return &DnaSample{
		Size: len(dna),
		Dna:  dna,
	}
}

// checks for possible mutation in dna sequence and set _Mutant_ attribute of DnaSample
func (d *DnaSample) CheckMutation() {
	var sequences = 0
	// go through every nitrogen base and check for a sequence in any of the 4 possible directions
	for i := 0; i < d.Size; i++ {
		for j := 0; j < d.Size; j++ {
			if sequences > 1 {
				d.Mutant = sequences > 1
				return
			}
			if string(d.Dna[i][j]) == "X" {
				continue
			}
			if d.checkDown(i, j) {
				sequences++
				continue
			}
			if d.checkRight(i, j) {
				sequences++
				continue
			}
			if d.checkDiagonalLeft(i, j) {
				sequences++
				continue
			}
			if d.checkDiagonalRight(i, j) {
				sequences++
				continue
			}
		}
	}
	d.Mutant = sequences > 1
}

func (d *DnaSample) checkDown(i, j int) bool {
	// check if a 4 sized sequence even fits
	if i > d.Size-4 {
		return false
	}

	compare := string(d.Dna[i][j])
	for n := i + 1; n < i+4; n++ {
		if string(d.Dna[n][j]) != compare {
			return false
		}
	}

	// invalidate used nitrogen base (avoid sharing between sequences)
	for n := i; n < i+4; n++ {
		d.Dna[n] = d.Dna[n][:j] + "X" + d.Dna[n][j+1:]
	}

	return true
}

func (d *DnaSample) checkRight(i, j int) bool {
	// check if a 4 sized sequence even fits
	if j > d.Size-4 {
		return false
	}

	compare := string(d.Dna[i][j])
	for n := j + 1; n < j+4; n++ {
		if string(d.Dna[i][n]) != compare {
			return false
		}
	}

	// invalidate used nitrogen base (avoid sharing between sequences)
	d.Dna[i] = d.Dna[i][:j] + strings.Repeat("X", 4) + d.Dna[i][j+4:]

	return true
}

func (d *DnaSample) checkDiagonalRight(i, j int) bool {
	// check if a 4 sized sequence even fits
	if i > d.Size-4 || j > d.Size-4 {
		return false
	}

	compare := string(d.Dna[i][j])
	for n, m := i+1, j+1; n < j+4; n, m = n+1, m+1 {
		if string(d.Dna[n][m]) != compare {
			return false
		}
	}

	// invalidate used nitrogen base (avoid sharing between sequences)
	for n, m := i, j; n < i+4; n, m = n+1, m+1 {
		d.Dna[n] = d.Dna[n][:m] + "X" + d.Dna[n][m+1:]
	}

	return true
}

func (d *DnaSample) checkDiagonalLeft(i, j int) bool {
	// check if a 4 sized sequence even fits
	if i > d.Size-4 || j < 3 {
		return false
	}

	compare := string(d.Dna[i][j])
	for n, m := i+1, j-1; n < i+4; n, m = n+1, m-1 {
		if string(d.Dna[n][m]) != compare {
			return false
		}
	}

	// invalidate used nitrogen base (avoid sharing between sequences)
	for n, m := i+1, j-1; n < i+4; n, m = n+1, m-1 {
		d.Dna[n] = d.Dna[n][:m] + "X" + d.Dna[n][m+1:]
	}

	return true
}
