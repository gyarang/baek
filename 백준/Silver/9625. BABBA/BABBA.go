package main

import "fmt"

func main() {
	cache := make([][2]int, 46)
	cache[0] = [2]int{1, 0}
	cache[1] = [2]int{0, 1}

	var cnt int
	fmt.Scan(&cnt)

	if cnt < 2 {
		fmt.Println(cache[cnt][0], cache[cnt][1])
		return
	}

	for i := 2; i <= cnt; i++ {
		cache[i][0], cache[i][1] = cache[i - 1][0] + cache[i - 2][0], cache[i - 1][1] + cache[i - 2][1]
	}
	fmt.Println(cache[cnt][0], cache[cnt][1])
}
