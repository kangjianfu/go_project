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

//share_info 表 视频记录表

type Share_info struct {
	Id                string `orm:"pk"`
	Description       string
	Location          string
	Zbk_id            string
	Service_code      string
	Video_url         string
	Nick_name         string
	Avata             string
	Created_at        time.Time
	Last_update       time.Time
	Last_notify_type  int
	Duration          float64
	Width             int64
	Height            int64
	Access_cnt        int64
	Zan_cnt           int64
	Report_cnt        int64
	Last_report_time  time.Time
	History_video_url string
}

//举报视频返回的dagrid
type DataGrid_report struct {
	Total int            `json:"total"`
	Rows  []*Report_item `json:"rows"`
}
type Report_item struct {
	Share_id         string  `json:"share_id"`
	Video_url        string  `json:"url"`
	Duration         float64 `json:"duration"`
	Zbk_id           string  `json:"zbk_id"`
	Nick_name        string  `json:"nick_name"`
	Location         string  `json:"location"`
	Description      string  `json:"description"`
	Access_cnt       int64   `json:"access_cnt"`
	Report_cnt       int64   `json:"report_cnt"`
	Last_report_time string  `json:"last_report_time"`
	Resolution       string  `json:"resolution"`
	Zan_cnt          int64   `json:"zan_cnt"`
}

//微信配置信息
type Wei_xin_config_info struct {
	Id                 int64  `form:"-"`
	Wx_app_name        string `form:"wx_app_name"`    //微信应用名称
	Wx_open_appid      string `form:"wx_open_appid"`  //微信应用的id
	Wx_open_secret     string `form:"wx_open_secret"` //微信应用id 密钥
	Wx_mp_appid        string `form:"wx_mp_appid"`    //微信公众号id
	Wx_mp_secret       string `form:"wx_mp_secret"`   //微信公众号的密钥
	Service_code       string `form:"service_code"`   //服务码
	Wx_app_description string `form:"description"`    //微信公众号的密钥
	Created_at         time.Time
}

//获取举报视频列表
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
			item.Last_report_time = report_list[k].Last_report_time.Local().Format("2006-01-02 15:04:05 ")
			item.Zan_cnt = report_list[k].Zan_cnt
			report_datagrid.Rows = append(report_datagrid.Rows, item)
		}
		fmt.Print(len(report_list))
		//report_datagrid.Rows = append(report_datagrid.Rows, &report_list)
		report_datagrid.Total = len(report_list)
	}
	return report_datagrid
}

//服务码列表
type Service_codes struct {
	Service_code string `orm:"pk"`
	Secret_key   string `json:"secret_key"`
}

//服务码返回的dagrid
type DataGrid_service_codes struct {
	Total int              `json:"total"`
	Rows  []*Service_codes `json:"rows"`
}

//获取服务码列表的方法
func DataGrid_service_codes_list(page int, rows int, parmas map[string]string) *DataGrid_service_codes {
	codes_grids := new(DataGrid_service_codes)
	service_codes := make([]*Service_codes, 0)
	o := orm.NewOrm()
	qs := o.QueryTable("service_codes")
	if parmas == nil {
		fmt.Println("过滤参数为空")
	} else {
		fmt.Println("过滤参数不为空")
	}
	_, err := qs.Limit(rows, page-1).All(&service_codes)
	if len(service_codes) > 0 && err == nil {
		for _, v := range service_codes {
			codes_grids.Rows = append(codes_grids.Rows, v)
		}
		codes_grids.Total = len(service_codes)
	}
	return codes_grids
}

//查询对象
type Query_model struct {
	Page   int
	Rows   int
	Params map[string]string
}
type DataGrid_model struct {
	Total int64         `json:"total"`
	Rows  []interface{} `json:"rows"`
}

//微信配置信息的datagrid
type DataGrid_wx_infos struct {
	Total int                    `json:"total"`
	Rows  []*Wei_xin_config_info `json:"rows"`
}

//获取微信配置信息列表的方法
func DataGrid_Wx_infos_list(page int, rows int, parmas map[string]string) *DataGrid_wx_infos {
	wx_infos_grids := new(DataGrid_wx_infos)
	wx_infos := make([]*Wei_xin_config_info, 0)
	o := orm.NewOrm()
	qs := o.QueryTable("wei_xin_config_info")
	if parmas == nil {
		fmt.Println("过滤参数为空")
	} else {
		fmt.Println("过滤参数不为空")
	}
	_, err := qs.Limit(rows, page-1).All(&wx_infos)
	if len(wx_infos) > 0 && err == nil {
		for _, v := range wx_infos {
			wx_infos_grids.Rows = append(wx_infos_grids.Rows, v)
		}
		wx_infos_grids.Total = len(wx_infos)
	} else {
		wx_infos_grids.Rows = nil
	}
	return wx_infos_grids
}

func Save_wx_info(wx_info *Wei_xin_config_info) int64 {
	o := orm.NewOrm()
	wx_info.Created_at = time.Now().Local()
	id, err := o.Insert(wx_info)
	if err == nil {
		return id
	} else {
		return 0
	}

}

//获取所有视频信息的列表
func Get_all_videos(model *Query_model) *DataGrid_model {
	grids := new(DataGrid_model)
	//定义一个存放查询出结果的数组
	video_items := make([]*Share_info, 0)
	o := orm.NewOrm()
	qs := o.QueryTable("share_info")
	if model.Params != nil {
		fmt.Println("传入查询参数不为空")
	} else {
		fmt.Println("传入查询参数为空")
	}
	if model.Rows == 0 {
		model.Rows = 20
	}
	if model.Page == 0 {
		model.Page = 1
	}
	qs = qs.OrderBy("-created_at")
	_, err := qs.Limit(model.Rows, model.Page-1).All(&video_items)
	if len(video_items) > 0 && err == nil {
		for _, v := range video_items {
			v.Created_at = v.Created_at.Location()
			grids.Rows = append(grids.Rows, v)
		}

	}
	total, _ := qs.Count()
	grids.Total = total
	return grids

}

func RegisterDB() {
	orm.RegisterModel(new(Report_history), new(Share_info), new(Service_codes), new(Wei_xin_config_info))
	orm.RegisterDriver(_POSTGRES_DRIVER, orm.DRPostgres)
	orm.RegisterDataBase("default", _POSTGRES_DRIVER, "user=zbk password=1T1Po7OgTyvz dbname=zbk host=db.zhiboyun.com port=5432 sslmode=disable", 10)
	orm.SetMaxIdleConns("default", 10)
	orm.RegisterDataBase("zbk", _POSTGRES_DRIVER, "user=zbk password=1T1Po7OgTyvz dbname=zbk host=db.zhiboyun.com port=5432 sslmode=disable", 5)
	orm.SetMaxIdleConns("zbk", 5)
	orm.RunSyncdb("zbk", false, true)
}
