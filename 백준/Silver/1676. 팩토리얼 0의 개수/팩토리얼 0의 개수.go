package main

import (
	"fmt"
	"math/big"
)

func factorial(n int) *big.Int {
	result := big.NewInt(1)
	for i := 2; i <= n; i++ {
		result = result.Mul(result, big.NewInt(int64(i)))
	}
	return result
}

func main() {
	var num int
	fmt.Scanln(&num)

	str := factorial(num).String()

	cnt := 0
	for i:=len(str)-1; i > 0; i-- {
		if str[i] == '0' {
			cnt++
		} else {
			break
		}
	}

	fmt.Println(cnt)
}