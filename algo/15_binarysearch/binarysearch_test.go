package _5_binarysearch

import "testing"

func TestBinarySearch(t *testing.T) {
	a := []int{1, 3, 5, 6, 8}
	finds := []int{0, 1, 4, 5, 8, 10}
	for _, v := range finds {
		t.Logf("find %d index %d", v, BinarySearch(a, v))
	}
}

func TestBinarySearchRecursive(t *testing.T) {
	a := []int{1, 3, 5, 6, 8}
	finds := []int{0, 1, 4, 5, 8, 10}
	for _, v := range finds {
		t.Logf("find %d index %d", v, BinarySearchRecursive(a, v))
	}
}

func TestBinarySearchFirst(t *testing.T) {
	a := []int{1, 2, 2, 2, 4, 8}

	finds := []int{0, 1, 2, 4, 5, 10}
	for _, v := range finds {
		t.Logf("find %d index %d", v, BinarySearchFirst(a, v))
	}
}

func TestBinarySearchLast(t *testing.T) {
	a := []int{1, 2, 2, 2, 3, 4}
	finds := []int{0, 1, 2, 4, 5, 10}
	for _, v := range finds {
		t.Logf("find %d index %d", v, BinarySearchLast(a, v))
	}
}

func TestBinarySearchFirstGT(t *testing.T) {
	a := []int{1, 2, 2, 2, 3, 4}

	finds := []int{0, 1, 2, 4, 5, 10}
	for _, v := range finds {
		t.Logf("find %d index %d", v, BinarySearchFirstGT(a, v))
	}
}

func TestBinarySearchLastLT(t *testing.T) {
	a := []int{1, 2, 2, 2, 3, 4}
	finds := []int{0, 1, 2, 4, 5, 10}
	for _, v := range finds {
		t.Logf("find %d index %d", v, BinarySearchLastLT(a, v))
	}

}

func TestSearch(t *testing.T) {
	a := []int{4, 5, 6, 7, 0, 1, 2}
	finds := []int{-1, 3, 4, 5, 6, 7, 0, 1, 2, 9}
	for _, v := range finds {
		t.Logf("find %d index %d", v, search(a, v))
	}
}
