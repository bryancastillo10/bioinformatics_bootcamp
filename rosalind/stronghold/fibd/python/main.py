def mortal_fibonacci_rabbits(n: int, m: int) -> int:
    ages = [0] * m
    ages[0] = 1  # month 1: one newborn pair

    for _ in range(1, n):
        newborns = sum(ages[1:])
        ages = [newborns] + ages[:-1]

    return sum(ages)


def main():
    with open("../input.txt", "r") as file:
        n, m = map(int, file.read().strip().split())

    res = mortal_fibonacci_rabbits(n, m)
    print(res)


if __name__ == "__main__":
    main()
