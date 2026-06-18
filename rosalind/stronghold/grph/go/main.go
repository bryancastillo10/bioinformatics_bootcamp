package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func parseFASTA(text string) map[string]string {
	records := make(map[string]string)

	var label string
	var sequenceBuilder strings.Builder

	scanner := bufio.NewScanner(strings.NewReader(text))
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		if strings.HasPrefix(line, ">") {
			if label != "" {
				records[label] = sequenceBuilder.String()
			}
			label = line[1:]
			sequenceBuilder.Reset()
		} else {
			sequenceBuilder.WriteString(line)
		}
	}

	if label != "" {
		records[label] = sequenceBuilder.String()
	}

	return records
}

func buildOverlapGraph(records map[string]string, k int) [][2]string {
	var adjacency [][2]string

	for sLabel, sSeq := range records {
		suffix := sSeq[len(sSeq)-k:]

		for tLabel, tSeq := range records {
			if sLabel != tLabel && suffix == tSeq[:k] {
				adjacency = append(adjacency, [2]string{sLabel, tLabel})
			}
		}
	}

	return adjacency
}

func main() {
	data, err := os.ReadFile("../input.txt")
	if err != nil {
		panic(err)
	}

	records := parseFASTA(string(data))
	adjacency := buildOverlapGraph(records, 3)

	for _, edge := range adjacency {
		fmt.Println(edge[0], edge[1])
	}
}