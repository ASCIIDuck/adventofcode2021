from days.day01 import part_a, part_b

EXAMPLE_INPUT="""3   4
4   3
2   5
1   3
3   9
3   3
"""

def test_part_a():
    out = part_a(EXAMPLE_INPUT)
    assert out == 11

def test_part_b():
    out = part_b(EXAMPLE_INPUT)
    assert out == 31