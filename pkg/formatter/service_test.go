package formatter

import "testing"

func TestCurrencyFormat(t *testing.T) {
	type args struct {
		amount float64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case format currency 1000 into 1.000",
			args: args{
				amount: 1000,
			},
			want: "1.000",
		},
		{
			name: "case format currency 1000.01 into 1.000,01",
			args: args{
				amount: 1000.01,
			},
			want: "1.000,01",
		},
		{
			name: "case format currency 100 into 100",
			args: args{
				amount: 100,
			},
			want: "100",
		},
		{
			name: "case format currency 1234567 into 1.234.567",
			args: args{
				amount: 1234567,
			},
			want: "1.234.567",
		},
		{
			name: "case format currency 1234567.89 into 1.234.567,89",
			args: args{
				amount: 1234567.89,
			},
			want: "1.234.567,89",
		},
		{
			name: "case format currency 100.009 into 100,01",
			args: args{
				amount: 100.009,
			},
			want: "100,01",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CurrencyFormat(tt.args.amount); got != tt.want {
				t.Errorf("CurrencyFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}
