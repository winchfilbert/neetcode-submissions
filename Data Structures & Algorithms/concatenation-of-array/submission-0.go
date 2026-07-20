func getConcatenation(nums []int) []int {
    size := len(nums)
    ans := make([]int, size * 2)
    
    copy(ans[0:size], nums)
    copy(ans[size:], nums)
    
    return ans
}