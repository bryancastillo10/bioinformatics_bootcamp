# Rosalind Practice

This folder contains my Rosalind bioinformatics practice solutions and small reusable utilities.

## Goals
- Practice core bioinformatics problem solving
- Solve problems in both Python and Go
- Build a clean and consistent coding structure
- Gradually extract reusable helpers for sequence parsing and file handling

## Structure

```text
rosalind/
├── README.md
├── requirements.txt
├── go.mod
├── shared/
│   ├── python/
│   └── go/
└── stronghold/
    ├── dna/
    │   ├── input.txt
    │   ├── python/
    │   │   └── main.py
    │   └── go/
    │       └── main.go
    └── ...
```

## Environment and Module Setup

### Python virtual environment
This project uses a local Python virtual environment for package isolation.

```text
rosalind-problems
```

### Go Module
This project uses a single Go module for shared helpers and problem solutions

```text
module rosalind-problems
```