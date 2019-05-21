package problem91

import "testing"

func Test_numDecodings(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test #1",
			args: args{""},
			want: 0,
		},
		{
			name: "Test #2",
			args: args{"1"},
			want: 1,
		},
		{
			name: "Test #3",
			args: args{"11"},
			want: 2,
		},
		{
			name: "Test #4",
			args: args{"21"},
			want: 2,
		},
		{
			name: "Test #5",
			args: args{"27"},
			want: 1,
		},
		{
			name: "Test #6",
			args: args{"31"},
			want: 1,
		},
		{
			name: "Test #7",
			args: args{"111"},
			want: 3,
		},
		{
			name: "Test #8",
			args: args{"1111"},
			want: 5,
		},
		{
			name: "Test #9",
			args: args{"226"},
			want: 3,
		},
		{
			name: "Test #10",
			args: args{"0"},
			want: 0,
		},
		{
			name: "Test #10",
			args: args{"10"},
			want: 1,
		},
		{
			name: "Test #11",
			args: args{"012"},
			want: 0,
		},
		{
			name: "Test #12",
			args: args{"110"},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numDecodings(tt.args.s); got != tt.want {
				t.Errorf("numDecodings() = %v, want %v", got, tt.want)
			}
		})
	}
}
