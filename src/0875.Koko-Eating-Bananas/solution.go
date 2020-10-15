package Solution

/*
[小浩算法:二分查找](https://www.geekxh.com/1.9.%E4%BA%8C%E5%88%86%E6%B3%95%E7%B3%BB%E5%88%97/901.html#_03%E3%80%81%E9%A2%98%E7%9B%AE%E5%88%86%E6%9E%90)
*/
func MinEatingSpeed1(piles []int, H int) int {
	//根据(https://leetcode-cn.com/problems/koko-eating-bananas/solution/chao-jian-dan-fang-fa-da-dao-shuang-bai-100-by-pen/)
	sum := 0
	max := 0
	for _, v := range piles {
		sum += v
		if v > max {
			max = v
		}
	}
	a, b := (sum-1)/H+1, max
	// a, b := 1, int(math.Pow10(9))
	for a < b {
		mid := (a + b) >> 1
		if canEat(piles, H, mid) {
			b = mid
		} else {
			a = mid + 1
		}
	}
	return a
}

//返回true表示能吃完
func canEat(piles []int, H, speed int) bool {
	sum := 0
	for _, v := range piles {
		sum += (v-1)/speed + 1
	}
	return sum <= H
}

//超出时间限制,不能算答案
func MinEatingSpeed2(piles []int, H int) int {
	k := 1
	for {
		h := 0
		for _, v := range piles {
			h += (v-1)/k + 1
		}
		if h <= H {
			return k
		}
		k++
	}
}
