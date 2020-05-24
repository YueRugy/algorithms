package recursive

import (
	"fmt"
	"testing"
	"time"
)

const number = 45

func TestFib(t *testing.T) {
	fmt.Println(fib(6))
}

func TestFib1(t *testing.T) {
	fmt.Println(fib1(6))
}

func TestFib2(t *testing.T) {
	fmt.Println(fib2(6))
}

func TestFibTime(t *testing.T) {

	timePrint(fib, number)
	timePrint(fib1, number)
	timePrint(fib2, number)
	timePrint(fib3,number)
	timePrint(fib4,number)
}

func TestOneNumber(t *testing.T)  {
	arr:=&[]int{1,2,1,3,2,3,4}
	fmt.Println(oneNumber(arr))
}

func timePrint(f func(int) int, n int) {
	begin := time.Now()
	fmt.Println(f(n))
	println(time.Now().Sub(begin).String())
}

