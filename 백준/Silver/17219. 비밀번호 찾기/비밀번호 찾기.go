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

	urlCnt, findCnt := 0, 0
	fmt.Fscanln(reader, &urlCnt, &findCnt)

	passwords := make(map[string]string)
	for i := 0; i < urlCnt; i++ {
		url, pw := "" ,""
		fmt.Fscanln(reader, &url, &pw)
		passwords[url] = pw
	}

	for i := 0; i < findCnt; i++ {
		url := ""
		fmt.Fscanln(reader, &url)
		fmt.Fprintln(writer, passwords[url])
	}
}