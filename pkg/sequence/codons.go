package sequence

import (
	"fmt"
	"strings"
)

type CodonTable struct {
	ID          int                  `json:"id"`
	Name        string               `json:"name"`
	Description string               `json:"description"`
	Codons      map[string]AminoAcid `json:"codon_table"`
}

// GetCodonTable stub
func GetCodonTable(id int) (CodonTable, error) {
	for _, c := range CodonTables {
		if id == c.ID {
			return c.Copy(), nil
		}
	}

	return CodonTable{}, fmt.Errorf("codon table no. %d not found", id)
}

func (c *CodonTable) Copy() CodonTable {
	tableCopy := *c

	copiedCodons := make(map[string]AminoAcid, len(tableCopy.Codons))
	for k, v := range tableCopy.Codons {
		copiedCodons[k] = v
	}

	tableCopy.Codons = copiedCodons
	return tableCopy
}

// ModifyCodonUsage stub
func (c *CodonTable) ModifyCodonUsage(customCodons map[string]AminoAcid) error {
	// Validate custom codons
	for codon, aa := range customCodons {
		// Check if the codon is 3 characters long
		if len(codon) != 3 {
			return fmt.Errorf("invalid codon length for '%s': expected 3 characters", codon)
		}

		// Ensure all characters in the codon are uppercase RNA nucleotides or ambiguous symbols
		for _, ch := range codon {
			if !isValidRNANucleotide(Nucleotide(ch)) {
				return fmt.Errorf("invalid character '%c' in codon '%s': expected uppercase RNA nucleotide or ambiguous symbol", ch, codon)
			}
		}

		// Validate the amino acid
		if !isValidAminoAcid(aa) {
			return fmt.Errorf("invalid amino acid '%c' for codon '%s'", aa, codon)
		}
	}

	// Modify the codon table
	for codon, aa := range customCodons {
		c.Codons[codon] = aa
	}

	return nil
}

func isValidRNANucleotide(base Nucleotide) bool {
	_, isAmbiguous := AmbiguousNucleotidesMap[base]
	isBasic := false
	for _, b := range AmbiguousNucleotidesMap['N'] {
		if b == base {
			isBasic = true
		}
	}

	return isBasic || isAmbiguous
}

func isValidAminoAcid(aa AminoAcid) bool {
	_, isAmbiguous := AmbiguousAminoAcidsMap[aa]
	isBasic := false
	for _, a := range AmbiguousAminoAcidsMap['X'] {
		if a == aa {
			isBasic = true
		}
	}

	return isBasic || isAmbiguous
}

func (c *CodonTable) TranslateCodon(codon string) AminoAcid {
	possibleBases := make([][]Nucleotide, len(codon))

	codon = strings.ToUpper(codon)
	codonsCount := 1

	for i, b := range codon {
		base := Nucleotide(b)

		variants, isAmbiguous := AmbiguousNucleotidesMap[base]
		if isAmbiguous {
			possibleBases[i] = variants
			codonsCount = codonsCount * len(variants)
		} else {
			possibleBases[i] = []Nucleotide{base}
		}
	}

	possibleCodons := make([]string, 0)
	allCombinations(possibleBases, 0, make([]Nucleotide, len(possibleBases)), &possibleCodons)

	var result AminoAcid
	for _, codon := range possibleCodons {
		aa, found := c.Codons[codon]
		// codon not found or differs from previous
		if !found || (result != 0 && aa != result) {
			return 'X'
		}
		result = aa
	}
	return result
}

func allCombinations(bases [][]Nucleotide, index int, currCombination []Nucleotide, result *[]string) {
	//  f the index is equal to the length of the input,
	// it means we have processed all byte slices and built a valid combination
	if index == len(currCombination) {
		newCombination := make([]Nucleotide, len(currCombination))
		copy(newCombination, currCombination)
		*result = append(*result, string(newCombination))
		return
	}

	for _, b := range bases[index] {
		currCombination[index] = b
		allCombinations(bases, index+1, currCombination, result)
	}
}
