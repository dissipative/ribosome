package ncbi

import (
	"bufio"
	"strings"
	"testing"

	"github.com/dissipative/ribosome/pkg/sequence"
)

func TestParsePRTCodonTables(t *testing.T) {
	testInput := `--*************************************************************************

Genetic-code-table ::= {
{
  name "Standard" ,
  name "SGC0" ,
  id 1 ,
  ncbieaa  "FFLLSSSSYY**CC*WLLLLPPPPHHQQRRRRIIIMTTTTNNKKSSRRVVVVAAAADDEEGGGG",
  sncbieaa "---M------**--*----M---------------M----------------------------"
  -- Base1  TTTTTTTTTTTTTTTTCCCCCCCCCCCCCCCCAAAAAAAAAAAAAAAAGGGGGGGGGGGGGGGG
  -- Base2  TTTTCCCCAAAAGGGGTTTTCCCCAAAAGGGGTTTTCCCCAAAAGGGGTTTTCCCCAAAAGGGG
  -- Base3  TCAGTCAGTCAGTCAGTCAGTCAGTCAGTCAGTCAGTCAGTCAGTCAGTCAGTCAGTCAGTCAG
},
{
  name "Vertebrate Mitochondrial" ,
  name "SGC1" ,
  id 2 ,
  ncbieaa  "FFLLSSSSYY**CCWWLLLLPPPPHHQQRRRRIIMMTTTTNNKKSS**VVVVAAAADDEEGGGG",
  sncbieaa "----------**--------------------MMMM----------**---M------------"
  -- Base1  TTTTTTTTTTTTTTTTCCCCCCCCCCCCCCCCAAAAAAAAAAAAAAAAGGGGGGGGGGGGGGGG
  -- Base2  TTTTCCCCAAAAGGGGTTTTCCCCAAAAGGGGTTTTCCCCAAAAGGGGTTTTCCCCAAAAGGGG
  -- Base3  TCAGTCAGTCAGTCAGTCAGTCAGTCAGTCAGTCAGTCAGTCAGTCAGTCAGTCAGTCAGTCAG
},
{
  name "Mold Mitochondrial; Protozoan Mitochondrial; Coelenterate
 Mitochondrial; Mycoplasma; Spiroplasma" ,
  name "SGC3" ,
  id 4 ,
  ncbieaa  "FFLLSSSSYY**CCWWLLLLPPPPHHQQRRRRIIIMTTTTNNKKSSRRVVVVAAAADDEEGGGG",
  sncbieaa "--MM------**-------M------------MMMM---------------M------------"
  -- Base1  TTTTTTTTTTTTTTTTCCCCCCCCCCCCCCCCAAAAAAAAAAAAAAAAGGGGGGGGGGGGGGGG
  -- Base2  TTTTCCCCAAAAGGGGTTTTCCCCAAAAGGGGTTTTCCCCAAAAGGGGTTTTCCCCAAAAGGGG
  -- Base3  TCAGTCAGTCAGTCAGTCAGTCAGTCAGTCAGTCAGTCAGTCAGTCAGTCAGTCAGTCAGTCAG
},
 {
  name "Euplotid Nuclear" ,
  name "SGC9" ,
  id 10 ,
  ncbieaa  "FFLLSSSSYY**CCCWLLLLPPPPHHQQRRRRIIIMTTTTNNKKSSRRVVVVAAAADDEEGGGG",
  sncbieaa "----------**-----------------------M----------------------------"
  -- Base1  TTTTTTTTTTTTTTTTCCCCCCCCCCCCCCCCAAAAAAAAAAAAAAAAGGGGGGGGGGGGGGGG
  -- Base2  TTTTCCCCAAAAGGGGTTTTCCCCAAAAGGGGTTTTCCCCAAAAGGGGTTTTCCCCAAAAGGGG
  -- Base3  TCAGTCAGTCAGTCAGTCAGTCAGTCAGTCAGTCAGTCAGTCAGTCAGTCAGTCAGTCAGTCAG
 },
 {
  name "Bacterial, Archaeal and Plant Plastid" ,
  id 11 ,
  ncbieaa  "FFLLSSSSYY**CC*WLLLLPPPPHHQQRRRRIIIMTTTTNNKKSSRRVVVVAAAADDEEGGGG",
  sncbieaa "---M------**--*----M------------MMMM---------------M------------"
  -- Base1  TTTTTTTTTTTTTTTTCCCCCCCCCCCCCCCCAAAAAAAAAAAAAAAAGGGGGGGGGGGGGGGG
  -- Base2  TTTTCCCCAAAAGGGGTTTTCCCCAAAAGGGGTTTTCCCCAAAAGGGGTTTTCCCCAAAAGGGG
  -- Base3  TCAGTCAGTCAGTCAGTCAGTCAGTCAGTCAGTCAGTCAGTCAGTCAGTCAGTCAGTCAGTCAG
 }
}`

	standardTable, err := sequence.GetCodonTable(1)
	if err != nil {
		t.Error(err)
	}
	vmTable, err := sequence.GetCodonTable(2)
	if err != nil {
		t.Error(err)
	}
	moldMtTable, err := sequence.GetCodonTable(4)
	if err != nil {
		t.Error(err)
	}
	euplotidTable, err := sequence.GetCodonTable(10)
	if err != nil {
		t.Error(err)
	}
	bacterialTable, err := sequence.GetCodonTable(11)
	if err != nil {
		t.Error(err)
	}

	expectedResult := []sequence.CodonTable{
		{
			ID:          1,
			Name:        "Standard",
			Description: "SGC0",
			Codons:      standardTable.Codons,
		},
		{
			ID:          2,
			Name:        "Vertebrate Mitochondrial",
			Description: "SGC1",
			Codons:      vmTable.Codons,
		},
		{
			ID:          4,
			Name:        "Mold Mitochondrial; Protozoan Mitochondrial; Coelenterate Mitochondrial; Mycoplasma; Spiroplasma",
			Description: "SGC3",
			Codons:      moldMtTable.Codons,
		},
		{
			ID:          10,
			Name:        "Euplotid Nuclear",
			Description: "SGC9",
			Codons:      euplotidTable.Codons,
		},
		{
			ID:          11,
			Name:        "Bacterial, Archaeal and Plant Plastid",
			Description: "",
			Codons:      bacterialTable.Codons,
		},
	}

	scanner := bufio.NewScanner(strings.NewReader(testInput))
	result, err := ParsePRTCodonTables(scanner)

	if err != nil {
		t.Error(err)
	}

	if len(result) != len(expectedResult) {
		t.Fatalf("Expected %d codon tables, got %d", len(expectedResult), len(result))
	}

	for i, table := range result {
		expectedTable := expectedResult[i]
		if table.ID != expectedTable.ID || table.Name != expectedTable.Name || table.Description != expectedTable.Description {
			t.Errorf("Table #%d: expected %+v, got %+v", i+1, expectedTable, table)
		}

		for codon, aa := range table.Codons {
			expectedAA, ok := expectedTable.Codons[codon]
			if !ok {
				t.Errorf("Table #%d: unexpected codon %s in result", i+1, codon)
			} else if aa != expectedAA {
				t.Errorf("Table #%d: expected amino acid %v for codon %s, got %v", i+1, expectedAA, codon, aa)
			}
		}
	}
}
