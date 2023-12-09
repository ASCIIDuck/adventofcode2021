import re
from parameterized import parameterized
from dataclasses import dataclass


def individualSeeds(seed_str: str) -> list[int]:
    return [int(x) for x in seed_str.strip().split(" ")]


def _transform(puzzle: str) -> (list[int], list[list[str]]):
    lines = puzzle.split("\n")
    seed_str = ""
    seeds: list[int] = []
    maps: list[list[str]] = []
    i = 0
    while i < len(lines):
        if lines[i].startswith("seeds:"):
            (_, seed_str) = lines[i].split(":", 2)
        elif re.match(".* map:$", lines[i]):
            map: list[str] = []
            i += 1
            while i < len(lines) and lines[i] != "":
                map.append(lines[i])
                i += 1
            maps.append(map)
        i += 1
    seeds = individualSeeds(seed_str)
    return seeds, maps


def transformInput(f: callable) -> callable:
    def wrapper(puzzle: str) -> str:
        seeds, maps = _transform(puzzle)
        return f(seeds, maps)

    return wrapper


def translate(source: int, map: list[str]) -> int:
    for line in map:
        (d, s, l) = [int(x) for x in line.split(" ")]
        if source >= s and source < s + l:
            return d + (source - s)
    return source


@transformInput
def partA(seeds: list[int], maps: list[list[str]]) -> str:
    lowest = None
    for seed in seeds:
        current = seed
        for map in maps:
            current = translate(current, map)
        if lowest is None or current < lowest:
            lowest = current
    return lowest


@transformInput
def partB(seeds: list[int], maps: list[list[str]]) -> str:
    return ""
