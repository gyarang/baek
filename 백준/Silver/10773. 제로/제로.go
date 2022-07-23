package main

import (
	"bufio"
	"fmt"
	"os"
)

type Stack struct {
	stack  []int
	length int
}

func (s *Stack) Pop() int {
	if s.length == 0 {
		return 0
	}
	s.length--
	pop := s.stack[s.length]
	s.stack = s.stack[:s.length]
	return pop
}

func (s *Stack) Push(v int) {
	s.length++
	s.stack = append(s.stack, v)
	return
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var cnt int
	fmt.Fscanln(reader, &cnt)

	var input, sum int
	s := Stack{make([]int, 0, cnt), 0}
	for i := 0; i < cnt; i++ {
		fmt.Fscanln(reader, &input)
		if input == 0 {
			pop := s.Pop()
			sum -= pop
		} else {
			sum += input
			s.Push(input)
		}
	}
	fmt.Println(sum)
}
