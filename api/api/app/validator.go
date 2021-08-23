package app

import (
	"github.com/go-playground/validator/v10"
)

func GetValidateError(data interface{}) (errBool bool, errMsg string) {
	validate := validator.New()
	errValidate := validate.Struct(data)
	if errValidate != nil {
		for _, err := range errValidate.(validator.ValidationErrors) {
			if err != nil {
				switch err.Tag() {
				case "required":
					return false, err.Field() + " 是必填的"
				case "len":
					return false, err.Field() + " 长度等于" + err.Param()
				case "max":
					return false, err.Field() + " 最大是 " + err.Param()
				case "min":
					return false, err.Field() + " 最小是 " + err.Param()
				case "gt":
					return false, err.Field() + " 必须大于 " + err.Param()
				case "eq":
					return false, err.Field() + " 必须等于" + err.Param()
				case "gte":
					return false, err.Field() + " 必须大于等于 " + err.Param()
				case "lt":
					return false, err.Field() + " 必须小于 " + err.Param()
				case "lte":
					return false, err.Field() + " 必须小于等于" + err.Param()
				}
			}
		}
	}
	return true, ""
}
