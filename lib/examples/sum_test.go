package examples

import "testing"

func TestSum(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Case 1",
			args: args{
				a: 1,
				b: 10,
			},
			want: 11,
		},
		{
			name: "Case 2",
			args: args{
				a: -3,
				b: 10,
			},
			want: 7,
		},
		{
			name: "Case 3",
			args: args{
				a: 0,
				b: 10,
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sum(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}
