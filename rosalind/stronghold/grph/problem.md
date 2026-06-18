# Overlap Graphs

A graph whose nodes have all been labeled can be represented by an **adjacency list**, in which each row of the list contains the two node labels corresponding to a unique edge.

A **directed graph** (or **digraph**) is a graph containing directed edges, each of which has an orientation. That is, a directed edge is represented by an arrow instead of a line segment; the starting and ending nodes of an edge form its **tail** and **head**, respectively. The directed edge with tail `v` and head `w` is represented by `(v, w)`.

A **directed loop** is a directed edge of the form `(v, v)`.

For a collection of strings and a positive integer `k`, the **overlap graph** for the strings is a directed graph `O_k` in which each string is represented by a node, and string `s` is connected to string `t` with a directed edge when:

- the length `k` suffix of `s` matches the length `k` prefix of `t`
- and `s ≠ t`

The condition `s ≠ t` prevents directed loops in the overlap graph, although directed cycles may still occur.

## Given

A collection of DNA strings in FASTA format having total length at most 10 kbp.

## Return

The adjacency list corresponding to `O_3`.

## Notes

- Each DNA string is a node in the graph.
- Draw a directed edge from string `s` to string `t` if the last 3 characters of `s` match the first 3 characters of `t`.
- Do not connect a string to itself.


### Example Input
```text
>Rosalind_0498
AAATAAA

>Rosalind_2391
AAATTTT

>Rosalind_2323
TTTTCCC

>Rosalind_0442
AAATCCC

>Rosalind_5013
GGGTGGG
```

### Example Output
```text 

```