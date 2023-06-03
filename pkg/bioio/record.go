package bioio

type Record struct {
	ID          string
	Accession   string
	Version     int
	Organism    string
	Taxonomy    string
	Keywords    []string
	Source      string
	Description string
	Features    []Feature
	References  []Reference
	Sequence    string
}

type Feature struct {
	Type       string
	Location   string
	Qualifiers map[string]string
}

type Reference struct {
	Number  int
	Authors []string
	Title   string
	Journal string
	Medline string
	PubMed  string
	Remarks string
}
