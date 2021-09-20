// 2.案例: 数组中的最长山脉
// 写出一个方法，对于给定的整数数组，返回数组中最长山峰的长度

// 山峰在数组中被定义为：峰顶左侧的值，严格单调增；右侧的值，严格单调减。一个山峰至少要由三个数字组成

// 输入: [1, 2, 3, 3, 4, 0, 10, 6, 5, -1, -3, 2, 3]
// 输出: 6

package array

import "testing"

func TestLongestMountain(t *testing.T) {
	cases := []struct {
		arr []int
		res int
	}{
		{
			arr: []int{1, 2, 3, 3, 4, 0, 10, 6, 5, -1, -3, 2, 3},
			res: 6,
		},
		{
			arr: []int{1, 1, 1},
			res: 0,
		},
		{
			arr: []int{2, 1, 4, 7, 3, 2, 5},
			res: 5,
		},
	}

	for _, cc := range cases {
		t.Run("TestLongestMountain", func(t *testing.T) {
			res := longestMountain(cc.arr)
			if cc.res != res {
				t.Errorf("failed test case arr: %v, want %v but got %v", cc.arr, cc.res, res)
			}
			res = longestMountain2(cc.arr)
			if cc.res != res {
				t.Errorf("failed2 test case arr: %v, want %v but got %v", cc.arr, cc.res, res)
			}
			res = longestMountain3(cc.arr)
			if cc.res != res {
				t.Errorf("failed3 test case arr: %v, want %v but got %v", cc.arr, cc.res, res)
			}
		})
	}
}
