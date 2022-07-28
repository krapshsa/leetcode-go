package Climbing_stairs

func climbStairs(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}

	n1 := 1
	n2 := 2
	for i := 3; i < n; i++ {
		tmp := n2
		n2 = n1 + n2
		n1 = tmp
	}

	return n1 + n2
}
