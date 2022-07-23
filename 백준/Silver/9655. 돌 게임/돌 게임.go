package main

import "fmt"

func main() {
	var cnt int
	fmt.Scanln(&cnt)

	arr := make([]bool, cnt+3)
	arr[1] = true
	arr[2] = false
	arr[3] = true

	for i := 4; i <= cnt; i++ {
		if arr[i-1] && arr[i-3] {
			arr[i] = false
		} else {
			arr[i] = true
		}
	}

	if arr[cnt] {
		fmt.Println("SK")
	} else {
		fmt.Println("CY")
	}
}
