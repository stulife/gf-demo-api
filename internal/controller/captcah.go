package controller

import (
	"context"
	v1 "gf-demo-api/api/v1"
	"gf-demo-api/internal/service"
)

var Captcha = captchaController{}

type captchaController struct {
}

// Get 获取验证码
func (c *captchaController) Get(ctx context.Context, req *v1.CaptchaReq) (res *v1.CaptchaRes, err error) {
	var (
		idKeyC, base64stringC string
	)
	idKeyC, base64stringC, err = service.Captcha().GetVerifyImgString(ctx)
	res = &v1.CaptchaRes{
		Key: idKeyC,
		Img: base64stringC,
	}
	return
}
