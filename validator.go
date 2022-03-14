package xiren

import (
	`github.com/go-playground/validator/v10`
	`github.com/goexl/baozheng`
)

func initValidator(validate *validator.Validate) (err error) {
	if err = validate.RegisterValidation(`mobile`, func(fl validator.FieldLevel) bool {
		return baozheng.Mobile(fl.Field().String())
	}); nil != err {
		return
	}

	if err = validate.RegisterValidation(`password`, func(fl validator.FieldLevel) bool {
		return baozheng.Password(fl.Field().String())
	}); nil != err {
		return
	}

	if err = validate.RegisterValidation(`filename`, func(fl validator.FieldLevel) bool {
		return baozheng.Filename(fl.Field().String())
	}); nil != err {
		return
	}

	if err = validate.RegisterValidation(`start_with_alpha`, func(fl validator.FieldLevel) bool {
		return baozheng.StartWithAlpha(fl.Field().String())
	}); nil != err {
		return
	}

	if err = validate.RegisterValidation(`prefix_or_suffix_space`, func(fl validator.FieldLevel) bool {
		return baozheng.PrefixOrSuffixSpace(fl.Field().String())
	}); nil != err {
		return
	}

	return
}
