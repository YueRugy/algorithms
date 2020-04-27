package util

import (
	"fmt"
	"time"
)

func UseTime(message string, f func(arr []int), arg []int) {

	before := time.Now()
	f(arg)
	after := time.Now()
	str := after.Sub(before).String()
	fmt.Println(message + "  use  " + str)

}
