// Jonathan Guillotte-Blouin - 7900293
package utils

import (
	"fmt"
	"bufio"
	"io"
	"strings"
)

// define a FASTA struct
type FASTA struct {
	Header string
	Body   string
}

func (brand *FASTA) Transcript() *FASTA {
	transcripted := make([]rune, len(brand.Body))
	for i, char := range brand.Body {
		// for every 'T', change it to a 'U'
		if char == 'T' {
			char = 'U'
		}

		transcripted[i] = char
	}

	brand.Body = string(transcripted)
	return brand
}

func (brand *FASTA) Complement() *FASTA {
	complemented := make([]rune, len(brand.Body))

	for i, char := range brand.Body {
		// map complementary nucleotides
		complement := complements[char]

		// put the complement in the temporary slice
		complemented[i] = complement
	}

	brand.Body = string(complemented)
	return brand
}

func (brand *FASTA) Reverse() *FASTA {
	brandLen := len(brand.Body)
	reversed := make([]rune, brandLen)

	for i, char := range brand.Body {
		// put the complement in reverse order of input
		reversed[brandLen-1-i] = char
	}

	brand.Body = string(reversed)
	return brand
}

func (brand *FASTA) ReverseComplement() {
	brand.Reverse().Complement()
}

// given a FASTA file, returns a pointer to a FASTA struct containing the information of the file
func ReadFASTAFile(file io.Reader) *FASTA {
	scanner := bufio.NewScanner(file)

	// first line is the header
	scanner.Scan()
	header := scanner.Text()

	// append the rest of the lines as the body
	body := []string{}
	for scanner.Scan() {
		body = append(body, scanner.Text())
	}

	// return the pointer to the struct
	return &FASTA{
		header,
		strings.Join(body, ""), // joins the multiple lines of the body as one line
	}
}

func (brand *FASTA) ReadFrame(direction string, start int) *FASTA {
	header := fmt.Sprintf("> %s Frame %d", direction, start + 1)
	codon := []rune{}
	frame := []rune{}

	for i, char := range brand.Body {
		if i < start {
			continue
		}

		codon = append(codon, char)

		if len(codon) == 3 {
			aminoAcid := NucleotidesToAminoAcid[string(codon)]
			frame = append(frame, aminoAcid)
			codon = []rune{}
		}
		i++
	}

	return &FASTA{
		header,
		string(frame),
	}
}

func (brand *FASTA) String() string {
	return fmt.Sprintf("%s\n%s\n", brand.Header, brand.Body)
}

// map of complement DNA nucleotides
var complements = map[rune]rune{
	'A': 'T',
	'T': 'A',
	'C': 'G',
	'G': 'C',
}

// map of nucleotides triplets to amino acid
var NucleotidesToAminoAcid = map[string]rune{
	"ATT": 'I',
	"ATC": 'I',
	"ATA": 'I',
	"CTT": 'L',
	"CTC": 'L',
	"CTA": 'L',
	"CTG": 'L',
	"TTA": 'L',
	"TTG": 'L',
	"GTT": 'V',
	"GTC": 'V',
	"GTA": 'V',
	"GTG": 'V',
	"TTT": 'F',
	"TTC": 'F',
	"ATG": 'M',
	"TGT": 'C',
	"TGC": 'C',
	"GCT": 'A',
	"GCC": 'A',
	"GCA": 'A',
	"GCG": 'A',
	"GGT": 'G',
	"GGC": 'G',
	"GGA": 'G',
	"GGG": 'G',
	"CCT": 'P',
	"CCC": 'P',
	"CCA": 'P',
	"CCG": 'P',
	"ACT": 'T',
	"ACC": 'T',
	"ACA": 'T',
	"ACG": 'T',
	"TCT": 'S',
	"TCC": 'S',
	"TCA": 'S',
	"TCG": 'S',
	"AGT": 'S',
	"AGC": 'S',
	"TAT": 'Y',
	"TAC": 'Y',
	"TGG": 'W',
	"CAA": 'Q',
	"CAG": 'Q',
	"AAT": 'N',
	"AAC": 'N',
	"CAT": 'H',
	"CAC": 'H',
	"GAA": 'E',
	"GAG": 'E',
	"GAT": 'D',
	"GAC": 'D',
	"AAA": 'K',
	"AAG": 'K',
	"CGT": 'R',
	"CGC": 'R',
	"CGA": 'R',
	"CGG": 'R',
	"AGA": 'R',
	"AGG": 'R',
	"TAA": '*',
	"TAG": '*',
	"TGA": '*',
}
