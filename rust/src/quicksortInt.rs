fn quicksort_int(array: &mut [i32], start: usize, end: usize) {
    if start >= end {
        return;
    }
    let pivot = array[end];
    let mut left = 0;
    let mut right = 0;
    while left + right < end - start {
        let num = array[start + left];
        if num < pivot {
            left += 1;
        } else {
            array.swap(start + left, end - right - 1);
            array.swap(end - right - 1, end - right);
            right += 1;
        }
    }
    quicksort_int(array, start, start + left - 1);
    quicksort_int(array, start + left + 1, end);
}
