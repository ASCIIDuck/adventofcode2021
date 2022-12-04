use util::read_input;
use std::ops::Range;
use std::str::Split;

fn main() {
    let data = read_input("input.txt");
    let lines = data.split("\n");
    let strict_res: u32 = parse_and_count(lines.clone(), &is_strict_subset);
    let res: u32 = parse_and_count(lines.clone(), &is_subset);
    println!("Numer of strict subsets {}", strict_res);
    println!("Numer of subsets {}", res);
}

fn parse_and_count(lines: Split<&str>, f: &dyn Fn(Range<u32>, Range<u32>) -> bool) -> u32 {
    return lines.map(|l| parse_line(l))
    .map(|line| 
        f(line.0.clone(), line.1.clone()) 
        || f(line.1 , line.0))
    .map(|x| if x {1} else {0}).sum();
}

fn parse_line(line: &str) -> (Range<u32>, Range<u32>) {
    return realize_tuple(line.split_at(line.find(',').unwrap()));
} 

fn is_strict_subset(first: Range<u32>, second: Range<u32>) -> bool {
    return second.fold(true, |acc, item| first.contains(&item) && acc);
}

fn is_subset(first: Range<u32>, second: Range<u32>) -> bool {
    return second.fold(false, |acc, item| first.contains(&item) || acc);
}

fn realize_tuple((one, two): (&str, &str)) ->  (Range<u32>, Range<u32>) {
    return (realize_range(one.trim_matches(',')), realize_range(two.trim_matches(',')))
}

fn realize_range(input: &str) ->  Range<u32> {
    let parts = (*input).split("-").collect::<Vec<&str>>();
    let lower = (*parts[0]).parse::<u32>().unwrap();
    let upper = (*parts[1]).parse::<u32>().unwrap();
    return lower..(upper+1);
}

#[cfg(test)] 
mod tests {
    use crate::is_strict_subset;

    #[test]
    fn is_strict_subset_non_overlapping() {
        let res = is_strict_subset(0..2, 2..3);
        assert_eq!(res, false);
    }
    #[test]
    fn is_strict_subset_overlapping() {
        let res = is_strict_subset(0..2, 1..3);
        assert_eq!(res, false);
    }
    #[test]
    fn is_strict_subset_yes() {
        let res = is_strict_subset(0..2, 1..2);
        assert_eq!(res, true);
    }

    #[test]
    fn is_subset_yes() {
        let res = is_subset(0..2, 1..2);
        assert_eq!(res, true);
    }
}