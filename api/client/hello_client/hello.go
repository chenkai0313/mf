package hello_client

import (
	"github.com/gin-gonic/gin"

	"api/client"
	"api/client/proto/hello"
)


const detectService_name =  "srv.server"

type HelloClient struct{}

func (h HelloClient) SayHello(c *gin.Context, req *hello.SayHelloReq) (res *hello.SayHelloResp, err error) {
	cli := hello.NewHelloService(detectService_name, client.MicroService.Client())
	res, err = cli.SayHello(c, req)
	return
}
