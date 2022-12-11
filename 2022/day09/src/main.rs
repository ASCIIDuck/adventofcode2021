use std::ops::{Add, AddAssign};

use util::read_input;

#[derive(Copy, Clone, Debug, Eq, PartialEq, PartialOrd, Ord)]
struct XY {
    x: i32,
    y: i32,
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
        if self.x.abs() != 0 {
            out.x += self.x / self.x.abs();
            self.x -= self.x / self.x.abs();
        }
        if self.y.abs() != 0 {
            out.y += self.y / self.y.abs();
            self.y -= self.y / self.y.abs();
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
    println!("{:?}", rope_vectors);
    for inst in instructions {
        rope_vectors[0] += inst;
        while rope_vectors[0..rope_len]
            .iter()
            .any(|k| k.ends_not_touching())
        {
            for i in 0..rope_vectors.len() - 1 {
                let shift = rope_vectors[i].contract();
                rope_vectors[i + 1] += shift;
            }
            tail_pos_history.push(rope_vectors[rope_len]);
            println!("{:?}", rope_vectors);
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
    let part2_res = walk_rope(data.clone(), 10);
    println!("Part 2 answer: {}", part2_res);
}

#[cfg(test)]
mod tests {
    use crate::walk_head;
    use crate::walk_rope;
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
    fn sample_input() {
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
    fn sample_input_10() {
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
}
