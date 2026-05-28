package main

import (
	"fmt"
	"os"
	"strings"
)

func dominantPhenotypeProbability(k int, m int, n int) float64 {
	total := float64(k + m + n)

	kfloat := float64(k)
	mfloat := float64(m)
	nfloat := float64(n)

	_ = kfloat

	// Probability of selecting two heterozygous organisms: Aa x Aa
	hetHet := (mfloat/total) * ((mfloat -1 )/(total -1))

	// Probability of selecting heterozygous then recessive: Aa x aa
	hetRec := (mfloat/ total) * (nfloat/ (total-1))

	// Probability of selecting recessive then heterozygous: aa x Aa
	recHet := (nfloat/total) * (mfloat / (total -1))

	// Probability of selecting two recessive organisms: aa x aa
	recRec := (nfloat/ total) * ((nfloat -1 )/ (total -1))

	recessiveProbability := (hetHet * 0.25) +
		(hetRec * 0.50) + (recHet * 0.50) + (recRec * 1.00)

	return 1- recessiveProbability
}


func main() {
	data, err := os.ReadFile("../input.txt")
	if err != nil {
		panic(err)
	}

	input := strings.TrimSpace(string(data))

	var k,m,n int
	fmt.Sscanf(input, "%d %d %d", &k, &m, &n)

	res := dominantPhenotypeProbability(k,m,n)
	fmt.Printf("%.5f\n",res)
}