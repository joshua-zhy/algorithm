package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"time"
)

// 给你一个数组 nums 和一个值 val，你需要 原地 移除所有数值等于 val 的元素，并返回移除后数组的新长度。

// 不要使用额外的数组空间，你必须仅使用 O(1) 额外空间并原地修改输入数组。

// 元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。

// 示例 1: 给定 nums = [3,2,2,3], val = 3, 函数应该返回新的长度 2, 并且 nums 中的前两个元素均为 2。 你不需要考虑数组中超出新长度后面的元素。

const (
	maxArrLen = 10
	minValue  = 1
	maxValue  = 100
	testTime  = 10
)

func removeElement(nums []int, removeNum int) []int {
	L := len(nums)
	if L < 1 {
		return nil
	}

	slow := 0
	for fast := 0; fast < L; fast++ {
		if nums[fast] != removeNum {
			nums[slow] = nums[fast]
			slow++//移动指针
		}
	}
	return nums[:slow]
}

func randNums(maxArrLen, minValue, maxValue int) []int {

	rand.Seed(time.Now().UnixNano())

	nums := make([]int, 0)

	nums = append(nums, minValue)
	for i := 0; i < maxArrLen; i++ {
		num := rand.Intn(maxValue)

		nums = append(nums, num)
	}

	return nums
}

func compare() error {
	nums := randNums(maxArrLen, minValue, maxValue)

	randRemoveIndex := rand.Intn(maxArrLen - 1)

	randRemoveNum := nums[randRemoveIndex]
	res := removeElement(nums, randRemoveNum)

	for _, v := range res {
		if v == randRemoveNum {
			return errors.New("fuck")
		}
	}
	return nil
}

func main() {
	for i := 0; i < testTime; i++ {
		if err := compare(); err != nil {
			log.Fatalln(err)
		}
	}
	fmt.Println("ok")
}
