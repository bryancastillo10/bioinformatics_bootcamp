def find_substring_locations(s: str, t: str) -> list[int]:
    locations = []

    for i in range(len(s) - len(t) + 1):
        if s[i : i + len(t)] == t:
            locations.append(i + 1)

    return locations


def main():
    with open("../input.txt", "r") as file:
        lines = file.read().strip().splitlines()

    s = lines[0].strip()
    t = lines[1].strip()

    locations = find_substring_locations(s, t)
    print(" ".join(map(str, locations)))


if __name__ == "__main__":
    main()
