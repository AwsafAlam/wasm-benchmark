fn sum_double(array: &[f64], n: usize) -> f64 {
    let mut s = 0.0;
    for i in 0..n {
        s += array[i];
    }
    s
}
