# wasmCV

Rust bindings for wasmCV WebAssembly interfaces to computer vision systems.

See https://github.com/wasmvision/wasmcv for information about wasmCV.

[![Package](https://img.shields.io/crates/v/wasmcv.svg)](https://crates.io/crates/wasmcv)

## How to use

This example Rust module exports a `process()` function to the WASM host application. When the host calls the processor, it passes in the wasmCV image `Mat` to be processed. The wasmCV module then calls functions on that `Mat` which are handled by the host application, by calling OpenCV to actually perform the operations.

```rust
#![no_std]

extern crate core;
extern crate wee_alloc;
extern crate alloc;
extern crate wasmcv;

use alloc::string::String;
use alloc::string::ToString;
use wasmcv::wasm::cv;

#[no_mangle]
pub extern fn process(mat: cv::mat::Mat) -> cv::mat::Mat {
    println(&["Cols: ", &mat.cols().to_string(), " Rows: ", &mat.rows().to_string()].concat());

    return mat;
}

/// Print a message to the host [`_println`].
fn println(message: &String) {
    unsafe {
        let (ptr, len) = string_to_ptr(message);
        _println(ptr, len);
    }
}

#[link(wasm_import_module = "hosted")]
extern "C" {
    #[link_name = "println"]
    fn _println(ptr: u32, size: u32);
}

unsafe fn string_to_ptr(s: &String) -> (u32, u32) {
    return (s.as_ptr() as u32, s.len() as u32);
}

// Use `wee_alloc` as the global allocator...for now.
#[global_allocator]
static ALLOC: wee_alloc::WeeAlloc = wee_alloc::WeeAlloc::INIT;
```

You can then compile this module using the Rust compiler.

```shell
cargo build --target wasm32-unknown-unknown --release
```

The `wasm32-unknown-unknown` target can be used with wasmCV to produce very lightweight guest modules when combined with `no_std`. The example above compiles to around 14k, including debug information.

```shell
-rwxrwxr-x 1 ron ron 14488 sep 12 14:23 processrs.wasm
```
