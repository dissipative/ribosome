package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/dissipative/ribosome/pkg/bioio"
	"github.com/dissipative/ribosome/pkg/sequence"
)

// example runs:
//
//	go run cmd/example_app/main.go --input="test/mitochondrions.raw.fas" --table-id=5
//	go run cmd/example_app/main.go --tables
func main() {
	// Declare flags
	var inputFile string
	var formatString string
	var codonTable int
	var tablesInfo bool
	flag.StringVar(&inputFile, "input", "", "Input file path")
	flag.StringVar(&formatString, "format", "fasta", "Format of the input file (fasta or genbank)")
	flag.IntVar(&codonTable, "table-id", 1, "Codon table used for sequence translation")
	flag.BoolVar(&tablesInfo, "tables", false, "Display codon tables")
	flag.Parse()

	if tablesInfo {
		printCodonTables()
		return
	}

	if inputFile == "" || formatString == "" {
		fmt.Println("Please provide both --input and --format flags.")
		os.Exit(1)
	}

	err := processSequences(formatString, codonTable, inputFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func processSequences(formatString string, codonTableID int, inputFile string) error {
	// Determine file format
	var format bioio.Format
	switch formatString {
	case "fas":
		fallthrough
	case "fasta":
		format = bioio.Fasta
	case "gb":
		fallthrough
	case "genbank":
		format = bioio.Genbank
	default:
		return errors.New("invalid format; please use either 'fasta'/'fas' or 'genbank'/'gb'")
	}

	codonTable, err := sequence.GetCodonTable(codonTableID)
	if err != nil {
		return err
	}

	// Read sequences from file
	sequences, err := bioio.ReadFile(inputFile, format)
	if err != nil {
		return fmt.Errorf("error reading file: %v", err)
	}

	// Process each sequence
	for i, seq := range sequences {
		err := printSequenceInfo(i, seq, codonTable)
		if err != nil {
			return err
		}
	}

	return nil
}

func printSequenceInfo(i int, seq bioio.Record, codonTable sequence.CodonTable) error {
	fmt.Printf("Record %d: %s\n", i+1, seq.ID)

	// Convert to RNASequence
	dna, err := sequence.NewDNASequence(seq.Sequence)
	if err != nil {
		return err
	}

	// Genbank sequences are reversed?
	rna := dna.Reverse().Transcribe()

	// Find ORFs with length > 900 bases
	orfs, err := rna.FindORFs(300, &codonTable)
	if err != nil {
		return fmt.Errorf("error finding ORFs: %v", err)
	}
	fmt.Printf("  Found %d ORFs\n", len(orfs))

	// Calculate GC content
	gcContent := sequence.GCContent(rna)
	fmt.Printf("  GC content: %.2f\n", gcContent)

	// Find the longest protein and print
	if orfs == nil {
		return nil
	}
	sort.Slice(orfs, func(i, j int) bool {
		return len(orfs[i].ProteinSeq) < len(orfs[j].ProteinSeq)
	})
	fmt.Printf("  Longest protein length: %d bases / %d amino acids\n", orfs[0].End-orfs[0].Start, len(orfs[0].ProteinSeq))
	fmt.Printf("  Longest protein frame: %d\n", orfs[0].Frame)
	fmt.Printf("  Longest protein sequence: %s\n", orfs[0].ProteinSeq)

	return nil
}

func printCodonTables() {
	for _, table := range sequence.CodonTables() {
		fmt.Printf("Codon Table ID %d:\n", table.ID)
		fmt.Printf("  Name: %s\n", table.Name)

		if table.Description != "" {
			fmt.Printf("  Description: %s\n", table.Description)
		}

		// Print start codons
		fmt.Println("  Start Codons:")
		for codon, aa := range table.StartCodons {
			fmt.Printf("    %s: %c", codon, aa)
		}

		// Print stop codons
		fmt.Println("\n  Stop Codons:")
		for codon, _ := range table.StopCodons {
			fmt.Printf("    %s", codon)
		}
		fmt.Println("\n", strings.Repeat("-", 50))
	}
}
