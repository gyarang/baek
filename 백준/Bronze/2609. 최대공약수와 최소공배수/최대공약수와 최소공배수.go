package main

import "fmt"

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	min := gcd(a, b)
	max := a * b / min

	fmt.Printf("%d\n%d", min, max)
}
