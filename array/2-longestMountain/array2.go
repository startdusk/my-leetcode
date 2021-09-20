// 解法1通过添加辅助数组的方式，简化了算法的难度，但是消耗了额外的空间。
// 实际上我们可以只消耗O(1)的空间。
// 我们使用三个指针start, end, peak来指向每个山峰的起点、终点和峰值。
// 通过比较三个指针的位置来判断当前山峰是否符合要求。

// 假设我们有输入数组nums = [2, 1, 2, 3, 2, 1, 0, 2, 3, 1, 1]，我们可以画成下图。
// 在遍历数组的过程中，对于每个山峰，我们记下起点start、终点end和峰值peak的位置。
package array

func longestMountain2(arr []int) int {
	ans, end, length := 0, 0, len(arr)
	for end < length {
		start := end // 记录起点
		for end+1 < length && arr[end] < arr[end+1] {
			end++
		}

		peek := end // 记录山顶
		for end+1 < length && arr[end] > arr[end+1] {
			end++
		}

		if start < peek && peek < end { // 判断是不是一座山
			l := 1 + (end - start)
			if ans < l {
				ans = l
			}
		}

		if start == end { // 碰到平坡
			end++
		}
	}
	return ans
}
