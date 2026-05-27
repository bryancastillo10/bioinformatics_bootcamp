def hamming_distance(s: str, t: str) -> int:
    distance = 0

    for i in range(len(s)):
        if s[i] != t[i]:
            distance += 1

    return distance


def main():
    with open("../input.txt", "r") as file:
        lines = file.read().strip().splitlines()

    s = lines[0].strip()
    t = lines[1].strip()

    res = hamming_distance(s, t)
    print(res)


if __name__ == "__main__":
    main()
