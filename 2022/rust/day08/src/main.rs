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
    scenic_score: i32,
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
            scenic_score: 0,
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

fn parse_trees(input: String) -> Vec<Vec<ElfTree>> {
    let trees: Vec<Vec<ElfTree>> = input
        .lines()
        .map(|line| {
            line.trim()
                .chars()
                .map(|c| ElfTree::new(c.to_string().parse::<i32>().unwrap()))
                .collect()
        })
        .collect();
    assert_eq!(trees.len(), trees[0].len());
    return trees;
}

fn mark_visible(trees: &mut Vec<Vec<ElfTree>>) {
    let max_dim = trees.len();
    for i in 1..max_dim {
        for j in 1..max_dim {
            if trees[i - 1][j].height > trees[i - 1][j].max_up {
                trees[i][j].max_up = trees[i - 1][j].height;
            } else {
                trees[i][j].max_up = trees[i - 1][j].max_up;
            }
            if trees[max_dim - i][j].height > trees[max_dim - i][j].max_down {
                trees[(max_dim - 1) - i][j].max_down = trees[max_dim - i][j].height;
            } else {
                trees[(max_dim - 1) - i][j].max_down = trees[max_dim - i][j].max_down;
            };
            if trees[i][j - 1].height > trees[i][j - 1].max_left {
                trees[i][j].max_left = trees[i][j - 1].height;
            } else {
                trees[i][j].max_left = trees[i][j - 1].max_left;
            }
            if trees[i][max_dim - j].height > trees[i][max_dim - j].max_right {
                trees[i][(max_dim - 1) - j].max_right = trees[i][max_dim - j].height;
            } else {
                trees[i][(max_dim - 1) - j].max_right = trees[i][max_dim - j].max_right;
            }
        }
    }
}

fn _caclulate_singe_scenic_score(trees: &Vec<Vec<ElfTree>>, row: usize, col: usize) -> i32 {
    let max_dim = trees.len();
    let (mut left, mut right, mut up, mut down) = (col, max_dim - col - 1, row, max_dim - row - 1);
    let cur_height = trees[row][col].height;
    for i in 1..max_dim {
        if row > i && i < up && trees[row - i][col].height >= cur_height {
            up = i;
        }
        if row + i < max_dim && i < down && trees[row + i][col].height >= cur_height {
            down = i;
        }
        if col > i && i < left && trees[row][col - i].height >= cur_height {
            left = i;
        }
        if col + i < max_dim && i < right && trees[row][col + i].height >= cur_height {
            right = i;
        }
    }
    return (left * right * up * down) as i32;
}

fn calculate_scenic_scores(trees: &mut Vec<Vec<ElfTree>>) {
    for i in 0..trees.len() {
        for j in 0..trees.len() {
            trees[i][j].scenic_score = _caclulate_singe_scenic_score(&trees, i, j);
        }
    }
}
fn do_work(input: String) {
    let mut trees = parse_trees(input);
    mark_visible(&mut trees);
    calculate_scenic_scores(&mut trees);
    let part1_res = trees.iter().fold(0, |acc, line| {
        acc + line
            .iter()
            .fold(0, |acc, t| if t.is_visible() { acc + 1 } else { acc })
    });
    println!("Part 1 is {}", part1_res);
    let part2_res = trees
        .iter()
        .map(|row| row.iter().map(|t| t.scenic_score).max().unwrap())
        .max()
        .unwrap();
    println!("Part 2 is {}", part2_res);
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
        let input = "
        30373
        25512
        65332
        33549
        35390";
        do_work(input.to_string());
    }
}
