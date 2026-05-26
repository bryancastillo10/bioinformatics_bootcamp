def pairings(n: int, k: int) -> int:
    if n <= 2:
        return 1

    prev = 1
    curr = 1

    for _ in range(3, n + 1):
        next_val = curr + (k * prev)
        prev = curr
        curr = next_val

    return curr


def main():
    with open("../input.txt", "r") as file:
        n, k = map(int, file.read().strip().split())

    res = pairings(n, k)
    print(res)


if __name__ == "__main__":
    main()
