package bioio

import (
	"bytes"
	"io"
	"reflect"
	"strings"
	"testing"
)

var inputSimple = `>sequence1
ATGCGAATTCAG
>sequence2
ATGGCACTGA
>sequence3
ATGCGTAGCATCAG`

var inputMultiline = `>sequence1
ATGCGAATTCAG
ATGGCACTGA
>sequence2
ATGGCACTGA
ATGCGTAGCATCAG
>sequence3
ATGCGTAGCATCAG

ATGCGAATTCAG`

func Test_readFASTA(t *testing.T) {
	tests := []struct {
		name    string
		file    io.Reader
		want    []Sequence
		wantErr bool
	}{
		{
			name: "read-simple",
			file: strings.NewReader(inputSimple),
			want: []Sequence{
				{
					ID:       "sequence1",
					Sequence: "ATGCGAATTCAG",
				},
				{
					ID:       "sequence2",
					Sequence: "ATGGCACTGA",
				},
				{
					ID:       "sequence3",
					Sequence: "ATGCGTAGCATCAG",
				},
			},
			wantErr: false,
		},
		{
			name: "read-multiline-sequences",
			file: strings.NewReader(inputMultiline),
			want: []Sequence{
				{
					ID:       "sequence1",
					Sequence: "ATGCGAATTCAGATGGCACTGA",
				},
				{
					ID:       "sequence2",
					Sequence: "ATGGCACTGAATGCGTAGCATCAG",
				},
				{
					ID:       "sequence3",
					Sequence: "ATGCGTAGCATCAGATGCGAATTCAG",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := readFASTA(tt.file)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadFile() got = %v, want %v", got, tt.want)
			}
		})
	}
}

var output = `>sequence1
ATGCGAATTCAGATGGCACTGA
>sequence2
ATGGCACTGAATGCGTAGCATCAG
>sequence3
ATGCGTAGCATCAGATGCGAATTCAG
`

func Test_writeFASTA(t *testing.T) {
	tests := []struct {
		name       string
		sequences  []Sequence
		wantWriter string
		wantErr    bool
	}{
		{
			name: "simple",
			sequences: []Sequence{
				{
					ID:       "sequence1",
					Sequence: "ATGCGAATTCAGATGGCACTGA",
				},
				{
					ID:       "sequence2",
					Sequence: "ATGGCACTGAATGCGTAGCATCAG",
				},
				{
					ID:       "sequence3",
					Sequence: "ATGCGTAGCATCAGATGCGAATTCAG",
				},
			},
			wantWriter: output,
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writer := &bytes.Buffer{}
			err := writeFASTA(writer, tt.sequences)
			if (err != nil) != tt.wantErr {
				t.Errorf("writeFASTA() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotWriter := writer.String(); gotWriter != tt.wantWriter {
				t.Errorf("writeFASTA() gotWriter = %v, want %v", gotWriter, tt.wantWriter)
			}
		})
	}
}
