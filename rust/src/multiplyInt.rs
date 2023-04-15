fn multiply_int(a: i32, b: i32, n: i32) -> i32 {
    let mut c = 1;
    for _ in 0..n {
        c = c * a * b;
    }
    c
}
