package main

import (
	"bufio"
	"fmt"
	"os"
)

func solution(input string) bool {
	stack := make([]int32, 0, len(input))
	for _, v := range input {
		if v == '(' || v == '[' {
			stack = append(stack, v)
			continue
		}
		if v == ')' {
			if len(stack) == 0 {
				return false
			}
			if stack[len(stack)-1] != '(' {
				return false
			} else {
				stack = stack[:len(stack)-1]
				continue
			}
		}
		if v == ']' {
			if len(stack) == 0 {
				return false
			}
			if stack[len(stack)-1] != '[' {
				return false
			} else {
				stack = stack[:len(stack)-1]
				continue
			}
		}
	}
	if len(stack) != 0 {
		return false
	} else {
		return true
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var input string
	for {
		input, _ = reader.ReadString('\n')
		if input == ".\n" {
			break
		}
		if solution(input) {
			fmt.Fprintln(writer, "yes")
		} else {
			fmt.Fprintln(writer, "no")
		}
	}
}
