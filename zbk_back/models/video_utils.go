package models

import (
	"log"
	"strconv"
	"strings"
	"time"
)

type DataGrid_video struct {
	Total int           `json:"total"`
	Rows  []*Video_info `json:"rows"`
}
type Video_info struct {
	File_name   string `json:"file_name"` //视频url
	Created_at  string `json:"created_at"`
	Resolution  string `json:"resolution"` //分辨率
	User_name   string `json:"user_name"`
	Access_cnt  int    `json:"access_cnt"`
	Zan_cnt     int    `json:"zan_cnt"`
	Location    string `json:"location"`
	Description string `json:"description"`
	Duration    string `json:"duration"`
	Url         string `json:"url"`
}

//把//api json 对象 转为 视频的dataGrid 返回
func Zby_model_to_datagrid_video(zby_model *Result_model) *DataGrid_video {
	var total = 0
	video_grid := &DataGrid_video{}
	for key, _ := range zby_model.Items {
		if strings.HasSuffix(zby_model.Items[key].File_name.S, "jpg") {
			continue
		} else {
			video_info := new(Video_info)
			video_info.File_name = zby_model.Items[key].File_name.S
			t, err := strconv.ParseInt(zby_model.Items[key].Created_at.N, 10, 64)
			if err != nil {
				log.Print("json 日期 转化失败")
			} else {
				tm := time.Unix(t, 10)
				video_info.Created_at = tm.Format("2006-01-02 03:04:05 PM")

			}
			video_info.User_name = zby_model.Items[key].User_name.S
			video_info.Url = zby_model.Items[key].Url.S
			video_info.Resolution = zby_model.Items[key].Width.N + " × " + zby_model.Items[key].Height.N
			video_info.Duration = zby_model.Items[key].Duration.N
			if zby_model.Items[key].Access_cnt.N != "" {
				access_cnt, err := strconv.Atoi(zby_model.Items[key].Access_cnt.N)
				if err != nil {
					log.Print("访问量 转为数字失败")
				} else {
					video_info.Access_cnt = access_cnt
				}
				zan_cnt, err := strconv.Atoi(zby_model.Items[key].Zan_cnt.N)
				if err != nil {
					log.Print("点赞量为 转为数字失败")
				} else {
					video_info.Zan_cnt = zan_cnt
				}
			}
			video_info.Location = zby_model.Items[key].Location.S
			video_info.Description = zby_model.Items[key].Description.S
			video_grid.Rows = append(video_grid.Rows, video_info)
			total++
		}
	}

	video_grid.Total = total
	return video_grid
}
