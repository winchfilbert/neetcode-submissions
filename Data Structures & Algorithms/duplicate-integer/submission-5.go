func hasDuplicate(nums []int) bool {
    hmap := make(map[int]bool)
	for i, n := range nums {
		if _, exist := hmap[nums[i]]; exist{
			return true
		}
		hmap[n]= false
	}
	return false
}


