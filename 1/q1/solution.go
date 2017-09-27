package main

import (
	"../utils"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	input := utils.ReadFASTAFile(file)

	var transcripted []string
	for _, rune := range input.Body {
		char := string(rune)
		if char == "T" {
			char = "U"
		}

		transcripted = append(transcripted, char)
	}

	fmt.Println(strings.Join(transcripted, ""))
}
