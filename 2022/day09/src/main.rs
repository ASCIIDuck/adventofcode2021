use std::ops::Add;

use util::read_input;

#[derive(Copy, Clone, Debug, Eq, PartialEq, PartialOrd, Ord)]
struct XY {
    x: i32,
    y: i32,
}

impl Add for XY {
    type Output = Self;
    fn add(self, other: Self) -> Self {
        return Self {
            x: self.x + other.x,
            y: self.y + other.y,
        };
    }
}

impl XY {
    fn ends_not_touching(&self) -> bool {
        return self.x.abs() > 1 || self.y.abs() > 1;
    }
    fn contract(&mut self) -> XY {
        let mut out = XY { x: 0, y: 0 };
        if self.x != 0 {
            out.x += self.x / self.x.abs();
            self.x -= self.x / self.x.abs();
        }
        if self.y != 0 {
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
fn main() {
    let data = read_input("input.txt");
    let res = walk_head(data);
    println!("Part 1: {}", res);
}

#[cfg(test)]
mod tests {
    use crate::walk_head;
    #[test]
    fn sample_input() {
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
}
