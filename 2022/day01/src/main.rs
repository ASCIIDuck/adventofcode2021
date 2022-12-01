use util::read_input;

fn main() {
    let data = read_input("input.txt");
    let lines = data.split("\n");
    for line in lines {
        println!("{}", line);
    }
}
