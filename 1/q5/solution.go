package main

import (
	"fmt"
	"strings"

	"../utils"
)

func main() {
	// part 1: print the associated codons
	fmt.Println("Part 1: Print the associated codons:\n")

	// make a map of aminoAcids to all the codons which generate it
	aminoAcidsToCodons := map[rune][]string{}
	for codon, aminoAcid := range utils.CodonToAminoAcid {
		aminoAcidsToCodons[aminoAcid] = append(aminoAcidsToCodons[aminoAcid], codon)
	}

	for aminoacid, codons := range aminoAcidsToCodons {
		fmt.Printf("%c: [%s]\n", aminoacid, strings.Join(codons, ", "))
	}

	fmt.Println("\n\nPart 2: Print the minimum hamming distances:\n")

	// get a slice of all the keys (e.g. only the amino acids)
	aminoAcids := []rune{}
	for aminoAcid := range aminoAcidsToCodons {
		aminoAcids = append(aminoAcids, aminoAcid)
	}

	// loop through all the pair of amino acids
	for i := 0; i < len(aminoAcids)-1; i++ {
		for j := i + 1; j < len(aminoAcids); j++ {
			// compute the minimum distance for that pair of amino acids
			minDistance := 4
			var minCodonA, minCodonB string
			aminoAcidA, aminoAcidB := aminoAcids[i], aminoAcids[j]
			codonsA, codonsB := aminoAcidsToCodons[aminoAcidA], aminoAcidsToCodons[aminoAcidB]

			// compare every pair of amino acids of both amino acids
			for _, codonA := range codonsA {
				for _, codonB := range codonsB {
					// if this pair of amino acids has the smallest hamming distance, remember it
					if dist := hammingDistance(codonA, codonB); dist < minDistance {
						minDistance = dist
						minCodonA, minCodonB = codonA, codonB
					}
				}
			}

			fmt.Printf("Pair %c-%c: Min distance: %d (%s-%s)\n", aminoAcidA, aminoAcidB, minDistance, minCodonA, minCodonB)
		}
	}

	fmt.Println("\n\nPart 3: Print all silent mutations:\n")

	nucleotides := []byte{'A', 'C', 'T', 'G'}
	for i := 0; i < 3; i++ {
		silentMutationsCtr := 0

		for codon, aminoAcid := range utils.CodonToAminoAcid {
			mutatedNucleotide := codon[i]

			for _, newNucleotide := range nucleotides {
				if newNucleotide == mutatedNucleotide {
					continue
				}

				mutatedCodonSlice := make([]byte, 3)
				for j := 0; j < 3; j++ {
					if i == j {
						mutatedCodonSlice[j] = newNucleotide
					} else {
						mutatedCodonSlice[j] = codon[j]
					}
				}

				mutatedCodon := string(mutatedCodonSlice)
				if aminoAcidMatched := utils.CodonToAminoAcid[mutatedCodon]; aminoAcidMatched == aminoAcid {
					silentMutationsCtr++
				}
			}
		}

		fmt.Printf("Number of silent mutations at index %d: %d (out of 192 possibilities)\n", i, silentMutationsCtr)
	}
}

func hammingDistance(a, b string) (ctr int) {
	if a[0] != b[0] {
		ctr++
	}
	if a[1] != b[1] {
		ctr++
	}
	if a[2] != b[2] {
		ctr++
	}

	return
}
