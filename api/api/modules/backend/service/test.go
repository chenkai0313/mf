package service

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"

	"api/api/api_errors"
	"api/api/internal/render"
	"api/api/modules/backend/form/request"
	"api/api/modules/backend/form/response"
	"api/app"
	"api/client/hello_client"
	"api/client/proto/hello"
)

type TestService struct {
}

func (t TestService) Test(c *gin.Context, req request.TestReq) render.Response {
	resp := response.TestResp{}
	cli := hello_client.HelloClient{}
	data := hello.SayHelloReq{
		Content: req.Content,
	}
	res, err := cli.SayHello(c, &data)
	result, _ := json.Marshal(res)
	app.ZapLog.Info("Test", string(result[:]))
	if err != nil {
		app.ZapLog.Error("Test", err.Error())
		return api_errors.InternalServerError("请求失败")
	}
	if res.Code != http.StatusOK {
		return api_errors.InternalServerError(res.Msg)
	}
	resp.Data = res.Data
	return render.Success(resp)
}
