package programmers

import (
	"sort"
	"math"
)

func checkPrime(num int) bool {
	sq_root := int(math.Sqrt(float64(num)))
	for i := 2; i <= sq_root; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func solution(nums []int) int {
    cnt := 0
	sort.Ints(nums[:])

	for i, v1 := range nums {
		for j, v2 := range nums[i+1:] {
			for _, v3 := range nums[i+j+2:] {
				sum := v1 + v2 + v3
				if checkPrime(sum) {
                    cnt +=1
				}
			}
		}
	}
	return cnt
}