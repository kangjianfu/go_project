package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	//自己实现handler mux
	//创建一个默认的服务
	mux := http.NewServeMux()
	//此时传入的 自己写的 handler myHandler 地址
	mux.Handle("/", &myHandler{})
	mux.HandleFunc("/hello", sayHello)

	//go 访问静态文件服务器配置
	//获取工作目录
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal("静态文件访问失败")
		return
	}
	//简易静态文件服务器的实现
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(wd))))

	//注册路由
	info := http.ListenAndServe(":8080", mux)

	if info != nil {
		log.Fatal(info)
	}
}

type myHandler struct{}

func (t *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "url"+r.URL.String())
}
func sayHello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "你好世界、你访问我了，我很开心呢。谢谢")
}
