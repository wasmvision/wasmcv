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
