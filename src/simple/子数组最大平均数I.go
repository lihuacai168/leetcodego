package simple

//给定 n 个整数，找出平均数最大且长度为 k 的连续子数组，并输出该最大平均数。
//
//
//
//示例：
//
//输入：[1,12,-5,-6,50,3], k = 4
//输出：12.75
//解释：最大平均数 (12-5-6+50)/4 = 51/4 = 12.75
//
//来源：力扣（LeetCode）
//链接：https://leetcode-cn.com/problems/maximum-average-subarray-i
//著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

func findMaxAverage(nums []int, k int) float64 {
	sum := 0
	for _, v := range nums[:k] {
		sum += v
	}
	maxTotal := sum
	for i := k; i < len(nums); i++ {
		sum = sum + nums[i] - nums[i-k]
		maxTotal = max(maxTotal, sum)
	}
	return float64(maxTotal) / float64(k)
}
