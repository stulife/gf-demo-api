package model

type LoginUserRes struct {
	Id           uint64 `json:"id"          description:"用户ID"`
	UserName     string `json:"userName"    description:"姓名"`
	UserNickname string `orm:"user_nickname"    json:"userNickname"` // 用户昵称
	UserPassword string `orm:"user_password"    json:"userPassword"` // 登录密码;cmf_password加密
	UserSalt     string `orm:"user_salt"        json:"userSalt"`     // 加密盐
	UserStatus   uint   `orm:"user_status"      json:"userStatus"`   // 用户状态;0:禁用,1:正常,2:未验证
	Age          int    `json:"age"         description:"年龄"`
	Sex          int    `json:"sex"         description:"性别（1男 2女 0未知）"`
}
