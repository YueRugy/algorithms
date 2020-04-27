package selection

func Selection(arr []int) {
	length := len(arr)

	for end := length - 1; end > 0; end-- {
		index := 0
		for begin := 0; begin <= end; begin++ {
			if arr[begin] > arr[index] {
				index = begin
			}
		}

		if index != end {
			arr[index] ^= arr[end]
			arr[end] ^= arr[index]
			arr[index] ^= arr[end]
		}

	}
}
