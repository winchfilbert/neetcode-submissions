func twoSum(nums []int, target int) []int {
	check := make(map[int]int)

	for i,n := range nums{
		remainder := target-nums[i]
		if _, exist := check[remainder]; exist{
			return []int{check[remainder], i}
		}
		check[n] = i
	}

	return nil

}
