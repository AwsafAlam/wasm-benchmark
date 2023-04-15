use std::f64;

struct Position {
    x: f64,
    y: f64,
    z: f64,
}

fn image_convolute(
    data: &mut Vec<u8>,
    data2: &mut Vec<u8>,
    width: i32,
    height: i32,
    weights: &Vec<f64>,
    wwidth: i32,
    wheight: i32,
) {
    let half_wwidth = wwidth / 2;
    let half_wheight = wheight / 2;
    for y in 0..height {
        for x in 0..width {
            let mut r = 0.0;
            let mut g = 0.0;
            let mut b = 0.0;
            let mut a = 0.0;
            for wy in 0..wheight {
                let sy = y + wy - half_wheight;
                if sy < 0 || sy >= height {
                    continue;
                }
                for wx in 0..wwidth {
                    let sx = x + wx - half_wwidth;
                    if sx < 0 || sx >= width {
                        continue;
                    }
                    let index = (sy * width + sx) as usize;
                    let weight = weights[(wy * wwidth + wx) as usize];
                    r += f64::from(data[index * 4 + 0]) * weight;
                    g += f64::from(data[index * 4 + 1]) * weight;
                    b += f64::from(data[index * 4 + 2]) * weight;
                    a += f64::from(data[index * 4 + 3]) * weight;
                }
            }
            let index = (y * width + x) as usize;
            data2[index * 4 + 0] = r as u8;
            data2[index * 4 + 1] = g as u8;
            data2[index * 4 + 2] = b as u8;
            data2[index * 4 + 3] = a as u8;
        }
    }
}
