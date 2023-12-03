import argparse
import importlib

import requests


def main():
    parser = argparse.ArgumentParser()
    parser.add_argument("--day", type=int)
    args = parser.parse_args()

    day_mod = importlib.import_module("days.day%02d" % args.day)
    puzzle = getInput(2023, args.day)
    partA, partB = "", ""
    try:
        partA = day_mod.partA(puzzle)
    except Exception as e:
        print(e)
    try:
        partB = day_mod.partB(puzzle)
    except Exception as e:
        print(e)
    print("Part A: ", partA)
    print("Part B: ", partB)


def getToken() -> str:
    try:
        with open(".token", "r") as f:
            return f.read().strip()
    except:
        return ""


def getInput(year: int, day: int) -> str:
    cache_path = ".cache/day%02d.input" % day
    try:
        with open(cache_path, "+r") as f:
            return f.read().strip()
    except FileNotFoundError:
        url = "https://adventofcode.com/%d/day/%d/input" % (year, day)
        sess = requests.Session()
        resp = sess.get(url, cookies={"session": getToken()})
        input = resp.content.decode()
        with open(cache_path, "+w") as f:
            f.write(input)
        return input


if __name__ == "__main__":
    main()
