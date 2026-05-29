package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func findSubstringLocations(s string, t string) []int {
	locations := []int{}

	for i := 0; i <= len(s) - len(t); i++ {
			if s[i:i+len(t)] == t {
				locations = append(locations, i + 1)
		}
	}

	return locations
}

func formatLocations(locations []int) string {
	result := []string{}

	for _, location := range locations {
		result = append(result, strconv.Itoa(location))
	}

	return strings.Join(result," ")
}

func main() {
	data, err := os.ReadFile("../input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(strings.TrimSpace(string(data)), "\n")

	s := strings.TrimSpace(lines[0])
	t := strings.TrimSpace(lines[1])

	locations := findSubstringLocations(s,t)
	res := formatLocations(locations)

	fmt.Println(res)
}