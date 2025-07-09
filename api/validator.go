package api

import (
	"github.com/go-playground/validator/v10"
	"simplebank/utils"
)

var validCurrency validator.Func = func(fieldlevel validator.FieldLevel) bool {
	if currency, ok := fieldlevel.Field().Interface().(string); ok {
		// check currency is supported
		return utils.IsSupportedCurrency(currency)
	}
	return false
}
