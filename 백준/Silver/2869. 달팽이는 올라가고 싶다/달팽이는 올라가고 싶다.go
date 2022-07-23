package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var speed, slip, height int
	fmt.Fscanf(reader, "%d %d %d", &speed, &slip, &height)

	day := (height - slip) / (speed - slip)
	if ((height - slip) % (speed - slip)) != 0 {
		day++
	}
	fmt.Println(day)
}
