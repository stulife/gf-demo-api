package controller

import (
	"context"
	"gf-demo-api/api/v1"
	"gf-demo-api/internal/model"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

var (
	Hello = cHello{}
)

type cHello struct{}

func (c *cHello) Hello(ctx context.Context, req *v1.HelloReq) (res *v1.HelloRes, err error) {
	err = gerror.NewCode(gcode.New(model.CodeBusinessFailed.Code(), model.CodeBusinessFailed.Desc(), nil))
	return
	//panic("我是错误处理语句")
	res = &v1.HelloRes{
		Name: "这是我的第一个go程序",
	}
	return
}
