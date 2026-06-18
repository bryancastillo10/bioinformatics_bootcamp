def parse_fasta(text: str) -> dict[str, str]:
    records = {}
    label = ""
    sequence_parts = []

    for line in text.strip().splitlines():
        line = line.strip()
        if not line:
            continue

        if line.startswith(">"):
            if label:
                records[label] = "".join(sequence_parts)
            label = line[1:]
            sequence_parts = []
        else:
            sequence_parts.append(line)

    if label:
        records[label] = "".join(sequence_parts)

    return records


def build_overlap_graph(records: dict[str, str], k: int = 3) -> list[tuple[str, str]]:
    adjacency = []

    for s_label, s_seq in records.items():
        suffix = s_seq[-k:]

        for t_label, t_seq in records.items():
            if s_label != t_label and suffix == t_seq[:k]:
                adjacency.append((s_label, t_label))

    return adjacency


def main():
    with open("../input.txt", "r") as file:
        data = file.read()

    records = parse_fasta(data)
    adjacency = build_overlap_graph(records, 3)

    for source, target in adjacency:
        print(source, target)


if __name__ == "__main__":
    main()
