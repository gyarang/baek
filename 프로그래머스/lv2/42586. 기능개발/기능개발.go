func remainDates(progress, speed int) int {
	remain := (100 - progress) / speed
	if (100 - progress)%speed != 0 {
		remain++
	}
	return remain
}

func solution(progresses []int, speeds []int) []int {
	var result []int

	remains := make([]int, len(progresses))
	for i := 0; i < len(progresses); i++ {
		remains[i] = remainDates(progresses[i], speeds[i])
	}

	for len(remains) > 0 {
		i := 1
		for ; i < len(remains); i++ {
			if remains[0] < remains[i] {
				break
			}
		}
		remains = remains[i:]
		result = append(result, i)
	}

	return result
}