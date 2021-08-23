package handler

import (
	"context"

	"server/app"
	"server/internal"
	"server/proto/hello"
	"server/repository"
	"server/validate"
)

type HelloHandler struct {
	HelloRepository repository.HelloRepositories
}

func (h HelloHandler) SayHello(ctx context.Context, req *hello.SayHelloReq, response *hello.SayHelloResp) error {
	var data validate.SayHelloReq
	if err := app.BindRequestParamsValidate(req, &data); err != nil {
		response.Code = internal.PARAMAS_INVALIDATE_ERROR
		response.Msg = err.Error()
		return nil
	}
	if err := h.HelloRepository.SayHello(ctx, req, response); err != nil {
		response.Code = internal.INTERNAL_SERVER_ERROR
		response.Msg = err.Error()
		return nil
	}
	response.Code = internal.Success
	response.Msg = "success"
	return nil
}
