package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var cnt int
	fmt.Fscanln(reader, &cnt)

	var input string
	fmt.Fscanln(reader, &input)

	cup := 1
	for i := 0; i < cnt; i++ {
		if input[i] == 'S' {
			cup++
			continue
		}
		if input[i] == 'L' {
			cup++
			i++
			continue
		}
	}
	if cup > cnt {
		fmt.Println(cnt)
	} else {
		fmt.Println(cup)
	}
}
