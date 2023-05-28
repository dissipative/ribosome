package sequence

import "testing"

func TestGCContent(t *testing.T) {
	testCases := []struct {
		name           string
		sequence       NucleotideSequence
		expectedResult float64
	}{
		{
			name:           "DNA all AT",
			sequence:       DNASequence("ATATATAT"),
			expectedResult: 0.0,
		},
		{
			name:           "DNA all GC",
			sequence:       DNASequence("GCGCGCGC"),
			expectedResult: 1.0,
		},
		{
			name:           "DNA mixed",
			sequence:       DNASequence("ATGCTAGCTAGGCGCG"),
			expectedResult: 0.625,
		},
		{
			name:           "RNA all AU",
			sequence:       RNASequence("AUAUAUAU"),
			expectedResult: 0.0,
		},
		{
			name:           "RNA all GC",
			sequence:       RNASequence("GCGCGCGC"),
			expectedResult: 1.0,
		},
		{
			name:           "RNA mixed",
			sequence:       RNASequence("AUGCUAGC"),
			expectedResult: 0.5,
		},
		{
			name:           "empty sequence",
			sequence:       DNASequence(""),
			expectedResult: 0.0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := GCContent(tc.sequence)
			if result != tc.expectedResult {
				t.Errorf("expected %f, got %f", tc.expectedResult, result)
			}
		})
	}
}
