package bioio

import (
	"bufio"
	"errors"
	"io"
	"strings"
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
