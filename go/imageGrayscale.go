package mypackage

func ImageGrayscale(data []byte, width int, height int) {
	for i := 0; i < width*height; i++ {
		r := data[i*4]
		g := data[i*4+1]
		b := data[i*4+2]
		gray := uint8(0.2126*float64(r) + 0.7152*float64(g) + 0.0722*float64(b))
		data[i*4] = gray
		data[i*4+1] = gray
		data[i*4+2] = gray
	}
}
