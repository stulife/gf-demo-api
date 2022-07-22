package service

import (
	"context"
	v1 "gf-demo-api/api/v1"
	"gf-demo-api/internal/dao"
	"gf-demo-api/internal/model"
	"gf-demo-api/utility"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

type IUser interface {
	GetAdminUserByUsernamePassword(ctx context.Context, req *v1.LoginReq) (user *model.LoginUserRes, err error)
}

type userImpl struct {
}

var (
	userService = userImpl{}
)

func User() IUser {
	return &userService
}

func (s *userImpl) GetAdminUserByUsernamePassword(ctx context.Context, req *v1.LoginReq) (user *model.LoginUserRes, err error) {
	user, err = s.GetUserByUsername(ctx, req.Username)
	if err != nil || user == nil || utility.EncryptPassword(req.Password, user.UserSalt) != user.UserPassword {
		err = gerror.NewCode(gcode.New(model.AccountPassWordFailed.Code(), model.AccountPassWordFailed.Desc(), nil))
		return
	}
	//账号状态
	if user.UserStatus == 0 {
		err = gerror.NewCode(gcode.New(model.AccountLock.Code(), model.AccountLock.Desc(), nil))
		return
	}
	return
}

// GetUserByUsername 通过用户名获取用户信息
func (s *userImpl) GetUserByUsername(ctx context.Context, userName string) (user *model.LoginUserRes, err error) {
	user = &model.LoginUserRes{}
	err = dao.User.Ctx(ctx).Fields(user).Where(dao.User.Columns().UserName, userName).Scan(user)
	return
}
