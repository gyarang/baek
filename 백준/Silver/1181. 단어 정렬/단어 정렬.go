package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var l int
	fmt.Fscanln(reader, &l)

	var input string
	arr := make([]string, 0, l)
	imap := make(map[string]struct{})
	for i := 0; i < l; i++ {
		fmt.Fscanln(reader, &input)
		if _, ok := imap[input]; !ok {
			imap[input] = struct{}{}
			arr = append(arr, input)
		}
	}

	sort.Slice(arr, func(i, j int) bool {
		if len(arr[i]) != len(arr[j]) {
			return len(arr[i]) < len(arr[j])
		}

		iStrArr := []uint8(arr[i])

		for index, v := range iStrArr {
			if v != arr[j][index] {
				return v < arr[j][index]
			}
		}
		return false
	})

	for _, v := range arr {
		fmt.Fprintln(writer, v)
	}
}
