package main

import (
	"bufio"
	"fmt"
	"os"
)

var arr [][]int
var visited []bool

func dfs(head int) bool {
	if visited[head] {
		return false
	}
	visited[head] = true
	for i := 0; i < len(arr[head]); i++ {
		if arr[head][i] == 1 && !visited[i] {
			dfs(i)
		}
	}
	return true
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var N, M int
	fmt.Fscanln(reader, &N, &M)

	visited = make([]bool, N+1)
	arr = make([][]int, N+1)
	for i := range arr {
		arr[i] = make([]int, N+1)
	}

	for i := 0; i < M; i++ {
		var a, b int
		fmt.Fscanln(reader, &a, &b)
		arr[a][b] = 1
		arr[b][a] = 1
	}

	result := 0
	for i := 1; i <= N; i++ {
		if dfs(i) {
			result++
		}
	}

	fmt.Fprintln(writer, result)
}
