func solution(n int, lost []int, reserve []int) int {
	uniforms := make([]int, n)
	for i := range uniforms {
		uniforms[i] = 1
	}
	for _, v := range lost {
		uniforms[v-1]--
	}
	for _, v := range reserve {
		uniforms[v-1]++
	}

	for i := 0; i < n; i++ {
		if i > 0 && uniforms[i] == 0 && uniforms[i-1] == 2 {
			uniforms[i] = 1
			uniforms[i-1] = 1
		} else if i < n-1 && uniforms[i] == 0 && uniforms[i+1] == 2{
			uniforms[i] = 1
			uniforms[i+1] = 1
		}
	}

	cnt := 0
	for _, v := range uniforms {
		if v != 0 {
			cnt ++
		}
	}
	return cnt
}