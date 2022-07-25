package first_bad_version

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 */

var badVersion int

func isBadVersion(version int) bool {
	return version >= badVersion
}

func firstBadVersion(n int) int {
	min := n
	left := 1
	right := n

	for left <= right {
		mid := (left + right) / 2
		if isBadVersion(mid) {
			min = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return min
}
