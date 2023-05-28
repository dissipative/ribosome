package sequence

import (
	"testing"
)

func TestRNASequence_Translate(t *testing.T) {
	standardTable, _ := GetCodonTable(1)

	tests := []struct {
		name         string
		rna          RNASequence
		table        *CodonTable
		expectedProt ProteinSequence
		expectedErr  error
	}{
		{
			name:         "Basic",
			rna:          "AUGUUUAGU",
			table:        &standardTable,
			expectedProt: "MFS",
			expectedErr:  nil,
		},
		{
			name:         "TooShort",
			rna:          "AU",
			table:        &standardTable,
			expectedProt: "",
			expectedErr:  ErrTooShortSequence,
		},
		{
			name:         "LongSequence",
			rna:          "AUGUUUAGUUAGGGCCAAAUG",
			table:        &standardTable,
			expectedProt: "MFS*GQM",
			expectedErr:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			prot, err := tt.rna.Translate(tt.table)
			if err != nil && err.Error() != tt.expectedErr.Error() {
				t.Errorf("Expected error: %v, got: %v", tt.expectedErr, err)
			}
			if prot != tt.expectedProt {
				t.Errorf("Expected protein sequence: %s, got: %s", tt.expectedProt, prot)
			}
		})
	}
}

func TestRNASequence_FindORFs(t *testing.T) {
	standartTable, _ := GetCodonTable(1)

	testCases := []struct {
		name      string
		r         RNASequence
		minCodons int
		expected  []ORF
	}{
		{
			name:      "ORFWithMinimumCodons",
			r:         "AUGCCCUAA",
			minCodons: 1,
			expected: []ORF{
				{
					Start:      0,
					End:        9,
					Codons:     3,
					Frame:      1,
					ProteinSeq: "MP*",
				},
			},
		},
		{
			name:      "NoORF",
			r:         "AUGCCCG",
			minCodons: 2,
			expected:  []ORF{},
		},
		{
			name:      "ORFShorterThanMinimumCodons",
			r:         "AUGCCCUAA",
			minCodons: 4,
			expected:  []ORF{},
		},
		{
			name:      "MultipleORFs",
			r:         "AUGCCCUAAAUGGGGUAA",
			minCodons: 1,
			expected: []ORF{
				{
					Start:      0,
					End:        9,
					Codons:     3,
					Frame:      1,
					ProteinSeq: "MP*",
				},
				{
					Start:      9,
					End:        18,
					Codons:     3,
					Frame:      1,
					ProteinSeq: "MG*",
				},
			},
		},
		{
			name:      "MultipleORFsInDifferentFrames",
			r:         "AUGCCCUAAAAUGGGGUAA",
			minCodons: 1,
			expected: []ORF{
				{
					Start:      0,
					End:        9,
					Codons:     3,
					Frame:      1,
					ProteinSeq: "MP*",
				},
				{
					Start:      10,
					End:        19,
					Codons:     3,
					Frame:      2,
					ProteinSeq: "MG*",
				},
			},
		},
		{
			name:      "ORFInSecondFrame",
			r:         "UAUGCCCUAA",
			minCodons: 3,
			expected: []ORF{
				{
					Start:      1,
					End:        10,
					Codons:     3,
					Frame:      2,
					ProteinSeq: "MP*",
				},
			},
		},
		{
			name:      "ORF in third frame",
			r:         "UUAUGCCCUAA",
			minCodons: 3,
			expected: []ORF{
				{
					Start:      2,
					End:        11,
					Codons:     3,
					Frame:      3,
					ProteinSeq: "MP*",
				},
			},
		},
		{
			name:      "No start codon",
			r:         "CCCUAA",
			minCodons: 1,
			expected:  []ORF{},
		},
		{
			name:      "No stop codon",
			r:         "AUGCCC",
			minCodons: 1,
			expected:  []ORF{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			orfs, err := tc.r.FindORFs(tc.minCodons, &standartTable)
			if err != nil {
				t.Fatalf("FindORFs failed: %v", err)
			}

			if len(orfs) != len(tc.expected) {
				t.Fatalf("Expected %v ORFs, got %v", len(tc.expected), len(orfs))
			}

			for i, orf := range orfs {
				if orf.Start != tc.expected[i].Start ||
					orf.End != tc.expected[i].End ||
					orf.Codons != tc.expected[i].Codons ||
					orf.Frame != tc.expected[i].Frame ||
					orf.ProteinSeq != tc.expected[i].ProteinSeq {
					t.Errorf("Unexpected ORF at index %v. Expected %v, got %v", i, tc.expected[i], orf)
				}
			}
		})
	}
}
