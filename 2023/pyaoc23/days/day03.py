import re
import functools


def transformInput(f: callable) -> list[list[str]]:
    def wrapper(puzzle: str):
        numbers: dict[(int, int), str] = {}
        symbols: dict[(int, int), str] = {}
        x, y = (0, 0)
        lines = puzzle.split("\n")

        while x < len(lines):
            y = 0
            if lines[0] == "":
                continue
            while y < len(lines[0]):
                curnum = extractNum(lines, x, y)
                figs = len(curnum)
                if figs > 0:
                    numbers[(x, y)] = curnum
                    y = y + figs
                elif lines[x][y] != ".":
                    symbols[(x, y)] = lines[x][y]
                    y = y + 1
                else:
                    y = y + 1
            x = x + 1
        return f(numbers, symbols)

    return wrapper


def extractNum(puzzle, x, y) -> str:
    num = ""
    for i in range(0, len(puzzle[x]) - y):
        if re.match("[0-9]", puzzle[x][y + i]):
            num = num + puzzle[x][y + i]
        else:
            break
    return num


@transformInput
def partA(numbers: dict[(int, int), str], symbols: dict[(int, int), str]) -> str:
    total = 0
    for (x, y), value in numbers.items():
        for i, j in [
            (0, -1),
            (0, len(value)),
        ] + [(n, m) for m in range(-1, len(value) + 1) for n in [1, -1]]:
            if (x + i, y + j) in symbols:
                total = total + int(value)
    return total


@transformInput
def partB(numbers: dict[(int, int), str], symbols: dict[(int, int), str]) -> str:
    total = 0
    adjacencies: dict[(int, int), list[int]] = {}
    for (x, y), value in numbers.items():
        for i, j in [
            (0, -1),
            (0, len(value)),
        ] + [(n, m) for m in range(-1, len(value) + 1) for n in [1, -1]]:
            if (x + i, y + j) in symbols:
                try:
                    adjacencies[(x + i, y + j)].append(int(value))
                except KeyError:
                    adjacencies[(x + i, y + j)] = [int(value)]

    for coords, value in symbols.items():
        if value == "*":
            try:
                if len(adjacencies[coords]) > 1:
                    total = total + functools.reduce(
                        lambda a, b: a * b, adjacencies[coords]
                    )
            except KeyError:
                pass
    return total
