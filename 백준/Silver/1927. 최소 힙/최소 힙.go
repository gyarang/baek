package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
)

type Item struct {
	value int
	priority int
	index int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool { return !(pq[i].value > pq[j].value) }

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(i interface{}) {
	n := len(*pq)
	item := i.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(*pq)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[:n-1]
	return item
}

func (pq *PriorityQueue) update(item *Item, value int, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}


func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	cnt := 0
	fmt.Fscanln(reader, &cnt)

	pq := PriorityQueue{}

	for i := 0; i < cnt; i++ {
		input := 0
		fmt.Fscanln(reader, &input)

		if input == 0 {
			if len(pq) == 0 {
				writer.WriteString("0\n")
			} else {
				val := heap.Pop(&pq)
				writer.WriteString(strconv.Itoa(val.(*Item).value) + "\n")
			}
		} else {
			item := &Item{value: input}
			heap.Push(&pq, item)
		}
	}
}