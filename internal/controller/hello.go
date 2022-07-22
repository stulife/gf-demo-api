package controller

import (
	"context"
	"gf-demo-api/api/v1"
)

var (
	Hello = cHello{}
)

type cHello struct{}

func (c *cHello) Hello(ctx context.Context, req *v1.HelloReq) (res *v1.HelloRes, err error) {
	//err = gerror.NewCode(gcode.New(utility.CodeInternalError.Code(), utility.CodeInternalError.Desc(), nil))
	//return
	//panic("我是错误处理语句")
	res = &v1.HelloRes{
		Key: "idKeyC",
		Img: "base64stringC",
	}
	return
}
