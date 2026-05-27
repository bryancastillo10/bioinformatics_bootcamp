package main

import (
	"fmt"
	"os"
	"strings"
)

func parseFasta(data string) map[string]string {
	records := make(map[string]string)
	currID := ""

	lines := strings.Split(strings.TrimSpace(data), "\n")

	for _, line := range lines {
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


func calcGCContent (dna string) float64 {
	gcCount := 0

	for _, base := range dna {
		if base == 'G' || base == 'C' {
			gcCount++
		}
	}

	return (float64(gcCount)/ float64(len(dna))) * 100
}

func findHighestGCContent(records map[string]string) (string, float64) {
	highestID := ""
	highestGC := 0.0

	for recordID, dna := range records {
		gcContent := calcGCContent(dna)

		if gcContent > highestGC {
			highestID = recordID
			highestGC = gcContent
		}
	}

	return  highestID, highestGC
}

func main() {
	data, err := os.ReadFile("../input.txt")
	if err != nil {
		panic(err)
	}

	records := parseFasta(string(data))
	highestID, highestGC := findHighestGCContent(records)

	fmt.Println(highestID)
	fmt.Printf("%.6f\n", highestGC)
}