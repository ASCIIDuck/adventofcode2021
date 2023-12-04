from dataclasses import dataclass
import math
import re


@dataclass(frozen=False)
class Card:
    id: int
    winners: list[int]
    numbers: list[int]
    multiplier: int = 1


def transformInput(f: callable) -> callable:
    card_regex = re.compile(
        "Card\s+(?P<id>\d+): (?P<winners>[0-9 ]+) \| (?P<numbers>[0-9 ]+)"
    )

    def wrapper(puzzle: str) -> str:
        cards: list[Card] = []
        for line in puzzle.split("\n"):
            match = card_regex.match(line)
            if not match:
                continue
            data = match.groupdict()
            c = Card(
                id=int(data["id"]),
                winners=[int(x) for x in data["winners"].strip().split(" ") if x != ""],
                numbers=[int(x) for x in data["numbers"].strip().split(" ") if x != ""],
            )
            cards.append(c)
        return f(cards)

    return wrapper


@transformInput
def partA(puzzle: list[Card]) -> str:
    score = 0
    for card in puzzle:
        matched_numbers = list(set(card.numbers).intersection(card.winners))
        score += int(math.pow(2, len(matched_numbers) - 1))
    return score


@transformInput
def partB(puzzle: list[Card]) -> str:
    card_count = 0
    for i in range(len(puzzle)):
        card = puzzle[i]
        matched_numbers = list(set(card.numbers).intersection(card.winners))
        score = len(matched_numbers)
        for j in range(i + 1, i + 1 + score):
            try:
                puzzle[j].multiplier += card.multiplier
            except IndexError:
                break
        card_count += card.multiplier

    return card_count
