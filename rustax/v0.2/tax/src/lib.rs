pub fn add(a: i32, b: i32) -> i32 {
    a + b
}

#[cfg(test)]
mod tests {
    #[test]
    fn add() {
        let result = crate::add(2, 2);
        assert_eq!(result, 4);
    }
}
