#![no_std]

extern crate core;
extern crate alloc;

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
