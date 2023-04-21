package sequence

import (
	"bytes"
)

type CodonTable struct {
	Name        string               `json:"name"`
	Description string               `json:"description"`
	Codons      map[string]AminoAcid `json:"codon_table"`
}

// GetCodonTable stub
func GetCodonTable(id int) (*CodonTable, error) {

	return &CodonTable{}, nil
}

// ModifyCodonUsage stub
func (c *CodonTable) ModifyCodonUsage(customCodons map[string]string) error {

	return nil
}

func (c *CodonTable) TranslateCodon(codon string) AminoAcid {
	possibleBases := make([][]Nucleotide, len(codon))

	codon = string(bytes.ToUpper([]byte(codon)))
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
