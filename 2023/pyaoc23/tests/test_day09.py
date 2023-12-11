from unittest import TestCase
from days.day09 import partA, partB


class TestDay09(TestCase):
    example = """0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45
"""

    def testPartA(self):
        res = partA(self.example)
        self.assertEqual(res, 114)

    def testPartB(self):
        res = partB(self.example)
        self.assertEqual(res, 2)
