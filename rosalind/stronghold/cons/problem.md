# CONS — Consensus and Profile

## Problem Description

A matrix is a rectangular table of values divided into rows and columns.

An `m × n` matrix has:

- `m` rows
- `n` columns

Given a matrix `A`, we write `A[i, j]` to represent the value found at row `i` and column `j`.

In this problem, we are given a collection of DNA strings. All DNA strings have the same length `n`.

From these DNA strings, we can build a **profile matrix**.

The profile matrix is a `4 × n` matrix that counts how many times each nucleotide appears at every position.

The four rows represent:

```text
A
C
G
T
```

For each column position `j`:

- `A[j]` counts how many times `A` appears at position `j`
- `C[j]` counts how many times `C` appears at position `j`
- `G[j]` counts how many times `G` appears at position `j`
- `T[j]` counts how many times `T` appears at position `j`

A **consensus string** is formed by choosing the most common nucleotide at each position.

If there is more than one most common nucleotide at a position, any one of them may be chosen.

## Given

A collection of at most 10 DNA strings in FASTA format.

Each DNA string:

- has equal length
- has length at most 1 kbp

## Return

A consensus string and the profile matrix for the collection.

## Sample Input

```text
>Rosalind_1
ATCCAGCT
>Rosalind_2
GGGCAACT
>Rosalind_3
ATGGATCT
>Rosalind_4
AAGCAACC
>Rosalind_5
TTGGAACT
>Rosalind_6
ATGCCATT
>Rosalind_7
ATGGCACT
```

## Sample Output

```text
ATGCAACT
A: 5 1 0 0 5 5 0 0
C: 0 0 1 4 2 0 6 1
G: 1 1 6 3 0 1 0 0
T: 1 5 0 0 0 1 1 6
```

## Explanation

At each position, we count the number of times each nucleotide appears.

For example, at position 1:

```text
A appears 5 times
C appears 0 times
G appears 1 time
T appears 1 time
```

So the first character of the consensus string is:

```text
A
```

Repeating this process for every position gives the consensus string:

```text
ATGCAACT
```