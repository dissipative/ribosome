package bioio

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func readFASTA(reader io.Reader) ([]Record, error) {
	var sequences []Record
	scanner := bufio.NewScanner(reader)

	var currentID string
	var currentSeq strings.Builder

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			continue
		}

		if line[0] == '>' {
			if currentID != "" {
				sequences = append(sequences, Record{ID: currentID, Sequence: currentSeq.String()})
				currentSeq.Reset()
			}
			currentID = line[1:]
		} else {
			currentSeq.WriteString(line)
		}
	}

	if currentID != "" {
		sequences = append(sequences, Record{ID: currentID, Sequence: currentSeq.String()})
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return sequences, nil
}

func writeFASTA(writer io.Writer, sequences []Record) error {
	for _, seq := range sequences {
		_, err := fmt.Fprintf(writer, ">%s\n%s\n", seq.ID, seq.Sequence)
		if err != nil {
			return err
		}
	}

	return nil
}
