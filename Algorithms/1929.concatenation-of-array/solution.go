package concatenationofarray

func getConcatenation(nums []int) []int {
	nlen := len(nums)
	arr := make([]int, nlen*2)

	for i := range nums {
		arr[i] = nums[i]
		arr[i+nlen] = nums[i]
	}

	return arr
}
