# wasmCV

C bindings for wasmCV WebAssembly interfaces to computer vision systems.

See https://github.com/wasmvision/wasmcv for information about wasmCV.

## How to use

This example C module exports a `process()` function to the WASM host application. When the host calls the processor, it passes in the wasmCV image `Mat` to be processed. The wasmCV module then calls functions on that `Mat` which are handled by the host application, by calling OpenCV to actually perform the operations.

```c
#include <string.h>
#include "wasmcv/imports.h"

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
        -o build/processc.wasm process.c itoa.c /path/to/wasmcv/import.c /path/to/wasmcv/import_component_type.o \
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
