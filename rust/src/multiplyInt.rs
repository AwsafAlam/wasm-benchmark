func multiplyInt(a, b, n int) int {
	c := 1
	for i := 0; i < n; i++ {
		c = c * a * b
	}
	return c
}
