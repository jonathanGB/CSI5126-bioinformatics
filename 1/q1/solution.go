// Jonathan Guillotte-Blouin - 7900293
package main

import (
	"../utils"
	"fmt"
	"log"
	"os"
)

/*
	Transcription
*/
func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// get a FASTA struct representing the input file
	input := utils.ReadFASTAFile(file)

	transcripted := make([]rune, len(input.Body))
	for i, char := range input.Body {
		// for every 'T', change it to a 'U'
		if char == 'T' {
			char = 'U'
		}

		transcripted[i] = char
	}

	fmt.Println(string(transcripted))
}
