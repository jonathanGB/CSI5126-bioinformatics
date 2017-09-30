// Jonathan Guillotte-Blouin - 7900293
package main

import (
	"fmt"
	"log"
	"os"

	"../utils"
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

	// get a FASTA struct representing the input file, then reverse-complement it
	input := utils.ReadFASTAFile(file)
	input.ReverseComplement()

	fmt.Println(input.Body)
}
