package main

import "fmt"

func main() {
	var t int
	fmt.Scan(&t)

	var a, b, c int

	for t > 0 {
		if t%300 != 0 {
			if t%60 != 0 {
				t -= 10
				c++
				continue
			}
			t -= 60
			b++
			continue
		}
		a = t / 300
		break
	}

	if t < 0 {
		fmt.Println(-1)
		return
	}
	fmt.Println(a, b, c)
}
