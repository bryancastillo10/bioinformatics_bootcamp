# PROT — Translating RNA into Protein

## Problem Description

The 20 commonly occurring amino acids are abbreviated using 20 letters from the English alphabet.

The letters used are all English alphabet letters except:

```text
B, J, O, U, X, and Z
```

Protein strings are constructed from these 20 amino acid symbols.

In bioinformatics, the term **genetic string** can refer to:

- DNA strings
- RNA strings
- Protein strings

The **RNA codon table** describes how RNA codons are translated into amino acids.

A codon is a sequence of three RNA bases. Each codon corresponds to either:

1. One amino acid
2. A stop signal

For example:

| RNA Codon | Amino Acid |
| --- | --- |
| `AUG` | `M` |
| `UUU` | `F` |
| `GCU` | `A` |
| `UAA` | Stop |

To translate an RNA string into a protein string, we read the RNA sequence three bases at a time and convert each codon into its corresponding amino acid.

Translation stops when a stop codon is reached.

## Given

An RNA string `s` corresponding to a strand of mRNA.

The length of `s` is at most 10 kbp.

## Return

The protein string encoded by `s`.

## Sample Input

```text
AUGGCCAUGGCGCCCAGAACUGAGAUCAAUAGUACCCGUAUUAACGGGUGA
```

## Sample Output

```text
MAMAPRTEINSTRING
```

## Explanation

The RNA string is read in codons:

```text
AUG GCC AUG GCG CCC AGA ACU GAG AUC AAU AGU ACC CGU AUU AAC GGG UGA
```

Using the RNA codon table:

```text
AUG → M
GCC → A
AUG → M
GCG → A
CCC → P
AGA → R
ACU → T
GAG → E
AUC → I
AAU → N
AGU → S
ACC → T
CGU → R
AUU → I
AAC → N
GGG → G
UGA → Stop
```

Therefore, the translated protein string is:

```text
MAMAPRTEINSTRING
```