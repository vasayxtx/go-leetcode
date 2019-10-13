package problem78

import (
	"reflect"
	"sort"
	"testing"
)

func sortSliceOfNums(s [][]int) {
	sort.Slice(s, func(i, j int) bool {
		nums1, nums2 := s[i], s[j]
		minLen := len(nums1)
		if len(nums2) < minLen {
			minLen = len(nums2)
		}
		for i := 0; i < minLen; i++ {
			if nums1[i] < nums2[i] {
				return true
			}
			if nums1[i] > nums2[i] {
				return false
			}
		}
		return len(nums1) == minLen
	})
}

func checkEqualSlicesOfNumsWithoutOrder(s1 [][]int, s2 [][]int) bool {
	sortSliceOfNums(s1)
	sortSliceOfNums(s2)
	return reflect.DeepEqual(s1, s2)
}

func Test_subsets(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "Test #1",
			args: args{[]int{}},
			want: [][]int{{}},
		},
		{
			name: "Test #2",
			args: args{[]int{1}},
			want: [][]int{{}, {1}},
		},
		{
			name: "Test #3",
			args: args{[]int{1, 2}},
			want: [][]int{{}, {1}, {2}, {1, 2}},
		},
		{
			name: "Test #4",
			args: args{[]int{1, 2, 3}},
			want: [][]int{{}, {1}, {2}, {3}, {1, 2}, {1, 3}, {2, 3}, {1, 2, 3}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subsets(tt.args.nums); !checkEqualSlicesOfNumsWithoutOrder(got, tt.want) {
				t.Errorf("subsets() = %v, want %v", got, tt.want)
			}
		})
	}
}
