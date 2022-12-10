use util::read_input;

#[derive(Copy, Clone)]
struct Node<T> {
    data: T,
    parent: Option<usize>,
    first_child: Option<usize>,
    last_child: Option<usize>,
    prev_sibling: Option<usize>,
    next_sibling: Option<usize>,
}

impl<T> Node<T> {
    fn new(data: T) -> Node<T> {
        return Node::<T> {
            data: data,
            parent: None,
            first_child: None,
            last_child: None,
            prev_sibling: None,
            next_sibling: None,
        };
    }
}

#[derive(Copy, Clone)]
struct File<'a> {
    name: &'a str,
    size: u32,
}
struct Tree<T> {
    nodes: Vec<Node<T>>,
    root: Option<usize>,
}

impl<T> Tree<T> {
    fn new() -> Tree<T> {
        return Tree {
            nodes: Vec::<Node<T>>::new(),
            root: None,
        };
    }
    fn add_node(&mut self, parent: Option<usize>, value: T) -> usize {
        let mut node: Node<T> = Node::<T>::new(value);
        let new_index = self.nodes.len();
        match parent {
            Some(p) => {
                node.parent = Some(p);
                let prev_sibling = self.nodes[p].last_child;
                node.prev_sibling = prev_sibling;
                self.nodes.push(node);
                match prev_sibling {
                    Some(prev_sibling) => self.nodes[prev_sibling].next_sibling = Some(new_index),
                    None => (),
                }
                match self.nodes[p].first_child {
                    Some(_) => (),
                    None => {
                        self.nodes[p].first_child = Some(new_index);
                    }
                }
                self.nodes[p].last_child = Some(new_index);
            }
            None => {
                self.nodes.push(node);
                self.root = Some(new_index);
            }
        }
        return new_index;
    }

    fn recursive_walk_and_talk(
        &mut self,
        starting: usize,
        selector: &dyn Fn(&Node<T>) -> Option<usize>,
        mutator: &dyn Fn(&mut T),
    ) {
        let cur = selector(&self.nodes[starting]);
        match cur {
            Some(selected) => {
                mutator(&mut self.nodes[selected].data);
                self.recursive_walk_and_talk(selected, selector, mutator);
            }
            None => (),
        }
    }
    fn recursive_walk_and_find<U>(
        &self,
        starting: usize,
        get_value: &dyn Fn(&T) -> U,
        aggregator: &dyn Fn(&Vec<U>) -> U,
    ) -> U {
        let mut results = Vec::<U>::new();
        results.push(get_value(&self.nodes[starting].data));
        match self.nodes[starting].next_sibling {
            Some(sib) => {
                results.push(self.recursive_walk_and_find(sib, get_value, aggregator));
            }
            None => (),
        }
        match self.nodes[starting].first_child {
            Some(sib) => {
                results.push(self.recursive_walk_and_find(sib, get_value, aggregator));
            }
            None => (),
        }

        return aggregator(&results);
    }
}

fn main() {
    //let data = read_input("input.txt");
    let data = read_input("sample.txt");
    let mut filesystem = Tree::<File>::new();
    let mut cwd = filesystem.add_node(None, File { name: "/", size: 0 });

    let lines = data.lines().collect::<Vec<&str>>();

    let mut line_no = 0;
    while line_no < lines.len() {
        let line = lines[line_no];
        if line.starts_with("$ cd ") {
            let target = &line[5..];
            if target == "/" {
                cwd = filesystem.root.unwrap();
            } else if target == ".." {
                cwd = filesystem.nodes[cwd].parent.unwrap();
            } else {
                let mut cur_child = filesystem.nodes[cwd].first_child;
                while cur_child != None {
                    if filesystem.nodes[cur_child.unwrap()].data.name == target {
                        cwd = cur_child.unwrap();
                        break;
                    }
                    cur_child = filesystem.nodes[cur_child.unwrap()].next_sibling;
                }
                if cur_child == None {
                    println!("Failed to cd to '{}'", target);
                }
            }
        } else if line.starts_with("$ ls") {
            line_no += 1;
            while line_no < lines.len() && !lines[line_no].starts_with("$") {
                let parts = lines[line_no].split_whitespace().collect::<Vec<&str>>();
                if parts[0] == "dir" {
                    filesystem.add_node(
                        Some(cwd),
                        File {
                            name: parts[1],
                            size: 0,
                        },
                    );
                } else {
                    let file_size = parts[0].parse::<u32>();
                    match file_size {
                        Ok(file_size) => {
                            let n = filesystem.add_node(
                                Some(cwd),
                                File {
                                    name: parts[1],
                                    size: file_size,
                                },
                            );
                            filesystem.recursive_walk_and_talk(
                                n,
                                &|cur: &Node<File>| -> Option<usize> { cur.parent },
                                &|f: &mut File| f.size += file_size,
                            );
                        }
                        Err(msg) => {
                            panic!("Failed to parse '{}': {}", parts[0], msg);
                        }
                    }
                }
                line_no += 1;
            }
            continue;
        }
        line_no += 1;
    }
    let part1_res = filesystem.recursive_walk_and_find(
        filesystem.root.unwrap(),
        &|file: &File| {
            if file.size <= 100000 {
                return file.size;
            } else {
                return 0;
            }
        },
        &|v: &Vec<u32>| {
            return v.iter().sum();
        },
    );
    println!("Part 1 is {}", part1_res);
}
