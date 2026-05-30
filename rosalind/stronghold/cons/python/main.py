def parse_fasta(data: str) -> list[str]:
    seq = []
    curr_seq = ""

    for line in data.strip().splitlines():
        line = line.strip()

        if line.startswith(">"):
            if curr_seq:
                seq.append(curr_seq)
                curr_seq = ""
        else:
            curr_seq += line
    if curr_seq:
        seq.append(curr_seq)

    return seq


def build_mat(sequences: list[str]) -> dict[str, list[int]]:
    seq_length = len(sequences[0])

    profile = {
        "A": [0] * seq_length,
        "C": [0] * seq_length,
        "G": [0] * seq_length,
        "T": [0] * seq_length,
    }

    for seq in sequences:
        for idx, nuc in enumerate(seq):
            profile[nuc][idx] += 1

    return profile


def build_cons_string(profile: dict[str, list[int]]) -> str:
    cons = ""
    nucleotides = ["A", "C", "G", "T"]
    seq_length = len(profile["A"])

    for idx in range(seq_length):
        most_common_nuc = "A"
        highest_count = profile["A"][idx]

        for nuc in nucleotides:
            count = profile[nuc][idx]

            if count > highest_count:
                highest_count = count
                most_common_nuc = nuc

        cons += most_common_nuc

    return cons


def main():
    with open("../input.txt", "r") as file:
        data = file.read()

    seq = parse_fasta(data)
    profile = build_mat(seq)
    cons = build_cons_string(profile)

    print(cons)

    for nuc in ["A", "G", "C", "T"]:
        counts = " ".join(map(str, profile[nuc]))
        print(f"{nuc}:{counts}")


if __name__ == "__main__":
    main()
