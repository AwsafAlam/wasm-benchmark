func quicksortInt(array []int, start int, end int) {
	if start >= end {
		return
	}
	pivot := array[end]
	left := 0
	right := 0
	for left+right < end-start {
		num := array[start+left]
		if num < pivot {
			left++
		} else {
			array[start+left] = array[end-right-1]
			array[end-right-1] = pivot
			array[end-right] = num
			right++
		}
	}
	quicksortInt(array, start, start+left-1)
	quicksortInt(array, start+left+1, end)
}
