func solution(n int, a int, b int) int {
	answer := 1

	if a > b {
		a, b = b, a
	}

	for b-a != 1 || a % 2 == 0{
		answer++
		a = (a + 1) / 2
		b = (b + 1) / 2
	}

	return answer
}