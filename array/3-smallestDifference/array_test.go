package array

import "testing"

func TestSmallestDifference(t *testing.T) {
	cases := []struct {
		arr1 []int
		arr2 []int
		res  int
	}{
		{
			arr1: []int{1, 3, 15, 11, 2},
			arr2: []int{23, 127, 235, 19, 8},
			res:  3,
		},
	}

	for _, cc := range cases {
		t.Run("TestSmallestDifference", func(t *testing.T) {
			res := smallestDifference(cc.arr1, cc.arr2)
			if cc.res != res {
				t.Errorf("failed test case arr1: %v arr2: %v, want %v but got %v", cc.arr1, cc.arr2, cc.res, res)
			}
		})
	}
}
