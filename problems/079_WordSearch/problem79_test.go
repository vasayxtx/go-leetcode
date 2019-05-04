package problem79

import "testing"

func Test_exist(t *testing.T) {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}

	type args struct {
		board [][]byte
		word  string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test #1",
			args: args{board, ""},
			want: true,
		},
		{
			name: "Test #2",
			args: args{nil, "A"},
			want: false,
		},
		{
			name: "Test #3",
			args: args{board, "ABCCED"},
			want: true,
		},
		{
			name: "Test #4",
			args: args{board, "SEE"},
			want: true,
		},
		{
			name: "Test #5",
			args: args{board, "ABCB"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := exist(tt.args.board, tt.args.word); got != tt.want {
				t.Errorf("exist() = %v, want %v", got, tt.want)
			}
		})
	}
}
