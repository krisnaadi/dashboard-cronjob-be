package translator

import (
	"fmt"
	"reflect"

	"github.com/go-playground/validator"
)

func TranslateError(err error, s interface{}) map[string]string {
	apiErrors := make(map[string]string)

	// Mapping translator errors
	if castedObject, ok := err.(validator.ValidationErrors); ok {
		for _, err := range castedObject {
			switch err.Tag() {
			case "required":
				apiErrors[err.Field()] = fmt.Sprintf("%s tidak boleh kosong", getTagName(reflect.TypeOf(s), err.Field()))
			case "email":
				apiErrors[err.Field()] = fmt.Sprintf("%s tidak valid dengan format email", getTagName(reflect.TypeOf(s), err.Field()))
			case "gte":
				apiErrors[err.Field()] = fmt.Sprintf("Nilai %s harus lebih besar atau sama dengan %s", getTagName(reflect.TypeOf(s), err.Field()), err.Param())
			case "lte":
				apiErrors[err.Field()] = fmt.Sprintf("Nilai %s harus lebih kecil atau sama dengan %s", getTagName(reflect.TypeOf(s), err.Field()), err.Param())
			case "date":
				apiErrors[err.Field()] = fmt.Sprintf("%s tidak valid dengan format YYYY-MM-DD", getTagName(reflect.TypeOf(s), err.Field()))
			case "oneof":
				apiErrors[err.Field()] = fmt.Sprintf("%s tidak dikenali", getTagName(reflect.TypeOf(s), err.Field()))
			}
		}
	}

	return apiErrors
}

func getTagName(t reflect.Type, fieldName string) string {
	field, _ := t.FieldByName(fieldName)
	name := field.Tag.Get("name")

	if name == "" {
		name = fieldName
	}

	return name
}
