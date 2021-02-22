package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
给定一个非负整数数组nums ，你最初位于数组的 第一个下标 。

数组中的每个元素代表你在该位置可以跳跃的最大长度。

判断你是否能够到达最后一个下标。

示例1：

输入：nums = [2,3,1,1,4]
输出：true
解释：可以先跳 1 步，从下标 0 到达下标 1, 然后再从下标 1 跳 3 步到达最后一个下标。
示例2：

输入：nums = [3,2,1,0,4]
输出：false
解释：无论怎样，总会到达下标为 3 的位置。但该下标的最大跳跃长度是 0 ， 所以永远不可能到达最后一个下标。
*/

func canJump(nums []int) bool {
	var (
		length     = len(nums)
		current    = 0
		currentMax = 0
		max        = length - 1
	)

	if length < 2 {
		return true
	}

	for index, val := range nums {
		current = index + val
		if current >= max {
			return true
		}

		if current > currentMax {
			currentMax = current
		}

		if currentMax <= index && index < max {
			return false
		} else if currentMax >= max {
			return true
		}
	}
	return false
}

func main() {
	rand.Seed(time.Now().UnixNano())
	var (
		max  = 1 + rand.Intn(30)
		nums = make([]int, max, max)
	)

	for index := 0; index < max; index++ {
		nums[index] = rand.Intn(5)
	}

	start := time.Now()
	can := canJump(nums)
	fmt.Printf("示例数据:%v，结果：%t，耗时：%s\n", nums, can, time.Now().Sub(start))
}
