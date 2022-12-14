use util::read_input;

trait Instruction {
    fn perform_instruction(&self, input: i32) -> i32;
    fn do_tick(&mut self) -> bool;
}

struct Noop {
    counter: i32,
}

impl Noop {
    fn new() -> Noop {
        return Noop { counter: 1 };
    }
}

impl Instruction for Noop {
    fn perform_instruction(&self, input: i32) -> i32 {
        return input;
    }
    fn do_tick(&mut self) -> bool {
        if self.counter <= 0 {
            return true;
        }
        self.counter -= 1;
        return false;
    }
}

struct Addx {
    counter: i32,
    x: i32,
}
impl Addx {
    fn new(x: i32) -> Addx {
        return Addx { counter: 2, x: x };
    }
}

impl Instruction for Addx {
    fn perform_instruction(&self, input: i32) -> i32 {
        return input + self.x;
    }
    fn do_tick(&mut self) -> bool {
        if self.counter <= 0 {
            return true;
        }
        self.counter -= 1;
        return false;
    }
}

fn evaluate_program(
    start_x: i32,
    input: String,
    selector: &dyn Fn(i32, i32) -> bool,
) -> Vec<(i32, i32)> {
    let mut tick = 0;
    let mut x = start_x;
    let mut out: Vec<(i32, i32)> = Vec::new();
    for line in input.lines() {
        let instruction_parts: Vec<&str> = line.trim().split_whitespace().collect();
        if instruction_parts.len() == 0 {
            println!("Blank line '{}'? Skipping!", line);
            continue;
        }
        let mut instruction: Box<dyn Instruction> = match instruction_parts[0] {
            "noop" => Box::new(Noop::new()),
            "addx" => {
                let x = instruction_parts[1].parse::<i32>().unwrap();
                Box::new(Addx::new(x))
            }
            &_ => continue,
        };
        while !instruction.do_tick() {
            tick += 1;
            if selector(tick, x) {
                out.push((tick, x));
            }
        }
        x = instruction.perform_instruction(x);
    }
    return out;
}

fn part_1(input: String) {
    let intervals = vec![20, 60, 100, 140, 180, 220];
    let result: i32 = evaluate_program(1, input, &|i, _| intervals.contains(&i))
        .iter()
        .map(|res| res.0 * res.1)
        .into_iter()
        .sum();
    println!("Part 1: {:?}", result);
}

fn part_2(input: String) {
    let result = evaluate_program(1, input, &|i, x| {
        let ret: bool = x <= i % 40 && i % 40 <= x + 2;
        return ret;
    });
    let mut screen: Vec<bool> = vec![false; 240];
    for r in result {
        screen[(r.0 as usize) - 1] = true;
    }
    let output = screen
        .chunks(40)
        .map(|line| {
            line.iter()
                .map(|&b| if b { "#" } else { " " })
                .collect::<Vec<&str>>()
                .join("")
        })
        .collect::<Vec<String>>()
        .join("\n");
    println!("Part 2:\n{}", output);
}
fn main() {
    let data = read_input("input.txt");
    part_1(data.clone());
    part_2(data.clone());
}

#[cfg(test)]
mod tests {
    use crate::evaluate_program;
    use crate::part_2;
    #[test]
    fn long_sample() {
        let input = "addx 15
        addx -11
        addx 6
        addx -3
        addx 5
        addx -1
        addx -8
        addx 13
        addx 4
        noop
        addx -1
        addx 5
        addx -1
        addx 5
        addx -1
        addx 5
        addx -1
        addx 5
        addx -1
        addx -35
        addx 1
        addx 24
        addx -19
        addx 1
        addx 16
        addx -11
        noop
        noop
        addx 21
        addx -15
        noop
        noop
        addx -3
        addx 9
        addx 1
        addx -3
        addx 8
        addx 1
        addx 5
        noop
        noop
        noop
        noop
        noop
        addx -36
        noop
        addx 1
        addx 7
        noop
        noop
        noop
        addx 2
        addx 6
        noop
        noop
        noop
        noop
        noop
        addx 1
        noop
        noop
        addx 7
        addx 1
        noop
        addx -13
        addx 13
        addx 7
        noop
        addx 1
        addx -33
        noop
        noop
        noop
        addx 2
        noop
        noop
        noop
        addx 8
        noop
        addx -1
        addx 2
        addx 1
        noop
        addx 17
        addx -9
        addx 1
        addx 1
        addx -3
        addx 11
        noop
        noop
        addx 1
        noop
        addx 1
        noop
        noop
        addx -13
        addx -19
        addx 1
        addx 3
        addx 26
        addx -30
        addx 12
        addx -1
        addx 3
        addx 1
        noop
        noop
        noop
        addx -9
        addx 18
        addx 1
        addx 2
        noop
        noop
        addx 9
        noop
        noop
        noop
        addx -1
        addx 2
        addx -37
        addx 1
        addx 3
        noop
        addx 15
        addx -21
        addx 22
        addx -6
        addx 1
        noop
        addx 2
        addx 1
        noop
        addx -10
        noop
        noop
        addx 20
        addx 1
        addx 2
        addx 2
        addx -6
        addx -11
        noop
        noop
        noop";
        let intervals = vec![20, 60, 100, 140, 180, 220];
        let results: Vec<(i32, i32)> =
            evaluate_program(1, input.to_string(), &|i, _| intervals.contains(&i));
        assert_eq!(
            results,
            [
                (20, 21),
                (60, 19),
                (100, 18),
                (140, 21),
                (180, 16),
                (220, 18)
            ]
        );
        part_2(input.to_string());
    }
}
