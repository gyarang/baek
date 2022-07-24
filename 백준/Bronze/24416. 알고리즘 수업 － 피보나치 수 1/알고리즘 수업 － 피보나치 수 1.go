package main

import "fmt"

func main() {
	var cnt int
	fmt.Scanln(&cnt)

	dpArr := make([]int, cnt+1)
	dpArr[1] = 1
	dpArr[2] = 1
	
	for i := 3; i <= cnt; i++ {
		dpArr[i] = dpArr[i-1] + dpArr[i-2]
	}

	dpCnt := cnt-2
	if dpCnt <= 0 {
		dpCnt = 1
	}

	fmt.Println(dpArr[cnt], dpCnt)
}