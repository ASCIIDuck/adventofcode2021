from util.input import transformInput
from dataclasses import dataclass
from enum import IntEnum


class HandValues(IntEnum):
    kind5 = 7
    kind4 = 6
    fullhouse = 5
    kind3 = 4
    two_pair = 3
    one_pair = 2
    high_card = 1


class Hand:
    def __init__(self, hand: str):
        pass

    def __gt__(self, other) -> bool:
        raise (NotImplementedError)


@dataclass
class StandardHand(Hand):
    LABEL_VALUES = {
        "2": 2,
        "3": 3,
        "4": 4,
        "5": 5,
        "6": 6,
        "7": 7,
        "8": 8,
        "9": 9,
        "T": 10,
        "J": 11,
        "Q": 12,
        "K": 13,
        "A": 14,
    }
    cards: dict[int, int]
    hand: str

    def __init__(self, hand: str):
        parsed = {}
        for c in hand:
            try:
                parsed[self.LABEL_VALUES[c]] += 1
            except KeyError:
                parsed[self.LABEL_VALUES[c]] = 1
        self.cards = parsed
        self.hand = [self.LABEL_VALUES[h] for h in hand]

    def labels_by_count(self, count: int) -> list[int]:
        return [label for label, n in self.cards if n == count]

    def n_of_kind(self) -> int:
        try:
            return max(self.cards.values())
        except ValueError:
            return 0

    def isFullHouse(self) -> bool:
        return all([n in self.cards.values() for n in [3, 2]])

    def numPairs(self) -> int:
        return len([x for x in self.cards.values() if x == 2])

    def value(self) -> str:
        if self.isFullHouse():
            return HandValues.fullhouse
        elif self.n_of_kind() == 5:
            return HandValues.kind5
        elif self.n_of_kind() == 4:
            return HandValues.kind4
        elif self.n_of_kind() == 3:
            return HandValues.kind3
        elif self.n_of_kind() == 2 and self.numPairs() == 2:
            return HandValues.two_pair
        elif self.n_of_kind() == 2:
            return HandValues.one_pair
        return HandValues.high_card

    def __gt__(self, other) -> bool:
        my_labels = list(
            reversed(sorted(self.cards.keys(), key=lambda k: self.cards[k] * 100 + k))
        )
        their_labels = list(
            reversed(sorted(other.cards.keys(), key=lambda k: other.cards[k] * 100 + k))
        )
        if self.value() > other.value():
            return True
        if self.value() == other.value():
            vals = list(zip(self.hand, other.hand))
            for a, b in vals:
                if a > b:
                    return True
                elif b > a:
                    break
        return False


@dataclass
class WildCardHand(StandardHand):
    LABEL_VALUES = {
        "J": 1,
        "2": 2,
        "3": 3,
        "4": 4,
        "5": 5,
        "6": 6,
        "7": 7,
        "8": 8,
        "9": 9,
        "T": 10,
        "Q": 11,
        "K": 12,
        "A": 13,
    }

    def __init__(self, hand: str):
        super().__init__(hand)
        self.jack_factor = self.cards.get(self.LABEL_VALUES["J"], 0)
        if self.jack_factor:
            del self.cards[self.LABEL_VALUES["J"]]

    def value(self) -> int:
        nkind = self.n_of_kind()
        if self.jack_factor >= 4:
            return HandValues.kind5
        elif self.jack_factor == 3:
            if nkind == 2:
                return HandValues.kind5
            else:
                return HandValues.kind4
        elif self.jack_factor == 2:
            if nkind == 3:
                return HandValues.kind5
            elif nkind == 2:
                return HandValues.kind4
            else:
                return HandValues.kind3
        elif self.jack_factor == 1:
            if nkind == 4:
                return HandValues.kind5
            if nkind == 3:
                return HandValues.kind4
            if nkind == 2 and self.numPairs() == 1:
                return HandValues.kind3
            if nkind == 2 and self.numPairs() == 2:
                return HandValues.fullhouse
            else:
                return HandValues.one_pair
        return super().value()


def parseHandsAndBets(puzzle: str, hand_cls: type) -> list[(Hand, int)]:
    def splitter(line: str) -> (str, int):
        hand, bet = line.strip().split(" ", 2)
        return hand_cls(hand), int(bet)

    output = [splitter(line) for line in puzzle.split("\n") if line != ""]
    return output


@transformInput(parseHandsAndBets, [StandardHand])
def partA(puzzle: list[(Hand, int)]) -> int:
    puzzle = sorted(puzzle, key=lambda l: l[0])
    payout = 0
    for i in range(len(puzzle)):
        _, bid = puzzle[i]
        payout += bid * (i + 1)
    return payout


@transformInput(parseHandsAndBets, [WildCardHand])
def partB(puzzle: list[(WildCardHand, int)]) -> int:
    puzzle = sorted(puzzle, key=lambda l: l[0])
    payout = 0
    for i in range(len(puzzle)):
        hand, bid = puzzle[i]
        value = hand.value()
        payout += bid * (i + 1)
    return payout
