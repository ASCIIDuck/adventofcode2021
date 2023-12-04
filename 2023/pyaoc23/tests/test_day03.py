from typing import Any
from unittest import TestCase
from days.day03 import partA, partB


class TestDay03(TestCase):
    sample_input = """467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598.."""

    def test_partA_sample(self):
        ans = partA(self.sample_input)
        self.assertEqual(4361, ans)

    def test_partB_sample(self):
        ans = partB(self.sample_input)
        self.assertEqual(467835, ans)
