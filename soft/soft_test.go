package soft

import "testing"

func TestSoft_SoftTest(t *testing.T) {
	arr := []int{12, 34, 5, 9, 79, 65, 23, 11}
	soft := &Soft{}
	soft.SoftTest(arr)
}
