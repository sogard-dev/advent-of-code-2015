type Input = Vec<Floor>;

#[derive(PartialEq)]
enum Floor {
    UP,
    DOWN,
}

impl Floor {
    fn to_int(&self) -> i32 {
        if self == &Floor::DOWN {
            return -1;
        }
        return 1;
    }
}

fn puzzle_1(input: Input) -> i32 {
    input.iter().fold(0_i32, |sum, f| sum + f.to_int())
}

fn puzzle_2(input: Input) -> i32 {
    let mut sum = 0;
    for (i, f) in input.iter().enumerate() {
        sum += f.to_int();
        if sum == -1 {
            return (i + 1) as i32;
        }
    }
    0
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(0, puzzle_1(parse("(())")));
        assert_eq!(0, puzzle_1(parse("()()")));
        assert_eq!(3, puzzle_1(parse("(((")));
        assert_eq!(3, puzzle_1(parse("(()(()(")));
        assert_eq!(-3, puzzle_1(parse(")())())")));
        assert_eq!(74, puzzle_1(parse(include_str!("puzzle_1.txt"))));

        assert_eq!(1, puzzle_2(parse(")")));
        assert_eq!(5, puzzle_2(parse("()())")));

        assert_eq!(1795, puzzle_2(parse(include_str!("puzzle_1.txt"))));
    }

    fn parse(s: &str) -> Input {
        s.chars().map(|c| {
            if c.eq(&'(') {
                Floor::UP
            } else {
                Floor::DOWN
            }
        }).collect()
    }
}