import sys
from typing import Tuple, List
from util.input import transform_input

def parse_lists(puzzle_input: str) -> Tuple[List[int], List[int]]:
    lines = puzzle_input.splitlines()
    left_list: List[int] = []
    right_list: List[int] = []
    for line in lines:
        (left, right) = line.split(' ',1)
        try:
            left_list.append(int(left))
            right_list.append(int(right))
        except ValueError:
            print("Failed to convert {} to int".format(line), file=sys.stderr)
    return left_list, right_list


@transform_input(parse_lists)
def part_a(left: List[int], right: List[int] ) -> int:
    left = sorted(left)
    right = sorted(right)
    distance = sum([ abs(l-r) for (l,r) in zip(left, right)])
    return distance

@transform_input(parse_lists)
def part_b(left: List[int], right: List[int] ) -> int:
    appearances = [ right.count(l) for l in left ]
    distance = sum([ l*a for (l,a) in zip(left, appearances)])
    return distance
