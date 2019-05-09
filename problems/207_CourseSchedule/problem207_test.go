package problem207

import "testing"

func Test_canFinish(t *testing.T) {
	type args struct {
		numCourses    int
		prerequisites [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test #1",
			args: args{0, [][]int{}},
			want: true,
		},
		{
			name: "Test #2",
			args: args{1, [][]int{}},
			want: true,
		},
		{
			name: "Test #3",
			args: args{2, [][]int{{1, 0}}},
			want: true,
		},
		{
			name: "Test #4",
			args: args{2, [][]int{{1, 0}, {0, 1}}},
			want: false,
		},
		{
			name: "Test #5",
			args: args{6, [][]int{{0, 1}, {0, 3}, {3, 1}, {1, 5}, {5, 2}, {4, 5}}},
			want: true,
		},
		{
			name: "Test #6",
			args: args{6, [][]int{{0, 1}, {0, 3}, {3, 1}, {1, 5}, {5, 2}, {4, 5}, {2, 0}}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canFinish(tt.args.numCourses, tt.args.prerequisites); got != tt.want {
				t.Errorf("canFinish() = %v, want %v", got, tt.want)
			}
		})
	}
}
