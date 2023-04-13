package main

import "fmt"

// 给定一个含有 n 个正整数的数组和一个正整数 s ，找出该数组中满足其和 ≥ s 的长度最小的 连续 子数组，并返回其长度。如果不存在符合条件的子数组，返回 0。

// 示例：

// 输入：s = 7, nums = [2,3,1,2,4,3] 输出：2 解释：子数组 [4,3] 是该条件下的长度最小的子数组。

func minSubArrayLen(nums []int, s int) int {
	start := 0      // 起始位置
	l := len(nums)  // 数组长度
	sum := 0        // 子数组之和
	result := l + 1 // 初始化返回长度为l+1，目的是为了判断“不存在符合条件的子数组，返回0”的情况

	for end := 0; end < l; end++ { //end 终止位置
		sum += nums[end]
		for sum >= s { //循环检测 子字符串的长度是不是最小
			subLength := end - start + 1 //找到的子字符串的长度

			if subLength < result {
				result = subLength //结果
			}
			sum -= nums[start]
			start++
		}
	}
	if result == l+1 {
		return 0
	} else {
		return result
	}
}

func main() {
	nums := []int{2, 3, 1, 2, 4, 3}
	res := minSubArrayLen(nums, 7)
	fmt.Println(res)
}
