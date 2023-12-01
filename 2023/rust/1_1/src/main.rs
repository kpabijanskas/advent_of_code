use std::fs;

fn main() {
    let content = fs::read_to_string("input").expect("file not found");

    let mut sum: u32 = 0;

    for line in content.lines() {
        let mut digit = "".to_string();

        for char in line.chars() {
            if char.is_digit(10) {
                digit.push(char);
                break;
            }
        }

        for char in line.chars().rev() {
            if char.is_digit(10) {
                digit.push(char);
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
