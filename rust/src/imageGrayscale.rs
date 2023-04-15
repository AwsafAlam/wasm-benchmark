fn image_grayscale(data: &mut [u8], width: usize, height: usize) {
    for i in 0..(width * height) {
        let r = data[i * 4];
        let g = data[i * 4 + 1];
        let b = data[i * 4 + 2];
        let gray = (0.2126 * r as f64 + 0.7152 * g as f64 + 0.0722 * b as f64) as u8;
        data[i * 4] = gray;
        data[i * 4 + 1] = gray;
        data[i * 4 + 2] = gray;
    }
}
