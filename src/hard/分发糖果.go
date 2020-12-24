package hard

func candy(ratings []int) int {
	n := len(ratings)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	pre := 0
	ans := 0
	for i := 1; i < n; i++ {
		if ratings[i-1] < ratings[i] {
			pre = 1
		} else {
			pre = 2
		}
		ans += pre
	}

	if ratings[n-1] == ratings[n-2] {
		ans += 1
	} else {
		ans += 2
	}

	return ans

}
