package sequence

import (
	"errors"
	"testing"
)

func TestRNASequence_Translate(t *testing.T) {
	tests := []struct {
		name         string
		rna          RNASequence
		tableID      int
		expectedProt ProteinSequence
		expectedErr  error
	}{
		{
			name:         "Basic",
			rna:          "AUGUUUAGU",
			tableID:      1,
			expectedProt: "MFS",
			expectedErr:  nil,
		},
		{
			name:         "TooShort",
			rna:          "AU",
			tableID:      1,
			expectedProt: "",
			expectedErr:  ErrTooShortSequence,
		},
		{
			name:         "InvalidTableID",
			rna:          "AUGUUUAGU",
			tableID:      999,
			expectedProt: "",
			expectedErr:  errors.New("codon table no. 999 not found"),
		},
		{
			name:         "LongSequence",
			rna:          "AUGUUUAGUUAGGGCCAAAUG",
			tableID:      1,
			expectedProt: "MFS*GQM",
			expectedErr:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			prot, err := tt.rna.Translate(tt.tableID)
			if err != nil && err.Error() != tt.expectedErr.Error() {
				t.Errorf("Expected error: %v, got: %v", tt.expectedErr, err)
			}
			if prot != tt.expectedProt {
				t.Errorf("Expected protein sequence: %s, got: %s", tt.expectedProt, prot)
			}
		})
	}
}
