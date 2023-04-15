func fibonacci(num int) {
	a := 1
	b := 0
	var temp int
	for num >= 0 {
		temp = a
		a = a + b
		b = temp
		num--
	}
}