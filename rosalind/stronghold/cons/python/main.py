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
		"T": [0] * seq_length
	}

	for seq in sequences:
		for idx, nucleotide in enumerate(seq):
			profile[nucleotide][index] += 1

	return profile