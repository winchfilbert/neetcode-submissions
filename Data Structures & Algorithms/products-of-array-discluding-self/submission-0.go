func productExceptSelf(nums []int) []int {
    arr := make([]int, len(nums))

    for i := 0 ; i < len(nums); i++{
        prod := 1
        for j := 0 ; j < len(nums); j++{
            if i == j {
                continue
            }
            prod *= nums[j]
        }
        arr[i] = prod
    }
    return arr
}
