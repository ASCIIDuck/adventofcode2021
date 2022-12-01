use std::fs::File;
use std::io::prelude::*;
use std::path::Path;

#[cfg(test)]
mod tests {
    #[test]
    fn it_works() {
        let result = 2 + 2;
        assert_eq!(result, 4);
    }
}

pub fn read_input(path: &str) -> String {
    let path = Path::new(path);
    let mut file = match File::open(&path) {
        Err(why) => panic!("Couldn't open {}: {}", path.display(), why),
        Ok(file) => file,
    };

    let mut s = String::new();
    match file.read_to_string(&mut s) {
        Err(why) => panic!("Couldn't read {}: {}", path.display(), why),
        Ok(i) => i
    };

    return s
}
