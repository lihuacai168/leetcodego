package simple

func removeDuplicates(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return n
	}
	// [0,0,0,1,1,1,2,2,3,3,4],
	// i=0, cur=3
	// [0,1,0,0,1,1,2,2,3,3,4],
	i := 0
	for j := 1; j < n; j++ {
		if nums[j] != nums[i] {
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1
}
