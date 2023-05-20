package runningsumof1darray

func runningSum(nums []int) []int {
	r := make([]int, len(nums))

	var t int
	for i, v := range nums {
		t += v
		r[i] = t
	}

	return r
}
