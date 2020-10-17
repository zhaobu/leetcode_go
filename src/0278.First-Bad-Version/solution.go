package Solution

import "fmt"

var (
	Versions []int
)

func isBadVersion(version int) bool {
	if version > len(Versions) {
		panic(fmt.Errorf("version=%d", version))
	}
	return Versions[version-1] == 1
}

/*
[二分查找 左闭右闭]
*/
func FirstBadVersion1(n int) int {
	i, j := 1, n
	for i <= j {
		mid := i + (j-i)>>1
		if isBadVersion(mid) {
			j = mid - 1
		} else {
			i = mid + 1
		}
	}
	return j + 1
}

/*
[二分查找 左闭右开]
*/
func FirstBadVersion2(n int) int {
	i, j := 1, n
	for i < j {
		mid := i + (j-i)>>1
		if isBadVersion(mid) {
			j = mid
		} else {
			i = mid + 1
		}
	}
	return i
}
