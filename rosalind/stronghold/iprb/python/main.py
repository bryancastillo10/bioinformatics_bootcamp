def dominant_phenotype_probability(k: int, m: int, n: int) -> float:
    total = k + m + n

    # Probability of selecting two heterozygous organisms: Aa x Aa
    het_het = (m / total) * ((m - 1) / (total - 1))

    # Probability of selecting heterozygous then recessive: Aa x aa
    het_rec = (m / total) * (n / (total - 1))

    # Probability of selecting recessive then heterozygous: aa x Aa
    rec_het = (n / total) * (m / (total - 1))

    # Probability of selecting recessive then recessive: aa x aa
    rec_rec = (n / total) * ((n - 1) / (total - 1))

    recessive_probability = (
        het_het * 0.25 + het_rec * 0.5 + rec_het * 0.5 + rec_rec * 1.00
    )

    return 1 - recessive_probability


def main():
    with open("../input.txt", "r") as file:
        k, m, n = map(int, file.read().strip().split())

    result = dominant_phenotype_probability(k, m, n)
    print(f"{result:.5f}")


if __name__ == "__main__":
    main()
