package main

func imageConvolute(data, data2 []byte, width, height int, weights []float64, wwidth, wheight int) {
	halfWWidth := wwidth / 2
	halfWHeight := wheight / 2
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			var r, g, b, a float64
			for wy := 0; wy < wheight; wy++ {
				sy := y + wy - halfWHeight
				if sy < 0 || sy >= height {
					continue
				}
				for wx := 0; wx < wwidth; wx++ {
					sx := x + wx - halfWWidth
					if sx < 0 || sx >= width {
						continue
					}
					index := sy*width*4 + sx*4
					weight := weights[wy*wwidth+wx]
					r += float64(data[index+0]) * weight
					g += float64(data[index+1]) * weight
					b += float64(data[index+2]) * weight
					a += float64(data[index+3]) * weight
				}
			}
			index := y*width*4 + x*4
			data2[index+0] = byte(r)
			data2[index+1] = byte(g)
			data2[index+2] = byte(b)
			data2[index+3] = byte(a)
		}
	}
}

func main() {
	// sample usage
	width, height := 10, 10
	data := make([]byte, width*height*4)
	data2 := make([]byte, width*height*4)
	weights := []float64{0.25, 0.5, 0.25, 0.5, 1, 0.5, 0.25, 0.5, 0.25}
	wwidth, wheight := 3, 3
	imageConvolute(data, data2, width, height, weights, wwidth, wheight)
}
