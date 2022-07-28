package mutant

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckMutation(t *testing.T) {
	testCases := []struct {
		name          string
		dna           []string
		checkResponse func(t *testing.T, d DnaSample)
	}{
		{
			name: "mutant - size 6",
			dna:  []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "CCCCTA", "TCACTG"},
			checkResponse: func(t *testing.T, d DnaSample) {
				assert.True(t, d.Mutant)
			},
		},
		{
			name: "mutant - size 6",
			dna:  []string{"ATGCGA", "CAGTGC", "TTATGT", "AGTAGG", "CTCCTA", "TCACTG"},
			checkResponse: func(t *testing.T, d DnaSample) {
				assert.True(t, d.Mutant)
			},
		},
		{
			name: "non mutant - size 6",
			dna:  []string{"TTGCGA", "CAGTGC", "TTATAT", "AGTAGG", "CGCCTA", "TCACTG"},
			checkResponse: func(t *testing.T, d DnaSample) {
				assert.False(t, d.Mutant)
			},
		},
		{
			name: "mutant - size 4",
			dna:  []string{"TTGC", "CCCC", "AAAA", "TTTT"},
			checkResponse: func(t *testing.T, d DnaSample) {
				assert.True(t, d.Mutant)
			},
		},
		{
			name: "non mutant - size 7",
			dna:  []string{"GTGCGAA", "CAGTGCG", "TTATGTA", "AGAACGT", "CCCCTAT", "TCACTGC", "TCACTGC"},
			checkResponse: func(t *testing.T, d DnaSample) {
				assert.False(t, d.Mutant)
			},
		},
		{
			name: "mutant - size 6",
			dna:  []string{"ATGCGA", "CAGTGC", "TTATGT", "AGAAGG", "ATCCCC", "TCACTG"},
			checkResponse: func(t *testing.T, d DnaSample) {
				assert.True(t, d.Mutant)
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			d := DnaSample{
				Size: len(tc.dna),
				Dna:  tc.dna,
			}
			d.CheckMutation()
			tc.checkResponse(t, d)
		})
	}
}
