from unittest import TestCase
from days.day06 import partA, partB


class TestDay06(TestCase):
    example = """Time:      7  15   30
Distance:  9  40  200
"""

    def testPartA(self):
        res = partA(self.example)
        self.assertEqual(res, 288)

    def testPartB(self):
        res = partB(self.example)
        self.assertEqual(res, 71503)
