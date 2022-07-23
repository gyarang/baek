package main

import (
	"fmt"
)

var factorialArray []int

func factorial(num int) int {
	if num < 1 {
		return 1
	}
	if factorialArray[num] != 0 {
		return factorialArray[num]
	}
	for i := 2; i <= num; i++ {
		factorialArray[i] = factorialArray[i-1] * i
	}
	return factorialArray[num]
}

func main() {
	factorialArray = make([]int, 11)
	factorialArray[1] = 1

	var n, k int
	fmt.Scanf("%d %d", &n, &k)

	result := factorial(n) / (factorial(k) * factorial(n-k))
	fmt.Println(result)
}
