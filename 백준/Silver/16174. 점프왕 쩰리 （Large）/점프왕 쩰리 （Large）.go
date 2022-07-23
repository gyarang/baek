package main

import (
	"bufio"
	"fmt"
	"os"
)

type Queue struct {
	arr    []Pair
	Length int
}

func (q *Queue) Push(x, y int) {
	q.arr = append(q.arr, Pair{x, y})
	q.Length += 1
}

func (q *Queue) Pop() Pair {
	p := q.arr[len(q.arr)-1]
	q.arr = q.arr[:len(q.arr)-1]
	q.Length -= 1
	return p
}

type Pair struct {
	x, y int
}

func main() {
	var size int
	var field [][]int
	var visited [64][64]bool

	reader := bufio.NewReader(os.Stdin)
	fmt.Fscan(reader, &size)

	field = make([][]int, size)
	for i := 0; i < size; i++ {
		row := make([]int, size)
		for j := 0; j < size; j++ {
			fmt.Fscan(reader, &row[j])
		}
		field[i] = row
	}

	q := Queue{
		arr: []Pair{},
	}
	q.Push(0, 0)

	for q.Length != 0 {
		p := q.Pop()
		value := field[p.x][p.y]
		visited[p.x][p.y] = true

		if value == -1 {
			fmt.Println("HaruHaru")
			return
		}

		if value == 0 {
			continue
		}

		if p.x+value < size && !visited[p.x+value][p.y] {
			q.Push(p.x+value, p.y)
		}
		if p.y+value < size && !visited[p.x][p.y+value] {
			q.Push(p.x, p.y+value)
		}
	}
	fmt.Println("Hing")
}
