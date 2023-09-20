use util::read_input;
fn main() {
    let data = read_input("input.txt");
    let lines: Vec<&str> = data.split("\n").collect::<Vec<&str>>();
    let mut initial_state: Vec<&str> = Vec::new();
    let mut directions: Vec<&str> = Vec::new();
    let mut break_point_found = false;
    for line in lines {
        if line == "" {
            break_point_found = true;
            continue;
        }
        if break_point_found {
            directions.push(line);
        } else {
            initial_state.push(line);
        }
    }
    let mut reverse_is_iter = initial_state.into_iter().rev();
    let num_stacks = reverse_is_iter
        .next()
        .unwrap()
        .split_whitespace()
        .collect::<Vec<&str>>()
        .len();
    let mut cur_state: Vec<Vec<&str>> = vec![Vec::new(); num_stacks];
    for line in reverse_is_iter {
        let parts = chunker(line);
        for i in 0..parts.len() {
            let elf_crate = parts[i].trim_matches(|c| c == '[' || c == ' ' || c == ']');
            if elf_crate != "" {
                cur_state[i].push(elf_crate);
            }
        }
    }

    part1(cur_state.clone(), directions.clone());
    part2(cur_state.clone(), directions.clone());
}

fn part1(mut cur_state: Vec<Vec<&str>>, directions: Vec<&str>) {
    for dir in directions {
        let parts: Vec<&str> = dir.split(" ").collect();
        let num_to_move: i32 = parts[1].parse().unwrap();
        let src_stack: usize = parts[3].parse::<usize>().unwrap() - 1;
        let dst_stack: usize = parts[5].parse::<usize>().unwrap() - 1;
        for _ in 0..num_to_move {
            match cur_state[src_stack].pop() {
                Some(c) => cur_state[dst_stack].push(c),
                None => (),
            }
        }
    }

    let output = cur_state
        .into_iter()
        .map(|stack| stack.last().copied().unwrap())
        .collect::<Vec<&str>>()
        .join("");
    println!("part 1 {}", output);
}

fn part2(mut cur_state: Vec<Vec<&str>>, directions: Vec<&str>) {
    for dir in directions {
        let parts: Vec<&str> = dir.split(" ").collect();
        let num_to_move: usize = parts[1].parse().unwrap();
        let src_stack: usize = parts[3].parse::<usize>().unwrap() - 1;
        let dst_stack: usize = parts[5].parse::<usize>().unwrap() - 1;
        let src_stack_len = cur_state[src_stack].len();
        let mut moving_crates = cur_state[src_stack].split_off(src_stack_len - num_to_move);
        cur_state[dst_stack].append(&mut moving_crates);
    }

    let output = cur_state
        .into_iter()
        .map(|stack| stack.last().copied().unwrap())
        .collect::<Vec<&str>>()
        .join("");
    println!("part 2 {}", output);
}

fn chunker(input: &str) -> Vec<&str> {
    let mut res: Vec<&str> = Vec::new();
    let width = 4;
    let mut i = 0;
    while i + width < input.len() {
        res.push(&input[i..i + width]);
        i += width;
    }
    res.push(&input[i..]);

    return res;
}
