package ncbi

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
	var parser = newTableParser()

	for scanner.Scan() {
		line := scanner.Text()

		// Ignore comments and genetic codes beginning
		if strings.Contains(line, "Genetic-code-table") || strings.Contains(line, "}") || (strings.HasPrefix(line, "--") && !strings.HasPrefix(line, "-- Base")) {
			continue
		}

		if strings.Contains(line, "{") {
			// refresh parser for each new table
			parser = newTableParser()
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

type tableParser struct {
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

func newTableParser() *tableParser {
	return &tableParser{
		table: &sequence.CodonTable{
			Codons:      make(map[string]sequence.AminoAcid),
			StartCodons: make(map[string]sequence.AminoAcid),
			StopCodons:  make(map[string]sequence.AminoAcid),
		},
		ncbiCodeData: &ncbiCodeData{},
	}
}

func (tp *tableParser) processNameAndDescription(line string) {
	val := removeTagAndTrim(line, "name")

	if strings.Contains(line, "name") {
		// multiline?
		if strings.Count(line, "\"") == 1 {
			tp.multiline = true
		}
		if tp.nameIsUsed {
			tp.table.Description = val
		} else {
			tp.table.Name = val
			if !tp.multiline {
				tp.nameIsUsed = true
			}
		}
	} else if tp.multiline {
		if tp.nameIsUsed {
			tp.table.Description = fmt.Sprintf("%s %s", tp.table.Description, val)
		} else {
			tp.table.Name = fmt.Sprintf("%s %s", tp.table.Name, val)
		}

		// name ending?
		if strings.Count(line, "\"") == 1 {
			tp.multiline = false
			if !tp.nameIsUsed {
				tp.nameIsUsed = true
			}
		}
	}
}

func (tp *tableParser) processID(line string) error {
	if !strings.Contains(line, " id") {
		return nil
	}

	var err error

	tp.table.ID, err = strconv.Atoi(removeTagAndTrim(line, "id"))
	if err != nil {
		return err
	}

	return nil
}

func (tp *tableParser) processCodons(line string) {
	if strings.Contains(line, "sncbieaa") {
		tp.ncbiCodeData.sncbieaa = removeTagAndTrim(line, "sncbieaa")
	}
	if strings.Contains(line, " ncbieaa") {
		tp.ncbiCodeData.ncbieaa = removeTagAndTrim(line, "ncbieaa")
	}
	if strings.Contains(line, "-- Base1") {
		tp.ncbiCodeData.base1 = removeTagAndTrim(line, "-- Base1")
	}
	if strings.Contains(line, "-- Base2") {
		tp.ncbiCodeData.base2 = removeTagAndTrim(line, "-- Base2")
	}
	if strings.Contains(line, "-- Base3") {
		tp.ncbiCodeData.base3 = removeTagAndTrim(line, "-- Base3")
	}
	if tp.ncbiCodeData != nil && tp.ncbiCodeData.sncbieaa != "" &&
		tp.ncbiCodeData.ncbieaa != "" && tp.ncbiCodeData.base1 != "" &&
		tp.ncbiCodeData.base2 != "" && tp.ncbiCodeData.base3 != "" {
		tp.parseNCBICodeData()
	}
}

func (tp *tableParser) parseNCBICodeData() {
	for i := 0; i < len(tp.ncbiCodeData.ncbieaa); i++ {
		codon := string([]byte{tp.ncbiCodeData.base1[i], tp.ncbiCodeData.base2[i], tp.ncbiCodeData.base3[i]})
		// !! DNA -> RNA
		codon = strings.ReplaceAll(codon, "T", "U")
		tp.table.Codons[codon] = sequence.AminoAcid(tp.ncbiCodeData.ncbieaa[i])

		if tp.ncbiCodeData.sncbieaa[i] == '-' {
			continue
		}
		if tp.ncbiCodeData.sncbieaa[i] == '*' {
			tp.table.StopCodons[codon] = sequence.AminoAcid(tp.ncbiCodeData.ncbieaa[i])
		}
		if tp.ncbiCodeData.sncbieaa[i] == 'M' {
			tp.table.StartCodons[codon] = sequence.AminoAcid(tp.ncbiCodeData.ncbieaa[i])
		}
	}
}
