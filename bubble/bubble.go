package main

import "fmt"

func main() {
	arr := []int{12, 34, 5, 9, 79, 65, 23, 11}
	bubble(arr)
	fmt.Println(arr)
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
