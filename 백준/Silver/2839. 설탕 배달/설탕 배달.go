package main

import "fmt"

func main() {
	var w, cnt int
	fmt.Scan(&w)

	for {
		if w%5 != 0 {
			if w < 3 {
				cnt = -1
				break
			}
			w -= 3
			cnt += 1
		} else {
			break
		}
	}

	if cnt != -1 {
		cnt += w / 5
	}
	fmt.Println(cnt)
}
