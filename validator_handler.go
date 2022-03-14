package xiren

import (
	`github.com/go-playground/validator/v10`
)

type validatorHandler func(fl validator.FieldLevel) bool
