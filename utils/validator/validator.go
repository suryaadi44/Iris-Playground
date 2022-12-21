package validator

import (
	"reflect"
	"strings"
	"suryaadi44/iris-playground/utils/response"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	en_trans "github.com/go-playground/validator/v10/translations/en"

	validation "github.com/go-playground/validator/v10"
)

type Validator interface {
	ValidateJSON(s interface{}) *response.ErrorValues
}

type CustomValidator struct {
	Validate *validation.Validate
	trans    ut.Translator
}

func NewValidator() Validator {
	validate := validation.New()
	english := en.New()
	uni := ut.New(english, english)
	trans, _ := uni.GetTranslator("en")
	_ = en_trans.RegisterDefaultTranslations(validate, trans)

	customValidator := &CustomValidator{
		Validate: validate,
		trans:    trans,
	}

	// customValidator.addTranslation("required_if", "{0} is required when {1}")

	return customValidator
}

func (c *CustomValidator) ValidateStruct(s interface{}) *response.ErrorValues {
	err := c.Validate.Struct(s)
	if err == nil {
		return nil
	}

	var errors response.ErrorValues
	for _, err := range err.(validation.ValidationErrors) {
		errors = append(errors, *response.NewErrorValue(
			err.Field(),
			err.Translate(c.trans),
		))
	}

	return &errors
}

func (c *CustomValidator) ValidateJSON(s interface{}) *response.ErrorValues {
	c.Validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	return c.ValidateStruct(s)
}

func (c *CustomValidator) addTranslation(tag string, message string) {
	registerFn := func(ut ut.Translator) error {
		return ut.Add(tag, message, true)
	}

	transFn := func(ut ut.Translator, fe validation.FieldError) string {
		param := fe.Param()
		tag := fe.Tag()

		t, err := ut.T(tag, fe.Field(), param)
		if err != nil {
			return fe.(error).Error()
		}
		return t
	}

	_ = c.Validate.RegisterTranslation(tag, c.trans, registerFn, transFn)
}
