package bioio

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
	"unicode"
)

func readGenbank(reader io.Reader) ([]Sequence, error) {
	var sequences []Sequence
	scanner := bufio.NewScanner(reader)
	var currentSeq *Sequence

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)

		if len(fields) == 0 {
			continue
		}

		switch fields[0] {
		case "LOCUS":
			if currentSeq != nil {
				sequences = append(sequences, *currentSeq)
			}
			currentSeq = &Sequence{}
			currentSeq.ID = fields[1]

		case "DEFINITION":
			if currentSeq == nil {
				currentSeq = &Sequence{}
			}

			description := strings.Join(fields[1:], " ")
			currentSeq.Description = strings.TrimSpace(description)

		case "ORIGIN":
			if currentSeq == nil {
				currentSeq = &Sequence{}
			}

			for scanner.Scan() {
				line := scanner.Text()
				if strings.HasPrefix(line, "//") {
					break
				}

				for _, char := range line {
					if char >= 'A' && char <= 'Z' || char >= 'a' && char <= 'z' {
						currentSeq.Sequence += string(char)
					}
				}
			}

		case "VERSION":
			if currentSeq == nil {
				currentSeq = &Sequence{}
			}

			// Assuming that the version information is always the second field and it includes accession number as well
			// Usually it's in the format "ACCESSION.VERSION"
			versionInfo := fields[1]
			parts := strings.Split(versionInfo, ".")
			if len(parts) > 1 {
				version, err := strconv.Atoi(parts[1])
				if err != nil {
					return nil, fmt.Errorf("invalid version number: %v", err)
				}
				currentSeq.Version = version
			}

		case "ORGANISM":
			if currentSeq == nil {
				currentSeq = &Sequence{}
			}

			// The organism name is usually the rest of the line after "ORGANISM"
			organism := strings.Join(fields[1:], " ")
			currentSeq.Organism = organism

			// The hierarchical classification is usually on the next lines,
			// until we reach a line that doesn't start with a whitespace
			var taxonomy string
			for scanner.Scan() {
				line := scanner.Text()
				if len(line) == 0 || !unicode.IsSpace(rune(line[0])) {
					break
				}

				// Remove leading and trailing whitespaces and add it to taxonomy
				taxonomy += " " + strings.TrimSpace(line)
			}

			currentSeq.Taxonomy = strings.TrimSpace(taxonomy)
		}
	}

	if currentSeq != nil {
		if currentSeq.ID == "" {
			return nil, errors.New("empty LOCUS field")
		}
		if currentSeq.Sequence == "" {
			return nil, errors.New("empty ORIGIN field")
		}

		sequences = append(sequences, *currentSeq)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return sequences, nil
}

func writeGenbank(writer io.Writer, sequences []Sequence) error {
	for _, seq := range sequences {
		_, err := fmt.Fprintf(writer, "LOCUS       %s\n", seq.ID)
		if err != nil {
			return err
		}

		_, err = fmt.Fprintf(writer, "DEFINITION  %s\n", seq.Description)
		if err != nil {
			return err
		}

		_, err = fmt.Fprint(writer, "ORIGIN\n")
		if err != nil {
			return err
		}

		for i := 0; i < len(seq.Sequence); i += 60 {
			end := i + 60
			if end > len(seq.Sequence) {
				end = len(seq.Sequence)
			}
			_, err = fmt.Fprintf(writer, "%9d %s\n", i+1, seq.Sequence[i:end])
			if err != nil {
				return err
			}
		}

		_, err = fmt.Fprint(writer, "//\n")
		if err != nil {
			return err
		}
	}

	return nil
}
