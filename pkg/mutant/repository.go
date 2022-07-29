package mutant

import (
	"crypto/md5"
	"database/sql"
	"fmt"
	"math"
	"strings"

	"github.com/coding-kiko/MutantCheckingApp/pkg/errors"
)

type laboratory struct {
	db *sql.DB
}

func NewLaboratory(d *sql.DB) Laboratory {
	return &laboratory{
		db: d,
	}
}

type Laboratory interface {
	StoreDna(sample DnaSample) error
	GetStats() (Stats, error)
}

var (
	storeDnaQuery    = `INSERT INTO dna(id, is_mutant) VALUES($1, $2)`
	updateStatsQuery = `UPDATE stats SET %s = %s + 1 WHERE id = 'main'`
	getStatsQuery    = `SELECT count_human_dna, count_mutant_dna FROM stats WHERE id = 'main'`
)

func (l *laboratory) GetStats() (Stats, error) {
	var humanCount, mutantCount int64
	var ratio float64

	err := l.db.QueryRow(getStatsQuery).Scan(&humanCount, &mutantCount)
	if err != nil {
		return Stats{}, err
	}

	if humanCount == 0 { // case of division by zero return undefined ratio
		return Stats{
			CountMutantDna: mutantCount,
			CountHumanDna:  humanCount,
			Ratio:          nil,
		}, nil
	}
	precision := math.Pow(10, 2.0)
	ratio = math.Round((float64(mutantCount)/float64(humanCount))*precision) / 2.0
	return Stats{
		CountMutantDna: mutantCount,
		CountHumanDna:  humanCount,
		Ratio:          &ratio,
	}, nil
}

// stores new dna record in the dna table
func (l *laboratory) StoreDna(sample DnaSample) error {
	id := GenerateDnaId(sample.Dna)
	_, err := l.db.Exec(storeDnaQuery, id, sample.Mutant)
	if err != nil {
		return errors.NewDuplicateDna(sample.Mutant)
	}

	err = l.UpdateStats(sample)
	if err != nil {
		return err
	}

	return nil
}

func (l *laboratory) UpdateStats(sample DnaSample) error {
	var column string

	if sample.Mutant {
		column = `count_mutant_dna`
	} else {
		column = `count_human_dna`
	}
	_, err := l.db.Exec(fmt.Sprintf(updateStatsQuery, column, column))
	if err != nil {
		return err
	}
	return nil
}

// generate id from concatenating all of the dna sequences and passsing through md5
func GenerateDnaId(dna []string) string {
	concatenatedDna := strings.Join(dna, "")
	return fmt.Sprintf("%x", md5.Sum([]byte(concatenatedDna)))
}
