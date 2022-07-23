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

	l := list.New()
	for i := 0; i < cnt; i++ {
		s := ""
		v := 0
		fmt.Fscanln(reader, &s, &v)

		switch s {
		case "push_front":
			l.PushFront(v)
		case "push_back":
			l.PushBack(v)
		case "pop_front":
			if l.Len() == 0 {
				buffer.WriteString("-1\n")
			} else {
				front := l.Front()
				buffer.WriteString(strconv.Itoa(front.Value.(int)))
				buffer.WriteRune('\n')
				l.Remove(front)
			}
		case "pop_back":
			if l.Len() == 0 {
				buffer.WriteString("-1\n")
			} else {
				back := l.Back()
				buffer.WriteString(strconv.Itoa(back.Value.(int)))
				buffer.WriteRune('\n')
				l.Remove(back)
			}
		case "size":
			buffer.WriteString(strconv.Itoa(l.Len()))
			buffer.WriteRune('\n')
		case "empty":
			if l.Len() == 0 {
				buffer.WriteString("1\n")
			} else {
				buffer.WriteString("0\n")
			}
		case "front":
			if l.Len() == 0 {
				buffer.WriteString("-1\n")
			} else {
				front := l.Front()
				buffer.WriteString(strconv.Itoa(front.Value.(int)))
				buffer.WriteRune('\n')
			}
		case "back":
			if l.Len() == 0 {
				buffer.WriteString("-1\n")
			} else {
				back := l.Back()
				buffer.WriteString(strconv.Itoa(back.Value.(int)))
				buffer.WriteRune('\n')
			}
		}
	}
	fmt.Print(buffer.String())
}
