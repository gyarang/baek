package main

func sum(a []int) int {
	var r int
    for _, v := range a {
        r += v
    }
    return r
}
