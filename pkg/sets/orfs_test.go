package sets

import (
	"github.com/dissipative/ribosome/pkg/bioio"
	"github.com/dissipative/ribosome/pkg/sequence"
	"testing"
)

func TestFindORFs(t *testing.T) {
	table, err := sequence.GetCodonTable(1)
	if err != nil {
		t.Fatalf("Unexpected error while getting codon table: %v", err)
	}

	tt := []struct {
		name        string
		minCodons   int
		sequences   []bioio.Record
		expectedLen int
		expectedErr error
	}{
		{
			name:      "TwoRecords",
			minCodons: 1,
			sequences: []bioio.Record{
				{
					ID:       "Seq1",
					Sequence: "AUGGUGAUGGUGUGA",
				},
				{
					ID:       "Seq2",
					Sequence: "AUGCCCAUGGUGUGA",
				},
			},
			expectedLen: 2,
		},
		{
			name:      "CatchError",
			minCodons: 5,
			sequences: []bioio.Record{
				{
					ID:       "Seq1",
					Sequence: "AUGGUGAUGGUGUGA",
				},
				{
					ID:       "Seq2",
					Sequence: "AUGCCCATGGTGUGA",
				},
				{
					ID:       "Seq3",
					Sequence: "AUGGUGAUGGUGUGA",
				},
				{
					ID:       "Seq4",
					Sequence: "AUGGUGAUGGUGUGA",
				},
				{
					ID:       "Seq5",
					Sequence: "AUGGUGAUGGUGUGA",
				},
			},
			expectedErr: sequence.ErrRNAContainsT,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			set := NewRNASet(tc.sequences)
			orfs, err := set.FindORFs(tc.minCodons, &table)
			if tc.expectedErr != nil && err == nil {
				t.Errorf("Expected error, but got no error")
			} else if tc.expectedErr == nil && err != nil {
				t.Errorf("Unexpected error: %v", err)
			}

			if tc.expectedErr == nil && len(orfs.mapped) != tc.expectedLen {
				t.Fatalf("Expected length of mapped ORFs: %d, got: %d", tc.expectedLen, len(orfs.mapped))
			}
		})
	}
}
