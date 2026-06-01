package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseFasta (data string) []string {
	sequences := []string{}
	currSeq := ""

	lines := strings.Split(strings.TrimSpace(data),"\n")

	for _, line := range lines {
		line = strings.TrimSpace(line)

		if strings.HasPrefix(line,">") {
			if currSeq != "" {
				sequences = append(sequences, currSeq)
				currSeq = ""
			}
		} else {
			currSeq += line
		}
	}

	if currSeq != "" {
		sequences = append(sequences, currSeq)
	}

	return sequences
}

func buildProfileMatrix(sequences []string) map[string][]int {
	seqLength := len(sequences[0])

	profile := map[string][]int {
		"A": make([] int, seqLength),
		"C": make([] int, seqLength),
		"G": make([] int, seqLength),
		"T": make([] int, seqLength),
	}

	for _, seq := range sequences {
		for idx, nuc := range seq {
			key := string(nuc)
			profile[key][idx]++
		}
	}

	return  profile
}

func buildConsensusString(profile map[string][]int) string {
	var consensus strings.Builder

	nucleotides := []string{"A","C","G","T",}
	seqLength := len(profile["A"])

	for idx := 0; idx < seqLength; idx++ {
		mostCommonNuc := "A"
		highestCount := profile["A"][idx]

		for _, nuc := range nucleotides {
			count := profile[nuc][idx]

			if count > highestCount {
				highestCount = count
				mostCommonNuc = nuc
			}
		}

		consensus.WriteString(mostCommonNuc)
	}

	return consensus.String()
}

func formatCounts(counts []int) string {
	res := []string{}

	for _, count := range counts {
		res = append(res, strconv.Itoa(count))
	}

	return strings.Join(res," ")
}

func main() {
	data, err := os.ReadFile("../input.txt")
	if err != nil {
		panic(err)
	}

	seq := parseFasta(string(data))
	profile := buildProfileMatrix(seq)
	consensus := buildConsensusString(profile)

	fmt.Println(consensus)

	for _, nuc := range []string{"A","C","G","T",}{
		fmt.Printf("%s: %s\n",nuc, formatCounts(profile[nuc]))
	}
}