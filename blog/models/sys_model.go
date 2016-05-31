package models

import (
	"time"
)

//用户表
type User struct {
	Id                 int64
	Name               string
	Titele             string
	Create_time        time.Time `orm:"index"`
	Update_time        time.Time
	Phone              int64  `orm:"index"` //手机号
	Email              string //邮箱
	Password           string //密码
	Topic_last_user_id int64  //最后文章发布人
}
