package main

import "fmt"

func main() {
	var cnt, a, b int
	fmt.Scan(&cnt)
	for i := 0; i < cnt; i++ {
		fmt.Scan(&a, &b)
		fmt.Println(a + b)
	}
}
