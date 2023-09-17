use regex::Regex;

pub fn parse_numbers(s: &str) -> Vec<isize> {
    let mut vec = Vec::new();
    let re = Regex::new(r"([-\d]+)").unwrap();
    for capture in re.captures_iter(&s) {
        for i in 1..capture.len() {
            if let Ok(num) = capture[i].parse::<isize>() {
                vec.push(num);
            }
        }
    }

    vec
}