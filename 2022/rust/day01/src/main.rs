use util::read_input;

fn main() {
    let data = read_input("input.txt");
    let lines = data.split("\n");
    let mut i: i32;
    let mut cur: i32 = 0;
    let mut score_board: [i32; 3] = [0, 0, 0];
    for line in lines {
        if line == "" {
            insert(&mut score_board, cur);
            cur = 0;
        }
        else {
            i = line.parse::<i32>().unwrap();
            cur += i;
        }
    }

    println!("Highest calorie elves:");
    let mut total: i32 = 0;
    for i in 0..score_board.len() {
        println!("\t{}", score_board[i]);
        total += score_board[i];
    }
    println!("Total: {}", total);
}

fn insert(scores: &mut [i32; 3], new_num: i32) -> &[i32] {
    let mut cur: i32 = new_num;
    for i in 0..scores.len() {
        if cur > scores[i] {
            let tmp = scores[i];
            scores[i]=cur;
            cur = tmp;
        }
    }
    return scores;
}
