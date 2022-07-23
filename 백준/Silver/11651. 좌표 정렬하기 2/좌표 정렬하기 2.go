package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type point struct {
	X int
	Y int
}

type pointArr []point

func (p pointArr) Len() int {
	return len(p)
}
func (p pointArr) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
func (p pointArr) Less(i, j int) bool {
	if p[i].Y < p[j].Y {
		return true
	} else if p[i].Y == p[j].Y {
		return p[i].X < p[j].X
	} else {
		return false
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var l int
	fmt.Fscanln(reader, &l)

	var x, y int
	points := make([]point, l)
	for i := 0; i < l; i++ {
		fmt.Fscanf(reader, "%d %d\n", &x, &y)
		points[i] = point{x, y}
	}

	sort.Sort(pointArr(points))
	for _, v := range points {
		fmt.Fprintf(writer, "%d %d\n", v.X, v.Y)
	}
}
