package main

import (
	"fmt"
	"os"
	"strings"
)

func watsonCrickPairing(base byte) byte {
	switch base{
		case 'A':
			return 'T'
		case 'T':
			return 'A'
		case 'G':
			return 'C'
		case 'C':
			return 'G'
		default:
			return base
	}
}

func reverseComplement(dna string) string {
	res := make([]byte, len(dna))

	for i := 0; i < len(dna); i++ {
		res[i] = watsonCrickPairing(dna[len(dna)-1-i])
	}

	return string(res)
}

func main() {
	data,err := os.ReadFile("../input.txt")
	if err != nil {
		panic(err)
	}

	seq := strings.TrimSpace(string(data))
	res := reverseComplement(seq)

	fmt.Println(res)
}