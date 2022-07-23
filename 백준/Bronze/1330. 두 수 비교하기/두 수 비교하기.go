package main

import "fmt"

func main() {
	var a, b int
	fmt.Scanf("%d %d", &a, &b)
	switch {
	case a > b:
		fmt.Printf("%s\n", ">")
	case a == b:
		fmt.Printf("%s\n", "==")
	default:
		fmt.Printf("%s\n", "<")
	}
}
