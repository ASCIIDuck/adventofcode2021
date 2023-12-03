import re
from typing import Callable


def solution(puzzle: str, check: str, transform: Callable) -> str:
    lines = puzzle.split("\n")
    subans = []
    for l in lines:
        l = l.strip()
        if not l:
            continue
        digits = re.findall(check, l)
        subans.append(transform(digits))
    return sum(subans)


def partA(puzzle: str) -> str:
    def transform(digits: list) -> int:
        return int(digits[0] + digits[-1])

    return solution(puzzle, "[1-9]", transform)


def partB(puzzle: str) -> str:
    def transform(digits: list) -> int:
        lookup = {
            "one": "1",
            "two": "2",
            "three": "3",
            "four": "4",
            "five": "5",
            "six": "6",
            "seven": "7",
            "eight": "8",
            "nine": "9",
        }
        a = ""
        b = ""
        try:
            a = lookup[digits[0]]
        except KeyError:
            a = digits[0]

        try:
            b = lookup[digits[-1]]
        except KeyError:
            b = digits[-1]
        return int(a + b)

    return solution(
        puzzle, "([1-9]|one|two|three|four|five|six|seven|eight|nine)", transform
    )
