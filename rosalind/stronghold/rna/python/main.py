def transcribe_dna_to_rna(dna: str) -> str:
    return dna.replace("T", "U")


def main():
    with open("../input.txt", "r") as file:
        dna = file.read().strip()

    rna = transcribe_dna_to_rna(dna)
    print(rna)


if __name__ == "__main__":
    main()