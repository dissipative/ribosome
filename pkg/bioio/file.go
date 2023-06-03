package bioio

import (
	"errors"
	"os"
)

type Format int

const (
	Fasta Format = iota
	Genbank
)

func ReadFile(filename string, format Format) ([]Record, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	switch format {
	case Fasta:
		return readFASTA(file)
	case Genbank:
		return readGenbank(file)
	default:
		return nil, errors.New("unknown file format")
	}
}

func WriteFile(filename string, format Format, sequences []Record) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	switch format {
	case Fasta:
		return writeFASTA(file, sequences)
	case Genbank:
		return writeGenbank(file, sequences)
	default:
		return errors.New("unknown file format")
	}
}
