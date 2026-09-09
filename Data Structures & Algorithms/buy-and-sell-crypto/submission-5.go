// kaya two pointers, but this is technically brute force
func maxProfit(prices []int) int {
	max := 0 // this is supposed to be mapped for profit


	for i := range prices{
		left := i
		right := len(prices)-1
			for left < right {
				profit := prices[right]-prices[left]
				if (max < profit){
					max = profit
				}
				right-- // why only right--, cuz the left should be the one we've checked
			}
	}
	return max
}
