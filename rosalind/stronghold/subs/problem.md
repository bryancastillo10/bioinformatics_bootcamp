# SUBS — Finding a Motif in DNA

## Problem 
Given two strings `s` and `t`, `t` is a substring of `s` if `t` appears inside `s` as a contiguous collection of symbols.

This means that all characters of `t` must appear together and in the same order within `s`.

For example, if:

```text
s = GATATATGCATATACTT
t = ATAT
```

Then `t` appears multiple times inside `s`.

The position of a symbol in a string is counted starting from `1`, not `0`.

For example, in the RNA string:

```text
AUGCUUCAGAAAGGUCUUACG
```

The positions of all occurrences of `U` are:

```text
2, 5, 6, 15, 17, 18
```

A substring can be represented as:

```text
s[j:k]
```

where `j` is the starting position and `k` is the ending position.

The location of a substring is its beginning position `j`.

In this problem, we need to find all starting positions where `t` appears as a substring of `s`.

Overlapping occurrences should also be counted.

## Given

Two DNA strings:

```text
s
t
```

Each string has length at most 1 kbp.

## Return

All locations of `t` as a substring of `s`.

The positions should be returned using 1-based indexing.

## Sample Input

```text
GATATATGCATATACTT
ATAT
```

## Sample Output

```text
2 4 10
```

## Explanation

For the sample input:

```text
s = GATATATGCATATACTT
t = ATAT
```

The substring `ATAT` appears starting at positions:

```text
2, 4, 10
```

Therefore, the answer is:

```text
2 4 10
```

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
Finding the same interval of DNA in the genomes of two different organisms (often taken from different species) is highly suggestive that the interval has the same function in both organisms.

We define a motif as such a commonly shared interval of DNA. A common task in molecular biology is to search an organism's genome for a known motif.

The situation is complicated by the fact that genomes are riddled with intervals of DNA that occur multiple times (possibly with slight modifications), called repeats. These repeats occur far more often than would be dictated by random chance, indicating that genomes are anything but random and in fact illustrate that the language of DNA must be very powerful (compare with the frequent reuse of common words in any human language).

The most common repeat in humans is the Alu repeat, which is approximately 300 bp long and recurs around a million times throughout every human genome (see Figure 1). However, Alu has not been found to serve a positive purpose, and appears in fact to be parasitic: when a new Alu repeat is inserted into a genome, it frequently causes genetic disorders.