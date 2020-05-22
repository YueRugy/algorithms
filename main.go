package main

import "fmt"

type ms []int

type HashSet struct {
	array map[string]struct{}
}

func NewHashSet() *HashSet {
	return &HashSet{
		array: make(map[string]struct{}, 16),
	}
}

func (hs *HashSet) Add(str string) {
	hs.array[str] = struct{}{}
}

func main() {

	list := make([]int, 10)
	for index := 0; index < 10; index++ {
		list[index] = index
	}
	foo(list)
	fmt.Println(list)
	//fmt.Println(1/2)
	//fmt.Println(math.Pow(3.0, 3.0))
	//hs:=NewHashSet()
	//hs.Add("a")
	//hs.Add("b")
	//fmt.Println(len(hs.array))
	//_,ok:=hs.array["a"]
	//fmt.Println(ok)
	//_,ok1:=hs.array["dd"]
	//fmt.Println(ok1)
}

func foo(list []int) {
	list[6]=1000
}
