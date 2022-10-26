package main

import (
	"fmt"
)

func main() {
	var x, y int
	fmt.Scanln(&x, &y)
	z := y * 100 / x
	if z >= 99 {
		fmt.Println("-1")
		return
	}

	left, right := 1, x

	for left <= right {
		mid := (left + right) / 2
		cur := (y + mid) * 100 / (x + mid)
		if z == cur {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	fmt.Println(left)
}
