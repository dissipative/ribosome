package bioio

import (
	"errors"
	"io"
	"reflect"
	"strings"
	"testing"
)

var genbankFileDNA = `LOCUS       TEST123                20 bp    DNA     linear   UNA 01-JAN-1980
DEFINITION  Test sequence.
ACCESSION   TEST123
VERSION     TEST123.1  GI:123456
KEYWORDS    .
SOURCE      synthetic construct
  ORGANISM  synthetic construct
            other sequences; artificial sequences.
FEATURES             Location/Qualifiers
ORIGIN
        1 cacacttgac tccactgtgc gagcgcgtgt gtgtgtgtgt tattttttga gaaaagcaaa
       61 gagagcaagt gagaaccaac accgatgcaa atcattttgt tggctttacg gttagttttt
//

LOCUS       TEST456                20 bp    DNA     linear   UNA 01-JAN-1980
DEFINITION  Test sequence 2.
ACCESSION   TEST456
VERSION     TEST456.1  GI:123456
KEYWORDS    .
SOURCE      synthetic construct
  ORGANISM  synthetic construct
            other sequences; artificial sequences.
FEATURES             Location/Qualifiers
ORIGIN
        1 accagtatga ctgcgctgcg accgtgcagg gtctcgtatt gtcattagaa aatctcacat
       61 attcattttt catttacaaa taaaaatagc cagcacgcag cctacaaatt agcagcgcac
//`

var genbankFileProtein = `LOCUS       TESTPROT              10 aa            linear   UNA 01-JAN-1980
DEFINITION  Test protein sequence.
ACCESSION   TESTPROT
VERSION     TESTPROT.1  GI:123457
KEYWORDS    .
SOURCE      synthetic construct
  ORGANISM  synthetic construct
            other sequences; artificial sequences.
FEATURES             Location/Qualifiers
ORIGIN
        1 MVMGRTPRTR
//`

var genbankInvalid1 = `DEFINITION  Test protein sequence.
ACCESSION   TESTPROT
VERSION     TESTPROT.1  GI:123457
KEYWORDS    .
SOURCE      synthetic construct
  ORGANISM  synthetic construct
            other sequences; artificial sequences.
FEATURES             Location/Qualifiers
ORIGIN
//`

var genbankInvalid2 = `LOCUS       TESTPROT              10 aa            linear   UNA 01-JAN-1980
DEFINITION  Test protein sequence.
ACCESSION   TESTPROT
VERSION     TESTPROT.1  GI:123457
KEYWORDS    .
SOURCE      synthetic construct
  ORGANISM  synthetic construct
            other sequences; artificial sequences.
FEATURES             Location/Qualifiers
ORIGIN
//`

func TestReadGenbank(t *testing.T) {
	var emptySeqs []Sequence

	emptyGenbankFile := ""

	testCases := []struct {
		name          string
		input         io.Reader
		expectedSeq   []Sequence
		expectedError error
	}{
		{
			name:  "dna-sequences",
			input: strings.NewReader(genbankFileDNA),
			expectedSeq: []Sequence{
				{
					ID:          "TEST123",
					Description: "Test sequence.",
					Sequence:    "cacacttgactccactgtgcgagcgcgtgtgtgtgtgtgttattttttgagaaaagcaaagagagcaagtgagaaccaacaccgatgcaaatcattttgttggctttacggttagttttt",
				},
				{
					ID:          "TEST456",
					Description: "Test sequence 2.",
					Sequence:    "accagtatgactgcgctgcgaccgtgcagggtctcgtattgtcattagaaaatctcacatattcatttttcatttacaaataaaaatagccagcacgcagcctacaaattagcagcgcac",
				},
			},
			expectedError: nil,
		},
		{
			name:  "protein-sequence",
			input: strings.NewReader(genbankFileProtein),
			expectedSeq: []Sequence{
				{
					ID:          "TESTPROT",
					Description: "Test protein sequence.",
					Sequence:    "MVMGRTPRTR",
				},
			},
			expectedError: nil,
		},
		{
			name:          "empty-file",
			input:         strings.NewReader(emptyGenbankFile),
			expectedSeq:   emptySeqs,
			expectedError: nil,
		},
		{
			name:          "invalid-genbank-file-1",
			input:         strings.NewReader(genbankInvalid1),
			expectedSeq:   nil,
			expectedError: errors.New("empty LOCUS field"),
		},
		{
			name:          "invalid-genbank-file-2",
			input:         strings.NewReader(genbankInvalid2),
			expectedSeq:   nil,
			expectedError: errors.New("empty ORIGIN field"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			seqs, err := readGenbank(tc.input)
			if tc.expectedError != nil && err.Error() != tc.expectedError.Error() {
				t.Errorf("Expected error '%v', got '%v'", tc.expectedError, err)
			}

			if !reflect.DeepEqual(seqs, tc.expectedSeq) {
				t.Errorf("Expected sequences '%v', got '%v'", tc.expectedSeq, seqs)
			}
		})
	}
}
