import functools


def compute_history(values: list[int], index: int = -1) -> list[int]:
    history = []
    while values and any([x != 0 for x in values]):
        history.append(values[index])
        new_values = []
        for i in range(1, len(values)):
            new_values.append(values[i] - values[i - 1])
        values = new_values
    return history


def partA(puzzle: str) -> str:
    ans = []
    for line in puzzle.split("\n"):
        values = [int(x) for x in line.strip().split(" ") if x != ""]
        history = compute_history(values, -1)
        next_val = functools.reduce(lambda x, y: x + y, history, 0)
        ans.append(next_val)
    return sum(ans)


def partB(puzzle: str) -> str:
    ans = []
    for line in puzzle.split("\n"):
        values = [int(x) for x in line.strip().split(" ") if x != ""]
        history = compute_history(values, 0)
        next_val = functools.reduce(lambda x, y: y - x, reversed(history), 0)
        ans.append(next_val)
    return sum(ans)
