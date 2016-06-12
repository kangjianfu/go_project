package models

import (
	"fmt"

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

//举报返回的dagrid
type DataGrid_report struct {
	Total int            `json:"total"`
	Rows  []*Report_item `json:"rows"`
}
type Report_item struct {
	Id       int64  `json:"id"`
	Share_id string `json:"share_id"`
	Reason   string `json:"reason"`
	Comment  string `json:"commment"`
}

//举报的视频

func Datagrid_report_history(page int, rows int) *DataGrid_report {
	report_datagrid := new(DataGrid_report)
	report_list := make([]*Report_history, 0)
	o := orm.NewOrm()
	//qs:=o.QueryTable("report_history")
	//o.Using("zbk")
	qs := o.QueryTable("report_history")
	_, err := qs.Limit(rows, page-1).All(&report_list)
	fmt.Println(err)
	if len(report_list) > 0 {
		for k, _ := range report_list {
			item := new(Report_item)
			item.Id = report_list[k].Id
			item.Share_id = report_list[k].Share_id
			item.Reason = report_list[k].Reason
			item.Comment = report_list[k].Comment
			report_datagrid.Rows = append(report_datagrid.Rows, item)
		}
		fmt.Print(len(report_list))
		//report_datagrid.Rows = append(report_datagrid.Rows, &report_list)
		report_datagrid.Total = len(report_list)
	}
	return report_datagrid
}

func RegisterDB() {
	orm.RegisterModel(new(Report_history))
	orm.RegisterDriver(_POSTGRES_DRIVER, orm.DRPostgres)
	orm.RegisterDataBase("default", _POSTGRES_DRIVER, "user=zbk password=1T1Po7OgTyvz dbname=zbk host=db.zhiboyun.com port=5432 sslmode=disable", 10)
	orm.SetMaxIdleConns("default", 10)
	orm.RegisterDataBase("zbk", _POSTGRES_DRIVER, "user=zbk password=1T1Po7OgTyvz dbname=zbk host=db.zhiboyun.com port=5432 sslmode=disable", 10)
	orm.SetMaxIdleConns("zbk", 10)
	orm.RunSyncdb("zbk", false, true)
}
