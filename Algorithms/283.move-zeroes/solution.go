package movezeroes

func moveZeroes(nums []int) {
	var z int

	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[z], nums[i] = nums[i], nums[z]
			z++
		}
	}
}
