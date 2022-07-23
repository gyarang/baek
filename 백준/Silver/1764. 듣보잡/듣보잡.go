package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var dl, bl int
	fmt.Fscanf(reader, "%d %d\n", &dl, &bl)

	ds := make(map[string]struct{})
	var name string
	for i := 0; i < dl; i++ {
		fmt.Fscan(reader, &name)
		ds[name] = struct{}{}
	}

	var buf bytes.Buffer
	dbs := make([]string, 0, len(ds))
	for i := 0; i < bl; i++ {
		fmt.Fscan(reader, &name)
		if _, ok := ds[name]; ok {
			dbs = append(dbs, name)
		}
	}

	sort.Sort(sort.StringSlice(dbs))
	buf.WriteString(strconv.Itoa(len(dbs)))
	buf.WriteRune('\n')
	for _, v := range dbs {
		buf.WriteString(v)
		buf.WriteRune('\n')
	}
	fmt.Print(buf.String())
}
