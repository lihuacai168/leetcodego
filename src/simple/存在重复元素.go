package simple

func containsDuplicate(nums []int) bool {

	m := map[int]int{}
	for _, num := range nums {
		if _, ok := m[num]; !ok {
			m[num] = 1
		} else {
			return true
		}
	}
	return false
}
