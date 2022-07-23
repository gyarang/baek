package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var cnt int
	var sum int64
	cntArr := make([]int, 8002)
	maxCnt := 4001

	fmt.Fscanln(reader, &cnt)
	arr := make([]int, cnt)
	for i := 0; i < cnt; i++ {
		var input int
		fmt.Fscanln(reader, &input)
		arr[i] = input
		sum += int64(input)
		cntArr[input+4000] += 1
		if cntArr[input+4000] > cntArr[maxCnt+4000] {
			maxCnt = input
		} else if cntArr[input+4000] == cntArr[maxCnt+4000] && input < maxCnt {
			maxCnt = input
		}
	}
	sort.Ints(arr)

	for i := maxCnt + 4001; i < 8001; i++ {
		if cntArr[i] == cntArr[maxCnt+4000] {
			maxCnt = i - 4000
			break
		}
	}

	ave := sum / int64(cnt)
	remain := sum % int64(cnt)
	if remain < 0 {
		remain *= -1
	}
	if remain > int64(cnt)/2 {
		if ave < 0 {
			ave--
		} else {
			ave++
		}
	}

	fmt.Fprintln(writer, ave)               // 평균
	fmt.Fprintln(writer, arr[(cnt/2)])      // 중앙값
	fmt.Fprintln(writer, maxCnt)            // 최빈값 (뒤에서 두번째)
	fmt.Fprintln(writer, arr[cnt-1]-arr[0]) // 범위
}
