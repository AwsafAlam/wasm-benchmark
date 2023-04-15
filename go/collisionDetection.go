package main

import "math"

type position struct {
	x, y, z float64
}

func collisionDetection(positions []position, radiuses []float64, res []byte, n int) int {
	count := 0
	for i := 0; i < n; i++ {
		p := positions[i]
		r := radiuses[i]
		collision := false
		for j := i + 1; j < n; j++ {
			p2 := positions[j]
			r2 := radiuses[j]
			dx := p.x - p2.x
			dy := p.y - p2.y
			dz := p.z - p2.z
			d := math.Sqrt(dx*dx + dy*dy + dz*dz)
			if r > d {
				collision = true
				count++
				break
			}
		}
		index := i / 8
		pos := 7 - (i % 8)
		if collision {
			res[index] |= 1 << pos
		} else {
			res[index] &= ^(1 << pos)
		}
	}
	return count
}

func main() {
	// sample usage
	positions := []position{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	radiuses := []float64{1, 2, 3}
	res := make([]byte, (len(positions)+7)/8)
	n := len(positions)
	count := collisionDetection(positions, radiuses, res, n)
	println(count)
}
