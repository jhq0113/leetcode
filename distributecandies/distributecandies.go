package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
给定一个偶数长度的数组，其中不同的数字代表着不同种类的糖果，每一个数字代表一个糖果。你需要把这些糖果平均分给一个弟弟和一个妹妹。返回妹妹可以获得的最大糖果的种类数。

示例 1:

输入: candies = [1,1,2,2,3,3]
输出: 3
解析: 一共有三种种类的糖果，每一种都有两个。
     最优分配方案：妹妹获得[1,2,3],弟弟也获得[1,2,3]。这样使妹妹获得糖果的种类数最多。
示例 2 :

输入: candies = [1,1,2,3]
输出: 2
解析: 妹妹获得糖果[2,3],弟弟获得糖果[1,1]，妹妹有两种不同的糖果，弟弟只有一种。这样使得妹妹可以获得的糖果种类数最多。
注意:

1.数组的长度为[2, 10,000]，并且确定为偶数。
2.数组中数字的大小在范围[-100,000, 100,000]内。
*/

func distributeCandies(candyType []int) int {
	if len(candyType) < 4 {
		return 1
	}

	halfLen := len(candyType)/2
	max := 0

	setValue := struct {}{}
	set := make(map[int16]struct{}, halfLen)
	for _,val := range candyType {
		if _, exists := set[int16(val)]; !exists {
			set[int16(val)] = setValue
			max++
			if max >= halfLen {
				return halfLen
			}
		}
	}

	return len(set)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	var (
		count int
		min = -10000
		max = 10000
		//保证偶数个数
		num = 2 + 2 * rand.Intn(4999)
		candyType = make([]int, num, num)
	)

	for index:= 0; index< num; index++ {
		//最大10000， 最小-10000
		candyType[ index ] = min + rand.Intn(2 * max)
	}

	start := time.Now()
	count = distributeCandies(candyType)
	fmt.Printf("共%d种，耗时%s\n", count, time.Now().Sub(start))
}
