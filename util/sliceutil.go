package util

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}
func CreateSlice(num, limit int) []int {
	if num <= 0 || limit <= 0 {
		return nil
	}
	if num > limit {
		return nil
	}

	sli := make([]int, num)
	for i := 0; i < len(sli); i++ {
		sli[i] = rand.Intn(limit)
	}
	return sli
}
