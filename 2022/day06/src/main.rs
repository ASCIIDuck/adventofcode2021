use std::collections::{VecDeque, HashMap};
use util::read_input;

struct Window {
    characters: VecDeque<char>,
    window_width: usize
}

impl Window {
    fn add_char(&mut self, c: char) {
        self.characters.push_back(c);
        while self.characters.len() > self.window_width {
            self.characters.pop_front();
        }
    }

    fn is_start(&self) -> bool {
        return self.characters.len() >= self.window_width && self.characters.iter()
          .fold(HashMap::<char, u32>::new(), |mut acc:HashMap<char, u32>, c: &char| {*acc.entry(*c).or_insert(0)+=1; acc} )
          .into_values().fold(true, |acc, n| (n==1) && acc);
    }
}
fn main() {
    let input: Vec<char> = read_input("input.txt").chars().collect();
    let part1_res = find_start(&input, 4);
    let part2_res = find_start(&input, 14);
    println!("Part 1: {}", part1_res);
    println!("Part 2: {}", part2_res);
}
fn find_start(input: &Vec<char>, window_width: usize) -> usize {
    let mut w: Window = Window{characters: VecDeque::<char>::new(), window_width: window_width};
    let mut result:usize  = 0;
    for i in 0..input.len() {
        w.add_char(input[i]);
        if w.is_start() {
            result = i;
            break;
        }
    }
    return result+1;
}