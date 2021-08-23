package app

import (
	"encoding/json"
	"errors"
	"reflect"

	"github.com/go-playground/validator/v10"
)

func BindParams(req interface{}, res interface{}) error {
	data, err := json.Marshal(req)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &res)
}

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

func BindRequestParamsValidate(req interface{}, data interface{}) error {
	if err := BindParams(req, data); err != nil {
		return err
	}
	if errValidateBool, errValidateMsg := GetValidateError(data); !errValidateBool {
		return errors.New(errValidateMsg)
	}
	return nil
}

func structAssign(binding interface{}, value interface{}) {
	bVal := reflect.ValueOf(binding).Elem() //获取reflect.Type类型
	vVal := reflect.ValueOf(value).Elem()   //获取reflect.Type类型
	vTypeOfT := vVal.Type()
	for i := 0; i < vVal.NumField(); i++ {
		// 在要修改的结构体中查询有数据结构体中相同属性的字段，有则修改其值
		name := vTypeOfT.Field(i).Name
		if ok := bVal.FieldByName(name).IsValid(); ok {
			bVal.FieldByName(name).Set(reflect.ValueOf(vVal.Field(i).Interface()))
		}
	}
}
