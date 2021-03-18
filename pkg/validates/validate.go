package validates

import (
	"errors"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
)

// https://github.com/go-playground/validator

// 验证函数，用于额外的校验
type ValidateFunc func(interface{}) error

var validate *validator.Validate

var trans ut.Translator

func init() {
	zh_ch := zh.New()
	uni := ut.New(zh_ch)               // 万能翻译器，保存所有的语言环境和翻译数据
	trans, _ = uni.GetTranslator("zh") // 翻译器
	validate = validator.New()
	_ = zh_translations.RegisterDefaultTranslations(validate, trans)
	// 添加额外翻译
	_ = validate.RegisterTranslation("required_without", trans, func(ut ut.Translator) error {
		return ut.Add("required_without", "{0} 为必填字段!", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required_without", fe.Field())
		return t
	})

	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := fld.Tag.Get("vName")
		if name == "-" {
			return ""
		}
		return name
	})
}

// 注册struct校验器
func RegisterStructValidation(fn validator.StructLevelFunc, types ...interface{}) {
	validate.RegisterStructValidation(fn, types)
}

// 验证Struct参数格式
func ValidateStruct(input interface{}, extraValidates ...ValidateFunc) error {
	err := validate.Struct(input)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		for _, e := range errs {
			// can translate each error one at a time.
			return errors.New(e.Translate(trans))
		}

	}
	for _, extraValidate := range extraValidates {
		err = extraValidate(input)
		if err != nil {
			return err
		}
	}
	return nil
}

// 验证参数格式
func ValidateVar(input interface{}, tag string, extraValidates ...ValidateFunc) error {
	err := validate.Var(input, tag)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		for _, e := range errs {
			// can translate each error one at a time.
			return errors.New(e.Translate(trans))
		}

	}
	for _, extraValidate := range extraValidates {
		err = extraValidate(input)
		if err != nil {
			return err
		}
	}
	return nil
}
