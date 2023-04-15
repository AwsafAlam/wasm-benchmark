fn fib(n: u32) -> u32 {
    match n {
        1 => 1,
        2 => 1,
        _ => fib(n - 1) + fib(n - 2),
    }
}

#[no_mangle]
pub extern "C" fn main(x: u32) -> u32 {
    fib(x)
}
