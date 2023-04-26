package sequence

import (
	"errors"
	"strings"
)

type RNASequence string

var ErrTooShortSequence = errors.New("sequence length must be at least 3")

func (r RNASequence) Translate(tableID int) (ProteinSequence, error) {
	seqLength := len(r)

	if seqLength < 3 {
		return "", ErrTooShortSequence
	}

	// Ignore any partial codon at the end of the sequence
	codonsToTranslate := seqLength / 3

	protein := make([]AminoAcid, codonsToTranslate)
	codonTable, err := GetCodonTable(tableID)
	if err != nil {
		return "", err
	}

	for i := 0; i < codonsToTranslate*3; i += 3 {
		codon := strings.ToUpper(string(r[i : i+3]))
		protein[i/3] = codonTable.TranslateCodon(codon)
	}

	return ProteinSequence(protein), nil
}
