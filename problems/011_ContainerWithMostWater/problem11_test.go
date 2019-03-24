package problem11

import "testing"

func test(t *testing.T, heights []int, expRes int) {
	actRes := maxArea(heights)
	if actRes != expRes {
		t.Errorf("Fail for %v, %d != %d", heights, actRes, expRes)
	}
}

func TestMaxArea(t *testing.T) {
	test(t, []int{1, 2}, 1)
	test(t, []int{3, 2}, 2)
	test(t, []int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49)
	test(t, []int{2, 3, 4, 5, 18, 17, 6}, 17)
}
