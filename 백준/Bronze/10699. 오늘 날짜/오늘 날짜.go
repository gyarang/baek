package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	location, _ := time.LoadLocation("Asia/Seoul")
	nowInSeoul := now.In(location)

	fmt.Printf("%d-%02d-%02d", nowInSeoul.Year(), nowInSeoul.Month(), nowInSeoul.Day())
}
