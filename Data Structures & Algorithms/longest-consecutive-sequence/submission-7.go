
func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
        return 0
    }

	sort.Ints(nums)

	longest := 1
	currentstreak := 1
	// take 1, to directly take from the element before
	for i := 1 ; i < len(nums); i++{
		if check := nums[i-1] + 1; check == nums[i]{
			currentstreak += 1
		} else if nums[i] == nums[i-1]{
			currentstreak += 0 // add 0, not initialize 0, because it's the same
		} else{
			// must check in here, or will miss by one number if the check at the top of if
			if currentstreak > longest{
				longest = currentstreak
			}
			currentstreak = 1
		}
	}

	// CRUCIAL: Check one last time in case the longest streak hit the end of the array
    if currentstreak > longest {
        longest = currentstreak
    }

	return longest
}

