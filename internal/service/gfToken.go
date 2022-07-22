package service

import (
	"context"
	"gf-demo-api/internal/consts"
	"gf-demo-api/internal/model"
	"gf-demo-api/utility"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/tiger1103/gfast-token/gftoken"
	"sync"
)

type IGfToken interface {
	GenerateToken(ctx context.Context, key string, data interface{}) (keys string, err error)
	Middleware(group *ghttp.RouterGroup) error
	ParseToken(r *ghttp.Request) (*gftoken.CustomClaims, error)
	IsLogin(r *ghttp.Request) (b bool, failed *gftoken.AuthFailed)
	GetRequestToken(r *ghttp.Request) (token string)
	RemoveToken(ctx context.Context, token string) (err error)
}

type gfTokenImpl struct {
	*gftoken.GfToken
}

var gT = gfTokenImpl{
	GfToken: gftoken.NewGfToken(),
}

func GetFToken(options *model.TokenOptions) IGfToken {
	var fun gftoken.OptionFunc
	if options.CacheModel == consts.CacheModelRedis {
		fun = gftoken.WithGRedis()
	} else {
		fun = gftoken.WithGCache()
	}
	gT.GfToken = gftoken.NewGfToken(
		gftoken.WithCacheKey(options.CacheKey),
		gftoken.WithTimeout(options.Timeout),
		gftoken.WithMaxRefresh(options.MaxRefresh),
		gftoken.WithMultiLogin(options.MultiLogin),
		gftoken.WithExcludePaths(options.ExcludePaths),
		fun)
	return &gT
}

type gft struct {
	options *model.TokenOptions
	gT      IGfToken
	lock    *sync.Mutex
}

var gftService = &gft{
	options: nil,
	gT:      nil,
	lock:    &sync.Mutex{},
}

func GfToken() IGfToken {
	if gftService.gT == nil {
		gftService.lock.Lock()
		defer gftService.lock.Unlock()
		if gftService.gT == nil {
			ctx := gctx.New()
			err := g.Cfg().MustGet(ctx, "gfToken").Struct(&gftService.options)
			utility.ErrIsNil(ctx, err)
			gftService.gT = GetFToken(gftService.options)
		}
	}
	return gftService.gT
}
