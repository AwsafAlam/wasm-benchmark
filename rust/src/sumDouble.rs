package main

func sumDouble(array []float64, n int) float64 {
	s := 0.0
	for i := 0; i < n; i++ {
		s += array[i]
	}
	return s
}
