package main

import (
	"bufio"
	"fmt"
	"os"
)

func isPrime(num int) bool {
	if num < 2 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	cnt := 0
	fmt.Scanln(&cnt)

	var num, result int
	for i := 0; i < cnt; i++ {
		fmt.Fscan(reader, &num)
		if isPrime(num) {
			result++
		}
	}
	fmt.Println(result)
}
