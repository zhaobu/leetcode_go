package _4_sorts

func CountingSort(a []int, n int) {
	if n <= 1 {
		return
	}

	var max int = a[0]
	for i := range a {
		if a[i] > max {
			max = a[i]
		}
	}

	//统计每个数的个数，c[i]表示a中i的个数
	c := make([]int, max+1)
	for i := range a {
		c[a[i]]++
	}

	//对数组c顺序求和，c[i]表示a中<=i的个数
	for i := 1; i <= max; i++ {
		c[i] += c[i-1]
	}

	r := make([]int, n)

	/*
		从后往前遍历数组a，c[a[i]]表示当前包括a[i]在内<=a[i]的个数,
		所以排序后a[i]在排好序的数组r中的位置应该为c[a[i]]-1
		同时应该把值<=a[i]的个数减掉一个
	*/
	for i := n - 1; i >= 0; i-- {
		index := c[a[i]] - 1
		r[index] = a[i]
		c[a[i]]--
	}

	copy(a, r)
}
