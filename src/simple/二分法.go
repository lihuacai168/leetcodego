package simple

func bisect(arr []int, target int) (targetIndex int) {
	n := len(arr)
	if n == 0 {
		return -1
	}
	begin, end := 0, n
	for begin <= end {
		mid := (begin + end) / 2
		midVal := arr[mid]
		if midVal == target {
			return mid
		}
		if midVal < target {
			begin = mid + 1
		} else {
			end = mid - 1
		}
	}
	return -1
}
