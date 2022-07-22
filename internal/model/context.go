package model

type Context struct {
	User *ContextUser
}
type ContextUser struct {
	Id       uint64 `json:"id"          description:"用户ID"`
	UserName string `json:"userName"    description:"姓名"`
	Age      int    `json:"age"         description:"年龄"`
	Sex      int    `json:"sex"         description:"性别（1男 2女 0未知）"`
}
