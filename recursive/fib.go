package recursive

func fib(n int) int {
	if n <= 2 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

func fib1(n int) int {
	if n <= 2 {
		return 1
	}
	array := make([]int, n+1)
	array[1], array[2] = 1, 1
	return fibArray(n, &array)
}

func fib2(n int) int {
	if n <= 2 {
		return 1
	}
	array := make([]int, n+1)
	array[1], array[2] = 1, 1
	for index := 3; index <= n; index++ {
		array[index] = array[index-1] + array[index-2]
	}
	return array[n]
}

func fib3(n int) int {
	if n <= 2 {
		return 1
	}
	rollArray := make([]int, 2, 2)

	for i := 0; i < 2; i++ {
		rollArray[i] = 1
	}

	for index := 3; index <= n; index++ { //index &(len(rollArray-1）len(rollArray)是2的n次幂时才能使用
		rollArray[index&1] = rollArray[(index-1)&1] + rollArray[(index-2)&1]
	}
	return rollArray[n&1]
}

func fib4(n int) int {
	if n <= 2 {
		return 1
	}

	first, second := 1, 1
	for i := 3; i <=n ; i++ {
		second=first+second
		first=second-first
	}
	return second
}

func fibArray(n int, array *[]int) int {
	if (*array)[n] == 0 && n > 2 {
		(*array)[n] = fibArray(n-1, array) + fibArray(n-2, array)
	}
	return (*array)[n]
}

func oneNumber(array *[]int) int {
	res := 0
	for _, number := range *array {
		res ^= number
	}
	return res
}
