fn multiply_int_vec(src1: &[i32], src2: &[i32], res: &mut [i32], n: usize) {
    for i in 0..n {
        res[i] = src1[i] * src2[i];
    }
}
