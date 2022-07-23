package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	buffer := bytes.Buffer{}

	var cnt, h, w, num int
	fmt.Fscanln(reader, &cnt)

	for i := 0; i < cnt; i++ {
		fmt.Fscanf(reader, "%d %d %d\n", &h, &w, &num)
		room := ((num-1)%h+1)*100 + ((num - 1) / h) + 1
		buffer.WriteString(strconv.Itoa(room))
		buffer.WriteByte('\n')
	}

	fmt.Print(buffer.String())
}
