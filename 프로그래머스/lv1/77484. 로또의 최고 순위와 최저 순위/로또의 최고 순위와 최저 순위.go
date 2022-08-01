func getRank(win_cnt int) int {
    rank := 0
	switch win_cnt {
	case 6:
		rank = 1
	case 5:
		rank = 2
	case 4:
		rank = 3
	case 3:
		rank = 4
	case 2:
		rank = 5
	default:
		rank = 6
	}
    return rank
}

func countUnknown(lottos []int) int {
	cnt := 0
	for _, n := range lottos {
		if n == 0 {
			cnt += 1
		}
	}
	return cnt
}

func contains(lottos []int, num int) bool {
	for _, n := range lottos {
		if n == num {
			return true
		}
	}
    return false
}

func solution(lottos []int, win_nums []int) []int {
	unknowns := countUnknown(lottos)
	win_cnt := 0

	for _, n := range win_nums {
		if contains(lottos, n) {
			win_cnt += 1
		}
	}

	min := getRank(win_cnt)
	max := min - unknowns
	if max <= 0 {
		max = 1
	}

	return []int{max, min}
}