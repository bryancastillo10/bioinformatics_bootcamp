from codon import RNA_CODON_TABLE


def translate_rna_to_protein(rna: str) -> str:
    protein = ""

    for i in range(0, len(rna), 3):
        codon = rna[i : i + 3]
        amino_acid = RNA_CODON_TABLE[codon]

        if amino_acid == "STOP":
            break

        protein += amino_acid

    return protein


def main():
    with open("../input.txt", "r") as file:
        rna = file.read().strip()

    res = translate_rna_to_protein(rna)
    print(res)


if __name__ == "__main__":
    main()
