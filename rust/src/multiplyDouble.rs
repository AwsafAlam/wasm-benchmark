func multiplyDouble(a, b float64, n int) float64 {
	c := 1.0
	for i := 0; i < n; i++ {
		c = c * a * b
	}
	return c
}
