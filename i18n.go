package xiren

import (
	`strings`

	`github.com/go-playground/validator/v10`
	`github.com/storezhang/gox`
)

var _ = Localization

// Localization 取得本地化错误信息
func Localization(lang string, errs validator.ValidationErrors) (translations validator.ValidationErrorsTranslations) {
	translations = getTranslations(lang, errs)
	// 得到的国际化字符串是一个带请求体的键值，类似于LoginReq.Password：错误消息
	// 而我们需要的是password: 错误消息
	newI18n := make(map[string]string, len(translations))
	for field, msg := range translations {
		newField := gox.InitialLowercase(gox.CamelName(field[strings.IndexRune(field, dot)+1:]))
		newI18n[newField] = msg
		// 删除原来的错误消息，避免前端混乱
		delete(translations, field)
	}
	translations = newI18n

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
	if translate, found := translator.GetTranslator(`zh`); found {
		translations = errs.Translate(translate)
	}

	return
}
