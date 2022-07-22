package cmd

import (
	"context"
	"gf-demo-api/internal/consts"
	"gf-demo-api/router"
	"gf-demo-api/utility"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/net/goai"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Group("/", func(group *ghttp.RouterGroup) {
				router.BindController(group)

			})
			enhanceOpenAPIDoc(s)
			s.Run()
			return nil
		},
	}
)

func enhanceOpenAPIDoc(s *ghttp.Server) {
	openapi := s.GetOpenApi()
	openapi.Config.CommonResponse = utility.DefaultResponse{}
	openapi.Config.CommonResponseDataField = `Data`
	securitySchemesMap := make(map[string]goai.SecuritySchemeRef)
	securitySchemesMap["Authorization"] = goai.SecuritySchemeRef{
		Value: &goai.SecurityScheme{
			Type:        "apiKey",
			In:          "header",
			Description: "访问授权",
			Name:        "Authorization",
		},
	}
	openapi.Components.SecuritySchemes = securitySchemesMap
	securityRequirement := goai.SecurityRequirement{
		"Authorization": {""},
	}
	openapi.Security = &goai.SecurityRequirements{
		securityRequirement,
	}
	// API description.
	openapi.Info = goai.Info{
		Title:       consts.OpenAPITitle,
		Description: consts.OpenAPIDescription,
		Contact: &goai.Contact{
			Name: consts.OpenAPIContactName,
			URL:  consts.OpenAPIContactUrl,
		},
		Version: "v1",
	}
}
