package sequence

import (
	"errors"
	"strings"
)

type RNASequence string

var ErrInvalidCodon = errors.New("invalid codon found in RNA sequence")
var ErrTooShortSequence = errors.New("sequence length must be at least 3")

func (r RNASequence) Translate() (ProteinSequence, error) {
	seqLength := len(r)

	if seqLength < 3 {
		return "", ErrTooShortSequence
	}

	// Ignore any partial codon at the end of the sequence
	codonsToTranslate := seqLength / 3

	protein := make([]AminoAcid, codonsToTranslate)
	codonTable, _ := GetCodonTable(1)

	for i := 0; i < codonsToTranslate*3; i += 3 {
		codon := strings.ToUpper(string(r[i : i+3]))

		aminoAcid, ok := codonTable.Codons[codon]
		if !ok {
			return "", ErrInvalidCodon
		}

		protein[i/3] = aminoAcid
	}

	return ProteinSequence(protein), nil
}
