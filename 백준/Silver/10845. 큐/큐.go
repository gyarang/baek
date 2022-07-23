package main

import (
	"bufio"
	"bytes"
	"container/list"
	"fmt"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	buffer := bytes.Buffer{}

	cnt := 0
	fmt.Fscanln(reader, &cnt)

	queue := list.New()
	method, val := "", 0
	for i := 0; i < cnt; i++ {
		fmt.Fscanln(reader, &method, &val)
		switch method {
		case "push":
			queue.PushBack(val)
		case "pop":
			pop := queue.Front()
			if pop == nil {
				buffer.WriteString("-1\n")
			} else {
				buffer.WriteString(strconv.Itoa(pop.Value.(int)))
				buffer.WriteRune('\n')
				queue.Remove(pop)
			}
		case "size":
			buffer.WriteString(strconv.Itoa(queue.Len()))
			buffer.WriteRune('\n')
		case "empty":
			if queue.Len() == 0 {
				buffer.WriteString("1\n")
			} else {
				buffer.WriteString("0\n")
			}
		case "front":
			pop := queue.Front()
			if pop == nil {
				buffer.WriteString("-1\n")
			} else {
				buffer.WriteString(strconv.Itoa(pop.Value.(int)))
				buffer.WriteRune('\n')
			}
		case "back":
			pop := queue.Back()
			if pop == nil {
				buffer.WriteString("-1\n")
			} else {
				buffer.WriteString(strconv.Itoa(pop.Value.(int)))
				buffer.WriteRune('\n')
			}
		}
	}
	fmt.Print(buffer.String())
}
