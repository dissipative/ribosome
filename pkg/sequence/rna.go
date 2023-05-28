package sequence

import (
	"errors"
	"strings"
)

// RNASequence is actually sequence of mRNA
type RNASequence string

var ErrTooShortSequence = errors.New("sequence length must be at least 3")

func (r RNASequence) String() string {
	return string(r)
}

func (r RNASequence) Length() int {
	return len(r)
}

func (r RNASequence) Translate(codonTable *CodonTable) (ProteinSequence, error) {
	seqLength := len(r)

	if seqLength < 3 {
		return "", ErrTooShortSequence
	}

	// Ignore any partial codon at the end of the sequence
	codonsToTranslate := seqLength / 3
	protein := make([]AminoAcid, codonsToTranslate)

	for i := 0; i < codonsToTranslate*3; i += 3 {
		codon := strings.ToUpper(string(r[i : i+3]))
		protein[i/3] = codonTable.TranslateCodon(codon)
	}

	return ProteinSequence(protein), nil
}

type ORF struct {
	Start      int
	End        int
	Codons     int
	Frame      int
	ProteinSeq ProteinSequence
}

func (r RNASequence) FindORFs(minCodons int, codonTable *CodonTable) ([]ORF, error) {
	var orfs []ORF
	if minCodons < 1 {
		minCodons = 1
	}

	seq := string(r)
	for frame := 0; frame < 3; frame++ {
		for i := frame; i <= len(seq)-3; i += 3 {
			_, isStartCodon := codonTable.StartCodons[seq[i:i+3]]
			if isStartCodon {
				for j := i + 3; j <= len(seq)-3; j += 3 {
					codon := seq[j : j+3]
					_, isStopCodon := codonTable.StopCodons[codon]
					if isStopCodon {
						length := (j + 3 - i) / 3
						if length >= minCodons {
							prot, err := r[i : j+3].Translate(codonTable)
							if err != nil {
								return nil, err
							}

							orf := ORF{
								Start:      i,
								End:        j + 3,
								Codons:     length,
								Frame:      frame + 1,
								ProteinSeq: prot,
							}
							orfs = append(orfs, orf)
						}
						break
					}
				}
			}
		}
	}
	return orfs, nil
}
