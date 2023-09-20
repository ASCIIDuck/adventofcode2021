extern crate itertools;
use util::read_input;
use std::collections::HashSet;
use itertools::Itertools;

fn main() {
    let data = read_input("input.txt");
    let rucksacks = data.split("\n");
    let mut common_priority = 0;
    let mut badge_priority = 0;
    let all_chars: HashSet<char> = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ".chars().collect();
    for ruck in rucksacks {
        let first: HashSet<char> = ruck[0..ruck.len()/2].chars().collect();
        let second: HashSet<char> = ruck[ruck.len()/2..ruck.len()].chars().collect();
        common_priority += first.intersection(&second).fold(0, |acc, i| acc+get_priority(i));
    }
    // rucksacks.clone().map(|s| (*s).chars().collect::<HashSet<char>>()).collect::<Vec<HashSet<char>>>().chunks(3).into_iter().fold(all_chars.clone(), |common: HashSet<char>, ruck: &HashSet<char>| common.intersection(ruck));
    let groups = data.split("\n").chunks(3);
    for group in &groups {
        let mut common = all_chars.clone();
        for ruck in group {
            let ruck_set: HashSet<char> = ruck.chars().collect();
            common = common.intersection(&ruck_set).copied().collect();
        }
        badge_priority += common.iter().fold(0, |acc, i| acc+get_priority(i));

    }
    println!("Common priority: {}", common_priority);
    println!("Badge priority: {}", badge_priority);
}

fn get_priority(item: &char) -> u32 {
    return if item.is_ascii_uppercase() { *item as u32 - 64 + 26} else { *item as u32 - 96}
}