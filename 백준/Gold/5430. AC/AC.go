package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func solution(code string, arr []int) string {
	isRevered := false

	for _, v := range code{
		switch v {
		case 'R':
			isRevered = !isRevered
		case 'D':
			if len(arr) == 0 {
				return "error"
			}
			if isRevered {
				arr = arr[:len(arr) - 1]
			} else {
				arr = arr[1:]
			}
		}
	}

	if isRevered {
		for i := 0; i < len(arr)/2; i++ {
			arr[i], arr[len(arr) - 1 - i] = arr[len(arr) -1 - i], arr[i]
		}
	}

	result, _ := json.Marshal(arr)
	return string(result)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var testCases int
	fmt.Fscanln(reader, &testCases)

	for testIndex := 0; testIndex < testCases; testIndex++ {
		var code, arrString string
		var cnt int

		fmt.Fscanln(reader, &code)
		fmt.Fscanln(reader, &cnt)
		fmt.Fscanln(reader, &arrString)

		arr := make([]int, cnt)
		_ = json.Unmarshal([]byte(arrString), &arr)

		result := solution(code, arr)
		fmt.Fprintln(writer, result)
	}
}
