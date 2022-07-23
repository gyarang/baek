package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type user struct {
	age  int
	name string
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscanln(reader, &n)

	var age int
	var name string
	users := make([]user, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d %s\n", &age, &name)
		users[i] = user{age, name}
	}

	sort.SliceStable(users, func(i, j int) bool {
		return users[i].age < users[j].age
	})

	for _, u := range users {
		fmt.Fprintln(writer, u.age, u.name)
	}
}
