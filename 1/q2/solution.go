// Jonathan Guillotte-Blouin - 7900293
package main

import (
	"../utils"
	"fmt"
	"log"
	"os"
)

/*
	Reverse complement
*/
func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// get a FASTA struct representing the input file
	input := utils.ReadFASTAFile(file)

	// map of complement nucleotides
	complements := map[rune]rune{
		'A': 'T',
		'T': 'A',
		'C': 'G',
		'G': 'C',
	}

	inputLen := len(input.Body)
	transcripted := make([]rune, inputLen)
	for i, char := range input.Body {
		// map complementary nucleotides
		complement := complements[char]

		// put the complement in reverse order of input
		transcripted[inputLen-1-i] = complement
	}

	fmt.Println(string(transcripted))
}
