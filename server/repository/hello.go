package repository

import (
	"context"

	"server/proto/hello"
)

type HelloRepositories interface {
	SayHello(ctx context.Context, req *hello.SayHelloReq, resp *hello.SayHelloResp) error
}

type HelloRepository struct {
}

func (h HelloRepository) SayHello(ctx context.Context, req *hello.SayHelloReq, resp *hello.SayHelloResp) error {
	reply := "hello world " + req.Content
	resp.Data = reply
	return nil
}
