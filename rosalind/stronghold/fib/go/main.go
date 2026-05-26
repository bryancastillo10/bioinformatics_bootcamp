package main

import (
	"fmt"
	"os"
	"strings"
)

func pairings(n int, k int) int {
	if n <= 2 {
		return 1
	}

	prev := 1
	curr := 1

	for month := 3; month <= n; month++ {
		nextVal := curr + (k * prev)
		prev = curr		
		curr = nextVal
	}

	return  curr
}

func main() {
	data, err := os.ReadFile("../input.txt")
	if err != nil {
		panic(err)
	}

	input := strings.TrimSpace(string(data))

	var n,k int
	fmt.Sscanf(input, "%d %d", &n, &k)

	res := pairings(n,k)

	fmt.Println(res)
}