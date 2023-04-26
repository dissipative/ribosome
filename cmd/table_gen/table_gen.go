package main

import (
	"bufio"
	"bytes"
	"fmt"
	"go/format"
	"log"
	"os"
	"text/template"

	"github.com/dissipative/ribosome/pkg/ncbi"
	"github.com/dissipative/ribosome/pkg/sequence"
	"github.com/jlaffaye/ftp"
)

const ftpAddr = "ftp.ncbi.nih.gov:21"
const ftpPath = "entrez/misc/data/gc.prt"
const anonymous = "anonymous"
const outputFile = "pkg/sequence/tables.go"

func main() {
	// Download gc.prt content
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
	tables, err := ncbi.ParsePRTCodonTables(scanner)
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
	err = os.WriteFile(outputFile, formatted, 0644)
	if err != nil {
		fmt.Printf("error writing to %s: %v\n", outputFile, err)
		return
	}

	fmt.Println("Generated tables.go")
}

func generateCode(wr *bytes.Buffer, tables []sequence.CodonTable) error {
	tmplFuncs := template.FuncMap{
		"aaToString": func(aa sequence.AminoAcid) string {
			return fmt.Sprintf("%q", aa)
		},
		"add": func(a, b int) int {
			return a + b
		},
	}

	tmpl, err := template.New("tables.tmpl").Funcs(tmplFuncs).ParseFiles("cmd/table_gen/tables.tmpl")
	if err != nil {
		return err
	}

	return tmpl.Execute(wr, tables)
}
