### Investigating WebAssembly Runtime Performance for C++, Rust and Go

WASM_Benchmark: A test suite for benchmarking C++, Rust and Go runtimes when converted to WebAssembly

---

WebAssembly has emerged as a new technology to address the problem of slow execution time of JavaScript-based web applications. WebAssembly allows developers to write high-performance code in languages other than JavaScript that can be compiled to WebAssembly bytecode. However, the choice of source language can affect the execution time of the resulting WebAssembly code. In this research paper, we compare the performance of WebAssembly code generated from C++, Rust and Go source languages. We use real-world use cases and benchmarks to evaluate the performance of each language and provide guidelines for selecting the appropriate source language for different types of web applications.

## Demo

[Online Demo](https://awsaf.dev/wasm-benchmark)

## Develop locally

```sh
$ git clone https://github.com/AwsafAlam/wasm-benchmark.git
$ cd WebAssembly-benchmark
$ ./scripts/run_server.sh
# access http://localhost:8080
```

## Emscripten Compilation

we use emscripted `emcc` for compilation from C, Rust and Go-lang to wasm modules.

## Running HTML files with emrun

emrun is a command line tool that can run generated HTML pages via a locally launched web server.
