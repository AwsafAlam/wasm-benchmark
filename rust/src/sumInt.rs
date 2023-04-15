fn sum_int(array: &[i32], n: usize) -> i32 {
    let mut s = 0;
    for i in 0..n {
        s += array[i];
    }
    s
}
