package main

// import (
// 	"fmt"
// 	"reflect"
// )
//
// type resume struct {
// 	Name string `info:"name" doc:"我的名字"`
// 	Sex  string `info:"sex"`
// }
//
// func findTag(arg interface{}) {
// 	// 获取元素
// 	t := reflect.TypeOf(arg).Elem()
// 	for i := 0; i < t.NumField(); i++ {
// 		info := t.Field(i).Tag.Get("info")
// 		doc := t.Field(i).Tag.Get("doc")
// 		fmt.Println("info : ", info, " doc : ", doc)
// 	}
// }
//
// func main() {
// 	var r resume
// 	findTag(&r)
// }
