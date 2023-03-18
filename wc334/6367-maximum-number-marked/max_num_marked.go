package maximum_number_marked

import "sort"

type myNum struct {
	value  int
	half   int
	marked bool
}

func maxNumOfMarkedIndices(nums []int) int {
	n := len(nums)
	sort.Ints(nums)
	ans := 0
	for i, j := 0, n-1; i < j; {
		if 2*nums[i] <= nums[j] {
			ans++
			i++
			j--
		} else {
			j--
		}
	}
	return ans
}

//func maxNumOfMarkedIndices(nums []int) int {
//	l := len(nums)
//	myNums := make([]myNum, l)
//
//	// init
//	for i := 0; i < l; i++ {
//		myNums[i] = myNum{nums[i], 0, false}
//	}
//	sort.Slice(myNums, func(i, j int) bool {
//		return myNums[i].value < myNums[j].value
//	})
//
//	for i := 0; i < l; i++ {
//		myNums[i].half = myNums[i].value / 2
//	}
//
//	for j := l - 1; j > 0; j-- {
//		if myNums[j].marked {
//			continue
//		}
//
//		// find first myNums[j].half >= myNums[i]
//		for i := j - 1; i >= 0; i-- {
//			if myNums[i].marked {
//				continue
//			}
//
//			if myNums[j].half >= myNums[i].value {
//				myNums[i].marked = true
//				myNums[j].marked = true
//				break
//			}
//		}
//	}
//
//	// count marked
//	cnt := 0
//	for i := 0; i < l; i++ {
//		if myNums[i].marked {
//			cnt++
//		}
//	}
//
//	return cnt
//}
