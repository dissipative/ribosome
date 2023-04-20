package bioio

import (
	"errors"
	"os"
)

type Format int

const (
	unknown Format = iota
	fasta
	//genbank
)

func ReadFile(filename string, format Format) ([]Sequence, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	switch format {
	case fasta:
		return readFASTA(file)
	default:
		return nil, errors.New("unknown file format")
	}
}

func WriteFile(filename string, format Format, sequences []Sequence) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	switch format {
	case fasta:
		return writeFASTA(file, sequences)
	default:
		return errors.New("unknown file format")
	}
}
