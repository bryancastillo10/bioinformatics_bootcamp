# HAMM — Counting Point Mutations

## Problem

Given two strings `s` and `t` of equal length, the **Hamming distance** between `s` and `t`, denoted as $d_H(s, t)$, is the number of corresponding symbols that are different.

In DNA strings, this means comparing two DNA sequences position by position and counting how many nucleotides differ.

For example:

```text
s = GAGCCTACTAACGGGAT
t = CATCGTAATGACGGCCT

Comparing each position:

G A G C C T A C T A A C G G G A T
C A T C G T A A T G A C G G C C T
^   ^   ^     ^   ^         ^ ^

There are 7 positions where the two strings differ.

```

Therefore:

 $d_H(s,t)$ =  **7**

Given: Two DNA strings s and t of equal length, not exceeding 1 kbp.

Return: The Hamming distance $d_H(s, t)$

## Files
- `input.txt` - sample/input DNA string
- `output.txt` - expected output
- `python/main.py` - Python solution
- `go/main.go` - Go solution

## Run

### Python
cd python

python main.py

### Go
cd go

go run main.go