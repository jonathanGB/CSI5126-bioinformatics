// Jonathan Guillotte-Blouin - 7900293
package utils

import (
	"bufio"
	"io"
	"strings"
)

// define a FASTA struct
type FASTA struct {
	Header string
	Body   string
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
