package main

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/dissipative/ribosome/internal/parser"
	"go/format"
	"log"
	"os"
	"text/template"
	"time"

	"github.com/dissipative/ribosome/pkg/sequence"
	"github.com/jlaffaye/ftp"
)

const (
	ftpAddr    = "ftp.ncbi.nih.gov:21"
	ftpPath    = "entrez/misc/data/gc.prt"
	anonymous  = "anonymous"
	outputFile = "pkg/sequence/codon_tables.go"
)

func main() {
	// Download codon tables from ftp://ftp.ncbi.nih.gov/entrez/misc/data/gc.prt
	conn, err := ftp.Dial(ftpAddr)
	if err != nil {
		log.Fatalf("error connecting to %s: %v\n", ftpAddr, err)
	}

	err = conn.Login(anonymous, anonymous)
	if err != nil {
		log.Fatalf("error at login to %s: %v\n", ftpAddr, err)
	}

	resp, err := conn.Retr(ftpPath)
	if err != nil {
		log.Fatalf("error downloading gc.prt: %v\n", err)
	}

	defer resp.Close()

	// Parse gc.prt content
	scanner := bufio.NewScanner(resp)
	tables, err := parser.ParsePRTCodonTables(scanner)
	if err != nil {
		fmt.Printf("error parsing gc.prt: %v\n", err)
		return
	}

	// Generate Go code for the codon tables
	var buf bytes.Buffer
	err = generateCode(&buf, tables)
	if err != nil {
		fmt.Printf("error generating Go code: %v\n", err)
		return
	}

	// Format the generated Go code
	formatted, err := format.Source(buf.Bytes())
	if err != nil {
		fmt.Printf("error formatting Go code: %v\n", err)
		return
	}

	// Write the generated code to the tables.go file
	err = os.WriteFile(outputFile, formatted, 0o644)
	if err != nil {
		fmt.Printf("error writing to %s: %v\n", outputFile, err)
		return
	}

	fmt.Printf("Generated %s\n", outputFile)
}

func generateCode(wr *bytes.Buffer, tables []sequence.CodonTable) error {
	tmplFuncs := template.FuncMap{
		"aaToString": func(aa sequence.AminoAcid) string {
			return fmt.Sprintf("%q", aa)
		},
		"add": func(a, b int) int {
			return a + b
		},
		"date": func() string {
			return fmt.Sprintf(" %s", time.Now().Format("2006-01-02 15:04"))
		},
	}

	tmpl, err := template.New("tables.tmpl").Funcs(tmplFuncs).ParseFiles("cmd/codon_tables_gen/tables.tmpl")
	if err != nil {
		return err
	}

	return tmpl.Execute(wr, tables)
}
