package validator

import (
	"bytes"
	"fmt"
	"strings"
	"unicode"

	"github.com/go-playground/validator/v10"
)

// Validator - go-playground/validator wrapper.
type Validator struct {
	Validator *validator.Validate
}

func (v Validator) Validate(i interface{}) (err error) {
	err = v.Validator.Struct(i)

	return
}

func ValidationErrors(err error) (message map[string]interface{}) {
	message = make(map[string]interface{})

	for _, e := range err.(validator.ValidationErrors) {
		switch e.Tag() {
		case "datetime":
			message[strings.ToLower(convertCase(e.Field(), '_'))] = fmt.Sprintf("Field %s must be date & time", convertCase(e.Field(), ' '))
		case "email":
			message[strings.ToLower(convertCase(e.Field(), '_'))] = "Input must be valid email address"
		case "max":
			message[strings.ToLower(convertCase(e.Field(), '_'))] = fmt.Sprintf("Field %s must be less than %s", convertCase(e.Field(), ' '), e.Param())
		case "min":
			message[strings.ToLower(convertCase(e.Field(), '_'))] = fmt.Sprintf("Field %s must be more than %s", convertCase(e.Field(), ' '), e.Param())
		case "required":
			message[strings.ToLower(convertCase(e.Field(), '_'))] = fmt.Sprintf("Field %s can not empty!", convertCase(e.Field(), ' '))
		}
	}

	return
}

func convertCase(t string, c rune) string {
	buf := &bytes.Buffer{}

	for i, r := range t {
		if unicode.IsUpper(r) && i > 0 {
			buf.WriteRune(c)
		}

		buf.WriteRune(r)
	}

	return buf.String()
}
