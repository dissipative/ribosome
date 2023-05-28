package sequence

import "strings"

type NucleotideSequence interface {
	String() string
	Length() int
}

func GCContent[S NucleotideSequence](seq S) float64 {
	if seq.Length() == 0 {
		return 0.0
	}

	gcCount := 0
	for _, nucleotide := range strings.ToUpper(seq.String()) {
		if nucleotide == 'G' || nucleotide == 'C' {
			gcCount++
		}
	}

	return float64(gcCount) / float64(seq.Length())
}
