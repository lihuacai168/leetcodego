package simple

//给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。
//
// 如果你最多只允许完成一笔交易（即买入和卖出一支股票一次），设计一个算法来计算你所能获取的最大利润。
//
// 注意：你不能在买入股票前卖出股票。
//
//
//
// 示例 1:
//
// 输入: [7,1,5,3,6,4]
//输出: 5
//解释: 在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5 。
//     注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格；同时，你不能在买入前卖出股票。
//
//
// 示例 2:
//
// 输入: [7,6,4,3,1]
//输出: 0
//解释: 在这种情况下, 没有交易完成, 所以最大利润为 0。
//
// Related Topics 数组 动态规划
// 👍 1349 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
// 暴力
//func maxProfit(prices []int) int {
//	n := len(prices)
//	begin := 0
//	ans := 0
//	for begin < n {
//		cur := begin + 1
//		for cur < n{
//			ans = max(ans, prices[cur] - prices[begin])
//			cur ++
//		}
//		begin ++
//	}
//	return ans
//}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// 官方解题思路
//func maxProfit(prices []int) int {
//	minprice := int(^uint(0) >> 1)
//	ans := 0
//	n := len(prices)
//	for i := 0; i < n; i++ {
//		ans = max(prices[i]-minprice, ans)
//		minprice = min(prices[i], minprice)
//	}
//	return ans
//}

// 动态规划
//定义状态转移方程
//给定转移方程初始值
//写代码递推实现转移方程
func maxProfit(prices []int) int {

	n := len(prices)
	if n < 1 {
		return 0
	}
	// buy 初始化买入的价钱
	// sell 初始化卖出的价钱
	buy, sell := -prices[0], 0
	for i := 0; i < n; i++ {
		// 卖出
		sell = max(sell, buy+prices[i])
		// 买入
		buy = max(buy, -prices[i])
	}
	return sell
}
