// 1.案例: 单调数列
// 判断数组是否具有单调性（单调递增或单调递减）

// 输入: [-1, -5, -10, -1100, -1100, -1101, -1102, -9001]
// 输出: true

// 输入: [1,3,2]
// 输出: false
package array

import "testing"

func TestIsMonotonic(t *testing.T) {
	cases := []struct {
		arr []int
		res bool
	}{
		{
			arr: []int{-1, -5, -10, -1100, -1100, -1101, -1102, -9001},
			res: true,
		},
		{
			arr: []int{1, 3, 2},
			res: false,
		},
		{
			arr: []int{4, 3, 2},
			res: true,
		},
	}

	for _, cc := range cases {
		t.Run("isMonotonic", func(t *testing.T) {
			res := isMonotonic(cc.arr)
			if cc.res != res {
				t.Errorf("failed test case arr: %v, want %v but got %v", cc.arr, cc.res, res)
			}
		})
	}
}
