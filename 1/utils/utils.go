package utils

import (
	"bufio"
	"io"
	"strings"
)

type FASTA struct {
	Header string
	Body   string
}

func ReadFASTAFile(file io.Reader) *FASTA {
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	header := scanner.Text()

	body := []string{}
	for scanner.Scan() {
		body = append(body, scanner.Text())
	}

	return &FASTA{
		header,
		strings.Join(body, ""),
	}
}
