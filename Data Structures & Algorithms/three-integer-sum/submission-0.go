import "slices"

func threeSum(nums []int) [][]int {
    slices.Sort(nums) // ts is important, as to use two pointers, it needed to be sort out first, i've got tricked

    sets := make([][]int, 0) // Fixed allocation size to avoid empty leading slots
    check_duplicates := make(map[string]bool, len(nums))

    // check per 3 numbers
    for check := 0 ; check < len(nums) - 2; check++{
        left := check + 1
        right := len(nums) - 1 // Start right at the very end

        for left < right {
            sum := nums[check] + nums[left] + nums[right]

            if sum == 0 {
                set_to_append := []int{nums[check], nums[left], nums[right]}
                sorted_set_for_key_map := GetFingerprint(set_to_append)
                
                if !slice_exists(sorted_set_for_key_map, check_duplicates) {
                    sets = append(sets, set_to_append)
                    check_duplicates[sorted_set_for_key_map] = true
                }
                left++
                right--
            } else if sum < 0 {
                left++
            } else {
                right--
            }
        }
    }

    return sets
}

func slice_exists(x string, check_duplicates map[string]bool)bool{
    return check_duplicates[x]
}

// GetFingerprint sorts the slice and turns it into a unique, comparable string key.
// Example: []int{3, 1, 2} -> "1,2,3"
func GetFingerprint(slice []int) string {
    copied := slices.Clone(slice)
    slices.Sort(copied)

    strVals := make([]string, len(copied))
    for i, v := range copied {
        strVals[i] = strconv.Itoa(v)
    }
    return strings.Join(strVals, ",")
}