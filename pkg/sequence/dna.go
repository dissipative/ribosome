package sequence

import (
	"errors"
	"strings"
)

type DNASequence string

var complementMapForDNA = map[Nucleotide]Nucleotide{
	'A': 'T', 'a': 'T',
	'T': 'A', 't': 'A',
	'C': 'G', 'c': 'G',
	'G': 'C', 'g': 'C',
	'R': 'Y', 'r': 'Y',
	'Y': 'R', 'y': 'R',
	'S': 'S', 's': 'S',
	'W': 'W', 'w': 'W',
	'K': 'M', 'k': 'M',
	'M': 'K', 'm': 'K',
	'B': 'V', 'b': 'V',
	'D': 'H', 'd': 'H',
	'H': 'D', 'h': 'D',
	'V': 'B', 'v': 'B',
	'N': 'N', 'n': 'N',
	'-': '-',
}

func NewDNASequence(input string) (DNASequence, error) {
	upper := strings.ToUpper(input)
	if strings.Contains(upper, "U") {
		return "", errors.New("string contains U base and is not valid DNA sequence")
	}

	return DNASequence(upper), nil
}

func (d DNASequence) String() string {
	return string(d)
}

func (d DNASequence) Length() int {
	return len(d)
}

func (d DNASequence) Reverse() DNASequence {
	seqLen := len(d)
	reversed := make([]Nucleotide, seqLen)

	for i, base := range []Nucleotide(d) {
		reversed[seqLen-i-1] = base
	}

	return DNASequence(reversed)
}

func (d DNASequence) Complement() DNASequence {
	complement := make([]Nucleotide, len(d))

	for i, base := range []Nucleotide(d) {
		complementaryBase, ok := complementMapForDNA[base]
		if !ok {
			complement[i] = '-'
		} else {
			complement[i] = complementaryBase
		}
	}

	return DNASequence(complement)
}

func (d DNASequence) ReverseComplement() DNASequence {
	return d.Reverse().Complement()
}

func (d DNASequence) Transcribe() RNASequence {
	transcriptionMap := complementMapForDNA
	transcriptionMap['A'] = 'U'
	transcriptionMap['a'] = 'U'

	rna := make([]Nucleotide, len(d))

	for i, base := range []Nucleotide(d) {
		rnaBase, ok := transcriptionMap[base]
		if !ok {
			rna[i] = '-'
		} else {
			rna[i] = rnaBase
		}
	}

	return RNASequence(rna)
}
