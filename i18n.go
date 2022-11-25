package xiren

import (
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/goexl/gox"
	"github.com/goexl/gox/field"

	"github.com/goexl/exc"
)

var _ = Localization

// Localization 取得本地化错误信息
func Localization(lang string, errs validator.ValidationErrors) (err error) {
	translations := getTranslations(lang, errs)
	// 得到的国际化字符串是一个带请求体的键值，类似于LoginReq.Password：错误消息
	// 而我们需要的是password: 错误消息
	fields := make([]gox.Field[any], 0, len(translations))
	for _field, msg := range translations {
		key := gox.Case(_field[strings.IndexRune(_field, dot)+1:])
		fields = append(fields, field.New(key.Camel(gox.CasePositionHead).String(), msg))
	}
	err = exc.NewFields(exceptionValidate, fields...)

	return
}

func getTranslations(lang string, errs validator.ValidationErrors) (translations validator.ValidationErrorsTranslations) {
	splits := strings.Split(lang, separator)
	for index := 0; index < len(splits); index++ {
		if translate, found := translator.FindTranslator(lang); found {
			translations = errs.Translate(translate)
		} else if last := strings.LastIndex(lang, separator); -1 != last {
			lang = lang[:last]
		}

		if nil != translations {
			break
		}
	}

	if nil != translations {
		return
	}

	// 默认使用中文
	if translate, found := translator.GetTranslator(langZh); found {
		translations = errs.Translate(translate)
	}

	return
}
