use std::f64;

struct Position {
    x: f64,
    y: f64,
    z: f64,
}

fn collision_detection(
    positions: &Vec<Position>,
    radiuses: &Vec<f64>,
    res: &mut Vec<u8>,
    n: usize,
) -> i32 {
    let mut count = 0;
    for i in 0..n {
        let p = &positions[i];
        let r = radiuses[i];
        let mut collision = false;
        for j in i + 1..n {
            let p2 = &positions[j];
            let r2 = radiuses[j];
            let dx = p.x - p2.x;
            let dy = p.y - p2.y;
            let dz = p.z - p2.z;
            let d = f64::sqrt(dx * dx + dy * dy + dz * dz);
            if r > d {
                collision = true;
                count += 1;
                break;
            }
        }
        let index = i / 8;
        let pos = 7 - (i % 8);
        if !collision {
            res[index] &= !(1 << pos);
        } else {
            res[index] |= 1 << pos;
        }
    }
    count
}
