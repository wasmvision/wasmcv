#![no_std]

extern crate core;
extern crate wee_alloc;
extern crate alloc;

use alloc::{string::String};
use alloc::string::ToString;
use alloc::format;

#[path = "../../../../../components/rust/cv.rs"]
mod cv;

#[no_mangle]
pub extern fn process(mat: cv::Mat) -> cv::Mat {
    let cols: String = String::from("Cols:");
    let rows: String = String::from("Rows:");

    let col: String = mat.cols().to_string();
    let row: String = mat.rows().to_string();

    let msg = format!("{cols} {col} {rows} {row}");
    println(&msg);

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

#[panic_handler]
fn handle_panic(_: &core::panic::PanicInfo) -> ! {
    core::arch::wasm32::unreachable()
}
