# <a id="imports"></a>World imports

wasmCV is a WebAssembly guest module interface for computer vision based on OpenCV.

 - Imports:
    - interface `wasm:cv/types`
    - interface `wasm:cv/mat`
    - interface `wasm:cv/cv`
    - interface `wasm:cv/dnn`
    - interface `wasm:cv/objdetect`
    - interface `wasm:cv/features2d`
 - Exports:
    - interface `wasm:cv/request`

## <a id="wasm_cv_types"></a>Import interface wasm:cv/types


----

### Types

#### <a id="error_result"></a>`type error-result`
`string`
<p>
#### <a id="size"></a>`record size`

size is a 2-element integer vector.
It represents a width and height.

##### Record Fields

- <a id="size.x"></a>`x`: `s32`
- <a id="size.y"></a>`y`: `s32`
#### <a id="point"></a>`type point`
[`size`](#size)
<p>point is a 2-element integer vector.
It represents a x and y coordinate.

#### <a id="scalar"></a>`record scalar`

scalar is a 4-element floating point vector.

##### Record Fields

- <a id="scalar.val1"></a>`val1`: `f32`
- <a id="scalar.val2"></a>`val2`: `f32`
- <a id="scalar.val3"></a>`val3`: `f32`
- <a id="scalar.val4"></a>`val4`: `f32`
#### <a id="rect"></a>`record rect`

rect is a rectangle with integer coordinates.
It is represented by the top-left corner and the bottom-right corner.

##### Record Fields

- <a id="rect.min"></a>`min`: [`size`](#size)
- <a id="rect.max"></a>`max`: [`size`](#size)
#### <a id="rgba"></a>`record RGBA`

RGBA is a color with red, green, blue, and alpha channels.

##### Record Fields

- <a id="rgba.r"></a>`r`: `u8`
- <a id="rgba.g"></a>`g`: `u8`
- <a id="rgba.b"></a>`b`: `u8`
- <a id="rgba.a"></a>`a`: `u8`
#### <a id="border_type"></a>`enum border-type`

border-type is a type of border to add to an image.

##### Enum Cases

- <a id="border_type.border_constant"></a>`border-constant`
- <a id="border_type.border_replicate"></a>`border-replicate`
- <a id="border_type.border_reflect"></a>`border-reflect`
- <a id="border_type.border_wrap"></a>`border-wrap`
- <a id="border_type.border_reflect101"></a>`border-reflect101`
- <a id="border_type.border_transparent"></a>`border-transparent`
- <a id="border_type.border_default"></a>`border-default`
- <a id="border_type.border_isolated"></a>`border-isolated`
#### <a id="adaptive_threshold_type"></a>`enum adaptive-threshold-type`

adaptive-threshold-type is a type of adaptive thresholding.

##### Enum Cases

- <a id="adaptive_threshold_type.adaptive_threshold_mean"></a>`adaptive-threshold-mean`
- <a id="adaptive_threshold_type.adaptive_threshold_gaussian"></a>`adaptive-threshold-gaussian`
#### <a id="threshold_type"></a>`enum threshold-type`

threshold-type is a type of thresholding.

##### Enum Cases

- <a id="threshold_type.threshold_binary"></a>`threshold-binary`
- <a id="threshold_type.threshold_binary_inv"></a>`threshold-binary-inv`
- <a id="threshold_type.threshold_trunc"></a>`threshold-trunc`
- <a id="threshold_type.threshold_to_zero"></a>`threshold-to-zero`
- <a id="threshold_type.threshold_to_zero_inv"></a>`threshold-to-zero-inv`
- <a id="threshold_type.threshold_mask"></a>`threshold-mask`
- <a id="threshold_type.threshold_otsu"></a>`threshold-otsu`
- <a id="threshold_type.tthreshold_triangle"></a>`tthreshold-triangle`
#### <a id="data_layout_type"></a>`enum data-layout-type`

data-layout-type is a type of data layout.

##### Enum Cases

- <a id="data_layout_type.data_layout_unknown"></a>`data-layout-unknown`
- <a id="data_layout_type.data_layout_nd"></a>`data-layout-nd`
- <a id="data_layout_type.data_layout_nchw"></a>`data-layout-nchw`
- <a id="data_layout_type.data_layout_ncdhw"></a>`data-layout-ncdhw`
- <a id="data_layout_type.data_layout_nhwc"></a>`data-layout-nhwc`
- <a id="data_layout_type.data_layout_ndhwc"></a>`data-layout-ndhwc`
- <a id="data_layout_type.data_layout_planar"></a>`data-layout-planar`
#### <a id="padding_mode_type"></a>`enum padding-mode-type`


##### Enum Cases

- <a id="padding_mode_type.padding_mode_null"></a>`padding-mode-null`
- <a id="padding_mode_type.padding_mode_crop_center"></a>`padding-mode-crop-center`
- <a id="padding_mode_type.padding_mode_letterbox"></a>`padding-mode-letterbox`
#### <a id="blob_params"></a>`record blob-params`


##### Record Fields

- <a id="blob_params.scale_factor"></a>`scale-factor`: `f32`
- <a id="blob_params.size"></a>`size`: [`size`](#size)
- <a id="blob_params.mean"></a>`mean`: [`scalar`](#scalar)
- <a id="blob_params.swap_rb"></a>`swap-RB`: `bool`
- <a id="blob_params.ddepth"></a>`ddepth`: `u8`
- <a id="blob_params.data_layout"></a>`data-layout`: [`data-layout-type`](#data_layout_type)
- <a id="blob_params.padding_mode"></a>`padding-mode`: [`padding-mode-type`](#padding_mode_type)
- <a id="blob_params.border"></a>`border`: [`scalar`](#scalar)
#### <a id="mix_max_loc_result"></a>`record mix-max-loc-result`


##### Record Fields

- <a id="mix_max_loc_result.min_val"></a>`min-val`: `f32`
- <a id="mix_max_loc_result.max_val"></a>`max-val`: `f32`
- <a id="mix_max_loc_result.min_loc"></a>`min-loc`: [`size`](#size)
- <a id="mix_max_loc_result.max_loc"></a>`max-loc`: [`size`](#size)
#### <a id="hershey_font_type"></a>`enum hershey-font-type`


##### Enum Cases

- <a id="hershey_font_type.hershey_font_simplex"></a>`hershey-font-simplex`
- <a id="hershey_font_type.hershey_font_plain"></a>`hershey-font-plain`
- <a id="hershey_font_type.hershey_font_duplex"></a>`hershey-font-duplex`
- <a id="hershey_font_type.hershey_font_complex"></a>`hershey-font-complex`
- <a id="hershey_font_type.hershey_font_triplex"></a>`hershey-font-triplex`
- <a id="hershey_font_type.hershey_font_complex_small"></a>`hershey-font-complex-small`
- <a id="hershey_font_type.hershey_font_script_simplex"></a>`hershey-font-script-simplex`
- <a id="hershey_font_type.hershey_font_script_complex"></a>`hershey-font-script-complex`
- <a id="hershey_font_type.hershey_font_italic"></a>`hershey-font-italic`
#### <a id="interpolation_type"></a>`enum interpolation-type`


##### Enum Cases

- <a id="interpolation_type.interpolation_nearest"></a>`interpolation-nearest`
- <a id="interpolation_type.interpolation_linear"></a>`interpolation-linear`
- <a id="interpolation_type.interpolation_cubic"></a>`interpolation-cubic`
- <a id="interpolation_type.interpolation_area"></a>`interpolation-area`
- <a id="interpolation_type.interpolation_lanczos4"></a>`interpolation-lanczos4`
#### <a id="color_coversion_type"></a>`enum color-coversion-type`


##### Enum Cases

- <a id="color_coversion_type.color_bgr_to_bgra"></a>`color-BGR-to-BGRA`
- <a id="color_coversion_type.color_rgb_to_rgba"></a>`color-RGB-to-RGBA`
- <a id="color_coversion_type.color_bgra_to_bgr"></a>`color-BGRA-to-BGR`
- <a id="color_coversion_type.color_rgba_to_rgb"></a>`color-RGBA-to-RGB`
- <a id="color_coversion_type.color_bgr_to_rgba"></a>`color-BGR-to-RGBA`
- <a id="color_coversion_type.color_rgb_to_bgra"></a>`color-RGB-to-BGRA`
- <a id="color_coversion_type.color_rgba_to_bgr"></a>`color-RGBA-to-BGR`
- <a id="color_coversion_type.color_bgra_to_rgb"></a>`color-BGRA-to-RGB`
- <a id="color_coversion_type.color_bgr_to_rgb"></a>`color-BGR-to-RGB`
- <a id="color_coversion_type.color_rgb_to_bgr"></a>`color-RGB-to-BGR`
- <a id="color_coversion_type.color_bgra_to_rgba"></a>`color-BGRA-to-RGBA`
- <a id="color_coversion_type.color_rgba_to_bgra"></a>`color-RGBA-to-BGRA`
- <a id="color_coversion_type.color_bgr_to_gray"></a>`color-BGR-to-gray`
- <a id="color_coversion_type.color_rgb_to_gray"></a>`color-RGB-to-gray`
- <a id="color_coversion_type.color_gray_to_bgr"></a>`color-gray-to-BGR`
- <a id="color_coversion_type.color_gray_to_rgb"></a>`color-gray-to-RGB`
- <a id="color_coversion_type.color_gray_to_bgra"></a>`color-gray-to-BGRA`
- <a id="color_coversion_type.color_gray_to_rgba"></a>`color-gray-to-RGBA`
- <a id="color_coversion_type.color_bgra_to_gray"></a>`color-BGRA-to-gray`
- <a id="color_coversion_type.color_rgba_to_gray"></a>`color-RGBA-to-gray`
#### <a id="morph_shape"></a>`enum morph-shape`


##### Enum Cases

- <a id="morph_shape.morph_rect"></a>`morph-rect`
- <a id="morph_shape.morph_cross"></a>`morph-cross`
- <a id="morph_shape.morph_ellipse"></a>`morph-ellipse`
#### <a id="key_point"></a>`record key-point`


##### Record Fields

- <a id="key_point.x"></a>`x`: `f32`
- <a id="key_point.y"></a>`y`: `f32`
- <a id="key_point.size"></a>`size`: `f32`
- <a id="key_point.angle"></a>`angle`: `f32`
- <a id="key_point.response"></a>`response`: `f32`
- <a id="key_point.octave"></a>`octave`: `s32`
- <a id="key_point.class_id"></a>`class-id`: `s32`
#### <a id="d_match"></a>`record d-match`

DMatch is data structure for matching keypoint descriptors.

For further details, please see:
https://docs.opencv.org/4.x/d4/de0/classcv_1_1DMatch.html#a546ddb9a87898f06e510e015a6de596e

##### Record Fields

- <a id="d_match.query_idx"></a>`query-idx`: `u32`
- <a id="d_match.train_idx"></a>`train-idx`: `u32`
- <a id="d_match.img_idx"></a>`img-idx`: `u32`
- <a id="d_match.distance"></a>`distance`: `f64`
## <a id="wasm_cv_mat"></a>Import interface wasm:cv/mat

mat resource is a matrix of bytes, wrapper around the cv::Mat type.

----

### Types

#### <a id="error_result"></a>`type error-result`
[`error-result`](#error_result)
<p>
#### <a id="mix_max_loc_result"></a>`type mix-max-loc-result`
[`mix-max-loc-result`](#mix_max_loc_result)
<p>
#### <a id="rect"></a>`type rect`
[`rect`](#rect)
<p>
#### <a id="mattype"></a>`enum mattype`


##### Enum Cases

- <a id="mattype.cv8u"></a>`cv8u`
- <a id="mattype.cv8s"></a>`cv8s`
- <a id="mattype.cv16u"></a>`cv16u`
- <a id="mattype.cv16s"></a>`cv16s`
- <a id="mattype.cv32s"></a>`cv32s`
- <a id="mattype.cv32f"></a>`cv32f`
- <a id="mattype.cv64f"></a>`cv64f`
#### <a id="mat"></a>`resource mat`

----

### Functions

#### <a id="constructor_mat"></a>`[constructor]mat: func`

Create a new Mat. id does not currently do anything.

##### Params

- <a id="constructor_mat.id"></a>`id`: `u32`

##### Return values

- <a id="constructor_mat.0"></a> own<[`mat`](#mat)>

#### <a id="static_mat_new_with_size"></a>`[static]mat.new-with-size: func`

Create a new Mat with the specified size and type.

##### Params

- <a id="static_mat_new_with_size.cols"></a>`cols`: `u32`
- <a id="static_mat_new_with_size.rows"></a>`rows`: `u32`
- <a id="static_mat_new_with_size.mattype"></a>`mattype`: [`mattype`](#mattype)

##### Return values

- <a id="static_mat_new_with_size.0"></a> own<[`mat`](#mat)>

#### <a id="method_mat_clone"></a>`[method]mat.clone: func`

Clone returns a cloned full copy of the Mat.

##### Params

- <a id="method_mat_clone.self"></a>`self`: borrow<[`mat`](#mat)>

##### Return values

- <a id="method_mat_clone.0"></a> own<[`mat`](#mat)>

#### <a id="method_mat_close"></a>`[method]mat.close: func`

Close the Mat

##### Params

- <a id="method_mat_close.self"></a>`self`: borrow<[`mat`](#mat)>

#### <a id="method_mat_cols"></a>`[method]mat.cols: func`

Cols returns the number of columns for this Mat.

##### Params

- <a id="method_mat_cols.self"></a>`self`: borrow<[`mat`](#mat)>

##### Return values

- <a id="method_mat_cols.0"></a> `u32`

#### <a id="method_mat_rows"></a>`[method]mat.rows: func`

Rows returns the number of rows for this Mat.

##### Params

- <a id="method_mat_rows.self"></a>`self`: borrow<[`mat`](#mat)>

##### Return values

- <a id="method_mat_rows.0"></a> `u32`

#### <a id="method_mat_region"></a>`[method]mat.region: func`

Region returns a new Mat that points to a region of this Mat. Changes made to the
region Mat will affect the original Mat, since they are pointers to the underlying
OpenCV Mat object.

##### Params

- <a id="method_mat_region.self"></a>`self`: borrow<[`mat`](#mat)>
- <a id="method_mat_region.rect"></a>`rect`: [`rect`](#rect)

##### Return values

- <a id="method_mat_region.0"></a> own<[`mat`](#mat)>

#### <a id="method_mat_copy_to"></a>`[method]mat.copy-to: func`

CopyTo copies Mat into destination Mat.

##### Params

- <a id="method_mat_copy_to.self"></a>`self`: borrow<[`mat`](#mat)>
- <a id="method_mat_copy_to.dst"></a>`dst`: borrow<[`mat`](#mat)>

#### <a id="method_mat_convert_to"></a>`[method]mat.convert-to: func`

ConvertTo converts Mat into destination Mat.

For further details, please see:
https://docs.opencv.org/4.x/d3/d63/classcv_1_1Mat.html#adf88c60c5b4980e05bb556080916978b

##### Params

- <a id="method_mat_convert_to.self"></a>`self`: borrow<[`mat`](#mat)>
- <a id="method_mat_convert_to.mattype"></a>`mattype`: [`mattype`](#mattype)

##### Return values

- <a id="method_mat_convert_to.0"></a> result<own<[`mat`](#mat)>, [`error-result`](#error_result)>

#### <a id="method_mat_convert_to_with_params"></a>`[method]mat.convert-to-with-params: func`

ConvertToWithParams converts Mat into destination Mat using additional params.

For further details, please see:
https://docs.opencv.org/4.x/d3/d63/classcv_1_1Mat.html#adf88c60c5b4980e05bb556080916978b

##### Params

- <a id="method_mat_convert_to_with_params.self"></a>`self`: borrow<[`mat`](#mat)>
- <a id="method_mat_convert_to_with_params.mattype"></a>`mattype`: [`mattype`](#mattype)
- <a id="method_mat_convert_to_with_params.alpha"></a>`alpha`: `f32`
- <a id="method_mat_convert_to_with_params.beta"></a>`beta`: `f32`

##### Return values

- <a id="method_mat_convert_to_with_params.0"></a> result<own<[`mat`](#mat)>, [`error-result`](#error_result)>

#### <a id="method_mat_mattype"></a>`[method]mat.mattype: func`

MatType returns the type of the Mat.

##### Params

- <a id="method_mat_mattype.self"></a>`self`: borrow<[`mat`](#mat)>

##### Return values

- <a id="method_mat_mattype.0"></a> [`mattype`](#mattype)

#### <a id="method_mat_size"></a>`[method]mat.size: func`

Size returns an array with one element for each dimension containing the size of that dimension for the Mat.

##### Params

- <a id="method_mat_size.self"></a>`self`: borrow<[`mat`](#mat)>

##### Return values

- <a id="method_mat_size.0"></a> list<`u32`>

#### <a id="method_mat_step"></a>`[method]mat.step: func`

Step returns the number of bytes each matrix row occupies.

##### Params

- <a id="method_mat_step.self"></a>`self`: borrow<[`mat`](#mat)>

##### Return values

- <a id="method_mat_step.0"></a> `u32`

#### <a id="method_mat_elemsize"></a>`[method]mat.elemsize: func`

ElemSize returns the matrix element size in bytes.

##### Params

- <a id="method_mat_elemsize.self"></a>`self`: borrow<[`mat`](#mat)>

##### Return values

- <a id="method_mat_elemsize.0"></a> `u32`

#### <a id="method_mat_empty"></a>`[method]mat.empty: func`

Empty returns true if the Mat is empty.

##### Params

- <a id="method_mat_empty.self"></a>`self`: borrow<[`mat`](#mat)>

##### Return values

- <a id="method_mat_empty.0"></a> `bool`

#### <a id="method_mat_get_float_at"></a>`[method]mat.get-float-at: func`

GetFloatAt returns the value at the specified row and column as a f32.

##### Params

- <a id="method_mat_get_float_at.self"></a>`self`: borrow<[`mat`](#mat)>
- <a id="method_mat_get_float_at.row"></a>`row`: `u32`
- <a id="method_mat_get_float_at.col"></a>`col`: `u32`

##### Return values

- <a id="method_mat_get_float_at.0"></a> `f32`

#### <a id="method_mat_set_float_at"></a>`[method]mat.set-float-at: func`

SetFloatAt sets the value at the specified row and column as a f32.

##### Params

- <a id="method_mat_set_float_at.self"></a>`self`: borrow<[`mat`](#mat)>
- <a id="method_mat_set_float_at.row"></a>`row`: `u32`
- <a id="method_mat_set_float_at.col"></a>`col`: `u32`
- <a id="method_mat_set_float_at.val"></a>`val`: `f32`

#### <a id="method_mat_get_uchar_at"></a>`[method]mat.get-uchar-at: func`

GetUCharAt returns the value at the specified row and column as a u8.

##### Params

- <a id="method_mat_get_uchar_at.self"></a>`self`: borrow<[`mat`](#mat)>
- <a id="method_mat_get_uchar_at.row"></a>`row`: `u32`
- <a id="method_mat_get_uchar_at.col"></a>`col`: `u32`

##### Return values

- <a id="method_mat_get_uchar_at.0"></a> `u8`

#### <a id="method_mat_set_uchar_at"></a>`[method]mat.set-uchar-at: func`

SetUCharAt sets the value at the specified row and column as a u8.

##### Params

- <a id="method_mat_set_uchar_at.self"></a>`self`: borrow<[`mat`](#mat)>
- <a id="method_mat_set_uchar_at.row"></a>`row`: `u32`
- <a id="method_mat_set_uchar_at.col"></a>`col`: `u32`
- <a id="method_mat_set_uchar_at.val"></a>`val`: `u8`

#### <a id="method_mat_get_int_at"></a>`[method]mat.get-int-at: func`

GetIntAt returns the value at the specified row and column as a s32.

##### Params

- <a id="method_mat_get_int_at.self"></a>`self`: borrow<[`mat`](#mat)>
- <a id="method_mat_get_int_at.row"></a>`row`: `u32`
- <a id="method_mat_get_int_at.col"></a>`col`: `u32`

##### Return values

- <a id="method_mat_get_int_at.0"></a> `s32`

#### <a id="method_mat_set_int_at"></a>`[method]mat.set-int-at: func`

SetIntAt sets the value at the specified row and column as a s32.

##### Params

- <a id="method_mat_set_int_at.self"></a>`self`: borrow<[`mat`](#mat)>
- <a id="method_mat_set_int_at.row"></a>`row`: `u32`
- <a id="method_mat_set_int_at.col"></a>`col`: `u32`
- <a id="method_mat_set_int_at.val"></a>`val`: `s32`

#### <a id="method_mat_get_float_at3"></a>`[method]mat.get-float-at3: func`

GetFloatAt3 returns the value at the specified x, y, z as a f32.

##### Params

- <a id="method_mat_get_float_at3.self"></a>`self`: borrow<[`mat`](#mat)>
- <a id="method_mat_get_float_at3.x"></a>`x`: `u32`
- <a id="method_mat_get_float_at3.y"></a>`y`: `u32`
- <a id="method_mat_get_float_at3.z"></a>`z`: `u32`

##### Return values

- <a id="method_mat_get_float_at3.0"></a> `f32`

#### <a id="method_mat_set_float_at3"></a>`[method]mat.set-float-at3: func`

SetFloatAt3 sets the value at the specified x, y, z as a f32.

##### Params

- <a id="method_mat_set_float_at3.self"></a>`self`: borrow<[`mat`](#mat)>
- <a id="method_mat_set_float_at3.x"></a>`x`: `u32`
- <a id="method_mat_set_float_at3.y"></a>`y`: `u32`
- <a id="method_mat_set_float_at3.z"></a>`z`: `u32`
- <a id="method_mat_set_float_at3.val"></a>`val`: `f32`

#### <a id="method_mat_get_uchar_at3"></a>`[method]mat.get-uchar-at3: func`

GetUCharAt3 returns the value at the specified x, y, z as a u8.

##### Params

- <a id="method_mat_get_uchar_at3.self"></a>`self`: borrow<[`mat`](#mat)>
- <a id="method_mat_get_uchar_at3.x"></a>`x`: `u32`
- <a id="method_mat_get_uchar_at3.y"></a>`y`: `u32`
- <a id="method_mat_get_uchar_at3.z"></a>`z`: `u32`

##### Return values

- <a id="method_mat_get_uchar_at3.0"></a> `u8`

#### <a id="method_mat_set_uchar_at3"></a>`[method]mat.set-uchar-at3: func`

SetUCharAt3 sets the value at the specified x, y, z as a u8.

##### Params

- <a id="method_mat_set_uchar_at3.self"></a>`self`: borrow<[`mat`](#mat)>
- <a id="method_mat_set_uchar_at3.x"></a>`x`: `u32`
- <a id="method_mat_set_uchar_at3.y"></a>`y`: `u32`
- <a id="method_mat_set_uchar_at3.z"></a>`z`: `u32`
- <a id="method_mat_set_uchar_at3.val"></a>`val`: `u8`

#### <a id="method_mat_get_int_at3"></a>`[method]mat.get-int-at3: func`

GetIntAt3 returns the value at the specified x, y, z as a s32.

##### Params

- <a id="method_mat_get_int_at3.self"></a>`self`: borrow<[`mat`](#mat)>
- <a id="method_mat_get_int_at3.x"></a>`x`: `u32`
- <a id="method_mat_get_int_at3.y"></a>`y`: `u32`
- <a id="method_mat_get_int_at3.z"></a>`z`: `u32`

##### Return values

- <a id="method_mat_get_int_at3.0"></a> `s32`

#### <a id="method_mat_set_int_at3"></a>`[method]mat.set-int-at3: func`

SetIntAt3 sets the value at the specified x, y, z as a s32.

##### Params

- <a id="method_mat_set_int_at3.self"></a>`self`: borrow<[`mat`](#mat)>
- <a id="method_mat_set_int_at3.x"></a>`x`: `u32`
- <a id="method_mat_set_int_at3.y"></a>`y`: `u32`
- <a id="method_mat_set_int_at3.z"></a>`z`: `u32`
- <a id="method_mat_set_int_at3.val"></a>`val`: `s32`

#### <a id="method_mat_get_vecb_at"></a>`[method]mat.get-vecb-at: func`

GetVecbAt returns a vector of bytes. Its size corresponds to the number of channels of the Mat.

##### Params

- <a id="method_mat_get_vecb_at.self"></a>`self`: borrow<[`mat`](#mat)>
- <a id="method_mat_get_vecb_at.row"></a>`row`: `u32`
- <a id="method_mat_get_vecb_at.col"></a>`col`: `u32`

##### Return values

- <a id="method_mat_get_vecb_at.0"></a> list<`u8`>

#### <a id="method_mat_get_vecf_at"></a>`[method]mat.get-vecf-at: func`

GetVecfAt returns a vector of floats. Its size corresponds to the number of channels of the Mat.

##### Params

- <a id="method_mat_get_vecf_at.self"></a>`self`: borrow<[`mat`](#mat)>
- <a id="method_mat_get_vecf_at.row"></a>`row`: `u32`
- <a id="method_mat_get_vecf_at.col"></a>`col`: `u32`

##### Return values

- <a id="method_mat_get_vecf_at.0"></a> list<`f32`>

#### <a id="method_mat_get_veci_at"></a>`[method]mat.get-veci-at: func`

GetVeciAt returns a vector of s32. Its size corresponds to the number of channels of the Mat.

##### Params

- <a id="method_mat_get_veci_at.self"></a>`self`: borrow<[`mat`](#mat)>
- <a id="method_mat_get_veci_at.row"></a>`row`: `u32`
- <a id="method_mat_get_veci_at.col"></a>`col`: `u32`

##### Return values

- <a id="method_mat_get_veci_at.0"></a> list<`s32`>

#### <a id="method_mat_add_uchar"></a>`[method]mat.add-uchar: func`

AddUChar adds a u8 value to each element in the Mat. Performs addition in place.

##### Params

- <a id="method_mat_add_uchar.self"></a>`self`: borrow<[`mat`](#mat)>
- <a id="method_mat_add_uchar.val"></a>`val`: `u8`

#### <a id="method_mat_subtract_uchar"></a>`[method]mat.subtract-uchar: func`

SubtractUChar subtracts a u8 value from each element in the Mat. Performs subtraction in place.

##### Params

- <a id="method_mat_subtract_uchar.self"></a>`self`: borrow<[`mat`](#mat)>
- <a id="method_mat_subtract_uchar.val"></a>`val`: `u8`

#### <a id="method_mat_multiply_uchar"></a>`[method]mat.multiply-uchar: func`

MultiplyUChar multiplies each element in the Mat by a u8 value. Performs multiplication in place.

##### Params

- <a id="method_mat_multiply_uchar.self"></a>`self`: borrow<[`mat`](#mat)>
- <a id="method_mat_multiply_uchar.val"></a>`val`: `u8`

#### <a id="method_mat_divide_uchar"></a>`[method]mat.divide-uchar: func`

DivideUChar divides each element in the Mat by a u8 value. Performs division in place.

##### Params

- <a id="method_mat_divide_uchar.self"></a>`self`: borrow<[`mat`](#mat)>
- <a id="method_mat_divide_uchar.val"></a>`val`: `u8`

#### <a id="method_mat_add_float"></a>`[method]mat.add-float: func`

AddFloat adds a float value to each element in the Mat. Performs addition in place.

##### Params

- <a id="method_mat_add_float.self"></a>`self`: borrow<[`mat`](#mat)>
- <a id="method_mat_add_float.val"></a>`val`: `f32`

#### <a id="method_mat_subtract_float"></a>`[method]mat.subtract-float: func`

SubtractFloat subtracts a float value from each element in the Mat. Performs subtraction in place.

##### Params

- <a id="method_mat_subtract_float.self"></a>`self`: borrow<[`mat`](#mat)>
- <a id="method_mat_subtract_float.val"></a>`val`: `f32`

#### <a id="method_mat_multiply_float"></a>`[method]mat.multiply-float: func`

MultiplyFloat multiplies each element in the Mat by a float value. Performs multiplication in place.

##### Params

- <a id="method_mat_multiply_float.self"></a>`self`: borrow<[`mat`](#mat)>
- <a id="method_mat_multiply_float.val"></a>`val`: `f32`

#### <a id="method_mat_divide_float"></a>`[method]mat.divide-float: func`

DivideFloat divides each element in the Mat by a float value. Performs division in place.

##### Params

- <a id="method_mat_divide_float.self"></a>`self`: borrow<[`mat`](#mat)>
- <a id="method_mat_divide_float.val"></a>`val`: `f32`

#### <a id="method_mat_reshape"></a>`[method]mat.reshape: func`

Reshape changes the shape and/or the number of channels of a 2D matrix without copying the data.

For further details, please see:
https://docs.opencv.org/4.x/d3/d63/classcv_1_1Mat.html#a4eb96e3251417fa88b78e2abd6cfd7d8

##### Params

- <a id="method_mat_reshape.self"></a>`self`: borrow<[`mat`](#mat)>
- <a id="method_mat_reshape.channels"></a>`channels`: `u32`
- <a id="method_mat_reshape.rows"></a>`rows`: `u32`

##### Return values

- <a id="method_mat_reshape.0"></a> result<own<[`mat`](#mat)>, [`error-result`](#error_result)>

#### <a id="method_mat_row_range"></a>`[method]mat.row-range: func`

RowRange creates a matrix header for the specified row span.

For further details, please see:
https://docs.opencv.org/4.x/d3/d63/classcv_1_1Mat.html#aa6542193430356ad631a9beabc624107

##### Params

- <a id="method_mat_row_range.self"></a>`self`: borrow<[`mat`](#mat)>
- <a id="method_mat_row_range.start"></a>`start`: `u32`
- <a id="method_mat_row_range.end"></a>`end`: `u32`

##### Return values

- <a id="method_mat_row_range.0"></a> result<own<[`mat`](#mat)>, [`error-result`](#error_result)>

#### <a id="method_mat_col_range"></a>`[method]mat.col-range: func`

ColRange creates a matrix header for the specified column span.

For further details, please see:
https://docs.opencv.org/4.x/d3/d63/classcv_1_1Mat.html#aadc8f9210fe4dec50513746c246fa8d9

##### Params

- <a id="method_mat_col_range.self"></a>`self`: borrow<[`mat`](#mat)>
- <a id="method_mat_col_range.start"></a>`start`: `u32`
- <a id="method_mat_col_range.end"></a>`end`: `u32`

##### Return values

- <a id="method_mat_col_range.0"></a> result<own<[`mat`](#mat)>, [`error-result`](#error_result)>

#### <a id="method_mat_min_max_loc"></a>`[method]mat.min-max-loc: func`

MinMaxLoc finds the global minimum and maximum in an array.

For further details, please see:
https://docs.opencv.org/trunk/d2/de8/group__core__array.html#gab473bf2eb6d14ff97e89b355dac20707

##### Params

- <a id="method_mat_min_max_loc.self"></a>`self`: borrow<[`mat`](#mat)>

##### Return values

- <a id="method_mat_min_max_loc.0"></a> result<[`mix-max-loc-result`](#mix_max_loc_result), [`error-result`](#error_result)>

#### <a id="method_mat_col"></a>`[method]mat.col: func`

col creates a matrix header for the specified matrix column.
The underlying data of the new matrix is shared with the original matrix.

##### Params

- <a id="method_mat_col.self"></a>`self`: borrow<[`mat`](#mat)>
- <a id="method_mat_col.col"></a>`col`: `u32`

##### Return values

- <a id="method_mat_col.0"></a> result<own<[`mat`](#mat)>, [`error-result`](#error_result)>

#### <a id="method_mat_row"></a>`[method]mat.row: func`

row creates a matrix header for the specified matrix row.
The underlying data of the new matrix is shared with the original matrix.

##### Params

- <a id="method_mat_row.self"></a>`self`: borrow<[`mat`](#mat)>
- <a id="method_mat_row.row"></a>`row`: `u32`

##### Return values

- <a id="method_mat_row.0"></a> result<own<[`mat`](#mat)>, [`error-result`](#error_result)>

#### <a id="static_mat_merge"></a>`[static]mat.merge: func`

Merge creates one multi-channel array out of several single-channel ones.

For further details, please see:
https://docs.opencv.org/4.x/d2/de8/group__core__array.html#ga7d7b4d6c6ee504b30a20b1680029c7b4

##### Params

- <a id="static_mat_merge.mv"></a>`mv`: list<own<[`mat`](#mat)>>

##### Return values

- <a id="static_mat_merge.0"></a> result<own<[`mat`](#mat)>, [`error-result`](#error_result)>

#### <a id="static_mat_zeros"></a>`[static]mat.zeros: func`

zeros returns a zero array of the specified size and type.

##### Params

- <a id="static_mat_zeros.cols"></a>`cols`: `u32`
- <a id="static_mat_zeros.rows"></a>`rows`: `u32`
- <a id="static_mat_zeros.mattype"></a>`mattype`: [`mattype`](#mattype)

##### Return values

- <a id="static_mat_zeros.0"></a> result<own<[`mat`](#mat)>, [`error-result`](#error_result)>

#### <a id="static_mat_ones"></a>`[static]mat.ones: func`

ones returns an array of the specified size and type populated with 1.

For further details, please see:
https://docs.opencv.org/4.x/d3/d63/classcv_1_1Mat.html#a5e10227b777425407986727e2d26fcdc

##### Params

- <a id="static_mat_ones.cols"></a>`cols`: `u32`
- <a id="static_mat_ones.rows"></a>`rows`: `u32`
- <a id="static_mat_ones.mattype"></a>`mattype`: [`mattype`](#mattype)

##### Return values

- <a id="static_mat_ones.0"></a> result<own<[`mat`](#mat)>, [`error-result`](#error_result)>

## <a id="wasm_cv_cv"></a>Import interface wasm:cv/cv


----

### Types

#### <a id="error_result"></a>`type error-result`
[`error-result`](#error_result)
<p>
#### <a id="border_type"></a>`type border-type`
[`border-type`](#border_type)
<p>
#### <a id="size"></a>`type size`
[`size`](#size)
<p>
#### <a id="point"></a>`type point`
[`point`](#point)
<p>
#### <a id="adaptive_threshold_type"></a>`type adaptive-threshold-type`
[`adaptive-threshold-type`](#adaptive_threshold_type)
<p>
#### <a id="threshold_type"></a>`type threshold-type`
[`threshold-type`](#threshold_type)
<p>
#### <a id="scalar"></a>`type scalar`
[`scalar`](#scalar)
<p>
#### <a id="rect"></a>`type rect`
[`rect`](#rect)
<p>
#### <a id="rgba"></a>`type RGBA`
[`RGBA`](#rgba)
<p>
#### <a id="hershey_font_type"></a>`type hershey-font-type`
[`hershey-font-type`](#hershey_font_type)
<p>
#### <a id="interpolation_type"></a>`type interpolation-type`
[`interpolation-type`](#interpolation_type)
<p>
#### <a id="color_coversion_type"></a>`type color-coversion-type`
[`color-coversion-type`](#color_coversion_type)
<p>
#### <a id="morph_shape"></a>`type morph-shape`
[`morph-shape`](#morph_shape)
<p>
#### <a id="mat"></a>`type mat`
[`mat`](#mat)
<p>
#### <a id="mattype"></a>`type mattype`
[`mattype`](#mattype)
<p>
----

### Functions

#### <a id="arrowed_line"></a>`arrowed-line: func`

drawing functions
ArrowedLine draws a arrow segment pointing from the first point to the second one.

For further details, please see:
https://docs.opencv.org/master/d6/d6e/group__imgproc__draw.html#ga0a165a3ca093fd488ac709fdf10c05b2

##### Params

- <a id="arrowed_line.img"></a>`img`: borrow<[`mat`](#mat)>
- <a id="arrowed_line.point1"></a>`point1`: [`point`](#point)
- <a id="arrowed_line.point2"></a>`point2`: [`point`](#point)
- <a id="arrowed_line.c"></a>`c`: [`RGBA`](#rgba)
- <a id="arrowed_line.thickness"></a>`thickness`: `u8`

##### Return values

- <a id="arrowed_line.0"></a> result<_, [`error-result`](#error_result)>

#### <a id="rectangle"></a>`rectangle: func`

Rectangle draws a simple, thick, or filled up-right rectangle.

For further details, please see:
https://docs.opencv.org/4.x/d6/d6e/group__imgproc__draw.html#ga07d2f74cadcf8e305e810ce8f3d1e1b7

##### Params

- <a id="rectangle.img"></a>`img`: borrow<[`mat`](#mat)>
- <a id="rectangle.r"></a>`r`: [`rect`](#rect)
- <a id="rectangle.c"></a>`c`: [`RGBA`](#rgba)
- <a id="rectangle.thickness"></a>`thickness`: `u8`

##### Return values

- <a id="rectangle.0"></a> result<_, [`error-result`](#error_result)>

#### <a id="circle"></a>`circle: func`

Circle draws a circle.

For further details, please see:
https://docs.opencv.org/master/d6/d6e/group__imgproc__draw.html#gaf10604b069374903dbd0f0488cb43670

##### Params

- <a id="circle.img"></a>`img`: borrow<[`mat`](#mat)>
- <a id="circle.center"></a>`center`: [`point`](#point)
- <a id="circle.radius"></a>`radius`: `u32`
- <a id="circle.c"></a>`c`: [`RGBA`](#rgba)
- <a id="circle.thickness"></a>`thickness`: `u8`

##### Return values

- <a id="circle.0"></a> result<_, [`error-result`](#error_result)>

#### <a id="line"></a>`line: func`

Line draws a line segment connecting two points.

For further details, please see:
https://docs.opencv.org/master/d6/d6e/group__imgproc__draw.html#ga7078a9fae8c7e7d13d24dac2520ae4a2

##### Params

- <a id="line.img"></a>`img`: borrow<[`mat`](#mat)>
- <a id="line.point1"></a>`point1`: [`point`](#point)
- <a id="line.point2"></a>`point2`: [`point`](#point)
- <a id="line.c"></a>`c`: [`RGBA`](#rgba)
- <a id="line.thickness"></a>`thickness`: `u8`

##### Return values

- <a id="line.0"></a> result<_, [`error-result`](#error_result)>

#### <a id="put_text"></a>`put-text: func`

PutText draws a text string.
It renders the specified text string into the img Mat at the location
passed in the "org" param, using the desired font face, font scale,
color, and line thinkness.

For further details, please see:
http://docs.opencv.org/master/d6/d6e/group__imgproc__draw.html#ga5126f47f883d730f633d74f07456c576

##### Params

- <a id="put_text.img"></a>`img`: borrow<[`mat`](#mat)>
- <a id="put_text.text"></a>`text`: `string`
- <a id="put_text.org"></a>`org`: [`point`](#point)
- <a id="put_text.font_face"></a>`font-face`: [`hershey-font-type`](#hershey_font_type)
- <a id="put_text.font_scale"></a>`font-scale`: `f64`
- <a id="put_text.c"></a>`c`: [`RGBA`](#rgba)
- <a id="put_text.thickness"></a>`thickness`: `s32`

##### Return values

- <a id="put_text.0"></a> result<_, [`error-result`](#error_result)>

#### <a id="adaptive_threshold"></a>`adaptive-threshold: func`

imgproc functions
AdaptiveThreshold applies a fixed-level threshold to each array element.

For further details, please see:
https://docs.opencv.org/master/d7/d1b/group__imgproc__misc.html#ga72b913f352e4a1b1b397736707afcde3

##### Params

- <a id="adaptive_threshold.src"></a>`src`: own<[`mat`](#mat)>
- <a id="adaptive_threshold.max_value"></a>`max-value`: `f32`
- <a id="adaptive_threshold.adaptive_type"></a>`adaptive-type`: [`adaptive-threshold-type`](#adaptive_threshold_type)
- <a id="adaptive_threshold.threshold_type"></a>`threshold-type`: [`threshold-type`](#threshold_type)
- <a id="adaptive_threshold.block_size"></a>`block-size`: `u32`
- <a id="adaptive_threshold.c"></a>`c`: `f32`

##### Return values

- <a id="adaptive_threshold.0"></a> result<own<[`mat`](#mat)>, [`error-result`](#error_result)>

#### <a id="blur"></a>`blur: func`

Blur blurs an image Mat using a normalized box filter.

For further details, please see:
https://docs.opencv.org/master/d4/d86/group__imgproc__filter.html#ga8c45db9afe636703801b0b2e440fce37

##### Params

- <a id="blur.src"></a>`src`: own<[`mat`](#mat)>
- <a id="blur.k_size"></a>`k-size`: [`size`](#size)

##### Return values

- <a id="blur.0"></a> result<own<[`mat`](#mat)>, [`error-result`](#error_result)>

#### <a id="box_filter"></a>`box-filter: func`

BoxFilter blurs an image using the box filter.

For further details, please see:
https://docs.opencv.org/master/d4/d86/group__imgproc__filter.html#gad533230ebf2d42509547d514f7d3fbc3

##### Params

- <a id="box_filter.src"></a>`src`: own<[`mat`](#mat)>
- <a id="box_filter.depth"></a>`depth`: `u32`
- <a id="box_filter.k_size"></a>`k-size`: [`size`](#size)

##### Return values

- <a id="box_filter.0"></a> result<own<[`mat`](#mat)>, [`error-result`](#error_result)>

#### <a id="canny"></a>`canny: func`

Canny finds edges in an image using the Canny algorithm.
The function finds edges in the input image image and marks
them in the output map edges using the Canny algorithm.
The smallest value between threshold1 and threshold2 is used
for edge linking. The largest value is used to
find initial segments of strong edges.
See http://en.wikipedia.org/wiki/Canny_edge_detector

For further details, please see:
http://docs.opencv.org/master/dd/d1a/group__imgproc__feature.html#ga04723e007ed888ddf11d9ba04e2232de

##### Params

- <a id="canny.src"></a>`src`: own<[`mat`](#mat)>
- <a id="canny.threshold1"></a>`threshold1`: `f32`
- <a id="canny.threshold2"></a>`threshold2`: `f32`

##### Return values

- <a id="canny.0"></a> result<own<[`mat`](#mat)>, [`error-result`](#error_result)>

#### <a id="cvt_color"></a>`cvt-color: func`

CvtColor converts an image from one color space to another.

For further details, please see:
http://docs.opencv.org/master/d7/d1b/group__imgproc__misc.html#ga4e0972be5de079fed4e3a10e24ef5ef0

##### Params

- <a id="cvt_color.src"></a>`src`: own<[`mat`](#mat)>
- <a id="cvt_color.code"></a>`code`: [`color-coversion-type`](#color_coversion_type)

##### Return values

- <a id="cvt_color.0"></a> result<own<[`mat`](#mat)>, [`error-result`](#error_result)>

#### <a id="dilate"></a>`dilate: func`

Dilate dilates an image by using a specific structuring element.

For further details, please see:
https://docs.opencv.org/4.x/d4/d86/group__imgproc__filter.html#ga4ff0f3318642c4f469d0e11f242f3b6c

##### Params

- <a id="dilate.src"></a>`src`: own<[`mat`](#mat)>
- <a id="dilate.kernel"></a>`kernel`: own<[`mat`](#mat)>

##### Return values

- <a id="dilate.0"></a> result<own<[`mat`](#mat)>, [`error-result`](#error_result)>

#### <a id="erode"></a>`erode: func`

Erode erodes an image by using a specific structuring element.

For further details, please see:
https://docs.opencv.org/master/d4/d86/group__imgproc__filter.html#gaeb1e0c1033e3f6b891a25d0511362aeb

##### Params

- <a id="erode.src"></a>`src`: own<[`mat`](#mat)>
- <a id="erode.kernel"></a>`kernel`: own<[`mat`](#mat)>

##### Return values

- <a id="erode.0"></a> result<own<[`mat`](#mat)>, [`error-result`](#error_result)>

#### <a id="equalize_hist"></a>`equalize-hist: func`

EqualizeHist normalizes the brightness and increases the contrast of the image.

For further details, please see:
https://docs.opencv.org/master/d6/dc7/group__imgproc__hist.html#ga7e54091f0c937d49bf84152a16f76d6e

##### Params

- <a id="equalize_hist.src"></a>`src`: own<[`mat`](#mat)>

##### Return values

- <a id="equalize_hist.0"></a> result<own<[`mat`](#mat)>, [`error-result`](#error_result)>

#### <a id="gaussian_blur"></a>`gaussian-blur: func`

GaussianBlur blurs an image using a Gaussian filter.

For further details, please see:
https://docs.opencv.org/4.x/d4/d86/group__imgproc__filter.html#gae8bdcd9154ed5ca3cbc1766d960f45c1

##### Params

- <a id="gaussian_blur.src"></a>`src`: own<[`mat`](#mat)>
- <a id="gaussian_blur.size"></a>`size`: [`size`](#size)
- <a id="gaussian_blur.sigma_x"></a>`sigma-x`: `f32`
- <a id="gaussian_blur.sigma_y"></a>`sigma-y`: `f32`
- <a id="gaussian_blur.border"></a>`border`: [`border-type`](#border_type)

##### Return values

- <a id="gaussian_blur.0"></a> result<own<[`mat`](#mat)>, [`error-result`](#error_result)>

#### <a id="get_structuring_element"></a>`get-structuring-element: func`

GetStructuringElement returns a structuring element of the specified size
and shape for morphological operations.

For further details, please see:
https://docs.opencv.org/master/d4/d86/group__imgproc__filter.html#gac342a1bb6eabf6f55c803b09268e36dc

##### Params

- <a id="get_structuring_element.shape"></a>`shape`: [`morph-shape`](#morph_shape)
- <a id="get_structuring_element.size"></a>`size`: [`size`](#size)

##### Return values

- <a id="get_structuring_element.0"></a> result<own<[`mat`](#mat)>, [`error-result`](#error_result)>

#### <a id="hough_lines"></a>`hough-lines: func`

HoughLines implements the standard or standard multi-scale Hough transform
algorithm for line detection. For a good explanation of Hough transform, see:
http://homepages.inf.ed.ac.uk/rbf/HIPR2/hough.htm

For further details, please see:
http://docs.opencv.org/master/dd/d1a/group__imgproc__feature.html#ga46b4e588934f6c8dfd509cc6e0e4545a

##### Params

- <a id="hough_lines.src"></a>`src`: own<[`mat`](#mat)>
- <a id="hough_lines.rho"></a>`rho`: `f64`
- <a id="hough_lines.theta"></a>`theta`: `f64`
- <a id="hough_lines.threshold"></a>`threshold`: `s32`

##### Return values

- <a id="hough_lines.0"></a> result<own<[`mat`](#mat)>, [`error-result`](#error_result)>

#### <a id="hough_lines_p"></a>`hough-lines-p: func`

HoughLinesP implements the probabilistic Hough transform
algorithm for line detection. For a good explanation of Hough transform, see:
http://homepages.inf.ed.ac.uk/rbf/HIPR2/hough.htm

For further details, please see:
http://docs.opencv.org/master/dd/d1a/group__imgproc__feature.html#ga8618180a5948286384e3b7ca02f6feeb

##### Params

- <a id="hough_lines_p.src"></a>`src`: own<[`mat`](#mat)>
- <a id="hough_lines_p.rho"></a>`rho`: `f64`
- <a id="hough_lines_p.theta"></a>`theta`: `f64`
- <a id="hough_lines_p.threshold"></a>`threshold`: `s32`

##### Return values

- <a id="hough_lines_p.0"></a> result<own<[`mat`](#mat)>, [`error-result`](#error_result)>

#### <a id="median_blur"></a>`median-blur: func`

MedianBlur blurs an image using the median filter.

For further details, please see:
https://docs.opencv.org/master/d4/d86/group__imgproc__filter.html#ga564869aa33e58769b4469101aac458f9

##### Params

- <a id="median_blur.src"></a>`src`: own<[`mat`](#mat)>
- <a id="median_blur.k_size"></a>`k-size`: [`size`](#size)

##### Return values

- <a id="median_blur.0"></a> result<own<[`mat`](#mat)>, [`error-result`](#error_result)>

#### <a id="resize"></a>`resize: func`

Resize resizes an image.
It resizes the image src down to or up to the specified size, storing the
result in dst. Note that src and dst may be the same image. If you wish to
scale by factor, an empty sz may be passed and non-zero fx and fy. Likewise,
if you wish to scale to an explicit size, a non-empty sz may be passed with
zero for both fx and fy.

For further details, please see:
https://docs.opencv.org/master/da/d54/group__imgproc__transform.html#ga47a974309e9102f5f08231edc7e7529d

##### Params

- <a id="resize.src"></a>`src`: own<[`mat`](#mat)>
- <a id="resize.size"></a>`size`: [`size`](#size)
- <a id="resize.fx"></a>`fx`: `f32`
- <a id="resize.fy"></a>`fy`: `f32`
- <a id="resize.interp"></a>`interp`: [`interpolation-type`](#interpolation_type)

##### Return values

- <a id="resize.0"></a> result<own<[`mat`](#mat)>, [`error-result`](#error_result)>

#### <a id="threshold"></a>`threshold: func`

Threshold applies a fixed-level threshold to each array element.

For further details, please see:
https://docs.opencv.org/3.3.0/d7/d1b/group__imgproc__misc.html#gae8a4a146d1ca78c626a53577199e9c57

##### Params

- <a id="threshold.src"></a>`src`: own<[`mat`](#mat)>
- <a id="threshold.thresh"></a>`thresh`: `f32`
- <a id="threshold.max_value"></a>`max-value`: `f32`
- <a id="threshold.threshold_type"></a>`threshold-type`: [`threshold-type`](#threshold_type)

##### Return values

- <a id="threshold.0"></a> result<own<[`mat`](#mat)>, [`error-result`](#error_result)>

#### <a id="transpose_nd"></a>`transpose-ND: func`

Transpose for n-dimensional matrices.

For further details, please see:
https://docs.opencv.org/4.x/d2/de8/group__core__array.html#gab1b1274b4a563be34cdfa55b8919a4ec

##### Params

- <a id="transpose_nd.src"></a>`src`: own<[`mat`](#mat)>
- <a id="transpose_nd.order"></a>`order`: list<`s32`>

##### Return values

- <a id="transpose_nd.0"></a> result<own<[`mat`](#mat)>, [`error-result`](#error_result)>

#### <a id="estimate_affine2d"></a>`estimate-affine2d: func`

estimate-affine2d computes an optimal affine transformation between two 2D point sets.

For further details, please see:
https://docs.opencv.org/4.0.0/d9/d0c/group__calib3d.html#ga27865b1d26bac9ce91efaee83e94d4dd

##### Params

- <a id="estimate_affine2d.frm"></a>`frm`: own<[`mat`](#mat)>
- <a id="estimate_affine2d.to"></a>`to`: own<[`mat`](#mat)>

##### Return values

- <a id="estimate_affine2d.0"></a> result<own<[`mat`](#mat)>, [`error-result`](#error_result)>

#### <a id="warp_affine"></a>`warp-affine: func`

warp-affine applies an affine transformation to an image.

For further details, please see:
https://docs.opencv.org/4.x/da/d54/group__imgproc__transform.html#ga0203d9ee5fcd28d40dbc4a1ea4451983

##### Params

- <a id="warp_affine.src"></a>`src`: own<[`mat`](#mat)>
- <a id="warp_affine.m"></a>`m`: own<[`mat`](#mat)>
- <a id="warp_affine.size"></a>`size`: [`size`](#size)

##### Return values

- <a id="warp_affine.0"></a> result<own<[`mat`](#mat)>, [`error-result`](#error_result)>

#### <a id="get_rotation_matrix2d"></a>`get-rotation-matrix2d: func`

get-rotation-matrix2d calculates an affine matrix of 2D rotation.

For further details, please see:
https://docs.opencv.org/4.x/da/d54/group__imgproc__transform.html#gafbbc470ce83812914a70abfb604f4326

##### Params

- <a id="get_rotation_matrix2d.center"></a>`center`: [`point`](#point)
- <a id="get_rotation_matrix2d.angle"></a>`angle`: `f64`
- <a id="get_rotation_matrix2d.scale"></a>`scale`: `f64`

##### Return values

- <a id="get_rotation_matrix2d.0"></a> result<own<[`mat`](#mat)>, [`error-result`](#error_result)>

#### <a id="add"></a>`add: func`

add calculates the per-element sum of two arrays.

For further details, please see:
https://docs.opencv.org/4.x/d2/de8/group__core__array.html#ga10ac1bfb180e2cfda1701d06c24fdbd6

##### Params

- <a id="add.src1"></a>`src1`: own<[`mat`](#mat)>
- <a id="add.src2"></a>`src2`: own<[`mat`](#mat)>

##### Return values

- <a id="add.0"></a> result<own<[`mat`](#mat)>, [`error-result`](#error_result)>

#### <a id="add_weighted"></a>`add-weighted: func`

add-weighted calculates the weighted sum of two arrays.

For further details, please see:
https://docs.opencv.org/4.x/d2/de8/group__core__array.html#gafafb2513349db3bcff51f54ee5592a19

##### Params

- <a id="add_weighted.src1"></a>`src1`: own<[`mat`](#mat)>
- <a id="add_weighted.alpha"></a>`alpha`: `f64`
- <a id="add_weighted.src2"></a>`src2`: own<[`mat`](#mat)>
- <a id="add_weighted.beta"></a>`beta`: `f64`
- <a id="add_weighted.gamma"></a>`gamma`: `f64`

##### Return values

- <a id="add_weighted.0"></a> result<own<[`mat`](#mat)>, [`error-result`](#error_result)>

#### <a id="divide"></a>`divide: func`

divide performs per-element division of two arrays.

For further details, please see:
https://docs.opencv.org/4.x/d2/de8/group__core__array.html#ga6db555d30115642fedae0cda05604874

##### Params

- <a id="divide.src1"></a>`src1`: own<[`mat`](#mat)>
- <a id="divide.src2"></a>`src2`: own<[`mat`](#mat)>

##### Return values

- <a id="divide.0"></a> result<own<[`mat`](#mat)>, [`error-result`](#error_result)>

#### <a id="exp"></a>`exp: func`

exp calculates the exponent of every array element.

For further details, please see:
https://docs.opencv.org/4.x/d2/de8/group__core__array.html#ga3e10108e2162c338f1b848af619f39e5

##### Params

- <a id="exp.src"></a>`src`: own<[`mat`](#mat)>

##### Return values

- <a id="exp.0"></a> result<own<[`mat`](#mat)>, [`error-result`](#error_result)>

#### <a id="hconcat"></a>`hconcat: func`

hconcat applies horizontal concatenation to given matrices.

For further details, please see:
https://docs.opencv.org/4.x/d2/de8/group__core__array.html#gaab5ceee39e0580f879df645a872c6bf7

##### Params

- <a id="hconcat.src1"></a>`src1`: own<[`mat`](#mat)>
- <a id="hconcat.src2"></a>`src2`: own<[`mat`](#mat)>

##### Return values

- <a id="hconcat.0"></a> result<own<[`mat`](#mat)>, [`error-result`](#error_result)>

#### <a id="vconcat"></a>`vconcat: func`

vconcat applies vertical concatenation to given matrices.

For further details, please see:
https://docs.opencv.org/4.x/d2/de8/group__core__array.html#ga744f53b69f6e4f12156cdde4e76aed27

##### Params

- <a id="vconcat.src1"></a>`src1`: own<[`mat`](#mat)>
- <a id="vconcat.src2"></a>`src2`: own<[`mat`](#mat)>

##### Return values

- <a id="vconcat.0"></a> result<own<[`mat`](#mat)>, [`error-result`](#error_result)>

#### <a id="lut"></a>`lut: func`

lut performs a look-up table transform of an array.

The function LUT fills the output array with values from the look-up table.
Indices of the entries are taken from the input array.

For further details, please see:
https://docs.opencv.org/4.x/d2/de8/group__core__array.html#gab55b8d062b7f5587720ede032d34156f

##### Params

- <a id="lut.src"></a>`src`: own<[`mat`](#mat)>
- <a id="lut.wblut"></a>`wblut`: own<[`mat`](#mat)>

##### Return values

- <a id="lut.0"></a> result<own<[`mat`](#mat)>, [`error-result`](#error_result)>

#### <a id="multiply"></a>`multiply: func`

multiply performs per-element multiplication of two arrays.

For further details, please see:
https://docs.opencv.org/4.x/d2/de8/group__core__array.html#ga979d898a58d7f61c53003e162e7ad89f

##### Params

- <a id="multiply.src1"></a>`src1`: own<[`mat`](#mat)>
- <a id="multiply.src2"></a>`src2`: own<[`mat`](#mat)>

##### Return values

- <a id="multiply.0"></a> result<own<[`mat`](#mat)>, [`error-result`](#error_result)>

#### <a id="reduce"></a>`reduce: func`

reduce reduces the matrix to a vector.

For further details, please see:
https://docs.opencv.org/4.x/d2/de8/group__core__array.html#ga4b78072a303f29d9031d56e5638da78e

##### Params

- <a id="reduce.src"></a>`src`: own<[`mat`](#mat)>
- <a id="reduce.dim"></a>`dim`: `u32`
- <a id="reduce.reduce_type"></a>`reduce-type`: `u32`
- <a id="reduce.depth_type"></a>`depth-type`: `u32`

##### Return values

- <a id="reduce.0"></a> result<own<[`mat`](#mat)>, [`error-result`](#error_result)>

#### <a id="reduce_arg_max"></a>`reduce-arg-max: func`

reduce-arg-max finds indices of max elements along provided axis.

For further details, please see:
https://docs.opencv.org/4.x/d2/de8/group__core__array.html#gaa87ea34d99bcc5bf9695048355163da0

##### Params

- <a id="reduce_arg_max.src"></a>`src`: own<[`mat`](#mat)>
- <a id="reduce_arg_max.axis"></a>`axis`: `u32`
- <a id="reduce_arg_max.last_index"></a>`last-index`: `bool`

##### Return values

- <a id="reduce_arg_max.0"></a> result<own<[`mat`](#mat)>, [`error-result`](#error_result)>

#### <a id="subtract"></a>`subtract: func`

subtract calculates the per-element sum of two arrays.

For further details, please see:
https://docs.opencv.org/4.x/d2/de8/group__core__array.html#gaa0f00d98b4b5edeaeb7b8333b2de353b

##### Params

- <a id="subtract.src1"></a>`src1`: own<[`mat`](#mat)>
- <a id="subtract.src2"></a>`src2`: own<[`mat`](#mat)>

##### Return values

- <a id="subtract.0"></a> result<own<[`mat`](#mat)>, [`error-result`](#error_result)>

#### <a id="normalize"></a>`normalize: func`

normalize normalizes the norm or value range of an array.

For further details, please see:
https://docs.opencv.org/4.x/d2/de8/group__core__array.html#ga87eef7ee3970f86906d69a92cbf064bd

##### Params

- <a id="normalize.src"></a>`src`: own<[`mat`](#mat)>
- <a id="normalize.alpha"></a>`alpha`: `f32`
- <a id="normalize.beta"></a>`beta`: `f32`
- <a id="normalize.norm_type"></a>`norm-type`: `u32`

##### Return values

- <a id="normalize.0"></a> result<own<[`mat`](#mat)>, [`error-result`](#error_result)>

#### <a id="norm"></a>`norm: func`

norm calculates the norm of an array.

For further details, please see:
https://docs.opencv.org/4.x/d2/de8/group__core__array.html#ga7c331fb8dd951707e184ef4e3f21dd33

##### Params

- <a id="norm.src"></a>`src`: own<[`mat`](#mat)>
- <a id="norm.norm_type"></a>`norm-type`: `u32`

##### Return values

- <a id="norm.0"></a> result<`f32`, [`error-result`](#error_result)>

## <a id="wasm_cv_dnn"></a>Import interface wasm:cv/dnn


----

### Types

#### <a id="mat"></a>`type mat`
[`mat`](#mat)
<p>
#### <a id="error_result"></a>`type error-result`
[`error-result`](#error_result)
<p>
#### <a id="size"></a>`type size`
[`size`](#size)
<p>
#### <a id="scalar"></a>`type scalar`
[`scalar`](#scalar)
<p>
#### <a id="rect"></a>`type rect`
[`rect`](#rect)
<p>
#### <a id="blob_params"></a>`type blob-params`
[`blob-params`](#blob_params)
<p>
#### <a id="data_layout_type"></a>`type data-layout-type`
[`data-layout-type`](#data_layout_type)
<p>
#### <a id="padding_mode_type"></a>`type padding-mode-type`
[`padding-mode-type`](#padding_mode_type)
<p>
#### <a id="net_backend_type"></a>`enum net-backend-type`


##### Enum Cases

- <a id="net_backend_type.net_backend_default"></a>`net-backend-default`
- <a id="net_backend_type.net_backend_halide"></a>`net-backend-halide`
- <a id="net_backend_type.net_backend_openvino"></a>`net-backend-openvino`
- <a id="net_backend_type.net_backend_opencv"></a>`net-backend-opencv`
- <a id="net_backend_type.net_backend_vkcom"></a>`net-backend-vkcom`
- <a id="net_backend_type.net_backend_cuda"></a>`net-backend-cuda`
#### <a id="net_target_type"></a>`enum net-target-type`


##### Enum Cases

- <a id="net_target_type.net_target_cpu"></a>`net-target-cpu`
- <a id="net_target_type.net_target_fp32"></a>`net-target-fp32`
- <a id="net_target_type.net_target_fp16"></a>`net-target-fp16`
- <a id="net_target_type.net_target_vpu"></a>`net-target-vpu`
- <a id="net_target_type.net_target_vulkan"></a>`net-target-vulkan`
- <a id="net_target_type.net_target_fpga"></a>`net-target-fpga`
- <a id="net_target_type.net_target_cuda"></a>`net-target-cuda`
- <a id="net_target_type.net_target_cuda_fp16"></a>`net-target-cuda-fp16`
#### <a id="layer"></a>`resource layer`

#### <a id="net"></a>`resource net`

----

### Functions

#### <a id="constructor_layer"></a>`[constructor]layer: func`


##### Return values

- <a id="constructor_layer.0"></a> own<[`layer`](#layer)>

#### <a id="method_layer_close"></a>`[method]layer.close: func`

Close the layer

##### Params

- <a id="method_layer_close.self"></a>`self`: borrow<[`layer`](#layer)>

#### <a id="method_layer_get_name"></a>`[method]layer.get-name: func`

GetName returns the name of the layer.

##### Params

- <a id="method_layer_get_name.self"></a>`self`: borrow<[`layer`](#layer)>

##### Return values

- <a id="method_layer_get_name.0"></a> result<`string`, [`error-result`](#error_result)>

#### <a id="static_net_read"></a>`[static]net.read: func`

ReadNet read deep learning network represented in one of the supported formats.

For further details, please see:
https://docs.opencv.org/4.x/d6/d0f/group__dnn.html#ga138439da76f26266fdefec9723f6c5cd

##### Params

- <a id="static_net_read.model"></a>`model`: `string`
- <a id="static_net_read.config"></a>`config`: `string`

##### Return values

- <a id="static_net_read.0"></a> result<own<[`net`](#net)>, [`error-result`](#error_result)>

#### <a id="static_net_read_from_onnx"></a>`[static]net.read-from-ONNX: func`

ReadNetFromONNX reads a network model stored in ONNX framework's format.

For further details, please see:
https://docs.opencv.org/4.x/d6/d0f/group__dnn.html#ga9198ecaac7c32ddf0aa7a1bcbd359567

##### Params

- <a id="static_net_read_from_onnx.model"></a>`model`: `string`

##### Return values

- <a id="static_net_read_from_onnx.0"></a> result<own<[`net`](#net)>, [`error-result`](#error_result)>

#### <a id="method_net_close"></a>`[method]net.close: func`

Close the network

##### Params

- <a id="method_net_close.self"></a>`self`: borrow<[`net`](#net)>

#### <a id="method_net_empty"></a>`[method]net.empty: func`

Empty returns true if there are no layers in the network.

For further details, please see:
https://docs.opencv.org/master/db/d30/classcv_1_1dnn_1_1Net.html#a6a5778787d5b8770deab5eda6968e66c

##### Params

- <a id="method_net_empty.self"></a>`self`: borrow<[`net`](#net)>

##### Return values

- <a id="method_net_empty.0"></a> `bool`

#### <a id="method_net_set_input"></a>`[method]net.set-input: func`

SetInput sets the new input value for the network.

For further details, please see:
https://docs.opencv.org/trunk/db/d30/classcv_1_1dnn_1_1Net.html#a672a08ae76444d75d05d7bfea3e4a328

##### Params

- <a id="method_net_set_input.self"></a>`self`: borrow<[`net`](#net)>
- <a id="method_net_set_input.input"></a>`input`: own<[`mat`](#mat)>
- <a id="method_net_set_input.name"></a>`name`: `string`

##### Return values

- <a id="method_net_set_input.0"></a> result<_, [`error-result`](#error_result)>

#### <a id="method_net_forward"></a>`[method]net.forward: func`

Forward runs forward pass to compute output of layer with name outputName.

For further details, please see:
https://docs.opencv.org/trunk/db/d30/classcv_1_1dnn_1_1Net.html#a98ed94cb6ef7063d3697259566da310b

##### Params

- <a id="method_net_forward.self"></a>`self`: borrow<[`net`](#net)>
- <a id="method_net_forward.output_name"></a>`output-name`: `string`

##### Return values

- <a id="method_net_forward.0"></a> result<own<[`mat`](#mat)>, [`error-result`](#error_result)>

#### <a id="method_net_forward_layers"></a>`[method]net.forward-layers: func`

ForwardLayers forward pass to compute outputs of layers listed in outBlobNames.

For further details, please see:
https://docs.opencv.org/4.x/db/d30/classcv_1_1dnn_1_1Net.html#afe22e099b60a2883e220645391f68d4c

##### Params

- <a id="method_net_forward_layers.self"></a>`self`: borrow<[`net`](#net)>
- <a id="method_net_forward_layers.output_names"></a>`output-names`: list<`string`>

##### Return values

- <a id="method_net_forward_layers.0"></a> result<list<own<[`mat`](#mat)>>, [`error-result`](#error_result)>

#### <a id="method_net_get_unconnected_out_layers"></a>`[method]net.get-unconnected-out-layers: func`

GetUnconnectedOutLayers returns indexes of layers with unconnected outputs.

For further details, please see:
https://docs.opencv.org/4.x/db/d30/classcv_1_1dnn_1_1Net.html#ae26f0c29b3733d15d0482098ef9053e3

##### Params

- <a id="method_net_get_unconnected_out_layers.self"></a>`self`: borrow<[`net`](#net)>

##### Return values

- <a id="method_net_get_unconnected_out_layers.0"></a> result<list<`u32`>, [`error-result`](#error_result)>

#### <a id="method_net_get_layer_names"></a>`[method]net.get-layer-names: func`

GetLayerNames returns names of layers in the network.

For further details, please see:
hhttps://docs.opencv.org/4.x/db/d30/classcv_1_1dnn_1_1Net.html#a38e67098ae4ae5906bf8d8ea72199c2e

##### Params

- <a id="method_net_get_layer_names.self"></a>`self`: borrow<[`net`](#net)>

##### Return values

- <a id="method_net_get_layer_names.0"></a> result<list<`string`>, [`error-result`](#error_result)>

#### <a id="method_net_get_layer"></a>`[method]net.get-layer: func`

GetLayer returns layer with specified id.

For further details, please see:
https://docs.opencv.org/4.x/db/d30/classcv_1_1dnn_1_1Net.html#ac944d7f2d3ead5ef9b1b2fa3885f3ff1

##### Params

- <a id="method_net_get_layer.self"></a>`self`: borrow<[`net`](#net)>
- <a id="method_net_get_layer.id"></a>`id`: `u32`

##### Return values

- <a id="method_net_get_layer.0"></a> result<own<[`layer`](#layer)>, [`error-result`](#error_result)>

#### <a id="blob_from_image"></a>`blob-from-image: func`

BlobFromImage creates 4-dimensional blob from image. Optionally resizes and crops image from center,
subtract mean values, scales values by scalefactor, swap Blue and Red channels.

For further details, please see:
https://docs.opencv.org/4.x/d6/d0f/group__dnn.html#ga29f34df9376379a603acd8df581ac8d7

##### Params

- <a id="blob_from_image.image"></a>`image`: own<[`mat`](#mat)>
- <a id="blob_from_image.scale_factor"></a>`scale-factor`: `f32`
- <a id="blob_from_image.size"></a>`size`: [`size`](#size)
- <a id="blob_from_image.mean"></a>`mean`: [`scalar`](#scalar)
- <a id="blob_from_image.swap_rb"></a>`swap-rb`: `bool`
- <a id="blob_from_image.crop"></a>`crop`: `bool`

##### Return values

- <a id="blob_from_image.0"></a> result<own<[`mat`](#mat)>, [`error-result`](#error_result)>

#### <a id="blob_from_image_with_params"></a>`blob-from-image-with-params: func`

BlobFromImageWithParams creates 4-dimensional blob from image. Optionally resizes and crops image from center,
subtract mean values, scales values by scalefactor, swap Blue and Red channels.

For further details, please see:
https://docs.opencv.org/4.x/d6/d0f/group__dnn.html#ga29f34df9376379a603acd8df581ac8d7

##### Params

- <a id="blob_from_image_with_params.image"></a>`image`: own<[`mat`](#mat)>
- <a id="blob_from_image_with_params.params"></a>`params`: [`blob-params`](#blob_params)

##### Return values

- <a id="blob_from_image_with_params.0"></a> result<own<[`mat`](#mat)>, [`error-result`](#error_result)>

#### <a id="blob_rects_to_image_rects"></a>`blob-rects-to-image-rects: func`

BlobRectsToImageRects converts blob rects to image rects.

For further details, please see:
https://docs.opencv.org/4.4.0/d6/d0f/group__dnn.html#ga9d118d70a1659af729d01b10233213ee

##### Params

- <a id="blob_rects_to_image_rects.params"></a>`params`: [`blob-params`](#blob_params)
- <a id="blob_rects_to_image_rects.blob_rects"></a>`blob-rects`: list<[`rect`](#rect)>
- <a id="blob_rects_to_image_rects.image_size"></a>`image-size`: [`size`](#size)

##### Return values

- <a id="blob_rects_to_image_rects.0"></a> result<list<[`rect`](#rect)>, [`error-result`](#error_result)>

#### <a id="nms_boxes"></a>`NMS-boxes: func`

NMSBoxes performs non maximum suppression given boxes and corresponding scores.

For futher details, please see:
https://docs.opencv.org/4.4.0/d6/d0f/group__dnn.html#ga9d118d70a1659af729d01b10233213ee

##### Params

- <a id="nms_boxes.bboxes"></a>`bboxes`: list<[`rect`](#rect)>
- <a id="nms_boxes.scores"></a>`scores`: list<`f32`>
- <a id="nms_boxes.score_threshold"></a>`score-threshold`: `f32`
- <a id="nms_boxes.nms_threshold"></a>`nms-threshold`: `f32`

##### Return values

- <a id="nms_boxes.0"></a> result<list<`u32`>, [`error-result`](#error_result)>

## <a id="wasm_cv_objdetect"></a>Import interface wasm:cv/objdetect


----

### Types

#### <a id="mat"></a>`type mat`
[`mat`](#mat)
<p>
#### <a id="error_result"></a>`type error-result`
[`error-result`](#error_result)
<p>
#### <a id="size"></a>`type size`
[`size`](#size)
<p>
#### <a id="rect"></a>`type rect`
[`rect`](#rect)
<p>
#### <a id="cascade_classifier"></a>`resource cascade-classifier`

CascadeClassifier is a cascade classifier class for object detection.
#### <a id="hog_descriptor"></a>`resource HOG-descriptor`

HOGDescriptor is a Histogram Of Gradiants (HOG) for object detection.

For further details, please see:
https://docs.opencv.org/master/d5/d33/structcv_1_1HOGDescriptor.html#a723b95b709cfd3f95cf9e616de988fc8
#### <a id="face_detector_yn"></a>`resource face-detector-YN`

#### <a id="face_distance_type"></a>`enum face-distance-type`


##### Enum Cases

- <a id="face_distance_type.face_distance_type_cosine"></a>`face-distance-type-cosine`
- <a id="face_distance_type.face_distance_norm_l2"></a>`face-distance-norm-l2`
#### <a id="face_recognizer_sf"></a>`resource face-recognizer-SF`

----

### Functions

#### <a id="constructor_cascade_classifier"></a>`[constructor]cascade-classifier: func`

NewCascadeClassifier returns a new CascadeClassifier.

For further details, please see:
https://docs.opencv.org/4.x/df/d20/classcv_1_1FaceDetectorYN.html#a5f7fb43c60c95ca5ebab78483de02516

##### Params

- <a id="constructor_cascade_classifier.name"></a>`name`: `string`

##### Return values

- <a id="constructor_cascade_classifier.0"></a> own<[`cascade-classifier`](#cascade_classifier)>

#### <a id="method_cascade_classifier_close"></a>`[method]cascade-classifier.close: func`

Close the CascadeClassifier

##### Params

- <a id="method_cascade_classifier_close.self"></a>`self`: borrow<[`cascade-classifier`](#cascade_classifier)>

#### <a id="method_cascade_classifier_load"></a>`[method]cascade-classifier.load: func`

Load cascade classifier from a file.

For further details, please see:
http://docs.opencv.org/master/d1/de5/classcv_1_1CascadeClassifier.html#a1a5884c8cc749422f9eb77c2471958bc

##### Params

- <a id="method_cascade_classifier_load.self"></a>`self`: borrow<[`cascade-classifier`](#cascade_classifier)>
- <a id="method_cascade_classifier_load.file"></a>`file`: `string`

##### Return values

- <a id="method_cascade_classifier_load.0"></a> `bool`

#### <a id="method_cascade_classifier_detect_multi_scale"></a>`[method]cascade-classifier.detect-multi-scale: func`

DetectMultiScale detects objects of different sizes in the input Mat image.
The detected objects are returned as a slice of image.Rectangle structs.

For further details, please see:
http://docs.opencv.org/master/d1/de5/classcv_1_1CascadeClassifier.html#aaf8181cb63968136476ec4204ffca498

##### Params

- <a id="method_cascade_classifier_detect_multi_scale.self"></a>`self`: borrow<[`cascade-classifier`](#cascade_classifier)>
- <a id="method_cascade_classifier_detect_multi_scale.image"></a>`image`: own<[`mat`](#mat)>

##### Return values

- <a id="method_cascade_classifier_detect_multi_scale.0"></a> result<list<[`rect`](#rect)>, [`error-result`](#error_result)>

#### <a id="method_cascade_classifier_detect_multi_scale_with_params"></a>`[method]cascade-classifier.detect-multi-scale-with-params: func`

DetectMultiScaleWithParams detects objects of different sizes in the input Mat image.
The detected objects are returned as a slice of image.Rectangle structs.

For further details, please see:
http://docs.opencv.org/master/d1/de5/classcv_1_1CascadeClassifier.html#aaf8181cb63968136476ec4204ffca498

##### Params

- <a id="method_cascade_classifier_detect_multi_scale_with_params.self"></a>`self`: borrow<[`cascade-classifier`](#cascade_classifier)>
- <a id="method_cascade_classifier_detect_multi_scale_with_params.image"></a>`image`: own<[`mat`](#mat)>
- <a id="method_cascade_classifier_detect_multi_scale_with_params.scale"></a>`scale`: `f64`
- <a id="method_cascade_classifier_detect_multi_scale_with_params.min_neighbors"></a>`min-neighbors`: `u32`
- <a id="method_cascade_classifier_detect_multi_scale_with_params.flags"></a>`flags`: `u32`
- <a id="method_cascade_classifier_detect_multi_scale_with_params.min_size"></a>`min-size`: [`size`](#size)
- <a id="method_cascade_classifier_detect_multi_scale_with_params.max_size"></a>`max-size`: [`size`](#size)

##### Return values

- <a id="method_cascade_classifier_detect_multi_scale_with_params.0"></a> result<list<[`rect`](#rect)>, [`error-result`](#error_result)>

#### <a id="constructor_hog_descriptor"></a>`[constructor]HOG-descriptor: func`

NewHOGDescriptor returns a new HOGDescriptor.

##### Params

- <a id="constructor_hog_descriptor.name"></a>`name`: `string`

##### Return values

- <a id="constructor_hog_descriptor.0"></a> own<[`HOG-descriptor`](#hog_descriptor)>

#### <a id="method_hog_descriptor_close"></a>`[method]HOG-descriptor.close: func`

Close the HOGDescriptor

##### Params

- <a id="method_hog_descriptor_close.self"></a>`self`: borrow<[`HOG-descriptor`](#hog_descriptor)>

#### <a id="method_hog_descriptor_detect_multi_scale"></a>`[method]HOG-descriptor.detect-multi-scale: func`

DetectMultiScale detects objects of different sizes in the input Mat image.
The detected objects are returned as a slice of image.Rectangle structs.

For further details, please see:
https://docs.opencv.org/master/d5/d33/structcv_1_1HOGDescriptor.html#a660e5cd036fd5ddf0f5767b352acd948

##### Params

- <a id="method_hog_descriptor_detect_multi_scale.self"></a>`self`: borrow<[`HOG-descriptor`](#hog_descriptor)>
- <a id="method_hog_descriptor_detect_multi_scale.image"></a>`image`: own<[`mat`](#mat)>

##### Return values

- <a id="method_hog_descriptor_detect_multi_scale.0"></a> result<list<[`rect`](#rect)>, [`error-result`](#error_result)>

#### <a id="method_hog_descriptor_detect_multi_scale_with_params"></a>`[method]HOG-descriptor.detect-multi-scale-with-params: func`

DetectMultiScaleWithParams detects objects of different sizes in the input Mat image.
The detected objects are returned as a slice of image.Rectangle structs.

For further details, please see:
https://docs.opencv.org/master/d5/d33/structcv_1_1HOGDescriptor.html#a660e5cd036fd5ddf0f5767b352acd948

##### Params

- <a id="method_hog_descriptor_detect_multi_scale_with_params.self"></a>`self`: borrow<[`HOG-descriptor`](#hog_descriptor)>
- <a id="method_hog_descriptor_detect_multi_scale_with_params.image"></a>`image`: own<[`mat`](#mat)>
- <a id="method_hog_descriptor_detect_multi_scale_with_params.hit_threshold"></a>`hit-threshold`: `f64`
- <a id="method_hog_descriptor_detect_multi_scale_with_params.win_stride"></a>`win-stride`: [`size`](#size)
- <a id="method_hog_descriptor_detect_multi_scale_with_params.padding"></a>`padding`: [`size`](#size)
- <a id="method_hog_descriptor_detect_multi_scale_with_params.scale"></a>`scale`: `f64`
- <a id="method_hog_descriptor_detect_multi_scale_with_params.final_threshold"></a>`final-threshold`: `f64`
- <a id="method_hog_descriptor_detect_multi_scale_with_params.use_meanshift_grouping"></a>`use-meanshift-grouping`: `bool`

##### Return values

- <a id="method_hog_descriptor_detect_multi_scale_with_params.0"></a> result<list<[`rect`](#rect)>, [`error-result`](#error_result)>

#### <a id="constructor_face_detector_yn"></a>`[constructor]face-detector-YN: func`

Creates an instance of face detector YN with given parameters.

For further details, please see:
https://docs.opencv.org/4.x/df/d20/classcv_1_1FaceDetectorYN.html#a5f7fb43c60c95ca5ebab78483de02516

##### Params

- <a id="constructor_face_detector_yn.model"></a>`model`: `string`
- <a id="constructor_face_detector_yn.config"></a>`config`: `string`
- <a id="constructor_face_detector_yn.input_size"></a>`input-size`: [`size`](#size)

##### Return values

- <a id="constructor_face_detector_yn.0"></a> own<[`face-detector-YN`](#face_detector_yn)>

#### <a id="static_face_detector_yn_new_with_params"></a>`[static]face-detector-YN.new-with-params: func`

Creates an instance of face detector YN with given parameters.

For further details, please see:
https://docs.opencv.org/4.x/df/d20/classcv_1_1FaceDetectorYN.html#a5f7fb43c60c95ca5ebab78483de02516

##### Params

- <a id="static_face_detector_yn_new_with_params.model"></a>`model`: `string`
- <a id="static_face_detector_yn_new_with_params.config"></a>`config`: `string`
- <a id="static_face_detector_yn_new_with_params.input_size"></a>`input-size`: [`size`](#size)
- <a id="static_face_detector_yn_new_with_params.score_threshold"></a>`score-threshold`: `f32`
- <a id="static_face_detector_yn_new_with_params.nms_threshold"></a>`nms-threshold`: `f32`
- <a id="static_face_detector_yn_new_with_params.top_k"></a>`top-k`: `u32`
- <a id="static_face_detector_yn_new_with_params.backend_id"></a>`backend-id`: `u32`
- <a id="static_face_detector_yn_new_with_params.target_id"></a>`target-id`: `u32`

##### Return values

- <a id="static_face_detector_yn_new_with_params.0"></a> own<[`face-detector-YN`](#face_detector_yn)>

#### <a id="method_face_detector_yn_close"></a>`[method]face-detector-YN.close: func`

Close the face detector

##### Params

- <a id="method_face_detector_yn_close.self"></a>`self`: borrow<[`face-detector-YN`](#face_detector_yn)>

#### <a id="method_face_detector_yn_detect"></a>`[method]face-detector-YN.detect: func`

Detects faces in the input image.

For further details, please see:
https://docs.opencv.org/4.x/df/d20/classcv_1_1FaceDetectorYN.html#ac05bd075ca3e6edc0e328927aae6f45b

##### Params

- <a id="method_face_detector_yn_detect.self"></a>`self`: borrow<[`face-detector-YN`](#face_detector_yn)>
- <a id="method_face_detector_yn_detect.input"></a>`input`: own<[`mat`](#mat)>

##### Return values

- <a id="method_face_detector_yn_detect.0"></a> result<own<[`mat`](#mat)>, [`error-result`](#error_result)>

#### <a id="method_face_detector_yn_get_input_size"></a>`[method]face-detector-YN.get-input-size: func`


##### Params

- <a id="method_face_detector_yn_get_input_size.self"></a>`self`: borrow<[`face-detector-YN`](#face_detector_yn)>

##### Return values

- <a id="method_face_detector_yn_get_input_size.0"></a> [`size`](#size)

#### <a id="method_face_detector_yn_get_nms_threshold"></a>`[method]face-detector-YN.get-nms-threshold: func`


##### Params

- <a id="method_face_detector_yn_get_nms_threshold.self"></a>`self`: borrow<[`face-detector-YN`](#face_detector_yn)>

##### Return values

- <a id="method_face_detector_yn_get_nms_threshold.0"></a> `f32`

#### <a id="method_face_detector_yn_get_score_threshold"></a>`[method]face-detector-YN.get-score-threshold: func`


##### Params

- <a id="method_face_detector_yn_get_score_threshold.self"></a>`self`: borrow<[`face-detector-YN`](#face_detector_yn)>

##### Return values

- <a id="method_face_detector_yn_get_score_threshold.0"></a> `f32`

#### <a id="method_face_detector_yn_get_topk"></a>`[method]face-detector-YN.get-topk: func`


##### Params

- <a id="method_face_detector_yn_get_topk.self"></a>`self`: borrow<[`face-detector-YN`](#face_detector_yn)>

##### Return values

- <a id="method_face_detector_yn_get_topk.0"></a> `u32`

#### <a id="method_face_detector_yn_set_input_size"></a>`[method]face-detector-YN.set-input-size: func`


##### Params

- <a id="method_face_detector_yn_set_input_size.self"></a>`self`: borrow<[`face-detector-YN`](#face_detector_yn)>
- <a id="method_face_detector_yn_set_input_size.size"></a>`size`: [`size`](#size)

#### <a id="method_face_detector_yn_set_nms_threshold"></a>`[method]face-detector-YN.set-nms-threshold: func`


##### Params

- <a id="method_face_detector_yn_set_nms_threshold.self"></a>`self`: borrow<[`face-detector-YN`](#face_detector_yn)>
- <a id="method_face_detector_yn_set_nms_threshold.threshold"></a>`threshold`: `f32`

#### <a id="method_face_detector_yn_set_score_threshold"></a>`[method]face-detector-YN.set-score-threshold: func`


##### Params

- <a id="method_face_detector_yn_set_score_threshold.self"></a>`self`: borrow<[`face-detector-YN`](#face_detector_yn)>
- <a id="method_face_detector_yn_set_score_threshold.threshold"></a>`threshold`: `f32`

#### <a id="method_face_detector_yn_set_topk"></a>`[method]face-detector-YN.set-topk: func`


##### Params

- <a id="method_face_detector_yn_set_topk.self"></a>`self`: borrow<[`face-detector-YN`](#face_detector_yn)>
- <a id="method_face_detector_yn_set_topk.topk"></a>`topk`: `u32`

#### <a id="constructor_face_recognizer_sf"></a>`[constructor]face-recognizer-SF: func`

Creates an instance of FaceRecognizerSF with given parameters.

For further details, please see:
https://docs.opencv.org/4.x/da/d09/classcv_1_1FaceRecognizerSF.html#a04df90b0cd7d26d350acd92621a35743

##### Params

- <a id="constructor_face_recognizer_sf.model"></a>`model`: `string`
- <a id="constructor_face_recognizer_sf.config"></a>`config`: `string`

##### Return values

- <a id="constructor_face_recognizer_sf.0"></a> own<[`face-recognizer-SF`](#face_recognizer_sf)>

#### <a id="static_face_recognizer_sf_new_with_params"></a>`[static]face-recognizer-SF.new-with-params: func`

Creates an instance of FaceRecognizerSF with given parameters.

For further details, please see:
https://docs.opencv.org/4.x/da/d09/classcv_1_1FaceRecognizerSF.html#a04df90b0cd7d26d350acd92621a35743

##### Params

- <a id="static_face_recognizer_sf_new_with_params.model"></a>`model`: `string`
- <a id="static_face_recognizer_sf_new_with_params.config"></a>`config`: `string`
- <a id="static_face_recognizer_sf_new_with_params.backend_id"></a>`backend-id`: `u32`
- <a id="static_face_recognizer_sf_new_with_params.target_id"></a>`target-id`: `u32`

##### Return values

- <a id="static_face_recognizer_sf_new_with_params.0"></a> own<[`face-recognizer-SF`](#face_recognizer_sf)>

#### <a id="method_face_recognizer_sf_close"></a>`[method]face-recognizer-SF.close: func`

Close the face FaceRecognizerSF

##### Params

- <a id="method_face_recognizer_sf_close.self"></a>`self`: borrow<[`face-recognizer-SF`](#face_recognizer_sf)>

#### <a id="method_face_recognizer_sf_align_crop"></a>`[method]face-recognizer-SF.align-crop: func`

Aligns detected face with the source input image and crops it.

For further details, please see:
https://docs.opencv.org/4.x/da/d09/classcv_1_1FaceRecognizerSF.html#a84492908abecbc9362b4ddc8d46b8345

##### Params

- <a id="method_face_recognizer_sf_align_crop.self"></a>`self`: borrow<[`face-recognizer-SF`](#face_recognizer_sf)>
- <a id="method_face_recognizer_sf_align_crop.src"></a>`src`: own<[`mat`](#mat)>
- <a id="method_face_recognizer_sf_align_crop.face_box"></a>`face-box`: own<[`mat`](#mat)>

##### Return values

- <a id="method_face_recognizer_sf_align_crop.0"></a> result<own<[`mat`](#mat)>, [`error-result`](#error_result)>

#### <a id="method_face_recognizer_sf_feature"></a>`[method]face-recognizer-SF.feature: func`

Feature extracts face feature from aligned image.

For further details, please see:
https://docs.opencv.org/4.x/da/d09/classcv_1_1FaceRecognizerSF.html#ab1b4a3c12213e89091a490c573dc5aba

##### Params

- <a id="method_face_recognizer_sf_feature.self"></a>`self`: borrow<[`face-recognizer-SF`](#face_recognizer_sf)>
- <a id="method_face_recognizer_sf_feature.aligned"></a>`aligned`: own<[`mat`](#mat)>

##### Return values

- <a id="method_face_recognizer_sf_feature.0"></a> result<own<[`mat`](#mat)>, [`error-result`](#error_result)>

#### <a id="method_face_recognizer_sf_match"></a>`[method]face-recognizer-SF.match: func`

Match calculates the distance between two face features.

For further details, please see:
https://docs.opencv.org/4.x/da/d09/classcv_1_1FaceRecognizerSF.html#a2f0362ca1e64320a1f3ba7e1386d0219

##### Params

- <a id="method_face_recognizer_sf_match.self"></a>`self`: borrow<[`face-recognizer-SF`](#face_recognizer_sf)>
- <a id="method_face_recognizer_sf_match.face1"></a>`face1`: own<[`mat`](#mat)>
- <a id="method_face_recognizer_sf_match.face2"></a>`face2`: own<[`mat`](#mat)>

##### Return values

- <a id="method_face_recognizer_sf_match.0"></a> result<`f32`, [`error-result`](#error_result)>

#### <a id="method_face_recognizer_sf_match_with_params"></a>`[method]face-recognizer-SF.match-with-params: func`

Match calculates the distance between two face features.

For further details, please see:
https://docs.opencv.org/4.x/da/d09/classcv_1_1FaceRecognizerSF.html#a2f0362ca1e64320a1f3ba7e1386d0219

##### Params

- <a id="method_face_recognizer_sf_match_with_params.self"></a>`self`: borrow<[`face-recognizer-SF`](#face_recognizer_sf)>
- <a id="method_face_recognizer_sf_match_with_params.face1"></a>`face1`: own<[`mat`](#mat)>
- <a id="method_face_recognizer_sf_match_with_params.face2"></a>`face2`: own<[`mat`](#mat)>
- <a id="method_face_recognizer_sf_match_with_params.distance"></a>`distance`: [`face-distance-type`](#face_distance_type)

##### Return values

- <a id="method_face_recognizer_sf_match_with_params.0"></a> result<`f32`, [`error-result`](#error_result)>

## <a id="wasm_cv_features2d"></a>Import interface wasm:cv/features2d


----

### Types

#### <a id="mat"></a>`type mat`
[`mat`](#mat)
<p>
#### <a id="error_result"></a>`type error-result`
[`error-result`](#error_result)
<p>
#### <a id="key_point"></a>`type key-point`
[`key-point`](#key_point)
<p>
#### <a id="d_match"></a>`type d-match`
[`d-match`](#d_match)
<p>
#### <a id="detector_result"></a>`record detector-result`

detector-result returns the keypoints and descripts for a detector.

##### Record Fields

- <a id="detector_result.kps"></a>`kps`: list<[`key-point`](#key_point)>
- <a id="detector_result.desc"></a>`desc`: own<[`mat`](#mat)>
#### <a id="akaze_detector"></a>`resource AKAZE-detector`

AKAZE-detector is a wrapper around the cv::AKAZE algorithm.
#### <a id="brisk_detector"></a>`resource BRISK-detector`

BRISK-detector is a wrapper around the cv::BRISK algorithm.
#### <a id="kaze_detector"></a>`resource KAZE-detector`

KAZE-detector is a wrapper around the cv::KAZE algorithm.
#### <a id="orb_score_type"></a>`enum ORB-score-type`


##### Enum Cases

- <a id="orb_score_type.orb_harris"></a>`ORB-HARRIS`
- <a id="orb_score_type.orb_fast"></a>`ORB-FAST`
#### <a id="orb_detector"></a>`resource ORB-detector`

ORB-detector is a wrapper around the cv::ORB algorithm.
#### <a id="sift_detector"></a>`resource SIFT-detector`

SIFT-detector is a wrapper around the cv::SIFT algorithm.
#### <a id="norm_type"></a>`enum norm-type`


##### Enum Cases

- <a id="norm_type.norm_none"></a>`NORM-NONE`
- <a id="norm_type.none_inf"></a>`NONE-INF`
- <a id="norm_type.norm_l1"></a>`NORM-L1`
- <a id="norm_type.norm_none2"></a>`NORM-NONE2`
- <a id="norm_type.norm_l2"></a>`NORM-L2`
- <a id="norm_type.norm_l2sqr"></a>`NORM-L2SQR`
- <a id="norm_type.norm_hamming"></a>`NORM-HAMMING`
- <a id="norm_type.norm_hamming2"></a>`NORM-HAMMING2`
- <a id="norm_type.norm_relative"></a>`NORM-RELATIVE`
#### <a id="bf_matcher"></a>`resource BF-matcher`

BF-matcher is a wrapper around the cv::BFMatcher algorithm.
#### <a id="flann_based_matcher"></a>`resource flann-based-matcher`

Flann-based-matcher is a wrapper around the cv::BFMatcher algorithm.
----

### Functions

#### <a id="constructor_akaze_detector"></a>`[constructor]AKAZE-detector: func`

Returns a new akaze-detector.

For further details, please see:
https://docs.opencv.org/4.x/d8/d30/classcv_1_1AKAZE.html

##### Params

- <a id="constructor_akaze_detector.name"></a>`name`: `string`

##### Return values

- <a id="constructor_akaze_detector.0"></a> own<[`AKAZE-detector`](#akaze_detector)>

#### <a id="method_akaze_detector_close"></a>`[method]AKAZE-detector.close: func`

Close the akaze-detector

##### Params

- <a id="method_akaze_detector_close.self"></a>`self`: borrow<[`AKAZE-detector`](#akaze_detector)>

#### <a id="method_akaze_detector_detect"></a>`[method]AKAZE-detector.detect: func`

Detect keypoints in an image using AKAZE.

For further details, please see:
https://docs.opencv.org/4.x/d0/d13/classcv_1_1Feature2D.html#aa4e9a7082ec61ebc108806704fbd7887

##### Params

- <a id="method_akaze_detector_detect.self"></a>`self`: borrow<[`AKAZE-detector`](#akaze_detector)>
- <a id="method_akaze_detector_detect.src"></a>`src`: own<[`mat`](#mat)>

##### Return values

- <a id="method_akaze_detector_detect.0"></a> result<list<[`key-point`](#key_point)>, [`error-result`](#error_result)>

#### <a id="method_akaze_detector_compute"></a>`[method]AKAZE-detector.compute: func`

Compute keypoints in an image using AKAZE.

For further details, please see:
https://docs.opencv.org/4.x/d0/d13/classcv_1_1Feature2D.html#ab3cce8d56f4fc5e1d530b5931e1e8dc0

##### Params

- <a id="method_akaze_detector_compute.self"></a>`self`: borrow<[`AKAZE-detector`](#akaze_detector)>
- <a id="method_akaze_detector_compute.src"></a>`src`: own<[`mat`](#mat)>
- <a id="method_akaze_detector_compute.mask"></a>`mask`: own<[`mat`](#mat)>
- <a id="method_akaze_detector_compute.kps"></a>`kps`: list<[`key-point`](#key_point)>

##### Return values

- <a id="method_akaze_detector_compute.0"></a> result<[`detector-result`](#detector_result), [`error-result`](#error_result)>

#### <a id="method_akaze_detector_detect_and_compute"></a>`[method]AKAZE-detector.detect-and-compute: func`

DetectAndCompute keypoints and compute in an image using AKAZE.

For further details, please see:
https://docs.opencv.org/4.x/d0/d13/classcv_1_1Feature2D.html#a8be0d1c20b08eb867184b8d74c15a677

##### Params

- <a id="method_akaze_detector_detect_and_compute.self"></a>`self`: borrow<[`AKAZE-detector`](#akaze_detector)>
- <a id="method_akaze_detector_detect_and_compute.src"></a>`src`: own<[`mat`](#mat)>
- <a id="method_akaze_detector_detect_and_compute.mask"></a>`mask`: own<[`mat`](#mat)>

##### Return values

- <a id="method_akaze_detector_detect_and_compute.0"></a> result<[`detector-result`](#detector_result), [`error-result`](#error_result)>

#### <a id="constructor_brisk_detector"></a>`[constructor]BRISK-detector: func`

Returns a new BRISK-detector.

For further details, please see:
https://docs.opencv.org/4.x/de/dbf/classcv_1_1BRISK.html

##### Params

- <a id="constructor_brisk_detector.name"></a>`name`: `string`

##### Return values

- <a id="constructor_brisk_detector.0"></a> own<[`BRISK-detector`](#brisk_detector)>

#### <a id="method_brisk_detector_close"></a>`[method]BRISK-detector.close: func`

Close the BRISK-detector

##### Params

- <a id="method_brisk_detector_close.self"></a>`self`: borrow<[`BRISK-detector`](#brisk_detector)>

#### <a id="method_brisk_detector_detect"></a>`[method]BRISK-detector.detect: func`

Detect keypoints in an image using BRISK.

For further details, please see:
https://docs.opencv.org/4.x/d0/d13/classcv_1_1Feature2D.html#aa4e9a7082ec61ebc108806704fbd7887

##### Params

- <a id="method_brisk_detector_detect.self"></a>`self`: borrow<[`BRISK-detector`](#brisk_detector)>
- <a id="method_brisk_detector_detect.src"></a>`src`: own<[`mat`](#mat)>

##### Return values

- <a id="method_brisk_detector_detect.0"></a> result<list<[`key-point`](#key_point)>, [`error-result`](#error_result)>

#### <a id="method_brisk_detector_compute"></a>`[method]BRISK-detector.compute: func`

Compute keypoints in an image using BRISK.

For further details, please see:
https://docs.opencv.org/4.x/d0/d13/classcv_1_1Feature2D.html#ab3cce8d56f4fc5e1d530b5931e1e8dc0

##### Params

- <a id="method_brisk_detector_compute.self"></a>`self`: borrow<[`BRISK-detector`](#brisk_detector)>
- <a id="method_brisk_detector_compute.src"></a>`src`: own<[`mat`](#mat)>
- <a id="method_brisk_detector_compute.mask"></a>`mask`: own<[`mat`](#mat)>
- <a id="method_brisk_detector_compute.kps"></a>`kps`: list<[`key-point`](#key_point)>

##### Return values

- <a id="method_brisk_detector_compute.0"></a> result<[`detector-result`](#detector_result), [`error-result`](#error_result)>

#### <a id="method_brisk_detector_detect_and_compute"></a>`[method]BRISK-detector.detect-and-compute: func`

DetectAndCompute keypoints and compute in an image using BRISK.

For further details, please see:
https://docs.opencv.org/4.x/d0/d13/classcv_1_1Feature2D.html#a8be0d1c20b08eb867184b8d74c15a677

##### Params

- <a id="method_brisk_detector_detect_and_compute.self"></a>`self`: borrow<[`BRISK-detector`](#brisk_detector)>
- <a id="method_brisk_detector_detect_and_compute.src"></a>`src`: own<[`mat`](#mat)>
- <a id="method_brisk_detector_detect_and_compute.mask"></a>`mask`: own<[`mat`](#mat)>

##### Return values

- <a id="method_brisk_detector_detect_and_compute.0"></a> result<[`detector-result`](#detector_result), [`error-result`](#error_result)>

#### <a id="constructor_kaze_detector"></a>`[constructor]KAZE-detector: func`

Returns a new KAZE-detector.

For further details, please see:
https://docs.opencv.org/4.x/d3/d61/classcv_1_1KAZE.html

##### Params

- <a id="constructor_kaze_detector.name"></a>`name`: `string`

##### Return values

- <a id="constructor_kaze_detector.0"></a> own<[`KAZE-detector`](#kaze_detector)>

#### <a id="method_kaze_detector_close"></a>`[method]KAZE-detector.close: func`

Close the KAZE-detector

##### Params

- <a id="method_kaze_detector_close.self"></a>`self`: borrow<[`KAZE-detector`](#kaze_detector)>

#### <a id="method_kaze_detector_detect"></a>`[method]KAZE-detector.detect: func`

Detect keypoints in an image using KAZE.

For further details, please see:
https://docs.opencv.org/4.x/d0/d13/classcv_1_1Feature2D.html#aa4e9a7082ec61ebc108806704fbd7887

##### Params

- <a id="method_kaze_detector_detect.self"></a>`self`: borrow<[`KAZE-detector`](#kaze_detector)>
- <a id="method_kaze_detector_detect.src"></a>`src`: own<[`mat`](#mat)>

##### Return values

- <a id="method_kaze_detector_detect.0"></a> result<list<[`key-point`](#key_point)>, [`error-result`](#error_result)>

#### <a id="method_kaze_detector_compute"></a>`[method]KAZE-detector.compute: func`

Compute keypoints in an image using KAZE.

For further details, please see:
https://docs.opencv.org/4.x/d0/d13/classcv_1_1Feature2D.html#ab3cce8d56f4fc5e1d530b5931e1e8dc0

##### Params

- <a id="method_kaze_detector_compute.self"></a>`self`: borrow<[`KAZE-detector`](#kaze_detector)>
- <a id="method_kaze_detector_compute.src"></a>`src`: own<[`mat`](#mat)>
- <a id="method_kaze_detector_compute.mask"></a>`mask`: own<[`mat`](#mat)>
- <a id="method_kaze_detector_compute.kps"></a>`kps`: list<[`key-point`](#key_point)>

##### Return values

- <a id="method_kaze_detector_compute.0"></a> result<[`detector-result`](#detector_result), [`error-result`](#error_result)>

#### <a id="method_kaze_detector_detect_and_compute"></a>`[method]KAZE-detector.detect-and-compute: func`

DetectAndCompute keypoints and compute in an image using KAZE.

For further details, please see:
https://docs.opencv.org/4.x/d0/d13/classcv_1_1Feature2D.html#a8be0d1c20b08eb867184b8d74c15a677

##### Params

- <a id="method_kaze_detector_detect_and_compute.self"></a>`self`: borrow<[`KAZE-detector`](#kaze_detector)>
- <a id="method_kaze_detector_detect_and_compute.src"></a>`src`: own<[`mat`](#mat)>
- <a id="method_kaze_detector_detect_and_compute.mask"></a>`mask`: own<[`mat`](#mat)>

##### Return values

- <a id="method_kaze_detector_detect_and_compute.0"></a> result<[`detector-result`](#detector_result), [`error-result`](#error_result)>

#### <a id="constructor_orb_detector"></a>`[constructor]ORB-detector: func`

Returns a new ORB-detector.

For further details, please see:
https://docs.opencv.org/4.x/db/d95/classcv_1_1ORB.html

##### Params

- <a id="constructor_orb_detector.name"></a>`name`: `string`

##### Return values

- <a id="constructor_orb_detector.0"></a> own<[`ORB-detector`](#orb_detector)>

#### <a id="static_orb_detector_new_with_params"></a>`[static]ORB-detector.new-with-params: func`

Returns a new ORB-detector.

For further details, please see:
https://docs.opencv.org/4.x/db/d95/classcv_1_1ORB.html

##### Params

- <a id="static_orb_detector_new_with_params.features"></a>`features`: `u32`
- <a id="static_orb_detector_new_with_params.scale"></a>`scale`: `f32`
- <a id="static_orb_detector_new_with_params.levels"></a>`levels`: `u32`
- <a id="static_orb_detector_new_with_params.edge_threshold"></a>`edge-threshold`: `u32`
- <a id="static_orb_detector_new_with_params.first"></a>`first`: `u32`
- <a id="static_orb_detector_new_with_params.wtak"></a>`WTAK`: `u32`
- <a id="static_orb_detector_new_with_params.score_type"></a>`score-type`: [`ORB-score-type`](#orb_score_type)
- <a id="static_orb_detector_new_with_params.patch_size"></a>`patch-size`: `u32`
- <a id="static_orb_detector_new_with_params.fast_threshold"></a>`fast-threshold`: `u32`

##### Return values

- <a id="static_orb_detector_new_with_params.0"></a> own<[`ORB-detector`](#orb_detector)>

#### <a id="method_orb_detector_close"></a>`[method]ORB-detector.close: func`

Close the ORB-detector

##### Params

- <a id="method_orb_detector_close.self"></a>`self`: borrow<[`ORB-detector`](#orb_detector)>

#### <a id="method_orb_detector_detect"></a>`[method]ORB-detector.detect: func`

Detect keypoints in an image using ORB.

For further details, please see:
https://docs.opencv.org/4.x/d0/d13/classcv_1_1Feature2D.html#aa4e9a7082ec61ebc108806704fbd7887

##### Params

- <a id="method_orb_detector_detect.self"></a>`self`: borrow<[`ORB-detector`](#orb_detector)>
- <a id="method_orb_detector_detect.src"></a>`src`: own<[`mat`](#mat)>

##### Return values

- <a id="method_orb_detector_detect.0"></a> result<list<[`key-point`](#key_point)>, [`error-result`](#error_result)>

#### <a id="method_orb_detector_compute"></a>`[method]ORB-detector.compute: func`

Compute keypoints in an image using ORB.

For further details, please see:
https://docs.opencv.org/4.x/d0/d13/classcv_1_1Feature2D.html#ab3cce8d56f4fc5e1d530b5931e1e8dc0

##### Params

- <a id="method_orb_detector_compute.self"></a>`self`: borrow<[`ORB-detector`](#orb_detector)>
- <a id="method_orb_detector_compute.src"></a>`src`: own<[`mat`](#mat)>
- <a id="method_orb_detector_compute.mask"></a>`mask`: own<[`mat`](#mat)>
- <a id="method_orb_detector_compute.kps"></a>`kps`: list<[`key-point`](#key_point)>

##### Return values

- <a id="method_orb_detector_compute.0"></a> result<[`detector-result`](#detector_result), [`error-result`](#error_result)>

#### <a id="method_orb_detector_detect_and_compute"></a>`[method]ORB-detector.detect-and-compute: func`

DetectAndCompute keypoints and compute in an image using ORB.

For further details, please see:
https://docs.opencv.org/4.x/d0/d13/classcv_1_1Feature2D.html#a8be0d1c20b08eb867184b8d74c15a677

##### Params

- <a id="method_orb_detector_detect_and_compute.self"></a>`self`: borrow<[`ORB-detector`](#orb_detector)>
- <a id="method_orb_detector_detect_and_compute.src"></a>`src`: own<[`mat`](#mat)>
- <a id="method_orb_detector_detect_and_compute.mask"></a>`mask`: own<[`mat`](#mat)>

##### Return values

- <a id="method_orb_detector_detect_and_compute.0"></a> result<[`detector-result`](#detector_result), [`error-result`](#error_result)>

#### <a id="constructor_sift_detector"></a>`[constructor]SIFT-detector: func`

Returns a new SIFT-detector.

For further details, please see:
https://docs.opencv.org/4.x/d7/d60/classcv_1_1SIFT.html

##### Params

- <a id="constructor_sift_detector.name"></a>`name`: `string`

##### Return values

- <a id="constructor_sift_detector.0"></a> own<[`SIFT-detector`](#sift_detector)>

#### <a id="method_sift_detector_close"></a>`[method]SIFT-detector.close: func`

Close the SIFT-detector

##### Params

- <a id="method_sift_detector_close.self"></a>`self`: borrow<[`SIFT-detector`](#sift_detector)>

#### <a id="method_sift_detector_detect"></a>`[method]SIFT-detector.detect: func`

Detect keypoints in an image using SIFT.

For further details, please see:
https://docs.opencv.org/4.x/d0/d13/classcv_1_1Feature2D.html#aa4e9a7082ec61ebc108806704fbd7887

##### Params

- <a id="method_sift_detector_detect.self"></a>`self`: borrow<[`SIFT-detector`](#sift_detector)>
- <a id="method_sift_detector_detect.src"></a>`src`: own<[`mat`](#mat)>

##### Return values

- <a id="method_sift_detector_detect.0"></a> result<list<[`key-point`](#key_point)>, [`error-result`](#error_result)>

#### <a id="method_sift_detector_compute"></a>`[method]SIFT-detector.compute: func`

Compute keypoints in an image using SIFT.

For further details, please see:
https://docs.opencv.org/4.x/d0/d13/classcv_1_1Feature2D.html#ab3cce8d56f4fc5e1d530b5931e1e8dc0

##### Params

- <a id="method_sift_detector_compute.self"></a>`self`: borrow<[`SIFT-detector`](#sift_detector)>
- <a id="method_sift_detector_compute.src"></a>`src`: own<[`mat`](#mat)>
- <a id="method_sift_detector_compute.mask"></a>`mask`: own<[`mat`](#mat)>
- <a id="method_sift_detector_compute.kps"></a>`kps`: list<[`key-point`](#key_point)>

##### Return values

- <a id="method_sift_detector_compute.0"></a> result<[`detector-result`](#detector_result), [`error-result`](#error_result)>

#### <a id="method_sift_detector_detect_and_compute"></a>`[method]SIFT-detector.detect-and-compute: func`

DetectAndCompute keypoints and compute in an image using SIFT.

For further details, please see:
https://docs.opencv.org/4.x/d0/d13/classcv_1_1Feature2D.html#a8be0d1c20b08eb867184b8d74c15a677

##### Params

- <a id="method_sift_detector_detect_and_compute.self"></a>`self`: borrow<[`SIFT-detector`](#sift_detector)>
- <a id="method_sift_detector_detect_and_compute.src"></a>`src`: own<[`mat`](#mat)>
- <a id="method_sift_detector_detect_and_compute.mask"></a>`mask`: own<[`mat`](#mat)>

##### Return values

- <a id="method_sift_detector_detect_and_compute.0"></a> result<[`detector-result`](#detector_result), [`error-result`](#error_result)>

#### <a id="constructor_bf_matcher"></a>`[constructor]BF-matcher: func`

Returns a new BF-matcher.

For further details, please see:
https://docs.opencv.org/4.x/d3/da1/classcv_1_1BFMatcher.html#abe0bb11749b30d97f60d6ade665617bd

##### Params

- <a id="constructor_bf_matcher.name"></a>`name`: `string`

##### Return values

- <a id="constructor_bf_matcher.0"></a> own<[`BF-matcher`](#bf_matcher)>

#### <a id="static_bf_matcher_new_with_params"></a>`[static]BF-matcher.new-with-params: func`

Returns a new BF-matcher.

For further details, please see:
https://docs.opencv.org/4.x/d3/da1/classcv_1_1BFMatcher.html#abe0bb11749b30d97f60d6ade665617bd

##### Params

- <a id="static_bf_matcher_new_with_params.norm"></a>`norm`: [`norm-type`](#norm_type)
- <a id="static_bf_matcher_new_with_params.cross_check"></a>`cross-check`: `bool`

##### Return values

- <a id="static_bf_matcher_new_with_params.0"></a> own<[`BF-matcher`](#bf_matcher)>

#### <a id="method_bf_matcher_close"></a>`[method]BF-matcher.close: func`

Close the BF-matcher

##### Params

- <a id="method_bf_matcher_close.self"></a>`self`: borrow<[`BF-matcher`](#bf_matcher)>

#### <a id="method_bf_matcher_match"></a>`[method]BF-matcher.match: func`

Match Finds the best match for each descriptor from a query set.

For further details, please see:
https://docs.opencv.org/4.x/db/d39/classcv_1_1DescriptorMatcher.html#a0f046f47b68ec7074391e1e85c750cba

##### Params

- <a id="method_bf_matcher_match.self"></a>`self`: borrow<[`BF-matcher`](#bf_matcher)>
- <a id="method_bf_matcher_match.query"></a>`query`: own<[`mat`](#mat)>
- <a id="method_bf_matcher_match.train"></a>`train`: own<[`mat`](#mat)>

##### Return values

- <a id="method_bf_matcher_match.0"></a> result<list<[`d-match`](#d_match)>, [`error-result`](#error_result)>

#### <a id="method_bf_matcher_knn_match"></a>`[method]BF-matcher.KNN-match: func`

KNNMatch finds the k best matches for each descriptor from a query set.

For further details, please see:
https://docs.opencv.org/4.x/db/d39/classcv_1_1DescriptorMatcher.html#aa880f9353cdf185ccf3013e08210483a

##### Params

- <a id="method_bf_matcher_knn_match.self"></a>`self`: borrow<[`BF-matcher`](#bf_matcher)>
- <a id="method_bf_matcher_knn_match.query"></a>`query`: own<[`mat`](#mat)>
- <a id="method_bf_matcher_knn_match.train"></a>`train`: own<[`mat`](#mat)>
- <a id="method_bf_matcher_knn_match.k"></a>`k`: `u32`

##### Return values

- <a id="method_bf_matcher_knn_match.0"></a> result<list<list<[`d-match`](#d_match)>>, [`error-result`](#error_result)>

#### <a id="constructor_flann_based_matcher"></a>`[constructor]flann-based-matcher: func`

Returns a new flann-based-matcher.

For further details, please see:
https://docs.opencv.org/4.x/dc/de2/classcv_1_1FlannBasedMatcher.html#ab9114a6471e364ad221f89068ca21382

##### Params

- <a id="constructor_flann_based_matcher.name"></a>`name`: `string`

##### Return values

- <a id="constructor_flann_based_matcher.0"></a> own<[`flann-based-matcher`](#flann_based_matcher)>

#### <a id="method_flann_based_matcher_close"></a>`[method]flann-based-matcher.close: func`

Close the flann-based-matcher

##### Params

- <a id="method_flann_based_matcher_close.self"></a>`self`: borrow<[`flann-based-matcher`](#flann_based_matcher)>

#### <a id="method_flann_based_matcher_knn_match"></a>`[method]flann-based-matcher.KNN-match: func`

KNNMatch finds the k best matches for each descriptor from a query set.

For further details, please see:
https://docs.opencv.org/4.x/db/d39/classcv_1_1DescriptorMatcher.html#aa880f9353cdf185ccf3013e08210483a

##### Params

- <a id="method_flann_based_matcher_knn_match.self"></a>`self`: borrow<[`flann-based-matcher`](#flann_based_matcher)>
- <a id="method_flann_based_matcher_knn_match.query"></a>`query`: own<[`mat`](#mat)>
- <a id="method_flann_based_matcher_knn_match.train"></a>`train`: own<[`mat`](#mat)>
- <a id="method_flann_based_matcher_knn_match.k"></a>`k`: `u32`

##### Return values

- <a id="method_flann_based_matcher_knn_match.0"></a> result<list<list<[`d-match`](#d_match)>>, [`error-result`](#error_result)>

## <a id="wasm_cv_request"></a>Export interface wasm:cv/request

----

### Types

#### <a id="mat"></a>`type mat`
[`mat`](#mat)
<p>
----

### Functions

#### <a id="process"></a>`process: func`


##### Params

- <a id="process.image"></a>`image`: own<[`mat`](#mat)>

##### Return values

- <a id="process.0"></a> own<[`mat`](#mat)>

