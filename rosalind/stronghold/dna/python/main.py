def count_dna_nucleotides(seq:str) -> str:
	a = seq.count("A")
	c = seq.count("C")
	g = seq.count("G")
	t = seq.count("T")
	return {
		"A": a,
		"C": c,
		"G": g,
		"T":t
	}

def main():
	with open("../input.txt") as file:
		seq = file.read().strip()

	res = count_dna_nucleotides(seq)
	print(res)

if __name__ == "__main__":
	main()