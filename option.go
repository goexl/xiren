package xiren

type (
	option interface {
		apply(options *options)
	}

	options struct {
		lang string
	}
)

func defaultOptions() *options {
	return &options{
		lang: langZh,
	}
}
