package parser

import (
	"bufio"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/dissipative/ribosome/pkg/sequence"
)

func ParsePRTCodonTables(scanner *bufio.Scanner) ([]sequence.CodonTable, error) {
	var tables []sequence.CodonTable
	parser := newGeneticCodesParser()

	for scanner.Scan() {
		line := scanner.Text()

		// Ignore comments and genetic codes beginning
		if strings.Contains(line, "Genetic-code-table") || strings.Contains(line, "}") || (strings.HasPrefix(line, "--") && !strings.HasPrefix(line, "-- Base")) {
			continue
		}

		if strings.Contains(line, "{") {
			// refresh parser for each new table
			parser = newGeneticCodesParser()
			continue
		}

		err := parser.processID(line)
		if err != nil {
			return nil, err
		}

		parser.processNameAndDescription(line)
		parser.processCodons(line)

		if parser.table.ID != 0 && parser.table.Name != "" && len(parser.table.Codons) > 0 {
			tables = append(tables, *parser.table)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	if len(tables) == 0 {
		return nil, errors.New("no codon tables found")
	}

	return tables, nil
}

func removeTagAndTrim(line string, tag string) string {
	line = strings.Replace(line, tag, "", 1)
	line = strings.ReplaceAll(line, "\"", "")
	line = strings.Trim(line, ",")
	return strings.TrimSpace(line)
}

type geneticCodesParser struct {
	nameIsUsed   bool
	multiline    bool
	table        *sequence.CodonTable
	ncbiCodeData *ncbiCodeData
}

type ncbiCodeData struct {
	ncbieaa  string
	sncbieaa string
	base1    string
	base2    string
	base3    string
}

func newGeneticCodesParser() *geneticCodesParser {
	return &geneticCodesParser{
		table: &sequence.CodonTable{
			Codons:      make(map[string]sequence.AminoAcid),
			StartCodons: make(map[string]sequence.AminoAcid),
			StopCodons:  make(map[string]sequence.AminoAcid),
		},
		ncbiCodeData: &ncbiCodeData{},
	}
}

func (gcp *geneticCodesParser) processNameAndDescription(line string) {
	val := removeTagAndTrim(line, "name")

	if strings.Contains(line, "name") {
		// multiline?
		if strings.Count(line, "\"") == 1 {
			gcp.multiline = true
		}
		if gcp.nameIsUsed {
			gcp.table.Description = val
		} else {
			gcp.table.Name = val
			if !gcp.multiline {
				gcp.nameIsUsed = true
			}
		}
	} else if gcp.multiline {
		if gcp.nameIsUsed {
			gcp.table.Description = fmt.Sprintf("%s %s", gcp.table.Description, val)
		} else {
			gcp.table.Name = fmt.Sprintf("%s %s", gcp.table.Name, val)
		}

		// name ending?
		if strings.Count(line, "\"") == 1 {
			gcp.multiline = false
			if !gcp.nameIsUsed {
				gcp.nameIsUsed = true
			}
		}
	}
}

func (gcp *geneticCodesParser) processID(line string) error {
	if !strings.Contains(line, " id") {
		return nil
	}

	var err error

	gcp.table.ID, err = strconv.Atoi(removeTagAndTrim(line, "id"))
	if err != nil {
		return err
	}

	return nil
}

func (gcp *geneticCodesParser) processCodons(line string) {
	if strings.Contains(line, "sncbieaa") {
		gcp.ncbiCodeData.sncbieaa = removeTagAndTrim(line, "sncbieaa")
	}
	if strings.Contains(line, " ncbieaa") {
		gcp.ncbiCodeData.ncbieaa = removeTagAndTrim(line, "ncbieaa")
	}
	if strings.Contains(line, "-- Base1") {
		gcp.ncbiCodeData.base1 = removeTagAndTrim(line, "-- Base1")
	}
	if strings.Contains(line, "-- Base2") {
		gcp.ncbiCodeData.base2 = removeTagAndTrim(line, "-- Base2")
	}
	if strings.Contains(line, "-- Base3") {
		gcp.ncbiCodeData.base3 = removeTagAndTrim(line, "-- Base3")
	}
	if gcp.ncbiCodeData != nil && gcp.ncbiCodeData.sncbieaa != "" &&
		gcp.ncbiCodeData.ncbieaa != "" && gcp.ncbiCodeData.base1 != "" &&
		gcp.ncbiCodeData.base2 != "" && gcp.ncbiCodeData.base3 != "" {
		gcp.parseNCBICodeData()
	}
}

func (gcp *geneticCodesParser) parseNCBICodeData() {
	for i := 0; i < len(gcp.ncbiCodeData.ncbieaa); i++ {
		codon := string([]byte{gcp.ncbiCodeData.base1[i], gcp.ncbiCodeData.base2[i], gcp.ncbiCodeData.base3[i]})
		// DNA -> RNA
		codon = strings.ReplaceAll(codon, "T", "U")
		gcp.table.Codons[codon] = sequence.AminoAcid(gcp.ncbiCodeData.ncbieaa[i])

		if gcp.ncbiCodeData.sncbieaa[i] == '-' {
			continue
		}
		if gcp.ncbiCodeData.sncbieaa[i] == '*' {
			gcp.table.StopCodons[codon] = sequence.AminoAcid(gcp.ncbiCodeData.ncbieaa[i])
		}
		if gcp.ncbiCodeData.sncbieaa[i] == 'M' {
			gcp.table.StartCodons[codon] = sequence.AminoAcid(gcp.ncbiCodeData.ncbieaa[i])
		}
	}
}
