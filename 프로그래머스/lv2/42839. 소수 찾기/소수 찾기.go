var cache [10000000]bool
var checked map[int]bool
var maxChecked = 0

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

func Eratosthenes(n int) bool {
	if n < 2 {
		return false
	}

	if n < maxChecked {
		return cache[n]
	}
	maxChecked = n

	for i := 2; i*i <= n; i++ {
		if cache[i] {
			for j := i * i; j <= n; j += i {
				cache[j] = false
			}
		}
	}

	return cache[n]
}

func dfs(res int, nums []int) {
	checked[res] = Eratosthenes(res)

	for i, v := range nums {
		newRes := res * 10 + v
		if _, exist := checked[newRes]; exist {
			continue
		}
		dfs(newRes, append(remove(nums, i)))
	}
}

func solution(numbers string) int {
	for i := range cache {
		cache[i] = true
	}

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