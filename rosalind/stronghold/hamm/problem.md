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

## Supplementary Information

## Evolution as a Sequence of Mistakes
A mutation is simply a mistake that occurs during the creation or copying of a nucleic acid, in particular DNA. Because nucleic acids are vital to cellular functions, mutations tend to cause a ripple effect throughout the cell. Although mutations are technically mistakes, a very rare mutation may equip the cell with a beneficial attribute. In fact, the macro effects of evolution are attributable by the accumulated result of beneficial microscopic mutations over many generations.

The simplest and most common type of nucleic acid mutation is a point mutation, which replaces one base with another at a single nucleotide. In the case of DNA, a point mutation must change the complementary base accordingly; see Figure 1.

Two DNA strands taken from different organism or species genomes are homologous if they share a recent ancestor; thus, counting the number of bases at which homologous strands differ provides us with the minimum number of point mutations that could have occurred on the evolutionary path between the two strands.

We are interested in minimizing the number of (point) mutations separating two species because of the biological principle of parsimony, which demands that evolutionary histories should be as simply explained as possible.