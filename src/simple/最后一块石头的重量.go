package simple

import "sort"

func lastStoneWeight(stones []int) int {
	n := len(stones)
	if n == 1 {
		return stones[0]
	}

	sort.Ints(stones)
	n -= 1
	for n > 0 {
		left := stones[n] - stones[n-1]
		if left == 0 {
			// 两个石头重量相等，全部粉碎
			stones[n-1] = 0
			n -= 2
		} else {
			i := n - 1
			stones[i] = left
			// 两个石头不相等，把粉碎后的石头放在升序排序的数组中
			for i > 0 && stones[i] < stones[i-1] {
				stones[i], stones[i-1] = stones[i-1], stones[i]
				i--
			}
			n--
		}
	}
	return stones[0]
}
