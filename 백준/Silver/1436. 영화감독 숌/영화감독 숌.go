package main

import (
	"fmt"
)

func main() {
	input := 0
	fmt.Scan(&input)

	cnt := 0
	result := 665
	for cnt < input {
		result++
		temp := result
		for temp >= 666 {
			if temp%1000 == 666 {
				cnt++
				break
			}
			temp /= 10
		}
	}
	fmt.Println(result)
}
