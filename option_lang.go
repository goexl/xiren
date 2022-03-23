package xiren

var (
	_        = Lang
	_ option = (*optionLang)(nil)
)

type optionLang struct {
	lang string
}

// Lang 语言
func Lang(lang string) *optionLang {
	return &optionLang{
		lang: lang,
	}
}

func (l *optionLang) apply(options *options) {
	options.lang = l.lang
}
