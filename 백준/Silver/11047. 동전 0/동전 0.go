package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var coinTypes, target int
	fmt.Fscanln(reader, &coinTypes, &target)
	coins := make([]int, coinTypes)

	for i := 0; i < coinTypes; i++ {
		fmt.Fscan(reader, &coins[i])
	}

	result := 0
	for i := len(coins)-1; i >= 0; i-- {
		if target == 0 {
			break
		}

		if coins[i] > target {
			continue
		}

		cnt := target/coins[i]
		target -= cnt * coins[i]
		result += cnt
	}

	fmt.Fprintln(writer, result)
}