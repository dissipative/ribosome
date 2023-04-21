package sequence

import (
	"testing"
)

func TestCodonTable_TranslateCodon(t *testing.T) {
	ct := StandardCodonTable

	testCases := []struct {
		name     string
		codon    string
		expected AminoAcid
	}{
		{"No codon", "", 'X'},
		{"Unambiguous codon", "AUG", 'M'},
		{"Ambiguous codon with single amino acid", "CUY", 'L'},
		{"Ambiguous stop codon", "URA", '*'},
		{"Ambiguous codon with multiple amino acids", "AAM", 'X'},
		{"Invalid codon", "AU-", 'X'},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := ct.TranslateCodon(tc.codon)
			if result != tc.expected {
				t.Errorf("Expected %c, got %c", tc.expected, result)
			}
		})
	}
}
