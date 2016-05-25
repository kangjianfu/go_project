package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	//设置路由 如果浏览器访问 / 则调用这sayHello 方法
	http.HandleFunc("/", sayHello)
	//监听信息
	info := http.ListenAndServe(":8080", nil)
	fmt.Println(info)
	// ListenAndServe always returns a non-nil error. 官网文档 listenAndserver 总是返回一个不是 nil的
	if info != nil {
		log.Fatal(info)
	}
}

//方法
func sayHello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "你好世界。this is version——1 这是第一个demo")
}
