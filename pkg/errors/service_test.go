package errors

import (
	golangErr "errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWrap(t *testing.T) {
	type args struct {
		cause error
	}
	tests := []struct {
		name string
		args args
		want ErrorWrapper
	}{
		{
			name: ";",
			args: args{
				cause: golangErr.New("test"),
			},
			want: &errorWrapper{
				err: golangErr.New("test"),
			},
		},
		{
			name: "_",
			args: args{
				cause: &errorWrapper{
					err: golangErr.New("test"),
				},
			},
			want: &errorWrapper{
				err: golangErr.New("test"),
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := Wrap(test.args.cause)
			assert.Equal(t, test.want, got)
		})
	}
}
