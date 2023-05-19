package searchinsertposition

func searchInsert(nums []int, target int) int {
	for i, v := range nums {
		if v == target {
			return i
		} else if v > target {
			return i
		}
	}

	return len(nums)
}
