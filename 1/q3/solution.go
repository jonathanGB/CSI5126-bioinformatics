// Jonathan Guillotte-Blouin - 7900293
package main

import (
	"fmt"
	"log"
	"os"

	"../utils"
)

/*
	All six reading frames
*/
func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// get a FASTA struct representing the input file
	input := utils.ReadFASTAFile(file)

	// read frames 1, 2, and 3 in direction 5'3'
	for i := 0; i < 3; i++ {
		fmt.Println(input.ReadFrame("5'3'", i))
	}

	input.ReverseComplement()

	// read frames 1, 2, and 3 in direction 3'5'
	for i := 0; i < 3; i++ {
		fmt.Println(input.ReadFrame("3'5'", i))
	}
}
