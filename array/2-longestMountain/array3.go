// 解法3
// 解法2使用了start, end, peak三根指针。
// 还有一种解法与之类似，但是使用了up和down两个状态，来代替三根指针。
// 通过比较up和down的值来判断当前状态。
package array

func longestMountain3(arr []int) int {
	ans, up, down := 0, 0, 0
	for i, al := 1, len(arr); i < al; i++ {
		if down > 0 && arr[i-1] < arr[i] || arr[i-1] == arr[i] { // 1.下坡结束
			up, down = 0, 0
		}

		if arr[i] > arr[i-1] { // 2.上坡
			up++
		}

		if arr[i] < arr[i-1] { // 3.下坡
			down++
		}

		if up > 0 && down > 0 {
			l := 1 + (up + down)
			if ans < l {
				ans = l
			}
		}
	}
	return ans
}
