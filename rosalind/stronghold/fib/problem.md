# FIB — Rabbits and Recurrence Relations

## Problem Description

A sequence is an ordered collection of objects, usually numbers, which are allowed to repeat. Sequences can be finite or infinite.

For example:

- A finite sequence: $(\pi, \sqrt{-2}, 0, \pi)$
- An infinite sequence of odd numbers: $(1, 3, 5, 7, 9, \dots)$

We use the notation $a_n$ to represent the $n$-th term of a sequence.

A recurrence relation is a way of defining the terms of a sequence based on the values of previous terms.

In the case of Fibonacci's rabbits, any given month will contain:

1. The rabbits that were alive during the previous month
2. The new offspring produced during the current month

A key observation is that the number of offspring in any month is equal to the number of rabbit pairs that were alive two months prior.

If $F_n$ represents the number of rabbit pairs alive after the $n$-th month, then the standard Fibonacci sequence is defined by:

$$
F_n = F_{n-1} + F_{n-2}
$$

with the starting values:

$$
F_1 = 1
$$

$$
F_2 = 1
$$

This problem modifies the original Fibonacci recurrence by allowing each reproduction-age rabbit pair to produce $k$ new rabbit pairs instead of only one.

Therefore, the recurrence relation becomes:

$$
F_n = F_{n-1} + kF_{n-2}
$$

## Given

Positive integers:

$$
n \leq 40
$$

and

$$
k \leq 5
$$

Where:

- $n$ is the number of months
- $k$ is the number of rabbit pairs produced by each reproduction-age rabbit pair

## Return

The total number of rabbit pairs that will be present after n months, if we begin with 1 pair and in each generation, every pair of reproduction-age rabbits produces a litter of k rabbit pairs (instead of only 1 pair).

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