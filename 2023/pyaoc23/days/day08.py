import re
from util.input import transformInput
import sys


def parseNode(line: str) -> (str, str, str):
    node_regex = re.compile("(?P<head>\w+) = \((?P<left>\w+), (?P<right>\w+)\)")
    if match := node_regex.match(line):
        vals = match.groupdict()
        return (vals["head"], vals["left"], vals["right"])


def parseDirectionsAndGraph(puzzle: str) -> (dict[str, (str, str)], str):
    lines = puzzle.split("\n")
    directions = lines.pop(0)
    graph: dict[str, (str, str)] = {}
    for line in lines:
        if line == "":
            continue
        head, left, right = parseNode(line)
        graph[head] = (left, right)

    return graph, directions


@transformInput(parseDirectionsAndGraph, [])
def partA(graph: dict[str, (str, str)], directions: str) -> str:
    cur = "AAA"
    target = "ZZZ"
    step_counter = 0
    l = len(directions)
    while cur != target:
        left, right = graph[cur]
        match directions[step_counter % l]:
            case "L":
                cur = left
            case "R":
                cur = right
            case _:
                print("Uh-oh")
                sys.exit(1)
        step_counter += 1
    return step_counter


@transformInput(parseDirectionsAndGraph)
def partB(graph: dict[str, (str, str)], directions: str) -> str:
    def choose(dir: str, opts: (str, str)):
        left, right = opts
        match dir:
            case "L":
                return left
            case "R":
                return right
            case _:
                print("Uh-oh")
                sys.exit(1)

    target_re = re.compile(r".*Z$")
    step_counter = 0
    l = len(directions)
    queue: list[str] = [k for k in graph.keys() if re.match(".*A$", k)]
    while queue:
        dir = directions[step_counter % l]
        queue = [choose(dir, graph[n]) for n in queue]
        step_counter += 1
        if all([target_re.match(s) for s in queue]):
            break
    return step_counter
