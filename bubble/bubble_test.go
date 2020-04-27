package bubble

import (
	"fmt"
	"sort"
	"testing"
)

func TestBubble_Sort(t *testing.T) {
	arr := []int{12, 34, 5, 9, 79, 65, 23, 11}
	bubble := NewBubble(arr)
	bubble.Sort()
	if sort.IntsAreSorted(bubble.Array) {
		t.Log("success")
	} else {
		t.Error("failed")
	}
	fmt.Println(bubble.Array)
}

func TestBubble1_Sort(t *testing.T) {
	arr := []int{12, 34, 5, 9, 79, 65, 23, 11}
	bubble := NewBubble1(arr)
	bubble.Sort()
	if sort.IntsAreSorted(bubble.Array) {
		t.Log("success")
	} else {
		t.Error("failed")
	}
	fmt.Println(bubble.Array)

}
func TestBubble2_Sort(t *testing.T) {
	arr := []int{12, 34, 5, 9, 79, 65, 23, 11}
	bubble := NewBubble2(arr)
	bubble.Sort()
	if sort.IntsAreSorted(bubble.Array) {
		t.Log("success")
	} else {
		t.Error("failed")
	}
	fmt.Println(bubble.Array)
}
