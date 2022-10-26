package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var cnt int
	fmt.Fscanln(reader, &cnt)

	for i := 0; i < cnt; i++ {
		var x1, y1, r1, x2, y2, r2 int
		fmt.Fscanln(reader, &x1, &y1, &r1, &x2, &y2, &r2)

		// 동일한 원
		if x1 == x2 && y1 == y2 && r1 == r2 {
			fmt.Fprintln(writer, "-1")
			continue
		}

		distance := (x1-x2)*(x1-x2) + (y1-y2)*(y1-y2) // 터렛 사이의 거리 ^2
		rDiffSquare := (r1 - r2) * (r1 - r2)          // 반지름 차 ^2
		rAddSquare := (r1 + r2) * (r1 + r2)           // 반지름 합 ^2

		if distance == rDiffSquare || distance == rAddSquare {
			fmt.Fprintln(writer, "1")
		} else if distance < rAddSquare && distance > rDiffSquare {
			fmt.Fprintln(writer, "2")
		} else {
			fmt.Fprintln(writer, "0")
		}
	}
}
