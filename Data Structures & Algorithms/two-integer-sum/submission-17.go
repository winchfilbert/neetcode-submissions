func twoSum(nums []int, target int) []int {
	hashset := make(map[int]int)
	for i, n := range nums{ 
		remainder := target - n
		if _, exist := hashset[remainder]; exist{
			return []int{hashset[remainder], i}
		}
		hashset[n] = i
	} 
	return nil
}
