type Input = Vec<isize>;

pub fn puzzle_1(input: Input) -> usize {
    7
}

pub fn puzzle_2(input: Input) -> usize {
    7
}

#[cfg(test)]
mod tests {
    use crate::util;
    use super::*;

    #[test]
    fn it_works() {
        assert_eq!(7, puzzle_1(parse(include_str!("puzzle_1_test.txt"))));
        assert_eq!(1655, puzzle_1(parse(include_str!("puzzle_1.txt"))));

        assert_eq!(5, puzzle_2(parse(include_str!("puzzle_1_test.txt"))));
        assert_eq!(1683, puzzle_2(parse(include_str!("puzzle_1.txt"))));
    }

    fn parse(s: &str) -> Input {
        s.lines().map(|s| util::parse_numbers(s)[0]).collect()
    }
}