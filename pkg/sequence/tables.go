package sequence

var Standard = CodonTable{
	ID:          1,
	Name:        "Standard",
	Description: "The Standard Code is the most common genetic code used by nuclear genes in eukaryotes, as well as by the vast majority of prokaryotes",
	Codons: map[string]AminoAcid{
		"AAA": 'K', "AAC": 'N', "AAG": 'K', "AAU": 'N',
		"ACA": 'T', "ACC": 'T', "ACG": 'T', "ACU": 'T',
		"AGA": 'R', "AGC": 'S', "AGG": 'R', "AGU": 'S',
		"AUA": 'I', "AUC": 'I', "AUG": 'M', "AUU": 'I',

		"CAA": 'Q', "CAC": 'H', "CAG": 'Q', "CAU": 'H',
		"CCA": 'P', "CCC": 'P', "CCG": 'P', "CCU": 'P',
		"CGA": 'R', "CGC": 'R', "CGG": 'R', "CGU": 'R',
		"CUA": 'L', "CUC": 'L', "CUG": 'L', "CUU": 'L',

		"GAA": 'E', "GAC": 'D', "GAG": 'E', "GAU": 'D',
		"GCA": 'A', "GCC": 'A', "GCG": 'A', "GCU": 'A',
		"GGA": 'G', "GGC": 'G', "GGG": 'G', "GGU": 'G',
		"GUA": 'V', "GUC": 'V', "GUG": 'V', "GUU": 'V',

		"UAA": '*', "UAC": 'Y', "UAG": '*', "UAU": 'Y',
		"UCA": 'S', "UCC": 'S', "UCG": 'S', "UCU": 'S',
		"UGA": '*', "UGC": 'C', "UGG": 'W', "UGU": 'C',
		"UUA": 'L', "UUC": 'F', "UUG": 'L', "UUU": 'F',
	},
}

var VertebrateMitochondrial = CodonTable{
	ID:          2,
	Name:        "Vertebrate Mitochondrial Code",
	Description: "The Vertebrate Mitochondrial Code is a variant of the genetic code used in the mitochondria of vertebrates",
	Codons: map[string]AminoAcid{
		"AAA": 'K', "AAC": 'N', "AAG": 'K', "AAU": 'N',
		"ACA": 'T', "ACC": 'T', "ACG": 'T', "ACU": 'T',
		"AGA": '*', "AGC": 'S', "AGG": '*', "AGU": 'S',
		"AUA": 'M', "AUC": 'I', "AUG": 'M', "AUU": 'I',

		"CAA": 'Q', "CAC": 'H', "CAG": 'Q', "CAU": 'H',
		"CCA": 'P', "CCC": 'P', "CCG": 'P', "CCU": 'P',
		"CGA": 'R', "CGC": 'R', "CGG": 'R', "CGU": 'R',
		"CUA": 'L', "CUC": 'L', "CUG": 'L', "CUU": 'L',

		"GAA": 'E', "GAC": 'D', "GAG": 'E', "GAU": 'D',
		"GCA": 'A', "GCC": 'A', "GCG": 'A', "GCU": 'A',
		"GGA": 'G', "GGC": 'G', "GGG": 'G', "GGU": 'G',
		"GUA": 'V', "GUC": 'V', "GUG": 'V', "GUU": 'V',

		"UAA": '*', "UAC": 'Y', "UAG": '*', "UAU": 'Y',
		"UCA": 'S', "UCC": 'S', "UCG": 'S', "UCU": 'S',
		"UGA": 'W', "UGC": 'C', "UGG": 'W', "UGU": 'C',
		"UUA": 'L', "UUC": 'F', "UUG": 'L', "UUU": 'F',
	},
}

var YeastMitochondrial = CodonTable{
	ID:          3,
	Name:        "Yeast Mitochondrial Code",
	Description: "The Yeast Mitochondrial Code is a variant of the genetic code used in the mitochondria of yeast, such as Saccharomyces cerevisiae",
	Codons: map[string]AminoAcid{
		"AAA": 'K', "AAC": 'N', "AAG": 'K', "AAU": 'N',
		"ACA": 'T', "ACC": 'T', "ACG": 'T', "ACU": 'T',
		"AGA": 'R', "AGC": 'S', "AGG": 'R', "AGU": 'S',
		"AUA": 'M', "AUC": 'I', "AUG": 'M', "AUU": 'I',

		"CAA": 'Q', "CAC": 'H', "CAG": 'Q', "CAU": 'H',
		"CCA": 'P', "CCC": 'P', "CCG": 'P', "CCU": 'P',
		"CGA": 'R', "CGC": 'R', "CGG": 'R', "CGU": 'R',
		"CUA": 'T', "CUC": 'L', "CUG": 'T', "CUU": 'L',

		"GAA": 'E', "GAC": 'D', "GAG": 'E', "GAU": 'D',
		"GCA": 'A', "GCC": 'A', "GCG": 'A', "GCU": 'A',
		"GGA": 'G', "GGC": 'G', "GGG": 'G', "GGU": 'G',
		"GUA": 'V', "GUC": 'V', "GUG": 'V', "GUU": 'V',

		"UAA": '*', "UAC": 'Y', "UAG": '*', "UAU": 'Y',
		"UCA": 'S', "UCC": 'S', "UCG": 'S', "UCU": 'S',
		"UGA": 'W', "UGC": 'C', "UGG": 'W', "UGU": 'C',
		"UUA": 'L', "UUC": 'F', "UUG": 'L', "UUU": 'F',
	},
}

var MoldProtozoanCoelenterateMitochondrial = CodonTable{
	ID:          4,
	Name:        "Mold, Protozoan, and Coelenterate Mitochondrial Code",
	Description: "The Mold, Protozoan, and Coelenterate Mitochondrial Code is a variant of the genetic code used in the mitochondria of molds, protozoans, and coelenterates, including some diplomonads and chlorophycean algae",
	Codons: map[string]AminoAcid{
		"AAA": 'K', "AAC": 'N', "AAG": 'K', "AAU": 'N',
		"ACA": 'T', "ACC": 'T', "ACG": 'T', "ACU": 'T',
		"AGA": 'S', "AGC": 'S', "AGG": 'S', "AGU": 'S',
		"AUA": 'M', "AUC": 'I', "AUG": 'M', "AUU": 'I',

		"CAA": 'Q', "CAC": 'H', "CAG": 'Q', "CAU": 'H',
		"CCA": 'P', "CCC": 'P', "CCG": 'P', "CCU": 'P',
		"CGA": 'R', "CGC": 'R', "CGG": 'R', "CGU": 'R',
		"CUA": 'L', "CUC": 'L', "CUG": 'L', "CUU": 'L',

		"GAA": 'E', "GAC": 'D', "GAG": 'E', "GAU": 'D',
		"GCA": 'A', "GCC": 'A', "GCG": 'A', "GCU": 'A',
		"GGA": 'G', "GGC": 'G', "GGG": 'G', "GGU": 'G',
		"GUA": 'V', "GUC": 'V', "GUG": 'V', "GUU": 'V',

		"UAA": '*', "UAC": 'Y', "UAG": '*', "UAU": 'Y',
		"UCA": 'S', "UCC": 'S', "UCG": 'S', "UCU": 'S',
		"UGA": 'W', "UGC": 'C', "UGG": 'W', "UGU": 'C',
		"UUA": 'L', "UUC": 'F', "UUG": 'L', "UUU": 'F',
	},
}

var InvertebrateMitochondrial = CodonTable{
	ID:          5,
	Name:        "Invertebrate Mitochondrial Code",
	Description: "The Invertebrate Mitochondrial Code is a variant of the genetic code used in the mitochondria of invertebrates, such as arthropods, mollusks, and echinoderms",
	Codons: map[string]AminoAcid{
		"AAA": 'K', "AAC": 'N', "AAG": 'K', "AAU": 'N',
		"ACA": 'T', "ACC": 'T', "ACG": 'T', "ACU": 'T',
		"AGA": 'S', "AGC": 'S', "AGG": 'S', "AGU": 'S',
		"AUA": 'M', "AUC": 'I', "AUG": 'M', "AUU": 'I',

		"CAA": 'Q', "CAC": 'H', "CAG": 'Q', "CAU": 'H',
		"CCA": 'P', "CCC": 'P', "CCG": 'P', "CCU": 'P',
		"CGA": 'R', "CGC": 'R', "CGG": 'R', "CGU": 'R',
		"CUA": 'L', "CUC": 'L', "CUG": 'L', "CUU": 'L',

		"GAA": 'E', "GAC": 'D', "GAG": 'E', "GAU": 'D',
		"GCA": 'A', "GCC": 'A', "GCG": 'A', "GCU": 'A',
		"GGA": 'G', "GGC": 'G', "GGG": 'G', "GGU": 'G',
		"GUA": 'V', "GUC": 'V', "GUG": 'V', "GUU": 'V',

		"UAA": '*', "UAC": 'Y', "UAG": '*', "UAU": 'Y',
		"UCA": 'S', "UCC": 'S', "UCG": 'S', "UCU": 'S',
		"UGA": 'W', "UGC": 'C', "UGG": 'W', "UGU": 'C',
		"UUA": 'L', "UUC": 'F', "UUG": 'L', "UUU": 'F',
	},
}

var CiliateDasycladaceanHexamitaNuclear = CodonTable{
	ID:          6,
	Name:        "Ciliate, Dasycladacean and Hexamita Nuclear Code",
	Description: "The Ciliate, Dasycladacean and Hexamita Nuclear Code is a variant of the genetic code used in the nuclear genes of ciliates, dasycladaceans, and hexamita species",
	Codons: map[string]AminoAcid{
		"AAA": 'K', "AAC": 'N', "AAG": 'K', "AAU": 'N',
		"ACA": 'T', "ACC": 'T', "ACG": 'T', "ACU": 'T',
		"AGA": 'R', "AGC": 'S', "AGG": 'R', "AGU": 'S',
		"AUA": 'I', "AUC": 'I', "AUG": 'M', "AUU": 'I',

		"CAA": 'Q', "CAC": 'H', "CAG": 'Q', "CAU": 'H',
		"CCA": 'P', "CCC": 'P', "CCG": 'P', "CCU": 'P',
		"CGA": 'R', "CGC": 'R', "CGG": 'R', "CGU": 'R',
		"CUA": 'L', "CUC": 'L', "CUG": 'L', "CUU": 'L',

		"GAA": 'E', "GAC": 'D', "GAG": 'E', "GAU": 'D',
		"GCA": 'A', "GCC": 'A', "GCG": 'A', "GCU": 'A',
		"GGA": 'G', "GGC": 'G', "GGG": 'G', "GGU": 'G',
		"GUA": 'V', "GUC": 'V', "GUG": 'V', "GUU": 'V',

		"UAA": 'Q', "UAC": 'Y', "UAG": 'Q', "UAU": 'Y',
		"UCA": 'S', "UCC": 'S', "UCG": 'S', "UCU": 'S',
		"UGA": 'W', "UGC": 'C', "UGG": 'W', "UGU": 'C',
		"UUA": 'L', "UUC": 'F', "UUG": 'L', "UUU": 'F',
	},
}

var EchinodermAndFlatwormMitochondrial = CodonTable{
	ID:          7,
	Name:        "Echinoderm and Flatworm Mitochondrial Code",
	Description: "The Echinoderm and Flatworm Mitochondrial Code is a variant of the genetic code used in the mitochondria of echinoderms, such as sea urchins and starfish, and flatworms, such as planarians.",
	Codons: map[string]AminoAcid{
		"AAA": 'K', "AAU": 'N', "AAC": 'N', "AAG": 'K',
		"AUA": 'I', "AUU": 'I', "AUC": 'I', "AUG": 'M',
		"ACA": 'T', "ACU": 'T', "ACC": 'T', "ACG": 'T',
		"AGA": 'S', "AGU": 'S', "AGC": 'S', "AGG": 'S',

		"CAA": 'Q', "CAU": 'H', "CAC": 'H', "CAG": 'Q',
		"CUA": 'L', "CUU": 'L', "CUC": 'L', "CUG": 'L',
		"CCA": 'P', "CCU": 'P', "CCC": 'P', "CCG": 'P',
		"CGA": 'R', "CGU": 'R', "CGC": 'R', "CGG": 'R',

		"GAA": 'E', "GAU": 'D', "GAC": 'D', "GAG": 'E',
		"GUA": 'V', "GUU": 'V', "GUC": 'V', "GUG": 'V',
		"GCA": 'A', "GCU": 'A', "GCC": 'A', "GCG": 'A',
		"GGA": 'G', "GGU": 'G', "GGC": 'G', "GGG": 'G',

		"UAA": '*', "UAU": 'Y', "UAC": 'Y', "UAG": '*',
		"UUA": 'L', "UUU": 'F', "UUC": 'F', "UUG": 'L',
		"UCA": 'S', "UCU": 'S', "UCC": 'S', "UCG": 'S',
		"UGA": 'W', "UGU": 'C', "UGC": 'C', "UGG": 'W',
	},
}

var EuplotidNuclear = CodonTable{
	ID:          8,
	Name:        "Euplotid Nuclear Code",
	Description: "The Euplotid Nuclear Code is a variant of the genetic code used in the nuclear genome of Euplotids, a group of ciliates.",
	Codons: map[string]AminoAcid{
		"AAA": 'K', "AAU": 'N', "AAC": 'N', "AAG": 'K',
		"AUA": 'I', "AUU": 'I', "AUC": 'I', "AUG": 'M',
		"ACA": 'T', "ACU": 'T', "ACC": 'T', "ACG": 'T',
		"AGA": 'S', "AGU": 'S', "AGC": 'S', "AGG": 'R',

		"CAA": 'Q', "CAU": 'H', "CAC": 'H', "CAG": 'Q',
		"CUA": 'L', "CUU": 'L', "CUC": 'L', "CUG": 'L',
		"CCA": 'P', "CCU": 'P', "CCC": 'P', "CCG": 'P',
		"CGA": 'R', "CGU": 'R', "CGC": 'R', "CGG": 'R',

		"GAA": 'E', "GAU": 'D', "GAC": 'D', "GAG": 'E',
		"GUA": 'V', "GUU": 'V', "GUC": 'V', "GUG": 'V',
		"GCA": 'A', "GCU": 'A', "GCC": 'A', "GCG": 'A',
		"GGA": 'G', "GGU": 'G', "GGC": 'G', "GGG": 'G',

		"UAA": 'Y', "UAU": 'Y', "UAC": 'Y', "UAG": 'Q',
		"UUA": 'L', "UUU": 'F', "UUC": 'F', "UUG": 'L',
		"UCA": 'S', "UCU": 'S', "UCC": 'S', "UCG": 'S',
		"UGA": 'C', "UGU": 'C', "UGC": 'C', "UGG": 'W',
	},
}

var EchinococcusAscarisMitochondrial = CodonTable{
	ID:          9,
	Name:        "Echinococcus and Ascaris Mitochondrial",
	Description: "Echinococcus and Ascaris Mitochondrial Translation Table",
	Codons: map[string]AminoAcid{
		"AAA": 'N', "AAC": 'N', "AAG": 'K', "AAU": 'N',
		"ACA": 'T', "ACC": 'T', "ACG": 'T', "ACU": 'T',
		"AGA": 'S', "AGC": 'S', "AGG": 'S', "AGU": 'S',
		"AUA": 'I', "AUC": 'I', "AUG": 'M', "AUU": 'I',

		"CAA": 'Q', "CAC": 'H', "CAG": 'Q', "CAU": 'H',
		"CCA": 'P', "CCC": 'P', "CCG": 'P', "CCU": 'P',
		"CGA": 'R', "CGC": 'R', "CGG": 'R', "CGU": 'R',
		"CUA": 'L', "CUC": 'L', "CUG": 'L', "CUU": 'L',

		"GAA": 'E', "GAC": 'D', "GAG": 'E', "GAU": 'D',
		"GCA": 'A', "GCC": 'A', "GCG": 'A', "GCU": 'A',
		"GGA": 'G', "GGC": 'G', "GGG": 'G', "GGU": 'G',
		"GUA": 'V', "GUC": 'V', "GUG": 'V', "GUU": 'V',

		"UAA": 'Y', "UAC": 'Y', "UAG": '*', "UAU": 'Y',
		"UCA": 'S', "UCC": 'S', "UCG": 'S', "UCU": 'S',
		"UGA": 'C', "UGC": 'C', "UGG": 'W', "UGU": 'C',
		"UUA": 'L', "UUC": 'F', "UUG": 'L', "UUU": 'F',
	},
}
