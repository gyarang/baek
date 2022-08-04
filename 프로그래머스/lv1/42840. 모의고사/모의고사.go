func max(s []int) int {
	max := 0
	for _, v := range s {
		if v > max {
			max = v
		}
	}
	return max
}

func checkPoint(answer, pattern []int) int {
	point := 0
	for i, v := range answer {
		if v == pattern[i%len(pattern)] {
			point++
		}
	}
	return point
}

func solution(answers []int) []int {
	patterens := [][]int{
		[]int{1, 2, 3, 4, 5},
		[]int{2, 1, 2, 3, 2, 4, 2, 5},
		[]int{3, 3, 1, 1, 2, 2, 4, 4, 5, 5}}

	points := make([]int, 3)
	for i, pattern := range patterens {
		points[i] = checkPoint(answers, pattern)
	}

	max := max(points)
	result := make([]int, 0, 3)
	for i, v := range points {
		if v == max {
			result = append(result, i+1)
		}
	}

	return result
}