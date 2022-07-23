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

	s := make(map[int]struct{})

	var input, cnt int
	for i := 0; i < 10; i++ {
		fmt.Fscan(reader, &input)
		if _, ok := s[input%42]; !ok {
			s[input%42] = struct{}{}
			cnt++
		}
	}
	fmt.Fprint(writer, cnt)
}
