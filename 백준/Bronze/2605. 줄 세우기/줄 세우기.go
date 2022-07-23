package main

import (
	"bufio"
	"fmt"
	"os"
)

var rd = bufio.NewReader(os.Stdin)
var wr = bufio.NewWriter(os.Stdout)

type ll struct {
	head   *node
	length int
}

func NewList() *ll {
	return &ll{}
}

func (l *ll) Get(cnt int) *node {
	currentNode := l.head
	for i := 0; i < cnt; i++ {
		currentNode = currentNode.next
	}
	return currentNode
}

func (l *ll) Add(name, num int) {
	n := &node{name: name}

	if num == l.length {
		n.next = l.head
		l.head = n
		l.length++
		return
	}

	var prev, next *node
	prev = l.Get(l.length - 1 - num)
	next = prev.next

	n.next = next
	prev.next = n
	l.length++
	return
}

func (l *ll) Print() {
	n := l.head
	for n != nil {
		fmt.Fprintf(wr, "%d ", n.name)
		n = n.next
	}
	fmt.Fprint(wr, "\n")
}

type node struct {
	name int
	next *node
}

func main() {
	defer wr.Flush()

	var cnt int
	fmt.Fscanln(rd, &cnt)

	list := NewList()
	var ticket int
	for i := 1; i <= cnt; i++ {
		fmt.Fscanf(rd, "%d ", &ticket)
		list.Add(i, ticket)
	}
	list.Print()
}
