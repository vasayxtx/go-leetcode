package problem39

import (
	"reflect"
	"testing"
)

func Test_combinationSum(t *testing.T) {
	type args struct {
		candidates []int
		target     int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "Test #1",
			args: args{
				candidates: nil,
				target:     0,
			},
			want: nil,
		},
		{
			name: "Test #2",
			args: args{
				candidates: []int{1},
				target:     1,
			},
			want: [][]int{{1}},
		},
		{
			name: "Test #3",
			args: args{
				candidates: []int{1, 2},
				target:     2,
			},
			want: [][]int{{2}, {1, 1}},
		},
		{
			name: "Test #3",
			args: args{
				candidates: []int{2, 3, 6, 7},
				target:     7,
			},
			want: [][]int{{7}, {2, 2, 3}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := combinationSum(tt.args.candidates, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("combinationSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
