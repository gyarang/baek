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

	for {
		name, age, weight := "", 0, 0
		fmt.Fscanln(reader, &name, &age, &weight)
		if age == 0 {
			return
		}
		if weight >= 80 || age > 17 {
			fmt.Fprintln(writer, name, "Senior")
		} else {
			fmt.Fprintln(writer, name, "Junior")
		}
	}
}
