package LinearSort

import "testing"

func TestBucketSort(t *testing.T) {
	a := make([]int, 0, 400)
	for i := -10; i <= 10; i++ {
		a = append(a, i)
		a = append(a, i)
		a = append(a, i)
	}
	BucketSort(a)
	t.Log(a)
}

func TestBucketSortSimple(t *testing.T) {
	a := []int{1, 6, 3, 5, 8, 6, 4}
	BucketSortSimple(a)
	t.Log(a)
}
