package main

import "strings"

func parseFasta (data string) map[string]string {
	records := make(map[string]string)
	currID := ""

	lines := strings.Split(strings.TrimSpace(data),"\n")

	for _,line := range lines {
		line = strings.TrimSpace(line)

		if strings.HasPrefix(line,">") {
			currID = strings.TrimPrefix(line,">")
			records[currID] = ""
		} else {
			records[currID] += line
		}
	}

	return  records
}

func buildProfileMatrix(sequences []string) map[string][] int {
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