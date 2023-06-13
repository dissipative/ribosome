# Ribosome

[![Go Report Card](https://goreportcard.com/badge/github.com/dissipative/ribosome)](https://goreportcard.com/report/github.com/dissipative/ribosome)

The Ribosome package is a Go library designed for efficient transcription and translation of DNA and RNA sequences, inspired by the real processes in living cells. 
The package provides an easy-to-use API for handling DNA, RNA, and Protein sequences, as well as functionality to work with different genetic code translation tables.

Disclaimer: Go language is not a popular choice for bioinformatics library such as this and there are much better solutions written in other languages. 
This library was written just for fun and as proof of concept. 

## Features

- DNA, RNA, and Protein sequence handling
- Support for ambiguous nucleotides and amino acids
- Transcription of DNA to RNA
- Translation of RNA to Protein
- Genetic code translation tables with custom codon usage modification
- ORF finding
- GC-content calculation

## Installation

To install the Ribosome package, use the following command:
```
go get github.com/dissipative/ribosome
```

## Usage
Import the Ribosome package in your Go project:

```go
import "github.com/dissipative/ribosome"
```

## DNA Sequence
Create a DNA sequence:

```go
dna, err := sequence.NewDNASequence("ATGCGAATTCAG")
```

## RNA Sequence
Transcribe a DNA sequence to RNA:

```go
rna, err := dna.Transcribe()
```

## Genetic Code Translation Tables
Get a translation table by its ID:

```go
codonTable, err := ribosome.GetCodonTable(1) // Get the standard genetic code (table 1)
```

## Protein Sequence
Translate an RNA sequence to a Protein sequence:

```go
protein, err := rna.Translate(&codonTable) 
```

## Modify a translation table with custom codon usage:

```go
customCodons := map[string]ribosome.AminoAcid{
    "ATA": 'M', // Change ATA codon to Methionine
}
err = table.ModifyCodonUsage(customCodons)
```

## Ambiguous Nucleotides and Amino Acids
The Ribosome package handles ambiguous nucleotides and amino acids with ease. 
For example, you can transcribe DNA sequences with ambiguous bases and translate RNA sequences with ambiguous codons to protein sequences with ambiguous amino acids.
