package medium

//func rotate1(nums []int, k int) {
//	n := len(nums)
//	if k > n {
//		k = k % n
//	}
//	for i := 0; i < k; i++ {
//		for j := n - 1; j > 0; j-- {
//			nums[j], nums[j-1] = nums[j-1], nums[j]
//		}
//	}
//}

// 还是很慢
func rotate1(nums []int, k int) {
	n := len(nums)
	if k > n {
		k = k % n
	}
	for i := 0; i < k; i++ {
		nums[i], nums[n-(k-i)] = nums[n-(k-i)], nums[i]
	}
	mid := n - k*2
	if mid < 0 {
		mid = -mid
		for i := 0; i < mid; i++ {
			for j := n - 1; j > k; j-- {
				nums[j], nums[j-1] = nums[j-1], nums[j]
			}
		}
	} else {
		for i := 0; i < mid; i++ {
			for j := k; j < n-1; j++ {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}

}
