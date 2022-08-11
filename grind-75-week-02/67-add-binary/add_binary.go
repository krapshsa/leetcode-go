package add_binary

import (
	"math"
	"strconv"
	"strings"
)

func reverseBits(lenMax int, bits []int) string {
	ans := make([]string, 0, lenMax+1)
	var j int
	for j = lenMax; j >= 0; j-- {
		if bits[j] == 1 {
			break
		}
	}
	for i := j; i >= 0; i-- {
		ans = append(ans, strconv.Itoa(bits[i]))
	}
	if len(ans) == 0 {
		return "0"
	} else {
		return strings.Join(ans, "")
	}
}

func addBinary(a string, b string) string {
	lenA := len(a)
	lenB := len(b)
	lenMax := int(math.Max(float64(lenA), float64(lenB)))
	bits := make([]int, 0, lenMax+1)

	for i := 0; i <= lenMax; i++ {
		bits = append(bits, 0)
	}
	for i := lenA - 1; i >= 0; i-- {
		bit := int(rune(a[i]) - '0')
		bits[lenA-1-i] += bit
	}
	for i := lenB - 1; i >= 0; i-- {
		bit := int(rune(b[i]) - '0')
		bits[lenB-1-i] += bit
	}
	for i := 0; i < len(bits); i++ {
		if bits[i] > 1 {
			bits[i] -= 2
			bits[i+1]++
		}
	}

	return reverseBits(lenMax, bits)
}
