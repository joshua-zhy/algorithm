package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"time"
)

// 给你一个按 非递减顺序 排序的整数数组 nums，返回 每个数字的平方 组成的新数组，要求也按 非递减顺序 排序。

// 示例 1： 输入：nums = [-4,-1,0,3,10] 输出：[0,1,9,16,100] 解释：平方后，数组变为 [16,1,0,9,100]，排序后，数组变为 [0,1,9,16,100]

// 示例 2： 输入：nums = [-7,-3,2,3,11] 输出：[4,9,9,49,121]

const (
	maxArrLen = 10
	minValue  = -10
	maxValue  = 100
	testTime  = 1000
)

func sortedSquares(nums []int) []int {
	L := len(nums)

	if L < 1 {
		return nil
	}

	l, r := 0, L-1 //原数组的左右指针

	newNums := make([]int, L) //新数组用来存在元素
	j := L - 1                //新数组的右指针 放最大的数开始

	for i := 0; i < L; i++ {
		ln, rn := nums[l]*nums[l], nums[r]*nums[r]
		if ln < rn {
			newNums[j] = rn
			r--
			j--
		} else {
			newNums[j] = ln
			l++
			j--
		}

	}

	return newNums
}

func randNums(maxArrLen, minValue, maxValue int) []int {

	rand.Seed(time.Now().UnixNano())

	nums := make([]int, 0)

	nums = append(nums, rand.Intn(maxArrLen+1)+minValue) //Intn 方法将返回一个非负伪随机数，其值范围在半开区间 [0, n)。如果传入的 n <= 0，将会 panic
	for i := 1; i < maxArrLen; i++ {
		num := rand.Intn(maxValue)

		for num <= nums[i-1] {
			num = rand.Intn(maxValue)
		}
		nums = append(nums, num)
	}
	return nums
}

// 对数器 for test 
func compare() error {
	nums := randNums(maxArrLen, minValue, maxValue)

	nums = sortedSquares(nums)

	// fmt.Println(nums)'
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
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
