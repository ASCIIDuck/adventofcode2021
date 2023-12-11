from enum import IntEnum
from util.input import transformInput
from dataclasses import dataclass
from pprint import pprint
import math


NORTH = 0b0001
SOUTH = 0b0010
EAST = 0b0100
WEST = 0b1000


@dataclass
class Node:
    adjacent: int
    visited: int = 0

    @property
    def east(self):
        return self.adjacent & EAST == EAST

    @property
    def west(self):
        return self.adjacent & WEST == WEST

    @property
    def north(self):
        return self.adjacent & NORTH == NORTH

    @property
    def south(self):
        return self.adjacent & SOUTH == SOUTH

    def visitedFrom(self, direction: int):
        return self.visted & direction


@dataclass
class StartNode(Node):
    pass


NodeFactory = {
    "|": lambda: Node(NORTH | SOUTH),
    "-": lambda: Node(EAST | WEST),
    "L": lambda: Node(NORTH | EAST),
    "J": lambda: Node(NORTH | WEST),
    "7": lambda: Node(SOUTH | WEST),
    "F": lambda: Node(SOUTH | EAST),
    ".": lambda: Node(0),
    "S": lambda: StartNode(NORTH | SOUTH | EAST | WEST),
}


def readTheRunes(puzzle: str) -> (dict[(int, int), Node], (int, int)):
    nodes = {}
    lines = puzzle.split("\n")
    starting_point = (0, 0)
    for i in range(len(lines)):
        for j in range(len(lines[i])):
            rune = lines[i][j]
            n = NodeFactory[rune]()
            nodes[(i, j)] = n
            if type(n) == StartNode:
                starting_point = (i, j)
    return nodes, starting_point


def adjacentNodes(
    puzzle: dict[(int, int), Node], one_point: (int, int), two_point: Node
) -> bool:
    one = puzzle.get(one_point, NodeFactory["."]())
    two = puzzle.get(two_point, NodeFactory["."]())
    i, j, *_ = [a - b for a, b in zip(one_point, two_point)]
    if (i, j) == (-1, 0) and one.south and two.north:
        return True
    if (i, j) == (1, 0) and one.north and two.south:
        return True
    if (i, j) == (0, 1) and one.west and two.east:
        return True
    if (i, j) == (0, -1) and one.east and two.west:
        return True
    return False


@transformInput(readTheRunes)
def partA(puzzle: dict[(int, int), Node], starting_point: (int, int)) -> str:
    queue = [[starting_point]]
    while queue:
        path = queue.pop(0)
        cur = path[-1]
        if cur == starting_point and len(path) > 1:
            return math.floor(len(path) / 2)

        dirMap = {
            (1, 0): (SOUTH, NORTH),
            (-1, 0): (NORTH, SOUTH),
            (0, 1): (EAST, WEST),
            (0, -1): (WEST, EAST),
        }

        for i, j in dirMap.keys():
            candidate = (cur[0] + i, cur[1] + j)
            (candDir, curDir) = dirMap[(i, j)]
            if (
                candidate in puzzle
                and not (puzzle[candidate].visited & curDir)
                and adjacentNodes(puzzle, cur, candidate)
            ):
                new_path = path.copy()
                new_path.append(candidate)
                queue.append(new_path)
                puzzle[cur].visited |= candDir
                break


def partB(puzzle: str) -> str:
    return ""
