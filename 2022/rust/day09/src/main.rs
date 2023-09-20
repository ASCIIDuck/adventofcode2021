use core::fmt;
use std::ops::{Add, AddAssign};

use util::read_input;

#[derive(Copy, Clone, Eq, PartialEq, PartialOrd, Ord)]
struct XY {
    x: i32,
    y: i32,
}
impl fmt::Debug for XY {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        write!(f, "({},{})", self.x, self.y)
    }
}

impl Add for XY {
    type Output = Self;
    fn add(self, rhs: Self) -> Self {
        return Self {
            x: self.x + rhs.x,
            y: self.y + rhs.y,
        };
    }
}
impl AddAssign for XY {
    fn add_assign(&mut self, rhs: Self) {
        self.x += rhs.x;
        self.y += rhs.y;
    }
}

impl XY {
    fn ends_not_touching(&self) -> bool {
        return self.x.abs() > 1 || self.y.abs() > 1;
    }
    fn contract(&mut self) -> XY {
        let mut out = XY { x: 0, y: 0 };
        if self.ends_not_touching() {
            let delta_x: i32 = match self.x.checked_div(self.x.abs()) {
                Some(res) => res,
                None => 0,
            };
            let delta_y = match self.y.checked_div(self.y.abs()) {
                Some(res) => res,
                None => 0,
            };
            out.x += delta_x;
            self.x -= delta_x;
            out.y += delta_y;
            self.y -= delta_y;
        }
        return out;
    }
}

fn walk_head(input: String) -> i32 {
    let lines = input.lines();
    let mut head_to_tail = XY { x: 0, y: 0 };
    let mut tail = XY { x: 0, y: 0 };
    let mut pos_history = Vec::<XY>::new();
    pos_history.push(tail);
    for line in lines {
        let parts = line.split_whitespace().collect::<Vec<&str>>();
        let dir = parts[0];
        let mag: i32 = parts[1].parse::<i32>().unwrap();
        match dir {
            "U" => {
                head_to_tail.y += mag;
            }
            "D" => {
                head_to_tail.y -= mag;
            }
            "L" => {
                head_to_tail.x -= mag;
            }
            "R" => {
                head_to_tail.x += mag;
            }
            &_ => (),
        }
        while head_to_tail.ends_not_touching() {
            let shift = head_to_tail.contract();
            tail = tail + shift;
            pos_history.push(tail);
        }
    }
    let mut uniq_pos = pos_history.clone();
    uniq_pos.sort();
    uniq_pos.dedup();
    let res = uniq_pos.len();
    return res as i32;
}

fn walk_rope(input: String, rope_len: usize) -> i32 {
    let instructions = input.lines().map(|l| {
        let (i, m) = l.trim().split_at(1);
        let mag = m.trim().parse::<i32>().unwrap();
        match i {
            "U" => {
                return XY { x: 0, y: 1 * mag };
            }
            "D" => {
                return XY { x: 0, y: -1 * mag };
            }
            "L" => {
                return XY { x: -1 * mag, y: 0 };
            }
            "R" => {
                return XY { x: 1 * mag, y: 0 };
            }
            &_ => {
                return XY { x: 0, y: 0 };
            }
        }
    });
    // by "vector" here I mean like the force vectors
    // Increase the size by one as the final vector will represent the
    // tail's position vector
    let mut rope_vectors = vec![XY { x: 0, y: 0 }; rope_len + 1];
    let mut tail_pos_history = Vec::<XY>::new();
    tail_pos_history.push(rope_vectors[rope_len]);
    for inst in instructions {
        rope_vectors[0] += inst;
        let mut breaker = 1000;
        while rope_vectors[0..rope_len]
            .iter()
            .any(|k| k.ends_not_touching())
        {
            for i in 0..rope_len {
                let shift = rope_vectors[i].contract();
                rope_vectors[i + 1] += shift;
            }
            tail_pos_history.push(rope_vectors[rope_len]);
            breaker -= 1;
            if breaker == 0 {
                break;
            }
        }
    }
    let mut uniq_pos = tail_pos_history.clone();
    uniq_pos.sort();
    uniq_pos.dedup();
    let res = uniq_pos.len();
    return res as i32;
}

fn main() {
    let data = read_input("input.txt");
    let part1_res = walk_rope(data.clone(), 1);
    println!("Part 1 answer: {}", part1_res);
    let part2_res = walk_rope(data.clone(), 9);
    println!("Part 2 answer: {}", part2_res);
}

#[cfg(test)]
mod tests {
    use crate::walk_head;
    use crate::walk_rope;
    use crate::XY;
    use util::read_input;
    #[test]
    fn test_walk_head() {
        let res = walk_head(
            "R 4
        U 4
        L 3
        D 1
        R 4
        D 1
        L 5
        R 2"
            .to_string(),
        );
        assert_eq!(res, 13);
    }
    #[test]
    fn test_walk_rope_1() {
        let res = walk_rope(
            "R 4
        U 4
        L 3
        D 1
        R 4
        D 1
        L 5
        R 2"
            .to_string(),
            1,
        );
        assert_eq!(res, 13);
    }
    #[test]
    fn test_walk_rope_10() {
        let res = walk_rope(
            "R 4
        U 4
        L 3
        D 1
        R 4
        D 1
        L 5
        R 2"
            .to_string(),
            10,
        );
        assert_eq!(res, 1);
    }

    #[test]
    fn test_walk_rope_10_sample2() {
        let res = walk_rope(
            "R 5
            U 8
            L 8
            D 3
            R 17
            D 10
            L 25
            U 20"
                .to_string(),
            9,
        );
        assert_eq!(res, 36);
    }

    #[test]
    fn run_input_part1() {
        let res = walk_rope(read_input("input.txt"), 1);
        assert_eq!(res, 6354);
    }

    #[test]
    fn test_ends_touching() {
        assert_eq!(XY { x: 0, y: 0 }.ends_not_touching(), false);
        assert_eq!(XY { x: 1, y: 0 }.ends_not_touching(), false);
        assert_eq!(XY { x: 0, y: 1 }.ends_not_touching(), false);
        assert_eq!(XY { x: 1, y: 1 }.ends_not_touching(), false);
        assert_eq!(XY { x: 2, y: 0 }.ends_not_touching(), true);
        assert_eq!(XY { x: 0, y: 2 }.ends_not_touching(), true);
        assert_eq!(XY { x: 2, y: 1 }.ends_not_touching(), true);
        assert_eq!(XY { x: 1, y: 2 }.ends_not_touching(), true);
        assert_eq!(XY { x: 2, y: 2 }.ends_not_touching(), true);
        assert_eq!(XY { x: 0, y: 0 }.ends_not_touching(), false);
        assert_eq!(XY { x: -1, y: 0 }.ends_not_touching(), false);
        assert_eq!(XY { x: 0, y: -1 }.ends_not_touching(), false);
        assert_eq!(XY { x: -1, y: -1 }.ends_not_touching(), false);
        assert_eq!(XY { x: -2, y: 0 }.ends_not_touching(), true);
        assert_eq!(XY { x: 0, y: -2 }.ends_not_touching(), true);
        assert_eq!(XY { x: -2, y: -1 }.ends_not_touching(), true);
        assert_eq!(XY { x: -1, y: -2 }.ends_not_touching(), true);
        assert_eq!(XY { x: -2, y: -2 }.ends_not_touching(), true);
    }
}
