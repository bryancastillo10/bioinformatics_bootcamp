# Mortal Fibonacci Rabbits

Recall the definition of the Fibonacci numbers from **“Rabbits and Recurrence Relations”**, which followed the recurrence relation:

\[
F_n = F_{n-1} + F_{n-2}
\]

and assumed that each pair of rabbits reaches maturity in one month and produces a single pair of offspring each subsequent month.

Our aim is to modify this recurrence relation to obtain a dynamic programming solution in the case that all rabbits die after a fixed number of months.

## Given

Two positive integers:

- \( n \leq 100 \)
- \( m \leq 20 \)

where:

- \( n \) is the number of months
- \( m \) is the lifespan of each rabbit pair in months

## Return

The total number of rabbit pairs that will remain after the \( n \)-th month if all rabbits live for \( m \) months.

## Example

### Input
```text
6 3
```

### Output
```
4 
```