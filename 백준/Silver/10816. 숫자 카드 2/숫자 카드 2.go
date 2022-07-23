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

	var inCnt, outCnt, input int
	fmt.Fscanln(reader, &inCnt)

	dict := make(map[int]int)
	for i := 0; i < inCnt; i++ {
		fmt.Fscan(reader, &input)
		dict[input]++
	}

	fmt.Fscanf(reader, "\n%d\n", &outCnt)
	for i := 0; i < outCnt; i++ {
		fmt.Fscan(reader, &input)
		buffer.WriteString(strconv.Itoa(dict[input]))
		buffer.WriteRune(' ')
	}
	fmt.Println(buffer.String())
}
