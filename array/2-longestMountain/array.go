// 2.案例: 数组中的最长山脉
// 写出一个方法，对于给定的整数数组，数组中的每一个数字代表山的高度，返回数组中最长山峰的长度

// 山峰在数组中被定义为：峰顶左侧的值，严格单调增；右侧的值，严格单调减。一个山峰至少要由三个数字组成

// 输入: [1, 2, 3, 3, 4, 0, 10, 6, 5, -1, -3, 2, 3]
// 输出: 6

package array

// 本题要求我们找到一个山峰peak。
// 对于peak而言，向左和向右都是递减的。
// 反过来说，找到两个序列，序列1从左往右递增，序列2从右往左递增，两个序列相交的点就是peak。
// 我们只需要找到所有peak中，两个序列长度和最长的那个。

// 我们可以通过添加辅助数组的方式，来简化解法。

// 创建一个辅助数组1，从左往右计算每个位置递增序列的长度
// 创建一个辅助数组2，从右往左计算每个位置递增序列的长度
// 在每个位置，将数组1的值和数组2的值相加，和最大的位置即为peak的位置
// 返回最大的和

// 分析: 时间复杂度O(n)，空间复杂度O(n)
// 大部分题目都可以通过添加有限个辅助数组来简化难度，这暗示了大部分数组题的空间消耗等价于输入数组的大小。
// 所以一般空间复杂度不超过O(n)。因此对于数组题的优化，大部分只要找到O(1)空间复杂度的解法。

func longestMountain(arr []int) int {
	leftPeek := leftToRight(arr)
	rightPeek := rightToLeft(arr)
	longest := 0
	for i, al := 0, len(arr); i < al; i++ {
		left := leftPeek[i]
		right := rightPeek[i]
		length := left + right + 1 // 补1，加上从最低的距离
		if left == 0 || right == 0 {
			continue
		}
		if longest < length {
			longest = length
		}
	}
	return longest
}

// 输入: [1, 2, 3, 3, 4, 0, 10, 6, 5, -1, -3, 2, 3]
// 从左往右计算，计算单调递增的长度集合，如果是递减则为0 => [0,1,2,0,1,0,1,0,0,0,0,1,2]
func leftToRight(arr []int) []int {
	ans := make([]int, len(arr))
	for i, al := 1, len(arr); i < al; i++ {
		if arr[i] > arr[i-1] {
			ans[i] = ans[i-1] + 1
		}
	}
	return ans
}

// 输入: [1, 2, 3, 3, 4, 0, 10, 6, 5, -1, -3, 2, 3]
// 从右往左计算，计算单调递减的长度集合，如果是递增则为0 => [0,0,0,0,1,0,4,3,2,1,0,0,0]
func rightToLeft(arr []int) []int {
	ans := make([]int, len(arr))
	for i := len(arr) - 2; i >= 0; i-- {
		if arr[i] > arr[i+1] {
			ans[i] = ans[i+1] + 1
		}
	}
	return ans
}
