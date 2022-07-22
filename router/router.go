package router

import (
	"gf-demo-api/internal/controller"
	"gf-demo-api/internal/service"
	"gf-demo-api/utility"
	"github.com/gogf/gf/v2/net/ghttp"
)

func BindController(group *ghttp.RouterGroup) {

	group.Group("/api/v1", func(group *ghttp.RouterGroup) {
		group.Middleware(utility.HandlerResponse)
		//group.GET("/swagger", func(r *ghttp.Request) {
		//	r.Response.Write(consts.SwaggerUITemplate)
		//})
		group.Group("/", func(group *ghttp.RouterGroup) {
			group.Middleware(service.Middleware().Cors)

			service.GfToken().Middleware(group)
			//context拦截器
			group.Middleware(service.Middleware().Ctx)
			group.Bind(controller.Hello)
			//登录验证拦截

		})
	})
}
