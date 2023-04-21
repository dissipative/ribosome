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
	'Y': {'C', 'U'},
	'S': {'G', 'C'},
	'W': {'A', 'U'},
	'K': {'G', 'U'},
	'M': {'A', 'C'},
	'B': {'C', 'G', 'U'},
	'D': {'A', 'G', 'U'},
	'H': {'A', 'C', 'U'},
	'V': {'A', 'C', 'G'},
	'N': {'A', 'C', 'G', 'U'},
}
