package recursive

import "fmt"

func hanoi(n int, p1, p2, p3 string) {
	if n < 1 {
		return
	}
	if n == 1 {
		move(n, p1, p3)
		return
	}
	hanoi(n-1,p1,p3,p2)
	move(n,p1,p3)
	hanoi(n-1,p2,p1,p3)
}



func move(n int, from, to string) {
	fmt.Printf("%d号盘子从%s挪到%s\n", n, from, to)
}
