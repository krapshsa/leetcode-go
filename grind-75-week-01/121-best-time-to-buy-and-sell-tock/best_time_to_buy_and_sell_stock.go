package best_time_to_buy_and_sell_stock

func maxProfit(prices []int) int {
	l := len(prices)

	if 0 == l {
		return 0
	}

	max := 0
	low := prices[0]
	for i := 1; i < l; i++ {
		// find the lowest price on the left-hand side, may be a good time to buy
		if prices[i] < low {
			low = prices[i]
			continue
		}

		profits := prices[i] - low
		if profits > max {
			max = profits
		}
	}

	return max
}
