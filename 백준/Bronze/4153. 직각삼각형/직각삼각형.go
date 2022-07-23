package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := make([]int, 3)
	for {
		for i := 0; i < 3; i++ {
			fmt.Scan(&arr[i])
			if arr[0] == 0 {
				return
			}
		}
		sort.Ints(arr)

		if arr[2]*arr[2] == arr[0]*arr[0]+arr[1]*arr[1] {
			fmt.Println("right")
		} else {
			fmt.Println("wrong")
		}
	}
}
