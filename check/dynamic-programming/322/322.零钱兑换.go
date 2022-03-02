package main

import (
	"fmt"
	"sort"
)

/*
 * @lc app=leetcode.cn id=322 lang=golang
 *
 * [322] 零钱兑换
 */

// @lc code=start
func coinChange(coins []int, amount int) int {
	if len(coins) < 1 {
		return -1
	}
	if amount < 1 {
		return 0
	}
	//先排序coins
	sort.Ints(coins)
	i := len(coins) - 1
	count := 0
	for i >= 0 && amount > 0 {
		for coins[i] <= amount {
			amount -= coins[i]
			count++
			fmt.Printf("amount=%d,count=%d,coins[%d]=%+v. \n", amount, count, i, coins[i])
		}
		i--
	}
	fmt.Printf("amount=%d,coins=%+v. \n", amount, coins)
	if amount != 0 {
		return -1
	}
	return count
}

// @lc code=end
