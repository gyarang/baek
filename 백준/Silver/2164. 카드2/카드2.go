package main

import (
	"container/ring"
	"fmt"
)

func main() {
	var cnt int
	fmt.Scan(&cnt)

	r := ring.New(cnt)

	for i := 1; i <= cnt; i++ {
		r.Value = i
		r = r.Next()
	}

	r = r.Move(-1)
	for cnt != 1 {
		cnt--
		r.Unlink(1)
		r = r.Move(1)
	}

	fmt.Println(r.Value.(int))
}
