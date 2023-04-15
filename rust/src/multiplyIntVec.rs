func multiplyIntVec(src1, src2, res []int, n int) {
	for i := 0; i < n; i++ {
		res[i] = src1[i] * src2[i]
	}
}
