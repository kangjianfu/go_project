package models

import (
	"time"
)

//用户表
type User struct {
	Id          int64
	Name        string
	Titele      string
	Enname      string    //英文名称
	Create_time time.Time `orm:"index"`
	Update_time time.Time
	Phone       int64  `orm:"index"` //手机号
	Email       string //邮箱
	Password    string //密码
	Dser_type   uint8  //用户类型
	Del_flag    uint8  //删除标记

}

//角色表
type Role struct {
	Id          int64
	Name        string
	Enname      string    //英文名称
	Users       []*User   `orm:"rel(m2m)"` //一个角色可以属于多个用户
	Create_time time.Time `orm:"index"`
	Update_time time.Time
	Resources   []*Resource `orm:"rel(m2m)"`
	Del_flag    uint8       //删除标记
}

//资源表
type Resource struct {
	Id            int64
	Uri           string
	Name          string
	Shut_uri      string //短名称。可以判断权限用
	Resource_type uint8  //类型、功能类型。菜单类型
	Pid           int64  //父id
	Icon          string //图标
	Del_flag      uint8  //删除标记
}

//字典表
type Dict struct {
	Id          int64
	Name        string
	Description string //描述
	Remarks     string //备注
	Pid         int    //父id
	Del_flag    string //删除标记
}
