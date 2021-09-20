// 3.案例: 最小差
// 给定两个数组，找到一对数字（两个数组各取一个），他们的差值最小。返回这两个数组。

// 输入:
// arrayOne = [-1, 5, 10, 20, 28, 3]
// arrayTwo = [26, 134, 135, 15, 17]
// 输出: 2

// 思路分析
// 我们需要在两个数组里，各取一个数字。这明显暗示了我们需要使用Two Pointers技巧。
// 本题需要用的Two Pointers技巧的一个变种，根据使用方法通常称为“谁小移谁”（滑稽）。

// 将两个数组分别排序
// 使用两根指针，分别指向数组第一位，同时开始遍历
// 比较两根指针指向的值，如果二者差值小于当前最小值，那么更新最小值
// 移动二者之间较小的那个
// 直到任意一个数组遍历结束

// 分析
// 时间复杂度O(n logn)，空间复杂度O(1)
// 由于使用了内置的排序方法，所以消耗了O(n logn)的时间复杂度

package array

import (
	"math"
	"sort"
)

func smallestDifference(arr1, arr2 []int) int {
	sort.Ints(arr1)
	sort.Ints(arr2)

	diff := math.MaxInt32
	for i, j := 0, 0; i < len(arr1) && j < len(arr2); {
		if int(math.Abs(float64(arr1[i]-arr2[j]))) < diff {
			diff = int(math.Abs(float64(arr1[i] - arr2[j]))) // 更新最小值
		} else if arr1[i] < arr2[j] { // 移动两根指针当中较小的那个
			i++
		} else {
			j++
		}
	}

	return diff
}
