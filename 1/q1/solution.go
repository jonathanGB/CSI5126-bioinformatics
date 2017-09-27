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

	// get a FASTA struct representing the input file, then transcript it
	input := utils.ReadFASTAFile(file)
	input.Transcript()

	fmt.Println(input.Body)
}
