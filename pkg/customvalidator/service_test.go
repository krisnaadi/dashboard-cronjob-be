package customvalidator

import (
	"testing"

	"github.com/go-playground/validator"
	"github.com/stretchr/testify/assert"
)

func TestValidateDateFormat(t *testing.T) {

	type args struct {
		date string
	}

	validate := validator.New()
	validate.RegisterValidation("date", ValidateDateFormat)

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "case valid format",
			args: args{
				date: "2022-01-01",
			},
			want: true,
		},
		{
			name: "case invalid day",
			args: args{
				date: "2022-01-32",
			},
			want: false,
		},
		{
			name: "case invalid month",
			args: args{
				date: "2022-13-01",
			},
			want: false,
		},
		{
			name: "case invalid leap year",
			args: args{
				date: "2022-02-29",
			},
			want: false,
		},
		{
			name: "case invalid format",
			args: args{
				date: "2022/02/01",
			},
			want: false,
		},
		{
			name: "case empty",
			args: args{
				date: "",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validate.Var(tt.args.date, "date")
			if tt.want {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
				assert.Equal(t, err.Error(), "Key: '' Error:Field validation for '' failed on the 'date' tag")
			}
		})
	}
}
