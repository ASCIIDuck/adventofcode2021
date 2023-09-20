use util::read_input;
use std::collections::HashMap;

#[derive(PartialEq)]
#[derive(Eq)]
#[derive(Hash)]
enum Throw {
    Rock,
    Paper,
    Scissors
}
#[derive(PartialEq)]
enum Guidance {
    Lose,
    Draw,
    Win
}



fn main() {
    let throw_translations  = HashMap::from ([
        ("A", Throw::Rock),
        ("B", Throw::Paper),
        ("C", Throw::Scissors),
        ("X", Throw::Rock),
        ("Y", Throw::Paper),
        ("Z", Throw::Scissors),
    ]);

    let guidances  = HashMap::from ([
        ("X", Guidance::Lose),
        ("Y", Guidance::Draw),
        ("Z", Guidance::Win),
    ]);

    let throw_values = HashMap::from([
        (Throw::Rock, 1),
        (Throw::Paper, 2),
        (Throw::Scissors, 3),
    ]);

    let winning_throw = HashMap::from([
        (Throw::Rock, Throw::Paper),
        (Throw::Paper, Throw::Scissors),
        (Throw::Scissors, Throw::Rock),
    ]);
    let losing_throw = HashMap::from([
        (Throw::Scissors, Throw::Paper),
        (Throw::Rock, Throw::Scissors),
        (Throw::Paper, Throw::Rock),
    ]);
    let draw_score = 3;
    let win_score = 6;
    let lose_score = 0;


    let data = read_input("input.txt");
    let lines = data.split("\n");
    let mut naive_score = 0;
    let mut true_score = 0;
    for line in lines {
        if line == "" {
            continue;
        }
        let parts: Vec<&str> = line.split(" ").collect();
        let opponent = throw_translations.get(parts[0]).unwrap();
        let my_naive_throw = throw_translations.get(parts[1]).unwrap();
        let guidance = guidances.get(parts[1]).unwrap();
        naive_score += throw_values.get(&my_naive_throw).unwrap();
        if *my_naive_throw == *opponent {
            naive_score += draw_score;
        }
        else if *my_naive_throw ==  *winning_throw.get(&opponent).unwrap() {
            naive_score += win_score;
        }
        else {
            naive_score += lose_score;
        }

        if *guidance == Guidance::Lose {
            true_score += lose_score;
            true_score += throw_values.get(losing_throw.get(&opponent).unwrap()).unwrap();
        }
        else if *guidance == Guidance::Draw {
            true_score += draw_score;
            true_score += throw_values.get(&opponent).unwrap();
        }
        else if *guidance == Guidance::Win {
            true_score += win_score;
            true_score += throw_values.get(winning_throw.get(&opponent).unwrap()).unwrap();
        }

    }
    println!("My naive score: {}", naive_score);
    println!("My true score: {}", true_score);
}
