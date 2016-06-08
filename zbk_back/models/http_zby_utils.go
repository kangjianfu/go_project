package models

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"log"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
)

//签名字符串时候是不带 ？ 。 uri 和参数之间没有？
var current_time = time.Now()
var timestamp = strconv.FormatInt(current_time.Unix()*1000, 10)
var zby_suffix = "service_code=" + beego.AppConfig.String("service_code")

func Get_task_list() string {
	//l := logs.NewConsole()
	var url = beego.AppConfig.String("aws_zby_host") + beego.AppConfig.String("aws_zby_task_list") + "?" + zby_suffix
	var uri = beego.AppConfig.String("aws_zby_task_list") + zby_suffix + timestamp
	req := httplib.Get(url)
	req.Header("xvs-timestamp", timestamp)
	req.Header("xvs-signature", Get_signature(uri))
	str, err := req.String()
	if err != nil {
		log.Fatal(err)

	}
	fmt.Println(str)
	return str

}

type List_files struct {
	Function string `json:"function"`
	Params   struct {
		Service_code string `json:"service_code"`
		Fields       string `json:"fields"`
		Limit        int    `json:"limit"`
		Order_asc    bool   `json:"order_asc"`
	} `json:"params"`
}

//api 返回的json 对象
type Result_model struct {
	Ret      int `json:"ret"`
	Item_cnt int `json:"item_cnt"`
	Items    []struct {
		File_name struct {
			S string
		} `json:"file_name"`
		Created_at struct {
			N string
		} `json:"created_at"`
		Width struct {
			N string
		} `json:"width"`
		Height struct {
			N string
		} `json:"height"`
		User_name struct {
			S string
		} `json:"user_name"`
		Duration struct {
			N string
		} `json:"duration"`
		Access_cnt struct {
			N string
		} `json:"access_cnt"`
		Zan_cnt struct {
			N string
		} `json:"zan_cnt"`
		Location struct {
			S string
		} `json:"location"`
		Url struct {
			S string
		} `json:"url"`
		Description struct {
			S string
		} `json:"description"`
	} `json:"items"`
}

//datagrid model for easyui

func Managent_api_list_files() *Result_model {
	reslut_mode := &Result_model{}
	files := &List_files{}
	client := &http.Client{}
	files.Function = "list_files"
	files.Params.Service_code = beego.AppConfig.String("service_code")
	files.Params.Fields = "file_name,created_at,duration,width,height,user_name,description,url,location,access_cnt,zan_cnt"
	files.Params.Limit = 29
	management_url := beego.AppConfig.String("aws_zby_host") + beego.AppConfig.String("aws_zby_manage_uri") + "?" + zby_suffix
	b, err := json.Marshal(files)
	if err != nil {
		log.Print("json 转化失败")
		log.Fatal(err)
		fmt.Println("encoding faild")
		return reslut_mode
	}
	jsonData := string(b)
	management_uri := beego.AppConfig.String("aws_zby_manage_uri") + zby_suffix + jsonData + timestamp
	req, err := http.NewRequest("POST", management_url, strings.NewReader(jsonData))
	if err != nil {
		log.Fatal(err)
	}
	singnature := Get_signature(management_uri)
	req.Header.Set("xvs-signature", singnature)
	req.Header.Set("xvs-timestamp", timestamp)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	reslut_data := string(body)
	err = json.Unmarshal([]byte(reslut_data), reslut_mode)
	if err != nil {
		log.Println("发生错误了， api 返回json 后，转为对象失败")

	} else {
		log.Println("success api 返回json 后 对象转换成功")
	}
	return reslut_mode

}

//用户列表的json
type User_list struct {
	Function string `json:"function"`
	Params   struct {
		Service_code string `json:"service_code"`
		Fields       string `json:"fields"`
		Page_index   int    `json:"page_index"`
		Per_page     int    `json:"per_page"`
	} `json:"params"`
}

//测试能不能获取用户的列表
func Test_management_user_list() {
	client := &http.Client{}
	management_url := beego.AppConfig.String("aws_zby_host") + beego.AppConfig.String("aws_zby_manage_uri") + "?" + zby_suffix
	list := &User_list{}
	list.Function = "list_users"
	list.Params.Service_code = beego.AppConfig.String("service_code")
	list.Params.Per_page = 10
	list.Params.Fields = "id,service_code,user_name"
	fmt.Println(list)
	b, err := json.Marshal(list)
	if err != nil {
		log.Print("json 转化失败")
		log.Fatal(err)
		fmt.Println("encoding faild")
		return
	}

	jsonData := string(b)
	management_uri := beego.AppConfig.String("aws_zby_manage_uri") + zby_suffix + jsonData + timestamp
	req, err := http.NewRequest("POST", management_url, strings.NewReader(jsonData))
	if err != nil {
		log.Fatal(err)

	}
	singnature := Get_signature(management_uri)
	req.Header.Set("xvs-signature", singnature)
	req.Header.Set("xvs-timestamp", timestamp)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))

}

//获取签名的方法
func Get_signature(content string) string {
	key := []byte(beego.AppConfig.String("code_key"))
	mac := hmac.New(sha256.New, key)
	mac.Write([]byte(content))
	signature := fmt.Sprintf("%x\n", mac.Sum(nil))
	return signature
}
