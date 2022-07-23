package main

import "fmt"

func main() {
	var cnt int
	fmt.Scanln(&cnt)

	if cnt%2 == 0 {
		fmt.Println("SK")
	} else {
		fmt.Println("CY")
	}
}
