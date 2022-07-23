package main

import (
	"fmt"
	"math/big"
)

func main() {
	var a, b big.Int
	fmt.Scanf("%d %d", &a, &b)

	div := new(big.Int)
	div = div.Div(&a, &b)

	mod := new(big.Int)
	mod = mod.Mod(&a, &b)

	fmt.Printf("%d\n", div)
	fmt.Printf("%d\n", mod)
}
