# wasmCV

wasmCV provides WebAssembly guest interface bindings for computer vision applications based on [OpenCV](https://github.com/opencv/opencv).

It includes [WIT](https://github.com/WebAssembly/component-model/blob/main/design/mvp/WIT.md) files defining the interface to be used between a WebAssembly host application and a WASM guest module intended to process OpenCV `Mat` image frames.

These interface definitions are then used to generate WASM bindings for TinyGo, Rust, and C. Those bindings can then be used in a WASM guest module to call OpenCV functions implemented by the host to obtain information or perform operations on OpenCV image frames.

```mermaid
flowchart LR
    subgraph host
        OpenCV
        Runtime[WASM Runtime]<-->OpenCV
    end
    subgraph guest
        Runtime<--wasmCV-->processor-go.wasm
        Runtime<--wasmCV-->processor-rust.wasm
        Runtime<--wasmCV-->processor-c.wasm
        processor-go.wasm
        processor-rust.wasm
        processor-c.wasm
    end
```

## Example wasmCV modules

### TinyGo

This TinyGo module exports a `process()` function to the WASM host application, which passes in the wasmCV image `Mat` to be processed. It then calls functions on that `Mat` which are handled by the host application by calling OpenCV to actually perform the operations.

```go
package main

import (
	"github.com/hybridgroup/mechanoid/convert"
	"wasmcv.org/wasm/cv/mat"
)

//go:wasmimport hosted println
func println(ptr, size uint32)

//export process
func process(image mat.Mat) mat.Mat {
	println(convert.StringToWasmPtr("Cols: " +
		convert.IntToString(int(image.Cols())) +
		" Rows: " +
		convert.IntToString(int(image.Rows())) +
		" Type: " +
		convert.IntToString(int(image.Type()))))

	return image
}
```

Install the `wasmcv` package into your Go package:

```shell
go get wasmcv.org/wasm/cv
```

You can then compile this module using the TinyGo compiler.

```shell
tinygo build -o processor.wasm -target=wasm-unknown processor.go
```

Note that the `wasm-unknown` target can be used with wasmCV to produce very lightweight guest modules. The example above compiles to around 31k, including debug information.

```shell
-rwxrwxr-x 1 ron ron 31248 sep 11 11:00 processor.wasm
```

See the [basic example application here](./examples/basic) to give it a try.

### Rust

This Rust module does the same thing as the TinyGo wasm module example. It exports a `process()` function to the WASM host application, which then passes in the wasmCV image `Mat` to be processed. The module then calls functions on that `Mat` which are handled by the host application by calling OpenCV to actually perform the operations.

```rust
#![no_std]

extern crate core;
extern crate alloc;

use alloc::string::String;
use alloc::string::ToString;
use wasmcv::wasm::cv;

#[no_mangle]
pub extern fn process(mat: cv::mat::Mat) -> cv::mat::Mat {
    println(&["Cols: ", &mat.cols().to_string(), " Rows: ", &mat.rows().to_string(), " Size: ", &mat.size().len().to_string()].concat());

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

#[no_mangle]
pub extern fn malloc(size: usize) -> *mut u8 {
    let layout = core::alloc::Layout::from_size_align(size, 1).unwrap();
    unsafe { alloc::alloc::alloc(layout) }
}
```

Install the `wasmcv` crate into your Rust project:

```shell
cargo add wasmcv
```

You can then compile this module using the Rust compiler.

```shell
cargo build --target wasm32-unknown-unknown --release
```

The `wasm32-unknown-unknown` target can be used with wasmCV to produce very lightweight guest modules when combined with `no_std`. The example above compiles to around 14k, including debug information.

```shell
-rwxrwxr-x 1 ron ron 14488 sep 12 14:23 processrs.wasm
```

See the [multi example application here](./examples/multi) to try wasmCV with Rust.

### C

This C module does the same thing as the TinyGo and Rust wasm module examples. It exports a `process()` function to the WASM host application, which then passes in the wasmCV image `Mat` to be processed. The module then calls functions on that `Mat` which are handled by the host application by calling OpenCV to actually perform the operations.

```c
#include <string.h>
#include "../../../../components/c/wasmcv/imports.h"

extern int itoa(int value, char *sp, int radix);

__attribute__((import_module("hosted"), import_name("println"))) void println(int32_t str, int32_t len);

wasm_cv_mat_borrow_mat_t process(wasm_cv_mat_borrow_mat_t image) {
    int32_t cols, rows;
    cols = wasm_cv_mat_method_mat_cols(image);
    rows = wasm_cv_mat_method_mat_rows(image);

    char buf[20];
    strcpy(buf, "Cols: ");
    itoa(cols, buf+6, 10);
    strcpy(buf+9, " Rows: ");
    itoa(rows, buf+16, 10);

    println((int32_t)buf, 20);

    return image;
}
```

You can then compile this module using the `clang` compiler.

```shell
/opt/wasi-sdk/bin/clang --target=wasm32-unknown-unknown -O3 \
        --sysroot="/path/to/lib/wasi-libc/sysroot" \
        -z stack-size=4096 -Wl,--initial-memory=65536 \
        -o ../processc.wasm process.c itoa.c ../../../../components/c/wasmcv/import.c ../../../../components/c/wasmcv/import_component_type.o \
        -Wl,--export=process \
        -Wl,--export=__data_end -Wl,--export=__heap_base \
        -Wl,--strip-all,--no-entry \
        -Wl,--unresolved-symbols=ignore-all \
        -nostdlib \
```

The `wasm32-unknown-unknown` target can be used with wasmCV to produce very lightweight guest modules. The example above compiles to just under 3k, including debug information.

```shell
-rwxrwxr-x 1 ron ron  2916 sep 13 20:03 processc.wasm
```

See the [multi example application here](./examples/multi) to try wasmCV with C.

## WASM Component Generation

WASM Guest bindings are generated using `wit-bindgen` v0.38 or above.

https://github.com/bytecodealliance/wit-bindgen

Go bindings are generated by `wit-bindgen-go` command line tool v0.5.0 or above.

https://github.com/bytecodealliance/go-modules

### TinyGo

```shell
wit-bindgen-go generate --out ./components/tinygo -w imports -p wasmcv.org ./wit/
```

Note that the TinyGo bindings are a git submodule. When regenerating the submodule must be updated in order to update the separate Go package repo.

### Rust

```shell
wit-bindgen rust --out-dir ./components/rust/wasmcv/src -w imports ./wit/
```

### C

```shell
wit-bindgen c --out-dir ./components/c/wasmcv/ -w imports ./wit/
```

