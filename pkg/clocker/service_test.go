package clocker

import (
	"reflect"
	"testing"
	"time"
)

func TestNow(t *testing.T) {

	tests := []struct {
		name string
		want string
	}{
		{
			name: "case success",
			want: "Asia/Makassar",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Now(); !reflect.DeepEqual(got.Location().String(), tt.want) {
				t.Errorf("Now() = %v, want %v", got.Location().String(), tt.want)
			}
		})
	}
}

func TestParse(t *testing.T) {
	type args struct {
		oldTime time.Time
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "case success",
			args: args{
				time.Now(),
			},
			want: "Asia/Makassar",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Parse(tt.args.oldTime); !reflect.DeepEqual(got.Location().String(), tt.want) {
				t.Errorf("Parse() = %v, want %v", got.Location().String(), tt.want)
			}
		})
	}
}
