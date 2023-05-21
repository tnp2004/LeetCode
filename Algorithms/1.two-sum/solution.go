package twosum

func twoSum(nums []int, target int) []int {
	hm := map[int]int{}

	for i, v := range nums {
		_, ok := hm[v]
		if ok {
			return []int{i, hm[v]}
		}
		hm[target-v] = i
	}
	return nil
}
