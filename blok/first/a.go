package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func (u User) Hello() {
	fmt.Println("hello world")
}
func main() {
	u := User{1, "张三那", 18}
	Info(u)
}
func Info(o interface{}) {
	//获取传入参数的类型
	t := reflect.TypeOf(o)
	fmt.Println("Type", t.Name())
	//获取传入参数所包含的字段
	v := reflect.ValueOf(o)
	fmt.Println("fileds")
	for i := 0; i < t.NumField(); i++ {
		//获取传入参数的自动名称
		f := t.Field(i)
		val := v.Field(i).Interface()
		fmt.Printf("%6s:%v =%v", f.Name, f.Type, val)
	}
}
