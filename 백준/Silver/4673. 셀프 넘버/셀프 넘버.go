package main

import (
	"bytes"
	"fmt"
	"strconv"
)

func main() {
	var buf bytes.Buffer

	nm := make(map[int]struct{})
	for i := 1; i <= 10000; i++ {
		num, sum := i, i
		for num > 0 {
			sum += num % 10
			num /= 10
		}

		if _, ex := nm[sum]; !ex {
			nm[sum] = struct{}{}
		}
		if _, ex := nm[i]; !ex {
			buf.WriteString(strconv.Itoa(i))
			buf.WriteRune('\n')
		}
	}
	fmt.Print(buf.String())
}
