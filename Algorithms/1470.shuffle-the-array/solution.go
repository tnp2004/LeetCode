package shufflethearray

func shuffle(nums []int, n int) []int {
	var a []int

	for i := 0; i < n; i++ {
		a = append(a, nums[i], nums[i+n])
	}
	return a
}
