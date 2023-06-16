package buildarrayfrompermutation

func buildArray(nums []int) []int {
	var newArr []int

	for i := range nums {
		newArr = append(newArr, nums[nums[i]])
	}

	return newArr
}
