package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	cnt := 0
	player := 0

	fmt.Fscanln(reader, &cnt)
	fmt.Fscan(reader, &player)

	arr := make([]int, cnt-1)
	for i := 0; i < cnt-1; i++ {
		fmt.Fscan(reader, &arr[i])
	}

	sort.Ints(arr)
	for _, v := range arr {
		if player <= v {
			fmt.Println("No")
			return
		}
		player += v
	}
	fmt.Println("Yes")
}
