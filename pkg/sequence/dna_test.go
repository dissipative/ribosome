package sequence

import (
	"reflect"
	"testing"
)

func TestDNASequence_Reverse(t *testing.T) {
	tests := []struct {
		name string
		d    DNASequence
		want DNASequence
	}{
		{
			name: "simple_revert",
			d:    DNASequence("ATGCGAATTCAG"),
			want: DNASequence("GACTTAAGCGTA"),
		},
		{
			name: "empty",
			d:    "",
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.Reverse(); got != tt.want {
				t.Errorf("Reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDNASequence_Complement(t *testing.T) {
	tests := []struct {
		name string
		d    DNASequence
		want DNASequence
	}{
		{
			name: "simple_sequence",
			d:    DNASequence("ATGCGAATTCAG"),
			want: DNASequence("TACGCTTAAGTC"),
		},
		{
			name: "ambiguous_nucleotides_and_gap",
			d:    DNASequence("ATGCGAATTCAGRYKMSWBDHV-N"),
			want: DNASequence("TACGCTTAAGTCYRMKSWVHDB-N"),
		},
		{
			name: "empty_sequence",
			d:    DNASequence(""),
			want: DNASequence(""),
		},
		{
			name: "unknown_base",
			d:    DNASequence("X"),
			want: DNASequence("-"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.Complement(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Complement() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDNASequence_Transcribe(t *testing.T) {
	tests := []struct {
		name        string
		dnaSeq      DNASequence
		expectedRNA RNASequence
	}{
		{
			name:        "simple_sequence",
			dnaSeq:      "ATGCGAATTCAG",
			expectedRNA: "UACGCUUAAGUC",
		},
		{
			name:        "mixed_case_sequence",
			dnaSeq:      "atgcgaattcag",
			expectedRNA: "UACGCUUAAGUC",
		},
		{
			name:        "ambiguous_bases",
			dnaSeq:      "ATGCGAATTCAGRYKMSWBDHV",
			expectedRNA: "UACGCUUAAGUCYRMKSWVHDB",
		},
		{
			name:        "empty",
			dnaSeq:      "",
			expectedRNA: "",
		},
		{
			name:        "invalid_base",
			dnaSeq:      "ATGCGAAX",
			expectedRNA: "UACGCUU-",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.dnaSeq.Transcribe(); got != tt.expectedRNA {
				t.Errorf("Transcribe() = %v, want %v", got, tt.expectedRNA)
			}
		})
	}
}
