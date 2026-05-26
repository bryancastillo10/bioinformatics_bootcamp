package main

import (
	"fmt"
	"os"
	"strings"
)

func transcribeDNAToRNA(dna string) string {
	return strings.ReplaceAll(dna, "T", "U")
}

func main() {
	data,err := os.ReadFile("../input.txt")

	if err != nil {
		panic(err)
	}

	dna := strings.TrimSpace(string(data))
	res := transcribeDNAToRNA(dna)

	fmt.Println(res)
}