from typing import Any
from unittest import TestCase
from days.day05 import partA, partB, translate, _transform
from parameterized import parameterized


class TestDay05(TestCase):
    example = """
seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4
"""

    @parameterized.expand(
        [
            (79, ["50 98 2", "52 50 48"], 81),
            (14, ["50 98 2", "52 50 48"], 14),
            (55, ["50 98 2", "52 50 48"], 57),
            (13, ["50 98 2", "52 50 48"], 13),
        ]
    )
    def testTranslate(self, source: int, map: str, expected: int):
        target = translate(source, map)
        self.assertEqual(target, expected)

    @parameterized.expand(
        [
            (
                example,
                [79, 14, 55, 13],
                [
                    [
                        "50 98 2",
                        "52 50 48",
                    ],
                    [
                        "0 15 37",
                        "37 52 2",
                        "39 0 15",
                    ],
                    [
                        "49 53 8",
                        "0 11 42",
                        "42 0 7",
                        "57 7 4",
                    ],
                    [
                        "88 18 7",
                        "18 25 70",
                    ],
                    [
                        "45 77 23",
                        "81 45 19",
                        "68 64 13",
                    ],
                    [
                        "0 69 1",
                        "1 0 69",
                    ],
                    [
                        "60 56 37",
                        "56 93 4",
                    ],
                ],
            ),
            (
                example,
                list(range(79, 79 + 14)) + list(range(55, 55 + 13)),
                [
                    [
                        "50 98 2",
                        "52 50 48",
                    ],
                    [
                        "0 15 37",
                        "37 52 2",
                        "39 0 15",
                    ],
                    [
                        "49 53 8",
                        "0 11 42",
                        "42 0 7",
                        "57 7 4",
                    ],
                    [
                        "88 18 7",
                        "18 25 70",
                    ],
                    [
                        "45 77 23",
                        "81 45 19",
                        "68 64 13",
                    ],
                    [
                        "0 69 1",
                        "1 0 69",
                    ],
                    [
                        "60 56 37",
                        "56 93 4",
                    ],
                ],
            ),
        ]
    )
    def testTransform(
        self,
        puzzle: str,
        expected_seeds: list[int],
        expected_maps: list[list[str]],
    ):
        seeds, maps = _transform(puzzle)
        self.assertEqual(seeds, expected_seeds)
        self.assertEqual(maps, expected_maps)

    def testPartA(self):
        result = partA(self.example)
        self.assertEqual(result, 35)

    def testPartB(self):
        result = partB(self.example)
        self.assertEqual(result, 46)
