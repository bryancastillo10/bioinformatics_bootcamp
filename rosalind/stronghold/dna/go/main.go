package main

import (
	"fmt"
	"os"
	"strings"
)

func countDNANucleotides(seq string) map[string]int {
	a := strings.Count(seq, "A")
	c := strings.Count(seq, "C")
	g := strings.Count(seq, "G")
	t := strings.Count(seq, "T")

	return map[string]int{
		"A": a,
		"C": c,
		"G": g,
		"T": t,
	}
}

func formatCounts(counts map[rune]int) string {
	return fmt.Sprintf("%d %d %d %d", counts['A'], counts['C'], counts['G'], counts['T'])
}

func main() {
	data, err := os.ReadFile("../input.txt")
	if err != nil {
		panic(err)
	}

	seq := strings.TrimSpace(string(data))
	res := countDNANucleotides(seq)
	fmt.Println(res)
}