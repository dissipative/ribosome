package sequence

type Nucleotide byte

type AminoAcid byte

var AmbiguousAminoAcidsMap = map[AminoAcid][]AminoAcid{
	'B': {'N', 'D'},
	'Z': {'Q', 'E'},
	'J': {'L', 'I'},
	'X': {'A', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'K', 'L', 'M', 'N', 'P', 'Q', 'R', 'S', 'T', 'V', 'W', 'Y'},
}

var AmbiguousNucleotidesMap = map[Nucleotide][]Nucleotide{
	'R': {'A', 'G'},
	'Y': {'C', 'T'},
	'S': {'G', 'C'},
	'W': {'A', 'T'},
	'K': {'G', 'T'},
	'M': {'A', 'C'},
	'B': {'C', 'G', 'T'},
	'D': {'A', 'G', 'T'},
	'H': {'A', 'C', 'T'},
	'V': {'A', 'C', 'G'},
	'N': {'A', 'C', 'G', 'T'},
}
