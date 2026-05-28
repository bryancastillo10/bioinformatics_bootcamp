package main

import (
	"fmt"
	"os"
	"strings"

	"rosalind-problems/stronghold/prot/go/codon"
)

func translateRNAToProtein(rna string) string {
	var protein strings.Builder

	for i := 0; i < len(rna); i += 3 {
		codonText := rna[i : i + 3]
		aminoAcid := codon.RnaCodonTable[codonText]

			if aminoAcid == "STOP" {
		break
		}

		protein.WriteString(aminoAcid)
	}

	return protein.String()
}

func main() {
	data, err := os.ReadFile("../input.txt")

	if err != nil {
		panic(err)
	}	

	rna := strings.TrimSpace(string(data))
	res := translateRNAToProtein(rna)

	fmt.Println(res)
}