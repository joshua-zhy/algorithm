package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

const (
	maxArrLen = 10
	minValue  = 1
	maxValue  = 100
	testTime  = 1000
)

// 给定一个 n 个元素有序的（升序）整型数组 nums 和一个目标值 target  ，写一个函数搜索 nums 中的 target，如果目标值存在返回下标，否则返回 -1。

// 示例 1:

// 输入: nums = [-1,0,3,5,9,12], target = 9
// 输出: 4
// 解释: 9 出现在 nums 中并且下标为 4

// 1.search 左闭右闭区间
func binarySearch(nums []int, target int) int {
	L := len(nums)

	if L < 1 {
		return -1
	}

	left := 0
	right := L - 1

	for left <= right {
		mid := left + ((right - left) >> 1)

		if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			return mid
		}
	}
	return -1

}

// 2.search
func binarySearch2(nums []int, target int) int {
	L := len(nums)

	if L < 1 {
		return -1
	}

	left := 0
	right := L

	for left < right {
		mid := left + ((right - left) >> 1)

		if nums[mid] > target {
			right = mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

func randNums(maxArrLen, minValue, maxValue int) []int {

	rand.Seed(time.Now().UnixNano())

	nums := make([]int, 0)

	nums = append(nums, minValue)
	for i := 1; i < maxArrLen; i++ {
		num := rand.Intn(maxValue)

		for num < nums[i-1] {
			num = rand.Intn(maxValue)
		}
		nums = append(nums, num)
	}
	return nums
}

// 对数器 for test
func compare() bool {
	nums := randNums(maxArrLen, minValue, maxValue)

	// res := binarySearch(nums, nums[index])

	randIndex := rand.Intn(maxArrLen)

	res := binarySearch(nums, nums[randIndex])

	// res := binarySearch2(nums, nums[randIndex])

	return nums[res] == nums[randIndex]

}

func main() {
	for i := 0; i < testTime; i++ {
		if res := compare(); !res {
			log.Fatalln("fuck")
		}
	}
	fmt.Println("ok")

}
