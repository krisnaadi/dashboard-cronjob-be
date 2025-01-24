package translator

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/tech-djoin/wallet-djoin-service/internal/constant"
	"github.com/tech-djoin/wallet-djoin-service/internal/pkg/customvalidator"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func CustomValidaton(validation map[string]func(fl validator.FieldLevel) bool) echo.Validator {
	v := validator.New()

	for name, function := range validation {
		v.RegisterValidation(name, function)
	}

	return &CustomValidator{validator: v}
}

func ValidateDateFormat(fl validator.FieldLevel) bool {
	dateStr := fl.Field().String()

	if dateStr == "" || fl.Field().IsZero() {
		return true
	}

	_, err := time.Parse(constant.DateFormatYYYYMMDD, dateStr)
	return err == nil
}

func TestTranslateError(t *testing.T) {
	type args struct {
		err error
		s   interface{}
	}

	type User struct {
		Username string `query:"username" name:"Username" validate:"required"`
	}

	type Member struct {
		Email string `query:"email" name:"Email" validate:"required,email"`
	}

	type Date struct {
		Date string `query:"date" name:"Date" validate:"date"`
	}

	type Exam struct {
		Grade int64 `query:"grade" validate:"gte=1,lte=10"`
	}

	type Gift struct {
		Tipe string `query:"tipe" name:"Tipe" validate:"required,oneof=buku buka"`
	}

	req := httptest.NewRequest(http.MethodGet, "/path", nil)
	rec := httptest.NewRecorder()

	// Create a new Echo instance
	e := echo.New()
	// Create a new Echo context using echo.NewContext
	c := e.NewContext(req, rec)

	//add cusstom validation
	customValidationMap := make(map[string]func(fl validator.FieldLevel) bool)
	customValidationMap["date"] = customvalidator.ValidateDateFormat

	customValidator := customvalidator.CustomValidaton(customValidationMap)
	e.Validator = customValidator

	tests := []struct {
		name string
		args args
		want map[string]string
	}{
		{
			name: "Case success mapping required validation error",
			args: args{
				err: c.Validate(User{}),
				s:   User{},
			},
			want: map[string]string{
				"Username": "Username tidak boleh kosong",
			},
		},
		{
			name: "Case success mapping gte validation error",
			args: args{
				err: c.Validate(Exam{Grade: 0}),
				s:   Exam{},
			},
			want: map[string]string{
				"Grade": "Nilai Grade harus lebih besar atau sama dengan 1",
			},
		},
		{
			name: "Case success mapping lte validation error",
			args: args{
				err: c.Validate(Exam{Grade: 11}),
				s:   Exam{},
			},
			want: map[string]string{
				"Grade": "Nilai Grade harus lebih kecil atau sama dengan 10",
			},
		},
		{
			name: "Case success mapping email validation error",
			args: args{
				err: c.Validate(Member{Email: "fake-email"}),
				s:   Member{},
			},
			want: map[string]string{
				"Email": "Email tidak valid dengan format email",
			},
		},
		{
			name: "Case success mapping date validation error",
			args: args{
				err: c.Validate(Date{Date: "05-05-2023"}),
				s:   Date{},
			},
			want: map[string]string{
				"Date": "Date tidak valid dengan format YYYY-MM-DD",
			},
		},
		{
			name: "Case success mapping oneof validation error",
			args: args{
				err: c.Validate(Gift{Tipe: "mainan"}),
				s:   Gift{},
			},
			want: map[string]string{
				"Tipe": "Tipe tidak dikenali",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TranslateError(tt.args.err, tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TranslateError() = %v, want %v", got, tt.want)
			}
		})
	}
}
