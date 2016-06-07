package models

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
)

//签名字符串时候是不带？好的 uri 和参数之间没有？
var current_time = time.Now()
var timestamp = strconv.FormatInt(current_time.Unix()*1000, 10)
var zby_suffix = "service_code=" + beego.AppConfig.String("service_code")

func Get_task_list() string {
	var url = beego.AppConfig.String("aws_zby_host") + beego.AppConfig.String("aws_zby_task_list") + "?" + zby_suffix
	var uri = beego.AppConfig.String("aws_zby_task_list") + zby_suffix + timestamp
	//	mac := hmac.New(sha256.New, key)
	//	mac.Write([]byte(uri))
	//	//fmt.Printf("%x\n", mac.Sum(nil))
	//	signature := fmt.Sprintf("%x\n", mac.Sum(nil))
	//	fmt.Println("signature       =" + signature)

	req := httplib.Get(url)
	req.Header("xvs-timestamp", timestamp)
	req.Header("xvs-signature", Get_signature(uri))
	str, err := req.String()
	if err != nil {
		//t.Fatal(err)
	}
	fmt.Println(str)
	return str

}

type User_list struct {
	Function string `json:"function"`
	Params   struct {
		Service_code string `json:"service_code"`
		Fields       string `json:"fields"`
		Page_index   int    `json:"page_index"`
		Per_page     int    `json:"per_page"`
	} `json:"params"`
}

func Test_management_user_list() {
	fmt.Println("timestamt     ___ =" + timestamp)
	//var t1 = "1465314169466"
	//fmt.Println("t1     ___ =" + t1)

	management_url := beego.AppConfig.String("aws_zby_host") + beego.AppConfig.String("aws_zby_manage_uri") + "?" + zby_suffix
	list := &User_list{}
	list.Function = "list_users"
	list.Params.Service_code = beego.AppConfig.String("service_code")
	list.Params.Per_page = 10
	list.Params.Fields = "id,service_code,user_name"
	fmt.Println(list)
	b, err := json.Marshal(list)
	if err != nil {
		fmt.Println("encoding faild")
		return
	}

	jsonData := string(b)
	management_uri := beego.AppConfig.String("aws_zby_manage_uri") + zby_suffix + jsonData + timestamp
	fmt.Println("management_uri   " + management_uri)
	req := httplib.Post(management_url)
	req.Debug(true)
	singnature := Get_signature(management_uri)
	fmt.Println("singnature                        =" + singnature)
	req.Header("xvs-signature", singnature)
	req.Header("xvs-timestamp", timestamp)
	req.JSONBody(jsonData)
	fmt.Println("jsonData       " + jsonData)
	str, err := req.String()
	if err != nil {
		//t.Fatal(err)
		return
	}
	fmt.Println("post 方法" + str)

	fmt.Println("management_url     ___ =" + management_url)
	//var jsonData = `{"function":"list_users","params":{"service_code":"ZBK_WEIXIN","fields":"id,service_code,user_name","page_index":0,"per_page":5}}`

}

//获取签名的方法
func Get_signature(content string) string {
	key := []byte(beego.AppConfig.String("code_key"))
	mac := hmac.New(sha256.New, key)
	mac.Write([]byte(content))
	signature := fmt.Sprintf("%x\n", mac.Sum(nil))
	return signature
}
