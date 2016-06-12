package models

import (
	"fmt"
	"strconv"
	"time"

	"github.com/astaxie/beego/orm"
	//前面添加下划线是指：只加载默认的函数。因为是驱动所以只需要进行驱动的注册。postgresql postgresql 数据驱动
	_ "github.com/lib/pq"
)

const (
	_DB_NAME         = "data/blog.db" // 数据库位置
	_POSTGRES_DRIVER = "postgres"     //驱动名称
)

type Report_history struct {
	Id       int64
	Share_id string
	Reason   string
	Comment  string
}
type Share_info struct {
	Id               string `orm:"pk"`
	Description      string
	Location         string
	Zbk_id           string
	Service_code     string
	Video_url        string
	Nick_name        string
	Avata            string
	Created_at       time.Time
	Last_update      time.Time
	Duration         float64
	Width            int64
	Height           int64
	Access_cnt       int64
	Report_cnt       int64
	Last_report_time time.Time
	Zan_cnt          int64
}

//举报返回的dagrid
type DataGrid_report struct {
	Total int            `json:"total"`
	Rows  []*Report_item `json:"rows"`
}
type Report_item struct {
	Share_id         string    `json:"share_id"`
	Video_url        string    `json:"url"`
	Duration         float64   `json:"duration"`
	Zbk_id           string    `json:"zbk_id"`
	Nick_name        string    `json:"nick_name"`
	Location         string    `json:"location"`
	Description      string    `json:"description"`
	Access_cnt       int64     `json:"access_cnt"`
	Report_cnt       int64     `json:"report_cnt"`
	Last_report_time time.Time `json:"last_report_time"`
	Resolution       string    `json:"resolution"`
	Zan_cnt          int64     `json:"zan_cnt"`
}

//举报的视频

func Datagrid_report_info(page int, rows int) *DataGrid_report {
	report_datagrid := new(DataGrid_report)
	report_list := make([]*Share_info, 0)
	o := orm.NewOrm()
	//qs:=o.QueryTable("report_history")
	//o.Using("zbk")
	qs := o.QueryTable("share_info")
	qs = qs.Filter("video_url__isnull", false)
	qs = qs.Filter("report_cnt__gt", 0)
	_, err := qs.Limit(rows, page-1).All(&report_list)
	if len(report_list) > 0 && err == nil {
		for k, _ := range report_list {
			item := new(Report_item)
			item.Share_id = report_list[k].Id
			item.Video_url = report_list[k].Video_url
			item.Zbk_id = report_list[k].Zbk_id
			item.Nick_name = report_list[k].Nick_name
			item.Duration = report_list[k].Duration
			item.Description = report_list[k].Description
			item.Location = report_list[k].Location
			item.Access_cnt = report_list[k].Access_cnt
			item.Report_cnt = report_list[k].Report_cnt
			item.Resolution = strconv.FormatInt(report_list[k].Width, 10) + " × " + strconv.FormatInt(report_list[k].Height, 10)
			item.Last_report_time = report_list[k].Last_report_time
			item.Zan_cnt = report_list[k].Zan_cnt
			report_datagrid.Rows = append(report_datagrid.Rows, item)
		}
		fmt.Print(len(report_list))
		//report_datagrid.Rows = append(report_datagrid.Rows, &report_list)
		report_datagrid.Total = len(report_list)
	}
	return report_datagrid
}

func RegisterDB() {
	orm.RegisterModel(new(Report_history), new(Share_info))
	orm.RegisterDriver(_POSTGRES_DRIVER, orm.DRPostgres)
	orm.RegisterDataBase("default", _POSTGRES_DRIVER, "user=zbk password=1T1Po7OgTyvz dbname=zbk host=db.zhiboyun.com port=5432 sslmode=disable", 10)
	orm.SetMaxIdleConns("default", 10)
	orm.RegisterDataBase("zbk", _POSTGRES_DRIVER, "user=zbk password=1T1Po7OgTyvz dbname=zbk host=db.zhiboyun.com port=5432 sslmode=disable", 5)
	orm.SetMaxIdleConns("zbk", 5)
	orm.RunSyncdb("zbk", false, true)
}
