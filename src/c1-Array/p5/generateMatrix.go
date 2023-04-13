package main

import "fmt"

/*
	给定一个正整数 n，生成一个包含 1 到 n^2 所有元素，且元素按顺时针顺序螺旋排列的正方形矩阵。
	输入: 3 输出: [ [ 1, 2, 3 ], [ 8, 9, 4 ], [ 7, 6, 5 ] ]
*/

func generateMatrix(n int) [][]int {
	//最终返回的二维数组
	matrix := make([][]int, n)

	//先给二维数组初始化一些空的一维数组，先把空填满，然后在空里进行赋值，就是这思路
	for i, _ := range matrix {
		matrix[i] = make([]int, n)
	}

	//初始化1,2,3,4 ...第一个数是从1开始的
	num := 1

	//左，右，上，下
	left, right, up, down := 0, n-1, 0, n-1

	for num <= n*n {
		//从左往右
		for i := left; i <= right; i++ {
			matrix[up][i] = num
			num++
		}
		up++ //下一条边 左边的第二个格子位置

		// fmt.Println(matrix)

		//从上往下
		for i := up; i <= down; i++ {
			matrix[i][right] = num
			num++
		}
		right--

		// if num == n*n {
		// 	return matrix
		// }

		// fmt.Println(matrix)

		//从右往左
		for i := right; i >= left; i-- {
			matrix[down][i] = num
			num++
		}

		down--
		// fmt.Println(matrix)

		//从下往上
		for i := down; i >= up; i-- {
			matrix[i][left] = num
			num++
		}

		left++
		// fmt.Println(matrix)
	}
	return matrix
}

func main() {
	fmt.Println(generateMatrix(3))
}
