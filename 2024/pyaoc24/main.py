import argparse
import importlib
import traceback

from util.input import get_input


def main():
    parser = argparse.ArgumentParser()
    parser.add_argument("--day", type=int)
    args = parser.parse_args()

    day_mod = importlib.import_module("days.day%02d" % args.day)
    puzzle = get_input(2024, args.day)
    part_a, part_b = "", ""
    try:
        part_a = day_mod.part_a(puzzle)
    except Exception as e:
        print(e)
        print(traceback.format_exc())
    try:
        part_b = day_mod.part_b(puzzle)
    except Exception as e:
        print(e)
        print(traceback.format_exc())
    print("Part A: ", part_a)
    print("Part B: ", part_b)


if __name__ == "__main__":
    main()