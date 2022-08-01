func solution(n int, left int64, right int64) []int {
	arr := make([]int, right - left + 1)
	for i := left; i < right + 1; i++ {
		x := i / int64(n)
		y := i % int64(n)
		var max int
		if x > y {
			max = int(x) + 1
		} else {
			max = int(y) + 1
		}
		arr[i - left] = max
	}
	return arr
}