from unittest import TestCase
from days.day10 import partA, partB, Node, adjacentNodes, readTheRunes
from parameterized import parameterized
from pprint import pprint


class TestDay10(TestCase):
    example = """.....
.S-7.
.|.|.
.L-J.
....."""

    def testPartA(self):
        res = partA(self.example)
        self.assertEqual(res, 4)

    def testPartB(self):
        res = partB(self.example)
        self.assertEqual(res, 0)

    @parameterized.expand(
        [
            [example, (1, 1), (1, 2), True],
            [example, (1, 1), (2, 1), True],
            [example, (1, 3), (1, 2), True],
            [example, (1, 3), (2, 3), True],
            [example, (1, 1), (2, 2), False],
            [example, (1, 1), (2, 3), False],
        ]
    )
    def testAdjacentNodes(
        self,
        puzzle_str: str,
        one: (int, int),
        two: (int, int),
        expected: bool,
    ):
        puzzle, _ = readTheRunes(puzzle_str)
        res = adjacentNodes(puzzle, one, two)
        self.assertEqual(res, expected)
