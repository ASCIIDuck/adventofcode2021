from unittest import TestCase
from parameterized import parameterized
from days.day08 import partA, partB, parseNode, parseDirectionsAndGraph


class TestDay08(TestCase):
    partAExample = """LLR

AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)
"""
    partBExample = """LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)
"""

    def testPartA(self):
        ans = partA(self.partAExample)
        self.assertEqual(ans, 6)

    def testPartB(self):
        ans = partB(self.partBExample)
        self.assertEqual(ans, 6)

    @parameterized.expand(
        [
            ("AAA = (BBB, BBB)", ("AAA", "BBB", "BBB")),
            ("BBB = (AAA, ZZZ)", ("BBB", "AAA", "ZZZ")),
            ("ZZZ = (ZZZ, ZZZ)", ("ZZZ", "ZZZ", "ZZZ")),
        ]
    )
    def testParseNode(self, line: str, expected: (str, str, str)):
        out = parseNode(line)
        self.assertEqual(out, expected)

    def testParseDirectionsAndGraph(self):
        graph, directions = parseDirectionsAndGraph(self.partAExample)
        self.assertEquals(directions, "LLR")
        self.assertEquals(
            graph,
            {
                "AAA": ("BBB", "BBB"),
                "BBB": ("AAA", "ZZZ"),
                "ZZZ": ("ZZZ", "ZZZ"),
            },
        )
