fn multiply_double(a: f64, b: f64, n: i32) -> f64 {
    let mut c = 1.0;
    for _ in 0..n {
        c = c * a * b;
    }
    return c;
}
