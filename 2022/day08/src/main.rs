use core::fmt;
use util::read_input;

const EDGE_VALUE: i32 = -1;

// Represents the plant sort of tree, not a the data structure
struct ElfTree {
    height: i32,
    max_right: i32,
    max_left: i32,
    max_up: i32,
    max_down: i32,
}

impl fmt::Debug for ElfTree {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        write!(f, "({}:{})", self.height, (self.is_visible()))
    }
}

impl ElfTree {
    fn new(height: i32) -> ElfTree {
        return ElfTree {
            height: height,
            max_right: EDGE_VALUE,
            max_left: EDGE_VALUE,
            max_up: EDGE_VALUE,
            max_down: EDGE_VALUE,
        };
    }

    fn is_visible(&self) -> bool {
        return self.height
            > *vec![self.max_down, self.max_up, self.max_left, self.max_right]
                .iter()
                .min()
                .unwrap();
    }
}

fn map_trees(input: String) -> Vec<Vec<ElfTree>> {
    let mut trees: Vec<Vec<ElfTree>> = input
        .lines()
        .map(|line| {
            line.trim()
                .chars()
                .map(|c| ElfTree::new(c.to_string().parse::<i32>().unwrap()))
                .collect()
        })
        .collect();

    assert_eq!(trees.len(), trees[0].len());
    let max_dim = trees.len();
    for i in 0..=max_dim / 2 {
        let lower = max_dim / 2 - i;
        let upper = max_dim / 2 + i;
        // Mark edge of new sub-plot as visible
        for j in 0..=2 * i {
            trees[lower][lower + j].max_up = EDGE_VALUE;
            trees[upper][lower + j].max_down = EDGE_VALUE;
            trees[lower + j][lower].max_left = EDGE_VALUE;
            trees[lower + j][upper].max_right = EDGE_VALUE;
        }
        for row in lower..=upper {
            for col in lower..=upper {
                if row > lower {
                    if trees[row - 1][col].max_up > trees[row - 1][col].height {
                        trees[row][col].max_up = trees[row - 1][col].max_up;
                    } else {
                        trees[row][col].max_up = trees[row - 1][col].height;
                    }
                }
                if row < upper {
                    if trees[row + 1][col].max_down > trees[row + 1][col].height {
                        trees[row][col].max_down = trees[row + 1][col].max_down;
                    } else {
                        trees[row][col].max_down = trees[row + 1][col].height;
                    }
                }
                if col > lower {
                    if trees[row][col - 1].max_left > trees[row][col - 1].height {
                        trees[row][col].max_left = trees[row][col - 1].max_left;
                    } else {
                        trees[row][col].max_left = trees[row][col - 1].height;
                    }
                }
                if col < upper {
                    if trees[row][col + 1].max_right > trees[row][col + 1].height {
                        trees[row][col].max_right = trees[row][col + 1].max_right;
                    } else {
                        trees[row][col].max_right = trees[row][col + 1].height;
                    }
                }
            }
        }
    }
    return trees;
}

fn do_work(input: String) {
    let trees = map_trees(input);
    let part1_res = trees.iter().fold(0, |acc, line| {
        acc + line
            .iter()
            .fold(0, |acc, t| if t.is_visible() { acc + 1 } else { acc })
    });
    for line in trees {
        for t in line {
            print!(
                "{}",
                if t.is_visible() {
                    t.height.to_string()
                } else {
                    ".".to_string()
                }
            );
        }
        println!("");
    }
    println!("Part 1 is {}", part1_res)
}
fn main() {
    let data = read_input("input.txt");
    do_work(data);
}

#[cfg(test)]
mod tests {
    use crate::do_work;

    #[test]
    fn test_one() {
        let input = "30373
        25512
        65332
        33549
        35390";
        do_work(input.to_string());
    }
    #[test]
    fn test_two() {
        let input = "3037361
        2551291
        6589811
        3373721
        3561531
        3451321
        6531321";
        do_work(input.to_string());
    }
}
