package simple

func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	dp := make([]int, n+1)
	// 因为使用dp[i-1]和dp[i-2]来表示动态方程，所以n应该取等号，才能访问到cost的最大值也就是i-1
	// 因为n取等，导致dp[]的最大容量变成n+1,否则dp[i]就会超出最大长度
	for i := 2; i <= n; i++ {
		dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}
	return dp[n]
}
