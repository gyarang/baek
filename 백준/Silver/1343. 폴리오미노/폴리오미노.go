package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	buffer := bytes.Buffer{}

	var input string
	fmt.Scan(&input)

	split := strings.Split(input, ".")
	result := make([]string, len(split))
	for i, v := range split {
		l := len(v)
		if l%2 != 0 {
			fmt.Println(-1)
			return
		}

		aCnt := l / 4
		l = l % 4
		bCnt := l / 2

		for i := 0; i < aCnt; i++ {
			buffer.WriteString("AAAA")
		}
		for i := 0; i < bCnt; i++ {
			buffer.WriteString("BB")
		}
		result[i] = buffer.String()
		buffer.Reset()
	}

	fmt.Println(strings.Join(result, "."))
}
