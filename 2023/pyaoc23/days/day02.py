import re


def parseBagForMax(bag: str) -> (int, int, int):
    def extractColor(bag: str, color: str) -> int:
        return max([int(x) for x in re.findall("([0-9]+) %s" % color, bag)])

    return (
        extractColor(bag, "red"),
        extractColor(bag, "green"),
        extractColor(bag, "blue"),
    )


def partA(puzzle: str) -> str:
    lines = puzzle.split("\n")
    valid_games = []
    for l in lines:
        r, g, b = parseBagForMax(l)
        if r <= 12 and g <= 13 and b <= 14:
            game_id = re.findall("Game (\d+):", l)[0]
            valid_games.append(int(game_id))
    return sum(valid_games)


def partB(puzzle: str) -> str:
    lines = puzzle.split("\n")
    powers = []
    for l in lines:
        r, g, b = parseBagForMax(l)
        powers.append(r * b * g)
    return sum(powers)
