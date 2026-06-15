package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func mortalFibonacciRabbits(n, m int) int64 {
	ages := make([]int64, m)
	ages[0] = 1

	for month := 1; month < n; month++ {
		var newborns int64
		for i := 1; i < m; i++ {
			newborns += ages[i]
		}

		newAges := make([]int64, m)
		newAges[0] = newborns
		copy(newAges[1:], ages[:m-1])
		ages = newAges
	}

	var total int64
	for _, count := range ages{
		total += count
	}

	return total
}

func main() {
	data,err := os.ReadFile("../input.txt")
	if err != nil {
		panic(err)
	}

	parts := strings.Fields(strings.TrimSpace(string(data)))
	n, err := strconv.Atoi(parts[0])
	if err != nil {
		panic(err)
	}

	m, err := strconv.Atoi(parts[1])
	if err != nil {
		panic(err)
	}

	res := mortalFibonacciRabbits(n,m)
	fmt.Println(res)
}