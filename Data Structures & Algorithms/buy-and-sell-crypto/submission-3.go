func maxProfit(prices []int) int {
	// days_visited := make(map[int]int)
	max := 0 // this is supposed to be mapped for profit


	for i := range prices{
		left := i
		right := len(prices)-1
			for left < right {
				profit := prices[right]-prices[left]
				if (max < profit){
					max = profit
				}
				right--
			}
	}
	return max
}
