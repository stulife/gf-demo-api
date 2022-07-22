package v1

import (
	"gf-demo-api/internal/model"
	"github.com/gogf/gf/v2/frame/g"
)

type LoginReq struct {
	g.Meta     `path:"/login" tags:"登录" method:"post" summary:"用户登录"`
	Username   string `p:"username" v:"required#用户名不能为空"`
	Password   string `p:"password" v:"required#密码不能为空"`
	VerifyCode string `p:"verifyCode" v:"required#验证码不能为空"`
	VerifyKey  string `p:"verifyKey"`
}

type LoginRes struct {
	g.Meta   `mime:"application/json"`
	UserInfo *model.LoginUserRes `json:"userInfo"`
	Token    string              `json:"token"`
}

type LoginOutReq struct {
	g.Meta        `path:"/loginOut" tags:"登录" method:"delete" summary:"退出登录"`
	Authorization string `p:"Authorization" in:"header" dc:"Bearer {{token}}"`
}

type LoginOutRes struct {
}
