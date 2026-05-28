# IPRB — Mendel's First Law

## Problem Description

Probability is the mathematical study of randomly occurring phenomena. We can model a random phenomenon using a **random variable**, which is a variable that can take different possible outcomes depending on the result of an underlying random process.

For example, suppose we have a bag containing:

- 3 red balls
- 2 blue balls

If we let `X` represent the color of a randomly drawn ball, then:

```text
Pr(X = red) = 3/5
Pr(X = blue) = 2/5
```

Random variables can also be combined to form new random variables.

For example, let `Y` represent the color of a second ball drawn from the same bag without replacing the first ball. The probability of `Y` being red or blue depends on the outcome of the first draw.

To represent all possible outcomes, we can use a **probability tree diagram**. The probability of each final outcome is found by multiplying the probabilities along the path from the start of the tree to the final outcome.

An **event** is a collection of outcomes. Since the outcomes are distinct, the probability of an event is the sum of the probabilities of all outcomes belonging to that event.

In this problem, we are given a population of organisms with three possible genotypes:

| Genotype | Meaning |
| --- | --- |
| `AA` | Homozygous dominant |
| `Aa` | Heterozygous |
| `aa` | Homozygous recessive |

We need to find the probability that two randomly selected organisms will produce an offspring with at least one dominant allele.

An offspring with genotype `AA` or `Aa` displays the dominant phenotype.

An offspring with genotype `aa` displays the recessive phenotype.

## Given

Three positive integers:

```text
k m n
```

Where:

- `k` is the number of homozygous dominant organisms, `AA`
- `m` is the number of heterozygous organisms, `Aa`
- `n` is the number of homozygous recessive organisms, `aa`

The total population size is:

```text
k + m + n
```

## Return

The probability that two randomly selected mating organisms will produce an individual possessing a dominant allele.

## Sample Input

```text
2 2 2
```

## Sample Output

```text
0.78333
```

## Explanation

For the sample input:

```text
k = 2
m = 2
n = 2
```

The population contains:

```text
2 AA organisms
2 Aa organisms
2 aa organisms
```

We can calculate the probability of producing a dominant phenotype by subtracting the probability of producing a recessive phenotype from `1`.

Only the following mating pairs can produce recessive offspring:

| Mating Pair | Probability of Recessive Offspring |
| --- | --- |
| `Aa × Aa` | `1/4` |
| `Aa × aa` | `1/2` |
| `aa × Aa` | `1/2` |
| `aa × aa` | `1` |

Therefore:

```text
Probability of dominant phenotype = 1 - Probability of recessive phenotype
```