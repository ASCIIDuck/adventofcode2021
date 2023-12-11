from unittest import TestCase
from days.day07 import partA, partB, StandardHand, HandValues, WildCardHand, Hand
from parameterized import parameterized


class TestDay07(TestCase):
    example = """32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483
"""

    def testPartA(self):
        res = partA(self.example)
        self.assertEqual(res, 6440)

    def testPartB(self):
        res = partB(self.example)
        self.assertEqual(res, 5905)


class TestStandardHand(TestCase):
    @parameterized.expand(
        [
            # n-of-kind comparisons
            [StandardHand("23QQQ"), StandardHand("234QQ"), True],  # 3oak vs 2 oak
            [StandardHand("234QQ"), StandardHand("234KK"), False],  # 2oak vs 3 oak
            [StandardHand("234QQ"), StandardHand("234KQ"), True],  # 2oak vs highcard
            [StandardHand("234KQ"), StandardHand("234QQ"), False],  # highcard vs 2oak
            [StandardHand("2QQQQ"), StandardHand("23QQQ"), True],  # 4 oak vs 3 oak
            [StandardHand("23QQQ"), StandardHand("2QQQQ"), False],  # 3oak vs 4 oak
            [StandardHand("QQQQQ"), StandardHand("2QQQQ"), True],  # 5 oak vs 4 oak
            [StandardHand("2QQQQ"), StandardHand("QQQQQ"), False],  # 4 oak vs 5 oak
            [StandardHand("AAAAA"), StandardHand("QQQQQ"), True],  # 5 oak vs 5 oak
            # FH to n-of-kind comparisons
            [StandardHand("KKQQQ"), StandardHand("2QQQQ"), False],  # FH vs 4oak
            [StandardHand("2QQQQ"), StandardHand("KKQQQ"), True],  # foak vs FH
            [StandardHand("KKQQQ"), StandardHand("23QQQ"), True],  # FH vs 3oak
            [StandardHand("23QQQ"), StandardHand("KKQQQ"), False],  # 3oak vs FH
            [StandardHand("QQQQQ"), StandardHand("KKQQQ"), True],  # 5 oak vs FH
            [StandardHand("KKQQQ"), StandardHand("QQQQQ"), False],  # FH vs 5 oak
            # Pairs
            [StandardHand("AAA23"), StandardHand("KK2QQ"), True],  # 3oak vs 2P
            [StandardHand("KK234"), StandardHand("KK2QQ"), False],  # P vs 2P
            [StandardHand("99288"), StandardHand("KK234"), True],  # 2P vs P
            [StandardHand("KK677"), StandardHand("KTJJT"), True],  # 2P vs 2P
            [StandardHand("KTJJT"), StandardHand("KK677"), False],  # 2P vs 2P
            [StandardHand("KK672"), StandardHand("KTJAT"), True],  # P vs P
            [StandardHand("KTJAT"), StandardHand("KK672"), False],  # P vs P
            # High Card
            [StandardHand("23456"), StandardHand("34567"), False],
            [StandardHand("34567"), StandardHand("23456"), True],
            # Wildcard comps
            [WildCardHand("QQQJA"), WildCardHand("KTJJT"), False],
            [WildCardHand("KTJJT"), WildCardHand("QQQJA"), True],
            [WildCardHand("JKKK2"), WildCardHand("QQQQ2"), False],
        ]
    )
    def testCmp(self, one: StandardHand, two: StandardHand, expected: bool):
        self.assertEqual(one > two, expected)

    @parameterized.expand([["23456", 0], ["23466", 1], ["22466", 2]])
    def testNumPairs(self, hand_str: str, expected: int):
        hand = StandardHand(hand_str)
        self.assertEqual(hand.numPairs(), expected)

    @parameterized.expand(
        [
            [StandardHand("23456"), HandValues.high_card],
            [StandardHand("23466"), HandValues.one_pair],
            [StandardHand("22466"), HandValues.two_pair],
            [StandardHand("22266"), HandValues.fullhouse],
            [StandardHand("KKQQQ"), HandValues.fullhouse],
            [StandardHand("23QQQ"), HandValues.kind3],
            [StandardHand("22267"), HandValues.kind3],
            [StandardHand("22227"), HandValues.kind4],
            [StandardHand("22222"), HandValues.kind5],
            [WildCardHand("22222"), HandValues.kind5],
            [WildCardHand("JJJJJ"), HandValues.kind5],
            [WildCardHand("4JJJJ"), HandValues.kind5],
            [WildCardHand("44JJJ"), HandValues.kind5],
            [WildCardHand("444JJ"), HandValues.kind5],
            [WildCardHand("4444J"), HandValues.kind5],
            [WildCardHand("43JJJ"), HandValues.kind4],
            [WildCardHand("434JJ"), HandValues.kind4],
            [WildCardHand("4344J"), HandValues.kind4],
            [WildCardHand("4324J"), HandValues.kind3],
            [WildCardHand("432JJ"), HandValues.kind3],
            [WildCardHand("J32JJ"), HandValues.kind4],
            [WildCardHand("4325J"), HandValues.one_pair],
            [WildCardHand("QJJQ2"), HandValues.kind4],
        ]
    )
    def testValue(self, hand: Hand, expected: int):
        self.assertEqual(hand.value(), expected)
