package merge_two_2d_array

import "sort"

func mergeArrays(nums1 [][]int, nums2 [][]int) [][]int {
	m := make(map[int]int)

	for i := 0; i < len(nums1); i++ {
		num := nums1[i]
		key := num[0]
		value := num[1]

		m[key] += value
	}
	for i := 0; i < len(nums2); i++ {
		num := nums2[i]
		key := num[0]
		value := num[1]

		m[key] += value
	}
	result := make([][]int, 0)
	for i, v := range m {
		result = append(result, []int{i, v})
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i][0] < result[j][0]
	})

	return result
}
