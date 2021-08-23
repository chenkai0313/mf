package api_errors

import (
	"api/api/app"
	"api/api/internal/render"
)

func BindRequestParamsValidate(req interface{}) render.Response {
	if errValidateBool, errValidateMsg := app.GetValidateError(req); !errValidateBool {
		return InvalidRequestParams(errValidateMsg)
	}
	return render.Success(nil)
}
