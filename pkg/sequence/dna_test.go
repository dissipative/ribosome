package sequence

import (
	"reflect"
	"testing"
)

func TestDNASequence_Reverse(t *testing.T) {
	tests := []struct {
		name     string
		source   DNASequence
		expected DNASequence
	}{
		{
			name:     "simple_revert",
			source:   DNASequence("ATGCGAATTCAG"),
			expected: DNASequence("GACTTAAGCGTA"),
		},
		{
			name:     "empty",
			source:   "",
			expected: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.source.Reverse(); got != tt.expected {
				t.Errorf("Reverse() = %v, expected %v", got, tt.expected)
			}
		})
	}
}

func TestDNASequence_Complement(t *testing.T) {
	tests := []struct {
		name     string
		source   DNASequence
		expected DNASequence
	}{
		{
			name:     "simple_sequence",
			source:   DNASequence("ATGCGAATTCAG"),
			expected: DNASequence("TACGCTTAAGTC"),
		},
		{
			name:     "ambiguous_nucleotides_and_gap",
			source:   DNASequence("ATGCGAATTCAGRYKMSWBDHV-N"),
			expected: DNASequence("TACGCTTAAGTCYRMKSWVHDB-N"),
		},
		{
			name:     "empty_sequence",
			source:   DNASequence(""),
			expected: DNASequence(""),
		},
		{
			name:     "unknown_base",
			source:   DNASequence("X"),
			expected: DNASequence("-"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.source.Complement(); !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("Complement() = %v, expected %v", got, tt.expected)
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
				t.Errorf("Transcribe() = %v, expected %v", got, tt.expectedRNA)
			}
		})
	}
}
