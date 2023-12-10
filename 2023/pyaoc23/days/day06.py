import math
import re


def _transformA(puzzle: str) -> list[(int, int)]:
    (time_str, dist_str) = re.sub(" +", " ", puzzle).strip().split("\n", 2)
    times = [int(x) for x in time_str.split(" ")[1:]]
    dists = [int(x) for x in dist_str.split(" ")[1:]]
    return zip(times, dists)


def _transformB(puzzle: str) -> list[(int, int)]:
    (time_str, dist_str) = re.sub(" +", " ", puzzle).strip().split("\n", 2)
    times = [int("".join(time_str.split(" ")[1:]))]
    dists = [int("".join(dist_str.split(" ")[1:]))]
    return zip(times, dists)


def transformInputA(f: callable) -> callable:
    def wrapper(puzzle: str) -> str:
        races = _transformA(puzzle)
        return f(races)

    return wrapper


def transformInputB(f: callable) -> callable:
    def wrapper(puzzle: str) -> str:
        races = _transformB(puzzle)
        return f(races)

    return wrapper


def solveZeroes(t: int, d: int) -> (int, int):
    b = math.sqrt(t * t - 4 * d)
    low = math.ceil(0.5 * (t - b))
    high = math.floor(0.5 * (t + b))

    if calculateDistance(low, t - low) <= d:
        low = low + 1
    if calculateDistance(high, t - high) <= d:
        high = high - 1

    return (low, high)


def calculateDistance(chargeTime: int, moveTime: int) -> int:
    return chargeTime * moveTime


@transformInputA
def partANiave(puzzle: list[(int, int)]) -> str:
    ans = 1
    for race in puzzle:
        time, dist_record = race
        race_solutions = 0
        for i in range(time):
            trial_dist = calculateDistance(i, time - i)
            if trial_dist > dist_record:
                race_solutions += 1
        ans = ans * race_solutions
    return ans


def _solve(puzzle: list[(int, int)]) -> str:
    ans = 1
    for race in puzzle:
        time, dist_record = race
        low, high = solveZeroes(time, dist_record)
        race_solutions = (high - low) + 1
        ans = ans * race_solutions
    return ans


@transformInputA
def partA(puzzle: list[(int, int)]) -> str:
    return _solve(puzzle)


@transformInputB
def partB(puzzle: list[(int, int)]) -> str:
    return _solve(puzzle)
