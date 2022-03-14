package xiren

import (
	`github.com/go-playground/locales/en`
	`github.com/go-playground/locales/zh`
	`github.com/go-playground/universal-translator`
	`github.com/go-playground/validator/v10`
	enLang `github.com/go-playground/validator/v10/translations/en`
	zhLang `github.com/go-playground/validator/v10/translations/zh`
)

var (
	_ = Struct
	_ = Var
	_ = New

	validate   *validator.Validate
	translator *ut.UniversalTranslator
)

// Struct 验证结构体
func Struct(obj interface{}) error {
	return validate.Struct(obj)
}

// Var 验证变量
func Var(field interface{}, tag string) error {
	return validate.Var(field, tag)
}

// New 创建新的验证器
func New() *validator.Validate {
	return validate
}

// 创建内置验证器
// 单例设计模式
func newValidate() (err error) {
	validate = validator.New()
	translator = ut.New(en.New(), zh.New())

	english, _ := translator.GetTranslator(`en`)
	if err = enLang.RegisterDefaultTranslations(validate, english); nil != err {
		return
	}
	chinese, _ := translator.GetTranslator(`zh`)
	if err = zhLang.RegisterDefaultTranslations(validate, chinese); nil != err {
		return
	}

	if err = initValidator(validate); nil != err {
		return
	}
	if err = initTranslation(validate, chinese); nil != err {
		return
	}

	return
}
