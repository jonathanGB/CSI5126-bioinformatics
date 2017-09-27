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
	input.ReverseComplement()

	fmt.Println(input.Body)
}
