package main

import (
	"fmt"
	"math/big"
)

func main() {
	var a, b big.Int
	fmt.Scanf("%d\n%d", &a, &b)

	add := new(big.Int)
	add = add.Add(&a, &b)

	sub := new(big.Int)
	sub = sub.Sub(&a, &b)

	mul := new(big.Int)
	mul = mul.Mul(&a, &b)

	fmt.Printf("%d\n", add)
	fmt.Printf("%d\n", sub)
	fmt.Printf("%d\n", mul)
}
