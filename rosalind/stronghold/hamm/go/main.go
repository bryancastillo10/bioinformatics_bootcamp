package main

import (
	"fmt"
	"os"
	"strings"
)

func hammingDistance(s string, t string) int {
	distance := 0

	for i := 0; i < len(s); i++ {
		if s[i] != t[i] {
			distance++
		}
	}

	return distance
}

func main() {
	data, err := os.ReadFile("../input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(strings.TrimSpace(string(data)),"\n")

	s := strings.TrimSpace(lines[0])
	t := strings.TrimSpace(lines[1])

	res := hammingDistance(s,t)
	fmt.Println(res)
}