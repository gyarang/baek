package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

type Stack struct {
	stack []int
	len   int
}

func NewStack(size int) Stack {
	return Stack{
		make([]int, 0, size),
		0,
	}
}

func (s *Stack) Push(val int) {
	s.stack = append(s.stack, val)
	s.len += 1
}

func (s *Stack) Pop() {
	if s.len == 0 {
		fmt.Fprintln(writer, "-1")
		return
	}
	p := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	s.len -= 1

	fmt.Fprintln(writer, p)
}

func (s *Stack) Size() {
	fmt.Fprintln(writer, s.len)
}

func (s *Stack) Empty() {
	if s.len == 0 {
		fmt.Fprintln(writer, "1")
	} else {
		fmt.Fprintln(writer, "0")
	}
}

func (s *Stack) Top() {
	if s.len == 0 {
		fmt.Fprintln(writer, "-1")
		return
	}
	fmt.Fprintln(writer, s.stack[len(s.stack)-1])
}

func (s *Stack) Message(m string, val int) {
	switch m {
	case "push":
		s.Push(val)
	case "pop":
		s.Pop()
	case "size":
		s.Size()
	case "empty":
		s.Empty()
	case "top":
		s.Top()
	}
}

func main() {
	defer writer.Flush()

	var n int
	fmt.Fscanln(reader, &n)

	s := NewStack(n)
	var method string
	var value int
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%s %d\n", &method, &value)
		s.Message(method, value)
	}
}
