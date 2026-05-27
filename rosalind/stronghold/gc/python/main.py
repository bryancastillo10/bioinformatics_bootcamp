def parse_fasta(data: str) -> dict[str, str]:
    records = {}
    curr_id = ""

    for line in data.strip().splitlines():
        line = line.strip()

        if line.startswith(">"):
            curr_id = line[1:]
            records[curr_id] = ""
        else:
            records[curr_id] += line

    return records


def calc_gc_content(dna: str) -> float:
    gc_count = dna.count("G") + dna.count("C")
    return (gc_count / len(dna)) * 100


def find_highest_gc_content(records: dict[str, str]) -> tuple[str, float]:
    highest_id = ""
    highest_gc = 0.0

    for record_id, dna in records.items():
        gc_content = calc_gc_content(dna)

        if gc_content > highest_gc:
            highest_id = record_id
            highest_gc = gc_content

    return highest_id, highest_gc


def main():
    with open("../input.txt", "r") as file:
        data = file.read()

    records = parse_fasta(data)
    highest_id, highest_gc = find_highest_gc_content(records)

    print(highest_id)
    print(f"{highest_gc:.6f}")


if __name__ == "__main__":
    main()
