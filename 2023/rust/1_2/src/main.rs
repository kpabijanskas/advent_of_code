use std::fs;

fn main() {
    let content = fs::read_to_string("input").expect("file not found");

    let mut sum: u32 = 0;

    for line in content.lines() {
        let mut digit = "".to_string();

        for i in 0..line.len() {
            if let Some(x) = get_digit(&line[i..].to_string()) {
                digit.push(x);
                break;
            }
        }

        for i in (0..line.len()).rev() {
            if let Some(x) = get_digit(&line[i..].to_string()) {
                digit.push(x);
                break;
            }
        }

        if let Ok(digit) = digit.parse::<u32>() {
            sum += digit
        } else {
            panic!("NaN")
        }
    }

    println!("{}", sum)
}

fn get_digit(s: &String) -> Option<char> {
    if s.chars().nth(0).unwrap().is_digit(10) {
        Some(s.chars().nth(0).unwrap())
    } else if s.starts_with("one") {
        Some('1')
    } else if s.starts_with("two") {
        Some('2')
    } else if s.starts_with("three") {
        Some('3')
    } else if s.starts_with("four") {
        Some('4')
    } else if s.starts_with("five") {
        Some('5')
    } else if s.starts_with("six") {
        Some('6')
    } else if s.starts_with("seven") {
        Some('7')
    } else if s.starts_with("eight") {
        Some('8')
    } else if s.starts_with("nine") {
        Some('9')
    } else {
        None
    }
}
