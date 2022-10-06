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

	var a, b []byte
	fmt.Fscanf(reader, "%s\n%s", &a, &b)

	// AND
	for i := 0; i < len(a); i++ {
		fmt.Fprint(writer, (a[i]-'0')&(b[i]-'0'))
	}
	fmt.Fprintln(writer, "")

	// OR
	for i := 0; i < len(a); i++ {
		fmt.Fprint(writer, (a[i]-'0')|(b[i]-'0'))
	}
	fmt.Fprintln(writer, "")

	// XOR
	for i := 0; i < len(a); i++ {
		fmt.Fprint(writer, (a[i]-'0')^(b[i]-'0'))
	}
	fmt.Fprintln(writer, "")

	// NOT
	for i := 0; i < len(a); i++ {
		fmt.Fprint(writer, 1^(a[i]-'0'))
	}
	fmt.Fprintln(writer, "")
	for i := 0; i < len(b); i++ {
		fmt.Fprint(writer, 1^(b[i]-'0'))
	}
}