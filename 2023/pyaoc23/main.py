import argparse
import importlib
import traceback

import requests
from util.input import getInput


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
        print(traceback.format_exc())
    try:
        partB = day_mod.partB(puzzle)
    except Exception as e:
        print(e)
        print(traceback.format_exc())
    print("Part A: ", partA)
    print("Part B: ", partB)


if __name__ == "__main__":
    main()
