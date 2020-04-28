package bs

func Index(arr []int, num int) int {
	if arr == nil || len(arr) == 0 {
		return -1
	}
	begin, end := 0, len(arr)
	for end > begin {
		mid := (begin + end) >> 1
		if arr[mid] == num {
			return mid
		}
		if arr[mid] > num {
			end = mid
		} else {
			begin = mid + 1
		}
	}
	return -1
}
