package main

import (
	"bufio"
	"fmt"
	"os"
)

var visited []bool
var graph [][]int
var result int

func dfs(start int) {
	visited[start] = true
	result++
	for i := 0; i < len(graph[start]); i++ {
		if graph[start][i] == 1 && !visited[i] {
			dfs(i)
		}
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var comCnt, edgeCnt int
	fmt.Fscanln(reader, &comCnt)

	visited = make([]bool, comCnt+1)
	graph = make([][]int, comCnt+1)
	for i := range graph {
		graph[i] = make([]int, comCnt+1)
	}

	fmt.Fscanln(reader, &edgeCnt)
	for i := 0; i < edgeCnt; i++ {
		var a, b int
		fmt.Fscanln(reader, &a, &b)
		graph[a][b] = 1
		graph[b][a] = 1
	}

	dfs(1)
	fmt.Fprintln(writer, result-1)
}
