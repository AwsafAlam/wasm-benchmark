func sumInt(array []int, n int) int {
	s := 0
	for i := 0; i < n; i++ {
		s += array[i]
	}
	return s
}
