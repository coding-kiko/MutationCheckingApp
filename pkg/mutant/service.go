package mutant

import (
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
	sample.CheckMutation()

	err := m.lab.StoreDna(*sample)
	if err != nil {
		return err
	}
	if !sample.Mutant {
		return errors.NewNotMutant()
	}
	return nil
}
