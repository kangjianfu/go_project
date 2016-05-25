package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

// 创建一个自己的 handlerfunc  定义一个map key 为string 类型 value 是一个函数类型

var mux map[string]func(w http.ResponseWriter, r *http.Request)

func main() {
	//需要获取地址
	ser := http.Server{
		Addr:        ":8080",
		Handler:     &myHandler{},
		ReadTimeout: 5 * time.Second,
	}
	//初始化map
	mux = make(map[string]func(w http.ResponseWriter, r *http.Request))
	//定义路径
	mux["/hello"] = sayHello
	mux["/bye"] = bye
	info := ser.ListenAndServe()
	if info != nil {
		log.Fatal(info)
	}

}

type myHandler struct{}

func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//判断路径是否存在 如果存在 则直接调用。运用了 map的 ok patten
	if h, ok := mux[r.URL.String()]; ok {
		h(w, r)
		return
	}
	io.WriteString(w, "url: "+r.URL.String())
}
func sayHello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "say hello")
}
func bye(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "say bye")
}
