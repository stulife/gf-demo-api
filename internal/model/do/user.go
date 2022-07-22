// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure of table tb_user for DAO operations like Where/Data.
type User struct {
	g.Meta      `orm:"table:tb_user, do:true"`
	Id          interface{} // 用户ID
	UserName    interface{} // 姓名
	Age         interface{} // 年龄
	Sex         interface{} // 性别（1男 2女 0未知）
	PhoneNumber interface{} // 手机号码
	CompanyId   interface{} // 单位id
	CompanyName interface{} // 单位
	Status      interface{} // 帐号状态（0正常 1停用）
	OpenId      interface{} // 来源账号id
	Avatar      interface{} // 图像路径
	SysUserId   interface{} // 中台用户ID
	SysPersonId interface{} // 中台人员ID
	AdminFlag   interface{} // 管理员标志，0否，1是
	DelFlag     interface{} // 删除标志（0代表存在 1代表删除）
	CreateBy    interface{} // 创建者
	CreateTime  *gtime.Time // 创建时间
	UpdateBy    interface{} // 更新者
	UpdateTime  *gtime.Time // 更新时间
}
