package main

import "fmt"

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
	m:=make(map[int]int,20)
	fmt.Println(len(m))
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
