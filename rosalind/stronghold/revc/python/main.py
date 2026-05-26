def reverse_complement(dna: str) -> str:
    watson_crick_pairing = {"A": "T", "T": "A", "G": "C", "C": "G"}

    rev_dna = dna[::-1]
    return "".join(watson_crick_pairing[base] for base in rev_dna)


def main():
    with open("../input.txt", "r") as file:
        seq = file.read().strip()

    res = reverse_complement(seq)
    print(res)


if __name__ == "__main__":
    main()
