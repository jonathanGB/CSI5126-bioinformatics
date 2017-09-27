package main

import (
	"../utils"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	input := utils.ReadFASTAFile(file)

	var transcripted []rune
	for _, char := range input.Body {
		if char == 'T' {
			char = 'U'
		}

		transcripted = append(transcripted, char)
	}

	fmt.Println(string(transcripted))
}
