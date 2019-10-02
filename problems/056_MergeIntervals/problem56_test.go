package problem56

import (
	"reflect"
	"testing"
)

func Test_merge(t *testing.T) {
	type args struct {
		intervals [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "Test #1",
			args: args{intervals: nil},
			want: nil,
		},
		{
			name: "Test #2",
			args: args{intervals: [][]int{{1, 10}}},
			want: [][]int{{1, 10}},
		},
		{
			name: "Test #3",
			args: args{intervals: [][]int{{11, 20}, {1, 10}}},
			want: [][]int{{1, 10}, {11, 20}},
		},
		{
			name: "Test #4",
			args: args{intervals: [][]int{{1, 10}, {2, 9}}},
			want: [][]int{{1, 10}},
		},
		{
			name: "Test #5",
			args: args{intervals: [][]int{{1, 10}, {2, 20}}},
			want: [][]int{{1, 20}},
		},
		{
			name: "Test #6",
			args: args{intervals: [][]int{{1, 4}, {4, 5}}},
			want: [][]int{{1, 5}},
		},
		{
			name: "Test #7",
			args: args{intervals: [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}},
			want: [][]int{{1, 6}, {8, 10}, {15, 18}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := merge(tt.args.intervals); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("merge() = %v, want %v", got, tt.want)
			}
		})
	}
}
