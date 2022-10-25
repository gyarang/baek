package main

import (
	"fmt"
)

func main() {
	var cnt int
	fmt.Scanln(&cnt)

	if cnt == 1 {
		fmt.Println("1")
		return
	}

	prev := 1
	now := 3
	for i := 2; i < cnt; i++ {
		prev, now = now, (prev*2+now)%10007
	}

	fmt.Println(now)
}
