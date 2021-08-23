package api_errors

import "api/api/internal/render"

func ErrorNotFound() render.Response {
	errCode := NOT_FOUND
	errMsg := "NOT FOUND"
	return render.Error(errCode, errMsg)
}

func InternalServerError(errorMsg ...string) render.Response {
	errCode := INTERNAL_SERVER_ERROR
	errMsg := "SYSTEM_BUSY"
	if len(errorMsg) != 0 {
		errMsg = ""
		for _, v := range errorMsg {
			errMsg += v + "\n"
		}
		errMsg = errorMsg[0]
	}
	return render.Error(errCode, errMsg)
}

func InvalidRequestParams(errorMsg ...string) render.Response {
	errCode := PARAMAS_INVALIDATE_ERROR
	noticeErrMsg := "request params error"
	if len(errorMsg) != 0 {
		noticeErrMsg = ""
		for _, v := range errorMsg {
			noticeErrMsg += v + "\n"
		}
		noticeErrMsg = errorMsg[0]
	}
	return render.Error(errCode, noticeErrMsg)
}

func ErrorTooManyRequests() render.Response {
	errCode := TOO_MANY_REQUESTS
	errMsg := "request too frequently"
	return render.Error(errCode, errMsg)
}

func ErrorNotAuthorization() render.Response {
	errCode := Not_Authorization
	errMsg := "api requires authorization"
	return render.Error(errCode, errMsg)
}

func ErrorNotSubmitConnectInfo() render.Response {
	errCode := Not_Submit_Connect
	errMsg := "未提交联系人信息或审核未通过"
	return render.Error(errCode, errMsg)
}

func ErrorNotCompanyAdminRole() render.Response {
	errCode := Not_Company_Admin
	errMsg := "当前用户不是该公司的管理员"
	return render.Error(errCode, errMsg)
}
