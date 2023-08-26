package main

import (
	"bufio"
	"fmt"
	"os"
)

func getResult(yourName string) [4]int {
	result := [4]int{0, 0, 0, 0}
	for _, v := range yourName {
		switch v {
		case 'L':
			result[0]++
		case 'O':
			result[1]++
		case 'V':
			result[2]++
		case 'E':
			result[3]++
		}
	}
	return result
}

func calc(myCnt, yourCnt [4]int) int {
	return (((myCnt[0] + yourCnt[0]) + (myCnt[1] + yourCnt[1])) *
		((myCnt[0] + yourCnt[0]) + (myCnt[2] + yourCnt[2])) *
		((myCnt[0] + yourCnt[0]) + (myCnt[3] + yourCnt[3])) *
		((myCnt[1] + yourCnt[1]) + (myCnt[2] + yourCnt[2])) *
		((myCnt[1] + yourCnt[1]) + (myCnt[3] + yourCnt[3])) *
		((myCnt[2] + yourCnt[2]) + (myCnt[3] + yourCnt[3]))) % 100
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var myName string
	var count int
	myCnt := [4]int{0, 0, 0, 0}

	max := 0
	maxTeam := "ㄱ"

	fmt.Fscanln(reader, &myName)
	for _, char := range myName {
		switch char {
		case 'L':
			myCnt[0]++
		case 'O':
			myCnt[1]++
		case 'V':
			myCnt[2]++
		case 'E':
			myCnt[3]++
		}
	}

	fmt.Fscanln(reader, &count)
	for i := 0; i < count; i++ {
		var name string
		fmt.Fscanln(reader, &name)
		yourCnt := getResult(name)
		result := calc(myCnt, yourCnt)
		if result > max {
			max = result
			maxTeam = name
			continue
		}
		if result == max {
			if name < maxTeam {
				maxTeam = name
			}
		}
	}

	fmt.Println(maxTeam)
}
