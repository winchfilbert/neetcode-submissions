func twoSum(nums []int, target int) []int {
    hmap := make(map[int]int)

    for i, n := range nums{
        remainder := target-nums[i]
        if _, exist := hmap[remainder]; exist{ // 2. cek kalo udah ada remaindernya
              return []int{hmap[remainder],i}
        }
        hmap[n] = i //1. misal buat kasus nums: 5,5 target 10, artinya value-index ke 0=5: index-ke 0
    }
    return nil

}
