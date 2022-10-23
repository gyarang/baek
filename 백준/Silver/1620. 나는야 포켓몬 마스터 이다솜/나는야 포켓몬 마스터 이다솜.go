package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var kinds, cnt int

	fmt.Fscanln(reader, &kinds, &cnt)

	pokemons := make([]string, kinds)
	byName := make(map[string]int)

	for i := 0; i < kinds; i++ {
		var poke string
		fmt.Fscanln(reader, &poke)
		pokemons[i] = poke
		byName[poke] = i
	}

	for i := 0; i < cnt; i++ {
		var input string
		fmt.Fscanln(reader, &input)
		if num, ok := byName[input]; ok {
			fmt.Fprintln(writer, num + 1)
		} else {
			num, _ = strconv.Atoi(input)
			fmt.Fprintln(writer, pokemons[num - 1])
		}
	}
}