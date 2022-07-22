package controller

import (
	"context"
	v1 "gf-demo-api/api/v1"
	"gf-demo-api/internal/model"
	"gf-demo-api/internal/service"
	"gf-demo-api/utility"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/gmode"
)

var (
	Login = loginController{}
)

type loginController struct {
}

func (c *loginController) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {
	var (
		user  *model.LoginUserRes
		token string
	)
	//判断验证码是否正确
	debug := gmode.IsDevelop()

	if !debug {
		if !service.Captcha().VerifyString(req.VerifyKey, req.VerifyCode) {
			err = gerror.NewCode(gcode.New(model.CaptchaInvalid.Code(), model.CaptchaInvalid.Desc(), nil))
			return
		}
	}

	user, err = service.User().GetAdminUserByUsernamePassword(ctx, req)
	if err != nil {
		return
	}
	ip := utility.GetClientIp(ctx)
	userAgent := utility.GetUserAgent(ctx)
	key := gconv.String(user.Id) + "-" + gmd5.MustEncryptString(user.UserName) + gmd5.MustEncryptString(user.UserPassword)
	if g.Cfg().MustGet(ctx, "gfToken.multiLogin").Bool() {
		key = gconv.String(user.Id) + "-" + gmd5.MustEncryptString(user.UserName) + gmd5.MustEncryptString(user.UserPassword+ip+userAgent)
	}
	user.UserPassword = ""
	token, err = service.GfToken().GenerateToken(ctx, key, user)
	if err != nil {
		return
	}

	res = &v1.LoginRes{
		UserInfo: user,
		Token:    token,
	}
	return
}

// LoginOut 退出登录
func (c *loginController) LoginOut(ctx context.Context, req *v1.LoginOutReq) (res *v1.LoginOutRes, err error) {
	err = service.GfToken().RemoveToken(ctx, service.GfToken().GetRequestToken(g.RequestFromCtx(ctx)))
	return
}
