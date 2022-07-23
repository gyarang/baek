package main

import "fmt"

func main() {
	input := 0
	fmt.Scanln(&input)

	size := 0
	temp := input
	for temp > 0 {
		temp /= 10
		size++
	}
	start := input - (size * 9)
	for i := start; i < input; i++ {
		sum := i
		temp := i
		for temp > 0 {
			sum += temp % 10
			temp /= 10
		}
		if sum == input {
			fmt.Println(i)
			return
		}
	}
	fmt.Println(0)
}
