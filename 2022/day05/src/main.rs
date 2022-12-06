use util::read_input;
fn main() {
    let data = read_input("input.txt");
    let lines: Vec<&str> =  data.split("\n").collect::<Vec<&str>>();
    let mut initial_state: Vec<&str> = Vec::new();
    for line in lines {
        if line == "" {
            break;
        }
        initial_state.push(line);
    }
    let mut reverse_is_iter = initial_state.into_iter().rev();
    let num_stacks = reverse_is_iter.next().unwrap().split_whitespace().collect::<Vec<&str>>().len();
    let mut cur_state: Vec<Vec<&str>> = vec![ Vec::new(); num_stacks];
    for line in reverse_is_iter {
        let parts = chunker(line);
        for i in 0..parts.len() {
            let elf_crate = parts[i].trim_matches(|c| c == '[' || c == ' ' || c == ']');
            if elf_crate != "" {
                cur_state[i].push(elf_crate);
            }
        }
    }
    for stack in cur_state {
        println!("{:?}", stack);
    }

}

fn chunker(input: &str) -> Vec<&str> {
    let mut res: Vec<&str> = Vec::new();
    let width = 4;
    let mut i = 0;
    while i+width < input.len() {
        res.push(&input[i..i+width]);
        i+=width;
    }
    res.push(&input[i..]);

    return res;
}

