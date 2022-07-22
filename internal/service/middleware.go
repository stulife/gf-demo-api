package service

import (
	"gf-demo-api/internal/model"
	utility "gf-demo-api/utility"
	"github.com/gogf/gf/v2/frame/g"

	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
)

type IMiddleware interface {
	Cors(r *ghttp.Request)
	Ctx(r *ghttp.Request)
}
type middlewareImpl struct{}

var middleService = middlewareImpl{}

func Middleware() IMiddleware {
	return &middleService
}

// Cors 跨域处理
func (s *middlewareImpl) Cors(r *ghttp.Request) {
	corsOptions := r.Response.DefaultCORSOptions()
	r.Response.CORS(corsOptions)
	r.Middleware.Next()
}

// Ctx 自定义上下文对象
func (s middlewareImpl) Ctx(r *ghttp.Request) {
	data, err := GfToken().ParseToken(r)
	if err != nil {
		failed := &utility.DefaultResponse{
			Code:    model.UnAuthorized.Code(),
			Message: model.UnAuthorized.Desc(),
		}
		r.Response.WriteJsonExit(failed)
	}
	if data != nil {
		context := new(model.Context)
		err = gconv.Struct(data.Data, &context.User)
		if err != nil {
			g.Log().Error(r.GetCtx(), err)
			failed := &utility.DefaultResponse{
				Code:    model.UnAuthorized.Code(),
				Message: model.UnAuthorized.Desc(),
			}
			r.Response.WriteJsonExit(failed)
		}
		Context().Init(r, context)
	}
	r.Middleware.Next()
}
