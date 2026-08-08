func hasDuplicate(nums []int) bool {
    num := make(map[int]bool)

    for i := range nums{
        if(num[nums[i]] == true){
            return true
        }
         num[nums[i]] = true
    }
    return false
}
