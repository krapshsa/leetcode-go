package minimum_operations

func minOperations(n int) int {
	zeroCnt := 0
	oneCnt := 0
	for n > 0 {
		m := n % 2
		if m == 0 {
			zeroCnt++
		} else {
			oneCnt++
		}
		n = n / 2
	}

	if zeroCnt >= oneCnt {
		return oneCnt
	} else {
		return zeroCnt + 2
	}
}
