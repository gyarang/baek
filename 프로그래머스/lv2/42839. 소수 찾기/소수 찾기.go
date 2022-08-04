var checked map[int]bool

func remove(s []int, index int) []int {
	newSlice := make([]int, len(s)-1)
	for i, v := range s {
		if i < index {
			newSlice[i] = v
		} else if i == index {
			continue
		} else {
			newSlice[i-1] = v
		}
	}
	return newSlice
}

func BruteForcePrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}


func dfs(res int, nums []int) {
	checked[res] = BruteForcePrime(res)

	for i, v := range nums {
		newRes := res * 10 + v
		if _, exist := checked[newRes]; exist {
			continue
		}
		dfs(newRes, append(remove(nums, i)))
	}
}

func solution(numbers string) int {
	checked = make(map[int]bool)

	nums := make([]int, len(numbers))
	for i, v := range numbers {
		nums[i] = int(v - '0')
	}

	dfs(0, nums)
	cnt := 0

	for _, isPrime := range checked {
		if isPrime {
			cnt++
		}
	}
	return cnt
}