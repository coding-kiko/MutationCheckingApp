package mutant

import (
	"fmt"

	"github.com/coding-kiko/MutantCheckingApp/pkg/errors"
)

type mutantService struct {
	lab Laboratory
}

func NewMutantService(l Laboratory) MutantService {
	return &mutantService{
		lab: l,
	}
}

type MutantService interface {
	AnalyzeDna(req DnaRequest) error
	GetStats() (Stats, error)
}

func (m *mutantService) GetStats() (Stats, error) {
	stats, err := m.lab.GetStats()
	if err != nil {
		return Stats{}, err
	}
	return stats, nil
}

func (m *mutantService) AnalyzeDna(req DnaRequest) error {
	sample := NewDnaSample(req.Dna)

	if sample.Size < 4 {
		return errors.NewBadRequest("DNA sample not long enough")
	}
	err := VerifySquareMatrix(sample.Dna)
	if err != nil {
		return err
	}

	sample.CheckMutation()

	err = m.lab.StoreDna(*sample)
	if err != nil {
		return err
	}
	if !sample.Mutant {
		return errors.NewNotMutant()
	}
	return nil
}

// Received and array of strings and returns an error if it is not NxN
func VerifySquareMatrix(matrix []string) error {
	rows := len(matrix)
	for i, row := range matrix {
		if len(row) != rows {
			return errors.NewBadRequest(fmt.Sprintf("DNA not represented by a square matrix: element %d", i+1))
		}
	}
	return nil
}
