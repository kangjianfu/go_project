package models

import (
	"github.com/Unknwon/com"
	"github.com/astaxie/beego/orm"
	//前面添加下划线是指：只加载默认的函数。因为是驱动所以只需要进行驱动的注册。
	//_ "github.com/lib/pq"//postgresql postgresql 数据驱动
	// _ "github.com/go-sql-driver/mysql" //mysql 数据库驱动
	"os"
	"path"
	"time"

	_ "github.com/mattn/go-sqlite3" // sqlite3 驱动。
)

const (
	_DB_NAME        = "data/blog.db" // 数据库位置
	_SQLITE3_DRIVER = "sqlite3"      //驱动名称
)

//创建分类表
type Category struct {
	Id                 int64
	Titele             string
	Create_time        time.Time `orm:"index"`
	Views              int64     `orm:"index"`
	Topic_time         time.Time `orm:"index"` //创建时间
	Topic_count        int64     `orm:"index"` //文章数
	Topic_last_user_id int64     //最后文章发布人
}

//创建 文章
type Topic struct {
	Id                 int64
	Uid                int64
	Title              string `orm:"size(2000)"`
	Content            string `orm:"size(5000)"` //文字内容
	Attachment         string
	Create_time        time.Time `orm:"index"`
	Update_time        time.Time `orm:"index"`
	Views              int64     `orm:"index"`
	Author             string
	Reply_time         time.Time //回复时间
	Reply_count        int64     `orm:"index"` //回复次数
	reply_last_user_id int64
}

//注册DB
func RegisterDB() {
	//检查文件是否存在
	if !com.IsExist(_DB_NAME) {
		os.MkdirAll(path.Dir(_DB_NAME), os.ModePerm) //创建当前目录 path.Dir(_DB_NAME):取出一个路径中的目录.os.ModePerm:文件权限
		os.Create(_DB_NAME)
	}
	//初始化数据库
	//注册模型
	orm.RegisterModel(new(Category), new(Topic), new(User), new(Role), new(Resource), new(Dict))
	// mysql / sqlite3 / postgres 这三种是默认已经注册过的，所以可以无需设置
	//注册驱动
	orm.RegisterDriver(_SQLITE3_DRIVER, orm.DRSqlite)
	//注册默认数据库 创建默认的数据库
	//注意：ORM 必须注册一个别名为 default 的数据库，作为默认使用
	orm.RegisterDataBase("default", _SQLITE3_DRIVER, _DB_NAME, 10)
	//根据数据库的别名，设置数据库的最大空闲连接
	orm.SetMaxIdleConns("default", 30)
}
