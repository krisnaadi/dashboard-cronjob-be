package converter

import "testing"

func TestBoolToInt(t *testing.T) {
	type args struct {
		value bool
	}
	tests := []struct {
		name string
		args args
		want int8
	}{
		{
			name: "case convert true to 1",
			args: args{
				value: true,
			},
			want: 1,
		},
		{
			name: "case convert false to 0",
			args: args{
				value: false,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BoolToInt(tt.args.value); got != tt.want {
				t.Errorf("CurrencyFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}
