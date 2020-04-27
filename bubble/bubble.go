package main

import (
	"github.com/algorithms/util"
)

func main() {
	arr := util.CreateSlice(1000, 10000)
	arr1 := make([]int, len(arr))
	copy(arr1, arr)
	test(arr,arr)
	//arr:=util.CreateSlice(10,20)
	//fmt.Println(arr)
	//bubble1(arr)
	//fmt.Println(arr)
	//arr := []int{12, 34, 5, 9, 79, 65, 23, 11}
	//sli:=make([]int,len(arr))
	//copy(sli,arr)
	////test(arr, arr)
	////bubble(arr)
	//bubble1(arr)
	//fmt.Println(arr)
	//fmt.Println(sli)
}

func test(arr, arr1 []int) {
	util.UseTime("bubble", bubble, arr)
	util.UseTime("bubble1", bubble1, arr1)
}

func bubble(arr []int) {
	length := len(arr)
	for index := 0; index < length-1; index++ {
		for j := 0; j < length-index-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j] ^= arr[j+1] //a=a^b
				arr[j+1] ^= arr[j] //b=a^b^b =a
				arr[j] ^= arr[j+1] //a=a^b^a =b
			}
		}
	}
}
func bubble1(arr []int) {
	length := len(arr)

	for i := 0; i < length-1; i++ {
		flag := true
		for j := 0; j < length-i-1; j++ {
			flag = false
			if arr[j] > arr[j+1] {
				arr[j] ^= arr[j+1]
				arr[j+1] ^= arr[j]
				arr[j] ^= arr[j+1]
			}
		}
		if flag {
			break
		}
	}
}
