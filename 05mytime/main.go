package main

import (
	"fmt"
	"time"
)

func main() {
	// loc, _ := time.LoadLocation("Asia/Dhaka")
	// presentTime := time.Now().In(loc)
	presentTime := time.Now()
	fmt.Println(presentTime)
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))
	fmt.Println(presentTime.Location())

	createDate := time.Date(1995, time.January, 10, 12, 15, 19, 0, time.Local)
	fmt.Println(createDate.Format("01-02-2006 15:04:05 Monday"))

}
