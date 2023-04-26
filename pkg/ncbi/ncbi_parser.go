package ncbi

import (
	"bufio"
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/dissipative/ribosome/pkg/sequence"
)

type ncbiCodeData struct {
	ncbieaa string
	base1   string
	base2   string
	base3   string
}

func ParsePRTCodonTables(scanner *bufio.Scanner) ([]sequence.CodonTable, error) {
	var tables []sequence.CodonTable
	var codons *ncbiCodeData
	var parser = &tableParser{}

	for scanner.Scan() {
		line := scanner.Text()

		// Ignore comments and genetic codes beginning
		if strings.Contains(line, "Genetic-code-table") || (strings.HasPrefix(line, "--") && !strings.HasPrefix(line, "-- Base")) {
			continue
		}

		if strings.Contains(line, "{") {
			codons = &ncbiCodeData{}
			parser = &tableParser{
				table: &sequence.CodonTable{
					Codons: make(map[string]sequence.AminoAcid),
				},
			}
			continue
		}

		parser.processName(line)

		// Match key-value pairs
		re := regexp.MustCompile(`\s*(?:([a-zA-Z0-9]+)\s+(?:"([^"]+)"\s*,?|(\d+)\s*,?)|--\s+([a-zA-Z0-9]+)\s+([a-zA-Z0-9]+))`)
		matches := re.FindStringSubmatch(line)

		if len(matches) == 6 {
			key := matches[1]
			textVal := matches[2]
			numVal := matches[3]
			baseKey := matches[4]
			baseVal := matches[5]

			if key == "" {
				key = baseKey
			}

			switch key {
			case "id":
				fmt.Sscanf(numVal, "%d", &parser.table.ID)
			case "ncbieaa":
				codons.ncbieaa = textVal
			case "Base1":
				codons.base1 = baseVal
			case "Base2":
				codons.base2 = baseVal
			case "Base3":
				codons.base3 = baseVal
				parser.table.Codons = codons.parse()
			}

			if parser.table.ID != 0 && parser.table.Name != "" && len(parser.table.Codons) > 0 {
				tables = append(tables, *parser.table)
			}
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

func (c *ncbiCodeData) parse() map[string]sequence.AminoAcid {
	codonTable := make(map[string]sequence.AminoAcid)

	for i := 0; i < len(c.ncbieaa); i++ {
		codon := string([]byte{c.base1[i], c.base2[i], c.base3[i]})
		// !! DNA -> RNA
		codon = strings.ReplaceAll(codon, "T", "U")
		codonTable[codon] = sequence.AminoAcid(c.ncbieaa[i])
	}

	return codonTable
}

func removeTagAndTrim(line string, tag string) string {
	line = strings.Replace(line, tag, "", 1)
	line = strings.ReplaceAll(line, "\"", "")
	line = strings.Trim(line, ",")
	return strings.TrimSpace(line)
}

type tableParser struct {
	nameIsUsed bool
	multiline  bool
	table      *sequence.CodonTable
}

func (tp *tableParser) processName(line string) {
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
