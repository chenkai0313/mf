package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"api/api/api_errors"
	"api/api/modules/backend/form/request"
	"api/api/modules/backend/service"
)

var testService service.TestService

func Test(c *gin.Context) {
	var req request.TestReq
	_ = c.ShouldBind(&req)

	if errRep := api_errors.BindRequestParamsValidate(req); errRep.IsFail() {
		c.JSON(http.StatusOK, errRep)
		return
	}

	res := testService.Test(c, req)
	c.JSON(http.StatusOK, res)
	return
}
