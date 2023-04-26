package sequence

import (
	"reflect"
	"testing"
)

func TestCodonTable_TranslateCodon(t *testing.T) {
	standardTable := CodonTables[0]

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
			result := standardTable.TranslateCodon(tc.codon)
			if result != tc.expected {
				t.Errorf("Expected %c, got %c", tc.expected, result)
			}
		})
	}
}

func TestCodonTable_ModifyCodonUsage(t *testing.T) {
	codonTable := &CodonTable{
		// Set up a simple codon table
		Codons: map[string]AminoAcid{
			"UUU": 'F', "UUC": 'F',
			"UUA": 'L', "UUG": 'L',
			// More codons
		},
	}

	tests := []struct {
		name          string
		customCodons  map[string]AminoAcid
		expectedError bool
	}{
		{
			name: "Valid modification",
			customCodons: map[string]AminoAcid{
				"UUU": 'L',
			},
			expectedError: false,
		},
		{
			name: "Invalid codon",
			customCodons: map[string]AminoAcid{
				"ZZZ": 'L',
			},
			expectedError: true,
		},
		{
			name: "Invalid amino acid",
			customCodons: map[string]AminoAcid{
				"UUU": 'O',
			},
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := codonTable.ModifyCodonUsage(tt.customCodons)

			if tt.expectedError && err == nil {
				t.Errorf("Expected error, but got no error")
			} else if !tt.expectedError && err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
		})
	}
}

func TestGetCodonTable(t *testing.T) {
	t.Run("edit-obtained-table", func(t *testing.T) {
		standardTable := CodonTables[0]

		got, err := GetCodonTable(standardTable.ID)
		if err != nil {
			t.Errorf("GetCodonTable() error = %v,", err)
			return
		}
		if !reflect.DeepEqual(got, standardTable) {
			t.Errorf("GetCodonTable() got = %v, want %v", got, standardTable)
		}

		err = got.ModifyCodonUsage(map[string]AminoAcid{"AAA": 'T'})
		if err != nil {
			t.Errorf("GetCodonTable() error = %v", err)
			return
		}

		if reflect.DeepEqual(AminoAcid('T'), standardTable.Codons["AAA"]) {
			t.Errorf("user must not be able to change exported variables. AAA codon is changet to %s", string(standardTable.Codons["AAA"]))
		}

	})

}

func TestCodonTable_Copy(t *testing.T) {
	tests := []struct {
		name string
		c    CodonTable
	}{
		{
			name: "Standard",
			c:    CodonTables[0],
		},
		{
			name: "Vertebrate Mitochondrial",
			c:    CodonTables[1],
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			originalCodons := make(map[string]AminoAcid)
			for k, v := range tt.c.Codons {
				originalCodons[k] = v
			}

			copied := tt.c.Copy()

			// Ensure that the original and copied tables are not the same
			if &tt.c == &copied {
				t.Errorf("Copy() did not create a new instance of CodonTable")
			}

			// Ensure that the original and copied Codons maps are not the same
			if &tt.c.Codons == &copied.Codons {
				t.Errorf("Copy() did not create a new instance of Codons map")
			}

			// Ensure that the content of the original and copied Codons maps are the same
			if !reflect.DeepEqual(tt.c.Codons, copied.Codons) {
				t.Errorf("Copy() produced a different Codons map")
			}

			// Modify the copied Codons map
			for k := range copied.Codons {
				copied.Codons[k] = 'X'
			}

			// Ensure that the original Codons map is not affected
			if reflect.DeepEqual(tt.c.Codons, copied.Codons) {
				t.Errorf("Copy() did not create an independent copy of Codons map")
			}

			// Ensure that the original Codons map is not changed
			if !reflect.DeepEqual(tt.c.Codons, originalCodons) {
				t.Errorf("Original Codons map was modified")
			}
		})
	}
}
