# wasmcv

WASM/WASI bindings for computer vision

## WASM Component Generation

WASM Guest bindings are generated using `wit-bindgen` v0.31 or above

https://github.com/bytecodealliance/wit-bindgen

### TinyGo

```shell
wit-bindgen tiny-go --out-dir ./components/tinygo/ ./wit/
```

### Rust

```shell
wit-bindgen rust --out-dir ./components/rust/ ./wit/
```

### C

```shell
wit-bindgen c --out-dir ./components/c/ ./wit/
```

